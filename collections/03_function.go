package collections

import "fmt"

// Go 支持全局变量
var globalVar = 12

func TestFunction() {
	// 函数声明是以 func 开头的
	// 对于调用多个返回值的函数时, 如果不需要对应的返回值, 使用 _ 空白标识符接收即可
	_, ret := Test1()
	fmt.Println("ret = ", ret)

	// defer 是延迟语句, 在函数返回前执行
	// defer 执行是逆序的
	// 先打印 defer 2, 后打印 defer 1
	i := 1
	// 参数值是当前的参数值, 不是执行时的参数值
	defer fmt.Println("defer ", i)
	i = 2
	defer fmt.Println("defer ", i)

}

func Test1() (string, int) {
	// go 的函数可以有多个返回值, 多个返回值时, return 的值用逗号隔开
	var retStr = "Hello"
	var retInt = 1
	return retStr, retInt
}
