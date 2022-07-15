package basic

import "fmt"

func TestDynamicArray() {
	n := 10
	// 第一个参数指定了返回类型, 第二个参数指定了数组的长度
	// 这个动态数组是一个切片, 底层实现是数组
	arr := make([]int, n)
	size := len(arr)
	fmt.Println("current array length is ", size)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	PrintSliceStatus(arr)

	// 切片也可以从数组中产生, 对切片的修改, 修改的是底层的数组
	// 左开右闭
	var tmpArr = [5]int{1, 2, 3, 4, 5}
	var slice = tmpArr[2:5]
	fmt.Println(tmpArr)
	for i := range slice {
		slice[i]++
	}
	fmt.Println(tmpArr)

	var src []int
	src = append(src, 0, 1, 2)
	dst := make([]int, 3, 3)
	copy(dst, src)
	fmt.Println("src = ", src)
	fmt.Println("dst = ", dst)
}

func PrintSliceStatus(slice []int) {
	fmt.Println("len: ", len(slice), " capacity: ", cap(slice))
}
