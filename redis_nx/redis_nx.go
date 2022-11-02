package redis_nx

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// DistributeLockRedis 基于redis的分布式可重入锁，自动续租
type DistributeLockRedis struct {
	key       string             // 锁的key
	expire    int64              // 锁超时时间
	status    bool               // 上锁成功标识
	cancelFun context.CancelFunc // 用于取消自动续租携程
	rd        *redis.Client       // redis句柄
}

// NewDistributeLockRedis 创建可重入锁
func NewDistributeLockRedis(rd *redis.Client, key string, expire int64) *DistributeLockRedis {
	return &DistributeLockRedis{
		rd:     rd,
		key:    key,
		expire: expire,
	}
}

// TryLock 上锁
func (dl *DistributeLockRedis) TryLock() (err error) {
	if err = dl.lock(); err != nil {
		return err
	}
	ctx, cancelFun := context.WithCancel(context.Background())
	dl.cancelFun = cancelFun
	dl.startWatchDog(ctx) // 创建守护协程，自动对锁进行续期
	dl.status = true
	return nil
}

// lock 竞争锁
func (dl *DistributeLockRedis) lock() error {
	if _, err := dl.rd.SetNX(context.Background(), dl.key, true, time.Duration(dl.expire)*time.Second).Result(); err != nil {
		return err
	}
	return nil
}

// startWatchDog 创建守护协程，自动续期
func (dl *DistributeLockRedis) startWatchDog(ctx context.Context) {
	go func() {
		for {
			select {
			// Unlock通知结束
			case <-ctx.Done():
				return
			default:
				// 否则只要开始了，就自动重入（续租锁）
				if dl.status {
					if _, err := dl.rd.ExpireNX(context.Background(), dl.key, time.Duration(dl.expire)*time.Second).Result(); err != nil {
						return
					}
					// 续租时间为 expire/2 秒
					time.Sleep(time.Duration(dl.expire/2) * time.Second)
				}
			}
		}
	}()
}

// Unlock 释放锁
func (dl *DistributeLockRedis) Unlock() (err error) {
	// 这个重入锁必须取消，放在第一个地方执行
	if dl.cancelFun != nil {
		dl.cancelFun() // 释放成功，取消重入锁
	}
	var res int64
	if dl.status {
		if res, err = dl.rd.Del(context.Background(), dl.key).Result(); err != nil {
			return fmt.Errorf("释放锁失败")
		}
		if res == 1 {
			dl.status = false
			return nil
		}
	}
	return fmt.Errorf("释放锁失败")
}

