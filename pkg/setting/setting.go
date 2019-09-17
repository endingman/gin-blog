package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	// 变量大写，外部包可以调用,相当于公有私有
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSecret    string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		log.Fatal("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	// github.com/go-ini/ini获取配置用法
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	// github.com/go-ini/ini获取指定分区===>分区应该是指app.ini里面的[XXXX]
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal("Fail to get section 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	//获取某个分区下的键：key := cfg.Section("").Key("key name")(忽略错误形式)
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	//获取某个分区下的键：key := cfg.Section("").Key("key name")(忽略错误形式)
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
