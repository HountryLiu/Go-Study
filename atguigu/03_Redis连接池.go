package main

/*
 * @Description:
 * @Author: Liuhongq
 * @Date: 2022-04-30 02:23:05
 * @LastEditTime: 2022-04-30 03:30:51
 * @LastEditors: Liuhongq
 * @Reference:
 */

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {

	pool = &redis.Pool{
		MaxIdle:     8,   //最大闲置连接数
		MaxActive:   0,   //最大活跃连接数，0代表无限
		IdleTimeout: 100, //闲置连接的超时时间
		Dial: func() (redis.Conn, error) { //定义拨号获得连接的函数
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("HMSET", "user1", "name", "wd2", "age", 11, "sex", "m")
	if err != nil {
		fmt.Println("redis set error:", err)
	}
	r, err := redis.StringMap(conn.Do("HGETALL", "user1"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Println(r)
	}
}
