package main

import "fmt"

// iota,特殊常量,可以认为是一个可以被编译器修改的常量.
// iota 在 const关键字出现时将被重置为 0 (const 内部的第一行之前),const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引).

// iota 仅作用于当前域

const a = iota // 0
const (
	b = iota // 0
)

// iota能够自动inc,inc的规则是作用域内每行+1.
// 所有注释行和空行全部忽略,无法使iota生效,iota不自增.
// 没有表达式的常量定义复用上一行的表达式
// 左移运算符 "<<" 是双目运算符,左移n位就是乘以2的n次方.其功能把"<<"左边的运算数的各二进位全部左移若干位,由"<<"右边的数指定移动的位数,高位丢弃,地位补0.
// 右移运算符 ">>" 是双目运算符,右移n位就是除以2的n次方.其功能把">>"左边的运算数的各二进位全部右移若干位,由">>"右边的数指定移动的位数.

const (
	c = 0       // iota = 0
	d = iota    // iota = 1
	e           // iota = 2
	f = "hello" // iota = 3
	// 这是一行注释
	g                 // iota = 4
	h    = iota       // iota = 5
	i                 // iota = 6
	j    = 0          // iota = 7
	k                 // iota = 8
	l, m = iota, iota // iota = 9
	n, o              // iota = 10
	
	p = iota + 1                  // iota = 11
	q                             // iota = 12
	_                             // iota = 13
	r = iota * iota               // iota = 14
	s                             // iota = 15
	t = r                         // iota = 16
	u                             // iota = 17
	v = 1 << iota                 // iota = 18
	w                             // iota = 19
	x               = iota * 0.01 // iota = 20
	y float32       = iota * 0.01 // iota = 21
	z                             // iota = 22
)

func main() {
	fmt.Printf("a : %T = %v\n", a, a) // 0
	fmt.Printf("b : %T = %v\n", b, b) // 0
	fmt.Printf("c : %T = %v\n", c, c) // 0
	fmt.Printf("d : %T = %v\n", d, d) // 1
	fmt.Printf("e : %T = %v\n", e, e) // 2
	fmt.Printf("f : %T = %v\n", f, f) // "hello"
	fmt.Printf("g : %T = %v\n", g, g) // "hello"
	fmt.Printf("h : %T = %v\n", h, h) // 5
	fmt.Printf("i : %T = %v\n", i, i) // 6
	fmt.Printf("j : %T = %v\n", j, j) // 0
	fmt.Printf("k : %T = %v\n", k, k) // 0
	fmt.Printf("l : %T = %v\n", l, l) // 9
	fmt.Printf("m : %T = %v\n", m, m) // 9
	fmt.Printf("n : %T = %v\n", n, n) // 10
	fmt.Printf("o : %T = %v\n", o, o) // 10
	fmt.Printf("p : %T = %v\n", p, p) // 12
	fmt.Printf("q : %T = %v\n", q, q) // 13
	fmt.Printf("r : %T = %v\n", r, r) // 14 * 14 = 196
	fmt.Printf("s : %T = %v\n", s, s) // 15 * 15 = 225
	fmt.Printf("t : %T = %v\n", t, t) // 14 * 14 = 196
	fmt.Printf("u : %T = %v\n", u, u) // 14 * 14 = 196
	fmt.Printf("v : %T = %v\n", v, v) // 1 * 2^18 = 262144
	fmt.Printf("w : %T = %v\n", w, w) // 1 * 2^19 = 524288
	fmt.Printf("x : %T = %v\n", x, x) // 0.2
	fmt.Printf("y : %T = %v\n", y, y) // 0.21
	fmt.Printf("z : %T = %v\n", z, z) // 0.22
}

// 参考文章：https://blog.wolfogre.com/posts/golang-iota/