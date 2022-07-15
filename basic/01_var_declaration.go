package basic

import (
	"fmt"
)

// 全局变量不能使用 :=
// c := 14

// 自动类型推断
var d = 13

func TestVarDeclaration() {
	// 先声明后赋值
	var a int
	a = 12

	// 直接赋值
	b := 15

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("d = ", d)
	var addrA = &a
	var addrB = &b
	var addrD = &d
	// 0xc0000180a8
	fmt.Println("addrA = ", addrA)
	// 0xc0000180c0
	fmt.Println("addrB = ", addrB)
	// addrB - addrA = 8, 正好是 int 的数据大小
	fmt.Println("addrD = ", addrD)

	//多重赋值
	var i, j, k = 0, 1, 2
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
	fmt.Println("k = ", k)
}
