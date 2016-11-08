package session

import (
	"fmt"
)

const (
	//session中的字段
	USER_ID = "userId"
)

var GlobalSessions *Manager

func CookieConfig(config Config) {
	var err error
	GlobalSessions, err = NewManager("redis", fmt.Sprintf(`{
            "cookieName" : "%s",
            "maxLifeTime" : %d,
            "maxCookieLife" : %d
            }`, config.CookieName, config.MaxLifeTime, config.MaxCookieLife))

	if err != nil {
		panic(err)
	}
}

