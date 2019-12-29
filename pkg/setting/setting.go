package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HttpPost int
	ReadTimeOut time.Duration
	WriteTimeOut time.Duration

	PageSize int
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse conf: %v",err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase()  {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section server: %v",err)
	}

	HttpPost = sec.Key("HTTP_POST").MustInt(8000)
	ReadTimeOut = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60))
	WriteTimeOut = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60))
}

func LoadApp()  {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get app : %v",err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("123")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}