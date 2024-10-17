package variable

import (
	"douyin-backend/app/global/my_errors"
	"douyin-backend/app/utils/snow_flake/snowflake_interf"
	"douyin-backend/app/utils/yml_config/interf"
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

var (
	BasePath           string       // 定义项目的根目录
	EventDestroyPrefix = "Destroy_" //  程序退出时需要销毁的事件前缀
	ConfigKeyPrefix    = "Config_"  //  配置文件键值缓存时，键的前缀

	ZapLog          *zap.Logger               // 全局日志指针
	ConfigYml       interf.YmlConfigInterface // 全局配置文件指针
	ConfigGormv2Yml interf.YmlConfigInterface // 全局配置文件指针
	DateFormat      = "2024-01-02 15:04:05"   //  配置文件键值缓存时，键的前缀

	// GormDbMysql gorm 数据库客户端，如果您操作数据库使用的是gorm，请取消以下注释，在 bootstrap>init 文件，进行初始化即可使用
	GormDbMysql      *gorm.DB // 全局gorm的客户端连接
	GormDbSqlserver  *gorm.DB // 全局gorm的客户端连接
	GormDbPostgreSql *gorm.DB // 全局gorm的客户端连接

	//雪花算法全局变量
	SnowFlake snowflake_interf.InterfaceSnowFlake

	//websocket
	WebsocketHub              interface{}
	WebsocketHandshakeSuccess = `{"code":2000,"msg":"ws连接成功","data":""}`
	WebsocketServerPingMsg    = "Server->Ping->Client"

	//casbin 全局操作指针
	Enforcer *casbin.SyncedEnforcer

	//  用户自行定义其他全局变量 ↓
	SystemCreateKey = "system_menu_create" // 系统菜单数据编辑界面用户以 raw 格式提交的 json 存储在上下文的键
	SystemEditKey   = "system_menu_edit"
)

func init() {
	// 1.初始化程序根目录
	if path, err := os.Getwd(); err == nil {
		// 路径进行处理，兼容单元测试程序程序启动时的奇怪路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(path, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = path
		}
	} else {
		log.Fatal(my_errors.ErrorsBasePath)
	}
}
