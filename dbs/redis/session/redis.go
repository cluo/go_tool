package session

import (
	"bytes"
	"encoding/gob"
	"net/http"
	"gopkg.in/redis.v4"
	"time"
)

var (
	_RedisProvider *RedisProvider
	providers = map[int]Provider{}
)

//提供session的redis实现
type RedisProvider struct {
	redisClient *redis.Client
}

func init() {
}

func Connect(client *redis.Client) {
	_RedisProvider = &RedisProvider{
		redisClient: client,
	}
	providers[ProviderRedisType] = _RedisProvider

}

func (this *RedisProvider) Init(sid string) (session *Session) {
	session = NewSession(sid, make(map[string]interface{}), this)
	return
}

func (this *RedisProvider) Release(session *Session, w http.ResponseWriter) (err error) {
	if session.data == nil {
		return
	}
	var buffer = bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buffer)
	err = enc.Encode(session.data)
	if err != nil {
		return
	}
	this.redisClient.Set(session.sessionId, buffer.String(), time.Duration(session.ExpireSeconds))
	return
}

func (this *RedisProvider) Read(sid string) (session *Session, err error) {
	var res = this.redisClient.Get(sid).Val()
	var mp = make(map[string]interface{})
	var buffer = bytes.NewBuffer([]byte(res))
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(&mp)
	if err != nil {
		return
	}
	session = NewSession(sid, mp, this)
	return
}

func (this *RedisProvider) Delete(sid string) {
	this.redisClient.Del(sid)
}

