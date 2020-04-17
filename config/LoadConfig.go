package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"time"
)

var cfg *ini.File

type App struct {
	StorageRootPath string
	RuntimeRootPath string
}

var AppSetting = &App{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	RedisHost        string
	RedisPassword    string
	RedisMaxidle     int
	RedisMaxActive   int
	RedisIdleTimeout time.Duration
}

var RedisSetting = &Redis{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize        int
	JwtSecret       string
	JwtAdminTimeout int64
	JwtApiTimeout   int64
	ViewUrl         string

	TimeZone     string
	SessionName  string
	SessionStore string

	LogPath string
	LogName string
}

var ServerSetting = &Server{}

type File struct {
	ImagePrefixUrl string
	ImageSavePath  string
	ImageMovePath  string
	ImageMaxSize   int
	ImageAllowExt  []string
}

var FileSetting = &File{}

func init() {
	var err error
	//读取配置
	cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("无法解析 'config/config.ini':%v ", err)
	}

	fmt.Println("cfg:", cfg)

	//加载服务
	loadServer()

	//服务超时设置
	serverTimeOut()
}

//加载服务
func loadServer() {
	mapToSection("app", AppSetting)
	mapToSection("server", ServerSetting)
	mapToSection("file", FileSetting)
	mapToSection("database", DatabaseSetting)
	mapToSection("redis", RedisSetting)
}

//服务超时设置
func serverTimeOut() {
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
	RedisSetting.RedisIdleTimeout = RedisSetting.RedisIdleTimeout * time.Second

}

//映射服务
func mapToSection(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("无法获取'server': %v", err)
	}
}
