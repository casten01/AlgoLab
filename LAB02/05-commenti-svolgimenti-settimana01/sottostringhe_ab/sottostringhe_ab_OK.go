package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	countA := 0
	count := 0

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	for i, ch := range sc.Text() {
		if ch == 'a' {
			countA++
		}
		if ch == 'b' {
			count += countA
		}
	}
	fmt.Println(count)
}
