package collections

import "fmt"

func TestPointer() {
	var iptr *int
	var i = 1
	iptr = &i
	fmt.Println("address of i is  ", iptr)
	fmt.Println("i = ", *iptr)

	var sptr *Student
	student := Student{
		sex: 1,
		age: 12,
	}
	sptr = &student
	fmt.Println("sptr.sex", &sptr.sex)
	fmt.Println("sptr.age = ", &sptr.age)
}

type Student struct {
	sex byte
	age int
}
