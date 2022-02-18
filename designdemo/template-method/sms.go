package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Sms 具体实现类
type Sms struct {
	notify Notify
}

func NewSms(notify Notify) INotify {
	return &Sms{
		notify: notify,
	}
}

func (s *Sms) genRandomCode(i int) string {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(1000000)))
	fmt.Println("code:", code)
	return code
}

func (s *Sms) sendVerifyCode(account, msg string, option ...interface{}) error {
	fmt.Println("Graceful Sms sendVerifyCode")
	return nil
}
