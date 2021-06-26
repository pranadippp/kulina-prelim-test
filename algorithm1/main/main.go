package main

import "fmt"

func main() {
	var n int

	fmt.Printf("fill number of sock: ")
	fmt.Scan(&n)
	fmt.Printf("fill the colors of each sock: ")

	// i is the color
	socks := make([]int, 101)
	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)

		socks[t]++
	}

	count := 0

	for _, v := range socks {
		count += v / 2
	}

	fmt.Printf("hasil= %v\n", count)
}
