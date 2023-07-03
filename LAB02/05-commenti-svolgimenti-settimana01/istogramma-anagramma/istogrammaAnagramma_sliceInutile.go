/* la deepEqual è utile quando si hanno strutture ricorsive, articolate.
In questo caso meglio fare un confronto elemento per elemento.

Inoltre non c'è nessun bisogno di costruire un nuovo vettore lettere2 e di riordinarlo poi con sort.
Anche se la complessità di queste operazioni è costante (dato che la dimensione della mappa è fissa
tanti elementi quante le lettere dell'alfabeto, 26), sono del tutto inutili. */

// Invece di:
  for key, _ := range lettere {
		lettere2 = append(lettere2, key)
	}
  sort.Strings(lettere2)
  for _,lett := range lettere2 {
    fmt.Print(lett, " ")
    fmt.Println(strings.Repeat("*", lettere[lett]))
  }

// si poteva semplicemente scrivere
  for lett := 'a'; lett < 'z'; lett++ {
    if (lettere[string(lett)] != 0 ) {
       fmt.Println(string(lett), strings.Repeat("*", lettere[string(lett)]))
    }
  }

// Infine, sembra più comoda una mappa che ha per chiavi rune, anziché lettere



package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	"unicode"
)

func quali(riga string) map[string]int {
	lettere := make(map[string]int)
	for _, char := range riga {
		if unicode.IsLetter(char) {
			lettere[string(char)]++
		}
	}
	return lettere
}

func anagrammi(s1, s2 string) bool {
	return reflect.DeepEqual(quali(s1), quali(s2))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	riga := strings.ToLower(scanner.Text())
	scanner.Scan()
	riga2 := strings.ToLower(scanner.Text())
	lettere := quali(riga)
	fmt.Println(anagrammi(riga, riga2))
	lettere2 := []string{}
	for key, _ := range lettere {
		lettere2 = append(lettere2, key)
	}
	sort.Strings(lettere2)
	for _, lett := range lettere2 {
		fmt.Print(lett, " ")
		fmt.Println(strings.Repeat("*", lettere[lett]))
	}
}
