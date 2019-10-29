package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

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
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

//func init() {
//	var err error
//	Cfg, err = ini.Load("conf/app.ini")
//
//	if err != nil {
//		log.Fatal("Fail to parse 'conf/app.ini': %v", err)
//	}
//
//	LoadBase()
//	LoadServer()
//	LoadApp()
//}
//
//func LoadBase() {
//	// github.com/go-ini/ini获取配置用法
//	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
//}
//
//func LoadServer() {
//	// github.com/go-ini/ini获取指定分区===>分区应该是指app.ini里面的[XXXX]
//	sec, err := Cfg.GetSection("server")
//	if err != nil {
//		log.Fatal("Fail to get section 'server': %v", err)
//	}
//
//	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
//	//获取某个分区下的键：key := cfg.Section("").Key("key name")(忽略错误形式)
//	HttpPort = sec.Key("HttpPort").MustInt(8000)
//	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
//	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
//}
//
//func LoadApp() {
//	sec, err := Cfg.GetSection("app")
//	if err != nil {
//		log.Fatalf("Fail to get section 'app': %v", err)
//	}
//	//获取某个分区下的键：key := cfg.Section("").Key("key name")(忽略错误形式)
//	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
//	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
//}
