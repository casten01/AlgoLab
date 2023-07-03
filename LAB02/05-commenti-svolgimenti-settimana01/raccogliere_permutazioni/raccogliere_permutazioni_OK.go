// questa soluzione ha complessità lineare (una sola scansione)

package main

import (
	"fmt"
)

func raccogli(v []int) int {
	count := 0
	occ := make(map[int]bool)
	for _, el := range v {
		occ[el] = true
		if occ[el+1] {
			count++
		}
	}
	fmt.Printf("Servono %d inversioni\n", count)
	return count
}

func main() {
	//var v = []int{4, 5, 1, 3, 6, 2}
	var v = make([]int, 0)
	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		v = append(v, n)
	}
	fmt.Println(v)
	fmt.Println(raccogli(v))
}
