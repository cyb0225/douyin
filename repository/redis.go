package repository

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

var (
	Client redis.Conn	
)

// InitRedis init redis conn
func InitRedis() (err error) {
	Client, err = redis.Dial("tcp", "127.0.0.1:6379")
    if err != nil {
        return 
    }

	log.Println("redis connect successfully")
	

	return nil
}