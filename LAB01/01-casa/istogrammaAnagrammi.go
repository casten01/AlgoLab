//ISTOGRAMMA AD ANAGRAMMI
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//1) Scrivete un programma che legga una riga di caratteri terminata da a-capo, poi legga un’altra lettera e stampi quante volte la lettera compare nella prima riga.

func readLine() string {
	// To create dynamic array
	scanner := bufio.NewScanner(os.Stdin)
	// Scans a line from Stdin(Console)
	scanner.Scan()
	// Holds the string that scanned
	text := scanner.Text()

	return text
}

func countLetter(text, lettera string) int {
	var counter int

	for _, char := range text {
		if strings.ToUpper(string(char)) == lettera {
			counter += 1
		}
	}

	return counter
}

//2) Scrivete un programma che legga una riga di caratteri terminata da a-capo e che visualizzi un istogramma con una barra per ogni carattere dell’alfabeto, lunga quanto il numero delle sue occorrenze. Il programma non deve visualizzare le barre delle lettere che non compaiono e non deve fare distinzione fra maiuscole e minuscole.

func mappingLetters(text string) map[string]int {
	istogramma := make(map[string]int)
	var stringLetter string

	for _, v := range text {
		stringLetter = strings.ToUpper(string(v))
		if stringLetter != " " {
			istogramma[stringLetter] = countLetter(text, stringLetter)
		}
	}

	return istogramma
}

func printIsto(istogramma map[string]int) {

	//definire ordine di stampa
	var ord []string
	var letterina rune
	for i := 0; i < 25; i++ {
		letterina = rune(65 + i)
		ord = append(ord, string(letterina))
	}

	//stampo istogramma (manca ciclo asterischi sbatti)
	for _, v := range ord {
		if istogramma[v] != 0 {
			fmt.Printf("%s : %d \n", v, istogramma[v])
		}
	}
}

// 3) Due parole costituiscono un anagramma se l’una si ottiene dall’altra permutando le lettere(es: attore, teatro). Scrivete un programma che legga due parole e verifichi se sono anagrammi.

func areAnagrammi(word1, word2 string) bool {
	map1 := mappingLetters(word1)
	map2 := mappingLetters(word2)

	for k := range map1 {
		if map1[k] != map2[k] {
			return false
		}
	}
	return true
}

func main() {

	//text := readLine()

	//printIsto(mappingLetters(text))

	boolino := areAnagrammi("teatro", "attor")
	fmt.Println(boolino)

}
