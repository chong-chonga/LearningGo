package collections

import (
	"fmt"
	"strings"
)

func TestString() {
	str := "Hello World!"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Println()

	// strings 是 string 的工具包
	// str1 == str2, ret = 0
	ret := strings.Compare("a", "a")
	fmt.Println("ret = ", ret)

	// str1 < str2, ret = -1
	ret = strings.Compare("a", "b")
	fmt.Println("ret = ", ret)

	// str1 > str2, ret = 1
	ret = strings.Compare("b", "a")
	fmt.Println("ret = ", ret)

	contains := strings.Contains("abc", "ab")
	if contains {
		fmt.Println("abc contains ab")
	}
}
