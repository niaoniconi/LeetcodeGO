package main

import (
	"fmt"
	"github.com/LeetcodeGO/arithmetic/aboutSort/findKs"
)

func main() {

	test := []int{2, 1}
	findKs.QuickSort(test)
	fmt.Println("answer: ", test)
}
