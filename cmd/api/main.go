package main

import "fmt"

type Numbers interface {
	float32 | float64 | int
}

func main() {
	fmt.Println(findMaxGeneric([]int{1, 2, 3, 4}))
	fmt.Println(findMaxGeneric([]float64{1.0, 2.0, 3.0, 4.5}))
}

func findMaxGeneric[Num Numbers](nums []Num) Num {
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	return maxNum
}
