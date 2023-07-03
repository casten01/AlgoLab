package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var count int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	var lettera rune
	fmt.Scanf("%c", &lettera)

	for i := 0; i < len(s); i++ { // meglio usare range
		if rune(s[i]) == lettera {
			count++
		}
	}
	fmt.Println("Il carattere", string(lettera), "compare", count, "volte")
}
