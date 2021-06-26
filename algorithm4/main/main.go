package main

import "fmt"

// %2 = 1 ganjil
// %2 = 0 genap
func main() {
	var a [101]int
	//j is the trip
	for j := 1; j <= 100; j++ {
		//k is the switch
		for k := 1; k <= 100; k++ {
			if k%j == 0 {
				a[k]++
			}
		}
	}

	for i := 1; i <= 100; i++ {
		if a[i]%2 == 1 {
			fmt.Printf("switch %v = nyala, dipencet %v kali\n", i, a[i])
		}
	}

}
