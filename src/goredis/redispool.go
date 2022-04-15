package goredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "39.107.92.94:6379", redis.DialPassword("hsuehly"))
		},
		MaxIdle:     8,   // 最大空闲连接数
		MaxActive:   0,   // 表示数据库的最大连接数
		IdleTimeout: 100, // 最大空闲时间

	}
}
func RedisPool() {
	c := pool.Get()
	// 创建完连接一定要关闭
	c.Close()
	fmt.Println("c")

	_, err := c.Do("hset", "user", "name", "hsuehly")
	if err != nil {
		println("set 失败", err)
		return
	}
	// 通过go 向redis 读取数据 返回的是一个interface{} 打印是byte切片
	//redis.Strings() 多个类型加s
	r1, err := redis.String(c.Do("hset", "100"))
	if err != nil {
		fmt.Println("获取失败", err)
	}
	fmt.Println("连接ok， 获取数据", r1)
	// 一旦关闭连接池那么后面就不能继续使用
	//pool.Close()
}
