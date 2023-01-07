package redis_nx

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestDistributeLockRedis(t *testing.T) {
	rd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	key := "test"
	// 创建可重入的分布式锁
	dl := NewDistributeLockRedis(rd, key, 10)
	// 争抢锁
	err := dl.TryLock()
	// if err != nil 没有抢到锁
	assert.Nil(t, err)
	// 抢到锁的记得释放锁
	defer func() {
		err := dl.Unlock()
		if err != nil {
			return
		}
	}()
	// 做真正的任务
	fmt.Println("I am running!!!")
}


func TestDistributeLockRedis_TryLock(t *testing.T) {
	type fields struct {
		key       string
		expire    int64
		status    bool
		cancelFun context.CancelFunc
		rd        *redis.Client
	}
	var tests []struct {
		name    string
		fields  fields
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dl := &DistributeLockRedis{
				key:       tt.fields.key,
				expire:    tt.fields.expire,
				status:    tt.fields.status,
				cancelFun: tt.fields.cancelFun,
				rd:        tt.fields.rd,
			}
			if err := dl.TryLock(); (err != nil) != tt.wantErr {
				t.Errorf("TryLock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDistributeLockRedis_Unlock(t *testing.T) {
	type fields struct {
		key       string
		expire    int64
		status    bool
		cancelFun context.CancelFunc
		rd        *redis.Client
	}
	var tests []struct {
		name    string
		fields  fields
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dl := &DistributeLockRedis{
				key:       tt.fields.key,
				expire:    tt.fields.expire,
				status:    tt.fields.status,
				cancelFun: tt.fields.cancelFun,
				rd:        tt.fields.rd,
			}
			if err := dl.Unlock(); (err != nil) != tt.wantErr {
				t.Errorf("Unlock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDistributeLockRedis_lock(t *testing.T) {
	type fields struct {
		key       string
		expire    int64
		status    bool
		cancelFun context.CancelFunc
		rd        *redis.Client
	}
	var tests []struct {
		name    string
		fields  fields
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dl := &DistributeLockRedis{
				key:       tt.fields.key,
				expire:    tt.fields.expire,
				status:    tt.fields.status,
				cancelFun: tt.fields.cancelFun,
				rd:        tt.fields.rd,
			}
			if err := dl.lock(); (err != nil) != tt.wantErr {
				t.Errorf("lock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDistributeLockRedis_startWatchDog(t *testing.T) {
	type fields struct {
		key       string
		expire    int64
		status    bool
		cancelFun context.CancelFunc
		rd        *redis.Client
	}
	type args struct {
		ctx context.Context
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestNewDistributeLockRedis(t *testing.T) {
	type args struct {
		rd     *redis.Client
		key    string
		expire int64
	}
	var tests []struct {
		name string
		args args
		want *DistributeLockRedis
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDistributeLockRedis(tt.args.rd, tt.args.key, tt.args.expire); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDistributeLockRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}
