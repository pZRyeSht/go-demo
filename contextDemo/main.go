package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Req struct {
	ID int `json:"id"`
}

type Resp struct {
	Name string `json:"name"`
}

// 并发请求
func main() {
	var req []*Req
	for i := 0; i < 5; i++ {
		temp := Req{ID: i}
		req = append(req, &temp)
	}
	resp, err := GetResp(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

	respOne, err := GetRespOne(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(respOne)
}

// context + channel
func GetResp(req []*Req) (resp []*Resp, err error) {
	// 获取一个10second超时的ctx
	_ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// 协程执行完成管理信号
	msgCh := make(chan struct{})
	// 等待队列
	waitQueue := len(req)
	// 初始化res
	res := make([]*Resp, len(req))

	// 遍历req，发起并发请求
	for k, v := range req {
		go func(ctx context.Context, index int, re *Req) {
			// 延迟执行
			defer func() {
				msgCh <- struct{}{}
			}()

			// ctx超时则退出
			select {
			case <-ctx.Done():
				return
			default:
			}

			// 获取resp
			resp, err := getSomeThing(v.ID)
			if err != nil {
				return
			}
			// 构建结果集
			res[index] = &Resp{Name: resp}
		}(_ctx, k, v)
	}

	for {
		select {
		// 如果已经到达deadline，返回一个被关闭的channel，则使用cancel递归调用所有children取消
		case <-_ctx.Done():
			cancel()
			return nil, errors.New("time out")
		// msgCh接收协程执行返回信号量，当等待队列waitQueue为0时，说明已经执行完毕，直接返回结果集
		case <-msgCh:
			waitQueue--
			if waitQueue == 0 {
				return res, nil
			}
		}
	}
}

// get resp
func getSomeThing(id int) (string, error) {
	return fmt.Sprintf("Return::%s", strconv.Itoa(id)), nil
}

// 最大单次请求数
const MaxQueryCount = 10

// waitGroup + context + channel
// 并发请求第三方接口
func GetRespOne(ctx context.Context, req []*Req) (resp []*Resp, err error) {
	if len(req) == 0 {
		return
	}

	// 初始化等待组，返回channel
	var (
		wg  sync.WaitGroup
		ch  = make(chan []*Resp, MaxQueryCount)
		res = make([]*Resp, 0, len(req))
	)

	// 如果请求大于单次最大请求，则并发请求
	if len(req) > MaxQueryCount {
		temp := len(req) / MaxQueryCount
		for i := 0; i < temp; i++ {
			wg.Add(1)
			go getSomeThingOne(ctx, ch, &wg, req[MaxQueryCount*i:MaxQueryCount])
		}
		// 余下不足MaxQueryCount条，继续请求
		if len(req[MaxQueryCount*temp:]) > 0 {
			wg.Add(1)
			go getSomeThingOne(ctx, ch, &wg, req[MaxQueryCount*temp:])
		}
	} else {
		wg.Add(1)
		go getSomeThingOne(ctx, ch, &wg, req)
	}

	// 多对一由生产者负责关闭通道
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 处理结果集
	for {
		select {
		case resp, ok := <-ch:
			if !ok {
				return nil, nil
			}
			res = append(res, resp...)
			// 设定超时时间
		case <-time.After(10 * time.Second):
			return nil, errors.New("time out")
		// 防止root context cancel
		case <-ctx.Done():
			return nil, errors.New("context done")
		}
	}
}

// 获取信息，最大支持单次请求MaxQueryCount条
func getSomeThingOne(ctx context.Context, ch chan<- []*Resp, wg *sync.WaitGroup, req []*Req) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		return
	default:
	}

	// get result
	ch <- getSomeThingTwe(req)

	return
}

func getSomeThingTwe(req []*Req) (resp []*Resp) {
	// ......
	return
}
