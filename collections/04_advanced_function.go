package collections

import "fmt"

func TestAdvancedFunction() {
	// Go 中的函数也是一种数据类型, 函数名称相当于函数代码的地址
	func1(1)
	// 打印函数名称是打印函数代码的地址
	fmt.Println("func1 = ", func1)

	// 函数可以赋给变量,可以通过变量调用函数
	f1 := func1
	fmt.Println("f1 = ", f1)
	f1(1)

	// 函数式编程: 匿名函数
	// 与 Java 不同的是, 这个匿名函数可以赋给变量, 从而多次使用
	f2 := func(arg1 int) int {
		return arg1
	}
	f2(1)

	// 另外, 函数可以作为一个函数参数, 也就是回调函数
	func2(f2)

	// 闭包函数中的内部函数使用到的局部变量不会随着函数执行完成而销毁
	inc := closure()
	// cnt = 1
	fmt.Println("cnt = ", inc())
	// cnt = 2
	fmt.Println("cnt = ", inc())
}

func func1(arg1 int) int {
	return arg1
}

func func2(f func(int) int) {
	f(1)
}

// 闭包函数, 以函数作为返回值, 且返回的函数使用了该函数的局部变量
func closure() func() int {
	// 内部函数使用了局部变量 cnt
	// 局部变量 cnt 不会随着函数返回而销毁, 会继续使用
	cnt := 0
	return func() int {
		cnt++
		return cnt
	}
}
