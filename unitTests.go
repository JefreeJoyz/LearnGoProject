package main

import "fmt"

func testUnitTest() {
	fmt.Println(Max([]int{3, 5, 19, -20, 44, 30}))
}

func Max(numbers []int) int {
	var max int

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}
