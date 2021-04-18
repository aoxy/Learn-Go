package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func main() {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("Set", "name2", "蓝猫Tom")
	if err != nil {
		fmt.Println("set error = ", err)
		return
	}

	value, err := redis.String(conn.Do("Get", "name2"))
	if err != nil {
		fmt.Println("get error = ", err)
		return
	}
	fmt.Println("get成功, value = ", value)
}

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "172.22.131.206:6379")
			conn.Do("AUTH", "redis2021")
			return conn, err
		},
	}
}
