package nosql

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)
func IsHas(tradeHash string) bool {
	result, err := redis.String(RedisConn.Do("Get",tradeHash))
	if err !=nil {
		log.Printf("dont have this txhash ,pass %s",tradeHash)
		return false
	}
	if len(result)==0{
		return false
	}else {
		return true
	}
}
func InsertHashToRedis(tradeHash string)bool  {
	_, err := RedisConn.Do("Set", tradeHash, time.Now())
	if err !=nil {
		log.Printf("write into redis fail :%s",err)
		return false
	}
	RedisConn.Do("expire", tradeHash, "999999999999999999")
	return true
}


