package basic

import (
	"bufio"
	"fmt"
	"os"
)

func TestStdOut() {
	a := 100           //int
	b := 3.14          //float64
	c := true          // bool
	d := "Hello World" //string
	f := 'A'           // char
	// %T 是打印变量的类型名称
	// 二进制形式输出 1100100
	fmt.Printf("%T,%b\n", a, a)
	fmt.Printf("%T,%f\n", b, b)
	fmt.Printf("%T,%t\n", c, c)
	fmt.Printf("%T,%s\n", d, d)
	fmt.Printf("%T,%d,%c\n", f, f, f)
	fmt.Println("-----------------------")
	// %v 以原样输出
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", f)
}

func TestStdIn() {
	var str string
	var n int
	fmt.Printf("请输入字符串: ")
	_, err := fmt.Scanln(&str)
	if err != nil {
		return
	}
	fmt.Printf("你输入的字符串是 %s\n", str)
	fmt.Printf("请输入整数: ")
	_, err = fmt.Scanln(&n)
	if err != nil {
		return
	}
	fmt.Printf("你输入的数字是 %d\n", n)
}

func TestBufIO() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please input a string: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("input = ", input)
}
