package basic

import "fmt"

func TestLoop() {
	// go 没有 while, for 循环可以替代 while
	count := 0

	for {
		if count > 1 {
			break
		}
		count++
	}
	fmt.Println("count = ", count)

	n := 100
	for n > 1 {
		n--
	}
	fmt.Println("n = ", n)

	dp := [100]int{1, 2, 3, 4, 5}
	for _, e := range dp {
		fmt.Printf("%d ", e)
	}
}
