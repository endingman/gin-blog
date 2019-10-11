package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

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

/**
编写与配置项保持一致的结构体（App、Server、Database）
使用 MapTo 将配置项映射到结构体上
对一些需特殊设置的配置项进行再赋值
*/
func Setup() {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
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
