package session

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

const (
	ProviderRedisType = iota + 1
	MemoryProviderType
)

//session manager
type Manager struct {
	Provider Provider
	Config   *Config
}

type Config struct {
	CookieName    string `json:"cookieName"`
	MaxLifeTime   int    `json:"maxLifeTime"`
	MaxCookieLife int    `json:"maxCookieLife"`
}

type Provider interface {
	Init(sid string) *Session
	Read(sid string) (*Session, error)
	Release(session *Session, w http.ResponseWriter) error
	Delete(sid string)
}

func NewManager(typ string, configStr string) (m *Manager, err error) {
	var config = &Config{}
	err = json.Unmarshal([]byte(configStr), config)
	if err != nil {
		return
	}
	m = &Manager{}
	switch typ {
	case "redis":
		m.Provider = providers[ProviderRedisType]
	default:
		err = fmt.Errorf("no such session type %s", typ)
	}
	m.Config = config
	return
}

func (m *Manager) Start(w http.ResponseWriter, r *http.Request) (session *Session) {
	//判断是否有sessionId
	c, err := r.Cookie(m.Config.CookieName)
	if c == nil || err != nil {

		//set sessionId
		sessionId := generateSessionId(r)
		cookie := &http.Cookie{
			Name:     m.Config.CookieName,
			Value:    sessionId,
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
			MaxAge:   m.Config.MaxCookieLife,
		}
		http.SetCookie(w, cookie)
		r.AddCookie(cookie)
		session = m.Provider.Init(sessionId)
	} else {
		sessionId, _ := url.QueryUnescape(c.Value)
		session, err = m.Provider.Read(sessionId)
		if err != nil || session == nil {
			session = m.Provider.Init(sessionId)
		}
	}
	session.ExpireSeconds = m.Config.MaxLifeTime
	return
}


func generateSessionId(r *http.Request) string {
	var sign = fmt.Sprintf("%s%d%d", r.RemoteAddr, time.Now().Nanosecond(), rand.Intn(100000))
	h := md5.New()
	h.Write([]byte(sign))
	return "hunterhug_session_" + hex.EncodeToString(h.Sum(nil))
}
