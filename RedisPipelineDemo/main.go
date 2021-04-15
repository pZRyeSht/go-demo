package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

// redis pipeline usage eg

func Init() *redis.Client {
	options := &redis.Options{
		Addr:     "127.0.0.1:6379",
		Username: "",
		Password: "",
		DB:       0,
	}
	return redis.NewClient(options)
}

func main() {
	cli := Init()
	// 1.non pipeline
	ctx := context.TODO()
	now := time.Now()
	for i := 0; i < 10000; i++ {
		cli.Set(ctx,strconv.Itoa(i), 1, 100 * time.Second)
	}
	fmt.Println("处理10000条数据，单条执行耗时：", time.Since(now))
	// 处理10000条数据，单条执行耗时： 653.327414ms

	// 2.pipeline
	//pipeCli := cli.Pipeline()
	//for i := 0; i < 10000; i++ {
	//	pipeCli.Set(ctx,strconv.Itoa(i), 1, 100 * time.Second)
	//}
	//_, err := pipeCli.Exec(ctx)
	//if err != nil {
	//	fmt.Println("level=error, ", err)
	//	return
	//}
	//fmt.Println("处理10000条数据，单条执行耗时：", time.Since(now))
	// 处理10000条数据，单条执行耗时： 30.777338ms
}
