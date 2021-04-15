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

// non pipeline
func EgNonPipeline(client *redis.Client) {
	ctx := context.TODO()
	now := time.Now()
	for i := 0; i < 10000; i++ {
		client.Set(ctx, strconv.Itoa(i), 1, 100*time.Second)
	}
	fmt.Println("NonPipeline 处理10000条数据，单条执行耗时：", time.Since(now))
}

func EgPipeline(client *redis.Client) {
	ctx := context.TODO()
	now := time.Now()
	pipeCli := client.Pipeline()
	for i := 0; i < 10000; i++ {
		pipeCli.Set(ctx, strconv.Itoa(i), 1, 100*time.Second)
	}
	_, err := pipeCli.Exec(ctx)
	if err != nil {
		fmt.Println("level=error, ", err)
		return
	}
	fmt.Println("Pipeline 处理10000条数据，单条执行耗时：", time.Since(now))
}

func main() {
	cli := Init()
	EgNonPipeline(cli)
	EgPipeline(cli)
}
