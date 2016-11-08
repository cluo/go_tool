package myredis

import (
	"errors"
	"gopkg.in/redis.v4"
	"time"
)

// redis tool

type RedisConfig struct {
	Host     string
	Password string
	DB       int
}

type MyRedis struct {
	Config RedisConfig
	Client *redis.Client
}

// return myredis
func NewRedis(config RedisConfig) (MyRedis, error) {
	myredis := MyRedis{Config: config}
	client := redis.NewClient(&redis.Options{
		Addr:       config.Host,
		Password:   config.Password, // no password set
		DB:         config.DB,       // use default DB
		MaxRetries: 2,               // fail command retry 2
		PoolSize:   20,              // redis pool size
		// another options is default
	})

	pong, err := client.Ping().Result()
	if err == nil && pong == "PONG" {
		myredis.Client = client
	}
	return myredis, err
}

// set key
func (db *MyRedis) Set(key string, value string, expire time.Duration) error {
	return db.Client.Set(key, value, expire).Err()
}

// get key
func (db *MyRedis) Get(key string) (string, error) {
	result, err := db.Client.Get(key).Result()
	if err == redis.Nil {
		return "", errors.New("redis key does not exists")
	} else if err != nil {
		return "", err
	} else {
		return result, err
	}
}

func (db *MyRedis) Lpush(key string, values ...interface{}) (int64, error) {
	return db.Client.LPush(key, values...).Result()
}

func (db *MyRedis) Lpushx(key string, values interface{}) (int64, error) {
	num, err := db.Client.LPushX(key, values).Result()
	if err != nil {
		return 0, err
	}
	if num == 0 {
		return 0, errors.New("Redis List not exist")
	} else {
		return num, err
	}
}

func (db *MyRedis) Rpush(key string, values ...interface{}) (int64, error) {
	return db.Client.RPush(key, values...).Result()
}

func (db *MyRedis) Rpushx(key string, values interface{}) (int64, error) {
	num, err := db.Client.RPushX(key, values).Result()
	if err != nil {
		return 0, err
	}
	if num == 0 {
		return 0, errors.New("Redis List not exist")
	} else {
		return num, err
	}
}

func (db *MyRedis) Llen(key string) (int64, error) {
	return db.Client.LLen(key).Result()
}

func (db *MyRedis) Rpop(key string) (string, error) {
	return db.Client.RPop(key).Result()
}

func (db *MyRedis) Lpop(key string) (string, error) {
	return db.Client.LPop(key).Result()
}

func (db *MyRedis) Brpop(timeout int, keys ...string) ([]string, error) {
	timeouts := time.Duration(timeout) * time.Second
	return db.Client.BRPop(timeouts, keys...).Result()
}

// if timeout is zero will be block until...
// and if  keys has many will return one such as []string{"pool","b"},pool is list,b is value
func (db *MyRedis) Blpop(timeout int, keys ...string) ([]string, error) {
	timeouts := time.Duration(timeout) * time.Second
	return db.Client.BLPop(timeouts, keys...).Result()
}

func (db *MyRedis) Brpoplpush(source, destination string, timeout int) (string, error) {
	timeouts := time.Duration(timeout) * time.Second
	return db.Client.BRPopLPush(source, destination, timeouts).Result()
}

func (db *MyRedis) Rpoplpush(source, destination string) (string, error) {
	return db.Client.RPopLPush(source, destination).Result()
}

func(db *MyRedis) Hexists(key, field string)(bool,error){
	return db.Client.HExists(key,field).Result()
}

func(db *MyRedis) Hget(key, field string)(string,error){
	return db.Client.HGet(key,field).Result()
}

func(db *MyRedis) Hset(key, field, value string)(bool,error){
	return db.Client.HSet(key,field,value).Result()
}
