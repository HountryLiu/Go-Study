package main

/*
 * @Description:
 * @Author: Liuhongq
 * @Date: 2018-08-24 12:09:22
 * @LastEditTime: 2022-04-30 04:41:28
 * @LastEditors: Liuhongq
 * @Reference:
 */

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {

	pool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲链接数
		MaxActive:   maxActive,   // 表示和数据库的最大链接数， 0 表示没有限制
		IdleTimeout: idleTimeout, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码， 链接哪个ip的redis
			return redis.Dial("tcp", address)
		},
	}
}
