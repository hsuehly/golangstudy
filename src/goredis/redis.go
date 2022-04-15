package goredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func Redis() {
	conn, err := redis.Dial("tcp", "39.107.92.94:6379", redis.DialPassword("hsuehly"))
	if err != nil {
		fmt.Println("redis 连接失败 错误信息=", err)
		return
	}
	defer conn.Close() //延迟关闭连接
	//_, err = conn.Do("hset", "users", "100", "{\"user_id\":100,\"user_pwd\":\"123456\",\"user_name\":\"hsuehly\"}")
	//if err != nil {
	//	fmt.Println("hset写入的失败", err)
	//}
	//r2, err := redis.String(conn.Do("hget", "users", "100"))
	//if err != nil {
	//	fmt.Println("获取hash数据失败", err)
	//}
	r2, err := redis.String(conn.Do("flushdb"))
	if err != nil {
		fmt.Println("获取hash数据失败", err)
	}
	fmt.Println("r2", r2)
	// 通过go 向redis 写入数据
	/*
		_, err = conn.Do("set", "name", "薛留阳hsuehly")
		if err != nil {
			println("set 失败", err)
			return
		}
		// 通过go 向redis 读取数据 返回的是一个interface{} 打印是byte切片
		//redis.Strings() 多个类型加s
		r1, err := redis.String(conn.Do("get", "name"))
		if err != nil {
			fmt.Println("获取失败", err)
		}

		// 通过go x写入hash 数据
		_, err = conn.Do("hset", "user", "name", "hsuehly")
		if err != nil {
			fmt.Println("hset写入的失败", err)
		}
		_, err = conn.Do("hset", "user", "age", "13")
		if err != nil {
			fmt.Println("hset写入的失败", err)
		}
		_, err = conn.Do("hmset", "user1", "name", "hsueh", "age", "12", "job", "golang")
		if err != nil {
			fmt.Println("hmset写入的失败", err)
		}
		// 读取数据
		r2, err := redis.String(conn.Do("hget", "user", "name"))
		if err != nil {
			fmt.Println("获取hash数据失败", err)
		}
		r3, err := redis.Strings(conn.Do("hmget", "user1", "name", "age"))
		if err != nil {
			fmt.Println("获取全部hash数据失败", err)
		}
		for i, s := range r3 {
			fmt.Printf("user1[%d]：%v \n", i, s)
		}
		r4, err := redis.StringMap(conn.Do("hgetall", "user1"))

		if err != nil {
			fmt.Println("获取全部hash数据失败", err)
		}
		fmt.Println("操作成功 获取数据", r1, r2, r3, r4)
	*/

}
