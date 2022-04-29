package main

/*
 * @Description:
 * @Author: Liuhongq
 * @Date: 2022-04-30 02:23:05
 * @LastEditTime: 2022-04-30 03:29:01
 * @LastEditors: Liuhongq
 * @Reference:
 */

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("HMSET", "user1", "name", "wd2", "age", 11, "sex", "m")
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
