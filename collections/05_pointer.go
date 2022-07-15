package collections

import "fmt"

func TestPointer() {
	var iptr *int

	// 指针未初始化时, 值为 nil
	if iptr == nil {
		fmt.Println("iptr == ", iptr)
	}

	var i = 1
	iptr = &i
	fmt.Println("address of i is  ", iptr)
	fmt.Println("i = ", *iptr)

	var sptr *Student
	student := Student{
		sex:  1,
		sex2: 2,
		age:  12,
	}
	sptr = &student
	// 0xc0000180c0
	fmt.Println("sptr.sex = ", &sptr.sex)
	// 0xc0000180c1
	fmt.Println("sptr.sex2 = ", &sptr.sex2)
	// 64 位机器中的 int 类型是 8个字节, 因此该类型变量的地址也应当是 8 的倍数, 故填充了 6 个缺少的字节
	// 0xc0000180c8
	fmt.Println("sptr.age = ", &sptr.age)

	// 将数组作为函数参数传递时, 如果需要修改数组中的值, 应当传递指针或者切片
	arr := [3]int{1, 2, 3}
	fmt.Printf("type of arr is %T\n", arr)
	fmt.Println("Before modifyArr is called, arr[0] = ", arr[0])
	modifyArr(arr[:])
	fmt.Println("After called modifyArr, arr[0] = ", arr[0])
	// 如此声明的数组是切片, 传递的是引用
	// 通过输出可知 arr1 是 []int 类型, 而 arr 是 [3]int 类型
	//两者根本不是同一个类型
	arr1 := []int{1, 2, 3}
	fmt.Printf("type of arr1 is %T\n", arr1)
	fmt.Println("Before modifyArr is called, arr1[0] = ", arr1[0])
	modifyArr(arr1)
	fmt.Println("After called modifyArr, arr1[0] = ", arr1[0])

	// 当然, go 支持双级指针
	var ipptr **int = &iptr
	fmt.Println("*ipptr = ", *ipptr)
	fmt.Println("**ipptr = ", **ipptr)
}

func modifyArr(arr []int) {
	arr[0] = 2
}

type Student struct {
	sex  byte
	sex2 byte
	age  int
}
