package main

import (
	"fmt"
	"math"
)

func main() {
	inputs := []int{1, 3, 4, 5, 6, 7, 9}
	for idx, input := range inputs {
		fmt.Printf("%d\n", input*int(math.Pow(10, float64(len(inputs)-idx-1))))
	}
}
