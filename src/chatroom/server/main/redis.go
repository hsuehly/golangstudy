package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

//"39.107.92.94:6379"
func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address, redis.DialPassword("hsuehly"))
		},
		MaxIdle:     maxIdle,     // 最大空闲连接数
		MaxActive:   maxActive,   // 表示数据库的最大连接数
		IdleTimeout: idleTimeout, // 最大空闲时间

	}
}
