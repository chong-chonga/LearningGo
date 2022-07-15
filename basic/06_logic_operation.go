package basic

import (
	"fmt"
	"go/types"
)

func TestLogicOperators() {
	// go 的 switch 不需要在 case 后面加上 break,如果需要执行下一个 case 的,则加上 fallthrough
	var grade = 90
	var c = 'c'
	switch c {
	case 'c':
		fmt.Println("the char is c")
	default:
		fmt.Println("I don't know")
	}

	switch grade {
	case 90, 91, 92:
		fmt.Println("grade is no less than 90")
	default:
		fmt.Println("grade is less than 90")
	}

	var x interface{}
	switch x.(type) {
	case int:
		fmt.Println("the type is int")
	case string:
		fmt.Println("the type is string")
	case types.Nil:
		fmt.Printf("the type is %t\n", x)
	}

}
