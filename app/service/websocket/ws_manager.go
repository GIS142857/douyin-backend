package websocket

import (
	"douyin-backend/app/utils/websocket/core"
	"sync"
)

// WsManager 管理 WebSocket 客户端连接
type WsManager struct {
	clients sync.Map // 存储所有的 WebSocket 连接（key: 用户ID, value: WebSocket连接）
}

// 初始化一个单例
var wsManager = &WsManager{
	clients: sync.Map{},
}

// AddClient 添加客户端连接
func (m *WsManager) AddClient(userID int64, conn *core.Client) {
	m.clients.Store(userID, conn)
}

// RemoveClient 删除客户端连接
func (m *WsManager) RemoveClient(userID int64) {
	m.clients.Delete(userID)
}

// GetClient 获取客户端连接
func (m *WsManager) GetClient(userID int64) (*core.Client, bool) {
	client, ok := m.clients.Load(userID)
	if ok {
		return client.(*core.Client), true
	}
	return nil, false
}
