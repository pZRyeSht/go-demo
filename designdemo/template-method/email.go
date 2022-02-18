package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Email 具体实现类
type Email struct {
	notify Notify
}

func NewEmail(notify Notify) INotify {
	return &Email{
		notify: notify,
	}
}

func (e *Email) genRandomCode(i int) string {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(1000000)))
	fmt.Println("code:", code)
	return code
}

func (e *Email) sendVerifyCode(account, msg string, option ...interface{}) error {
	fmt.Println("Graceful Email sendVerifyCode")
	return nil
}
