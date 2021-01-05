package main

import (
	"fmt"
	"golang.org/x/xerrors"
)

func main() {
	err := foo2()
	fmt.Printf("%v\n", err)
	fmt.Printf("%+v\n", err)

	Is()
	As()
}

// xerrors 最佳实践 Error Wrapping
// 嵌套error，往上层传，打印error时带上堆栈信息
// must least go 1.13
// 注意，xerrors.Errorf(“foo1 : %w”,foo())中必须以: %w的格式占位
var myerror = xerrors.New("myerror")

func foo() error {
	return myerror
}

func foo1() error {
	return xerrors.Errorf("foo1 : %w", foo())
}

func foo2() error {
	return xerrors.Errorf("foo2 : %w", foo1())
}

// 总是返回自定义错误类型以携带与错误上下文有关的信息，该类型必须实现 xerrors.Wrapper 接口
// 使用 errors.Is 来判断错误，使用 errors.As 来获取错误上下文
// 等值测试与type断言的类型测试
// xerrors的Is方法原型:func Is(err, target error) bool
// Is会将target与err中的error chain上的每个error 信息进行等值比较，如果相同，则返回true

func Is() {
	fmt.Println("------------Is Start------------")
	err1 := xerrors.New("1")
	err2 := xerrors.Errorf("wrap 2: %w", err1)
	err3 := xerrors.Errorf("wrap 3: %w", err2)

	erra := xerrors.New("a")

	b := xerrors.Is(err3, err1)
	fmt.Println("err3 is err1? -> ", b)

	b = xerrors.Is(err2, err1)
	fmt.Println("err2 is err1? -> ", b)

	b = xerrors.Is(err3, err2)
	fmt.Println("err3 is err2? -> ", b)

	b = xerrors.Is(erra, err1)
	fmt.Println("erra is err1? -> ", b)
	fmt.Println("------------Is End------------")
}

// As会将err中的error chain上的每个error type与target的类型做匹配，如果相同，则返回true，
// 并且将匹配的那个error var的地址赋值给target，相当于通过As的target将error chain中类型匹配的那个error变量析出
// xerrors的As方法原型:func As(err error, target interface{}) bool
type MyError struct{}

func (MyError) Error() string {
	return "MyError"
}

func As() {
	fmt.Println("------------As Start------------")
	err1 := MyError{}
	err2 := xerrors.Errorf("wrap 2: %w", err1)
	err3 := xerrors.Errorf("wrap 3: %w", err2)
	var err MyError

	b := xerrors.As(err3, &err)
	fmt.Println("err3 as MyError? -> ", b)
	fmt.Println("err is err1? -> ", xerrors.Is(err, err1))

	err4 := xerrors.Opaque(err3)
	b = xerrors.As(err4, &err)
	fmt.Println("err4 as MyError? -> ", b)
	b = xerrors.Is(err4, err3)
	fmt.Println("err4 is err3? -> ", b)
	fmt.Println("------------As End------------")
}
