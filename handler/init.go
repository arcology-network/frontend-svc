package handler

import (
	"github.com/astaxie/beego/session"
)

var (
	globalSessions  *session.Manager
	accessTokenHash string
)

func init() {
	globalSessions, _ = session.NewManager("cookie", &session.ManagerConfig{
		CookieName:      "gosessionid",
		Gclifetime:      3600,
		EnableSetCookie: true,
		ProviderConfig:  `{"cookieName":"gosessionid","securityKey":"cookiehashkey"}`,
	})
	go globalSessions.GC()
}

func Init(ath string) {
	accessTokenHash = ath
}
