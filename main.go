package main

import (
	"fmt"
	"github.com/LeetcodeGO/operateGraph/directedGraph"
)

func main() {

	fmt.Println("hello world")
	red := [][]int{
		[]int{0, 1},
		[]int{1, 2},
		[]int{2, 3},
		[]int{3, 4},
	}
	blue := [][]int{
		[]int{1, 2},
		[]int{2, 3},
		[]int{3, 1},
	}
	answer := directedGraph.ShortestAlternatingPaths(5, red, blue)
	fmt.Println("answer: ", answer)
}
