package main

import "fmt"

func main() {
	var counts, currentPos int32
	var steps int32
	var path string

	fmt.Printf("fill the steps: ")
	fmt.Scan(&steps)
	fmt.Printf("fill the path: ")
	fmt.Scan(&path)

	seaLevel := true

	for i := int32(0); i < steps; i++ {
		if path[i] == 'U' {
			currentPos++
		}
		if path[i] != 'U' {
			currentPos--
		}

		if seaLevel && currentPos == -1 {
			counts++
		}

		seaLevel = currentPos == 0
	}

	fmt.Printf("hasil= %v\n", counts)
}
