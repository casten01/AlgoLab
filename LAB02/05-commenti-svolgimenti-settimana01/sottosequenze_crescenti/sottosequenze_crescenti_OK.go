package main

import (
	"fmt"
	"math"
)

func main() {
	var curr, prev, succ int
	prev = math.MaxInt
	fmt.Scan(&curr)
	count := 0

	for {
		_, err := fmt.Scan(&succ)
		if err != nil {
			break
		}
		// Analizzo prev, curr, succ:
		if curr < prev && curr < succ {
			//inizio di salita
			count++
		}
		prev = curr
		curr = succ
	}
	fmt.Println(count)
}
