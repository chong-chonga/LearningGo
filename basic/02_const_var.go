package basic

import "fmt"

// 常量没有被使用是不会报错的
const (
	OK           = 200
	BAD_ARG      = 400
	SERVER_ERROR = 500
)

func TestConstVar() {
	var code = 200
	if code == OK {
		fmt.Println("the status is OK!")
	}

}
