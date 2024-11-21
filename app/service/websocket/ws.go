package websocket

import (
	"douyin-backend/app/global/consts"
	"douyin-backend/app/global/my_errors"
	"douyin-backend/app/global/variable"
	userstoken "douyin-backend/app/service/users/token"
	"douyin-backend/app/utils/websocket/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"strconv"
)

type Ws struct {
	WsClient *core.Client
}

type Msg struct {
	TxUid      string `json:"tx_uid"`      // 发送者 UID
	RxUid      string `json:"rx_uid"`      // 接收者 UID
	MsgType    int    `json:"msg_type"`    // 消息类型
	MsgData    string `json:"msg_data"`    // 消息内容
	ReadState  int    `json:"read_state"`  // 消息读取状态
	CreateTime int    `json:"create_time"` // 创建时间
	DeleteTime int    `json:"delete_time"` // 删除时间
}

// onOpen 事件函数
func (w *Ws) OnOpen(context *gin.Context) (*Ws, bool) {
	if client, ok := (&core.Client{}).OnOpen(context); ok {

		//token := context.GetString(consts.ValidatorPrefix + "token")
		//variable.ZapLog.Info("获取到的客户端上线时携带的唯一标记值：", zap.String("token", token))

		// 成功上线以后，开发者可以基于客户端上线时携带的唯一参数(这里用token键表示)
		// 在数据库查询更多的其他字段信息，直接追加在 Client 结构体上，方便后续使用
		//client.ClientMoreParams.UserParams1 = "123"
		//client.ClientMoreParams.UserParams2 = "456"
		//fmt.Printf("最终每一个客户端(client) 已有的参数：%+v\n", client)

		w.WsClient = client
		go w.WsClient.Heartbeat() // 一旦握手+协议升级成功，就为每一个连接开启一个自动化的隐式心跳检测包
		return w, true
	} else {
		return nil, false
	}
}

// OnMessage 处理业务消息
func (w *Ws) OnMessage(context *gin.Context) {
	// 当前用户的 UID
	var txUid int64
	token := context.GetString(consts.ValidatorPrefix + "token")
	tokenIsEffective := userstoken.CreateUserFactory().IsEffective(token)
	if tokenIsEffective {
		if customeToken, err := userstoken.CreateUserFactory().ParseToken(token); err == nil {
			txUid = customeToken.UID
		} else {
			variable.ZapLog.Error("Ws.OnMessage ParseToken Failed!", zap.Error(err))
			return
		}
	}
	// 注册当前用户的 WebSocket 连接
	wsManager.AddClient(txUid, w.WsClient)

	go w.WsClient.ReadPump(func(messageType int, receivedData []byte) {
		// 解析 JSON 消息
		var msg Msg
		if err := json.Unmarshal(receivedData, &msg); err != nil {
			variable.ZapLog.Error("消息解析失败", zap.Error(err))
			return
		}

		// 确保发送者 UID 为当前连接用户
		msg.TxUid = strconv.FormatInt(txUid, 10)

		// 查找接收者 WebSocket 客户端
		rxUid, _ := strconv.ParseInt(msg.RxUid, 10, 64)
		targetClient, exists := wsManager.GetClient(rxUid)
		if exists {
			// 转发消息
			responseMsg, _ := json.Marshal(msg)
			if err := targetClient.SendMessage(messageType, string(responseMsg)); err != nil {
				variable.ZapLog.Error("转发消息失败", zap.Error(err))
			} else {
				// 更新消息状态为 SENT
				msg.ReadState = 1
			}
		} else {
			// 接收者不在线，将消息存储到数据库
			fmt.Println("接收者不在线!")
		}

	}, w.OnError, func() {
		// 移除用户连接
		wsManager.RemoveClient(txUid)
		w.OnClose()
	})
}

// OnError 客户端与服务端在消息交互过程中发生错误回调函数
func (w *Ws) OnError(err error) {
	variable.ZapLog.Error("远端掉线、卡死、刷新浏览器等会触发该错误:", zap.Error(err))
	//fmt.Printf("远端掉线、卡死、刷新浏览器等会触发该错误: %v\n", err.Error())
}

// OnClose 客户端关闭回调，发生onError回调以后会继续回调该函数
func (w *Ws) OnClose() {
	w.WsClient.State = 0
	w.WsClient.Hub.UnRegister <- w.WsClient // 向hub管道投递一条注销消息，由hub中心负责关闭连接、删除在线数据
}

// GetOnlineClients 获取在线的全部客户端
func (w *Ws) GetOnlineClients() int {
	//fmt.Printf("在线客户端数量：%d\n", len(w.WsClient.Hub.Clients))
	return len(w.WsClient.Hub.Clients)
}

// BroadcastMsg 向全部在线客户端广播消息
// 广播函数可能被不同的逻辑同时调用，由于操作的都是 Conn ，因此为了保证并发安全，加互斥锁
func (w *Ws) BroadcastMsg(sendMsg string) {
	for onlineClient := range w.WsClient.Hub.Clients {
		//获取每一个在线的客户端，向远端发送消息
		if err := onlineClient.SendMessage(websocket.TextMessage, sendMsg); err != nil {
			variable.ZapLog.Error(my_errors.ErrorsWebsocketWriteMgsFail, zap.Error(err))
		}
	}
}
