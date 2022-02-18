package main

import "fmt"

// INotify 抽象类封装不变的部分，扩展可变部分
type INotify interface {
	genRandomCode(int) string                                        // 生成验证码
	sendVerifyCode(account, msg string, option ...interface{}) error // 发送验证码
}

// Notify 抽象类
type Notify struct {
	ify INotify
}

func NewNotify(notify INotify) Notify {
	return Notify{
		ify: notify,
	}
}

func (n *Notify) genRandomCode(int) string {
	return ""
}

func (n *Notify) sendVerifyCode(account, msg string, option ...interface{}) error {
	return nil
}

// 公共复用默认实现，由抽象类实现
func (n *Notify) print(msg string) error {
	fmt.Println(msg)
	return nil
}
