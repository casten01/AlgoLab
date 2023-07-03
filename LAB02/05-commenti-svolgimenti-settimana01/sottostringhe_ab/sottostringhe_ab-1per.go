// Bene, ma perché 1*cntA ????

package main

import "fmt"

func main() {
	var cnt, cntA int
	s := "cdbaagbabbgbab"

	for _, r := range s {
		if r == 'a' {
			cntA++
		}
		if r == 'b' {
			cnt = cnt + 1*cntA
		}
	}

	fmt.Println(cnt)

}
