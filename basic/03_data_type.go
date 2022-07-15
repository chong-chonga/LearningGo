package basic

import "fmt"

func TestDataType() {

	// 32位和64位实数和虚数
	var v_num1 complex64 = 1
	fmt.Println("v_num1 = ", v_num1)
	var v_num2 complex128 = 2
	fmt.Println("v_num2 = ", v_num2)

	// 类型转换,相较于C/C++更加严格,只有类型兼容的可以转换
	// Go 允许向下缩小转换,也就是允许精度损失
	var n1 int32 = 12
	var n2 = int64(n1)

	var n3 int64 = 100
	var n4 = int32(n3)
	fmt.Println("n3 = ", n3, ", n4 = ", n4)
	fmt.Println("n1 = ", n1, ", n2 = ", n2)

}
