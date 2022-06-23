package nosql

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"mubackend/config"
)

var  RedisConn redis.Conn

func InitRedis(){
	dial, err := redis.Dial("tcp", config.Cfg.RedisNetwork)
	if err !=nil {
		log.Printf("redis connect false :%s",err)
	}
	dial.Do("Auth",config.Cfg.RedisPassword)
	RedisConn = dial
}
