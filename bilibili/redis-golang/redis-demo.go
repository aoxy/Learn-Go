package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	//连接Redis(在WSL2中)
	conn, err := redis.Dial("tcp", "172.22.131.206:6379")
	if err != nil {
		fmt.Println("redis.Dial error = ", err)
		return
	}
	if _, err := conn.Do("AUTH", "redis2021"); err != nil {
		conn.Close()
		return
	}
	fmt.Println("connect success...", conn)
	defer conn.Close()

	//操作
	_, err = conn.Do("Set", "name", "橘猫Sally")
	if err != nil {
		fmt.Println("set error = ", err)
		return
	}

	value, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get error = ", err)
		return
	}

	fmt.Println("get成功, value = ", value)
}
