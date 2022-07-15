package basic

import "fmt"

func TestArray() {
	// 可以通过 len 函数来获取数组的长度
	arr1 := [10]int{1, 2, 3}
	arr2 := []int{1, 2, 3}
	len1 := len(arr1)
	len2 := len(arr2)
	fmt.Println("len1 = ", len1)
	fmt.Println("len2 = ", len2)

	arr3 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for _, ints := range arr3 {
		for _, e := range ints {
			fmt.Printf("%d ", e)
		}
		fmt.Println()
	}

	// 值传递
	arr4 := [3]int{1, 2, 3}
	arr5 := arr4
	arr5[0] = 2
	fmt.Println(arr4)
	fmt.Println(arr5)

	var arr6 = [3]int{1, 2, 3}
	arr7 := arr6
	arr7[0] = 2
	fmt.Println(arr6)
	fmt.Println(arr7)

	// 引用传递
	var arr8 = []int{1, 2, 3}
	arr9 := arr8
	arr9[0] = 2
	fmt.Println(arr8)
	fmt.Println(arr9)

}
