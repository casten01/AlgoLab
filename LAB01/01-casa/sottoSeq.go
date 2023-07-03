// SOMME DELLE SOTTOSEQUENZE CRESCENTI: Scrivere un programma che legge una sequenza di interi e stampa la somma degli interi in ciascuna sottosequenza crescente. Il programma interrompe la lettura quando riceve un segnale di EOF. Ad esempio, su input 1 2 13 0 7 8 9 -1 0 2 il programma stampa le somme 16, 24 e 1.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInteri() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	sequenza := strings.Split(text, " ")
	var seqInt []int
	for i := 0; i < len(sequenza); i++ {
		tmp, _ := strconv.Atoi(sequenza[i])
		seqInt = append(seqInt, tmp)
	}

	return seqInt
}

func sottoSeqCresc(sequenza []int) []int {

	var seqCresc []int
	var j int
	seqCresc = append(seqCresc, sequenza[0])

	for i := 1; i < len(sequenza); i++ {
		if sequenza[i] > sequenza[i-1] {
			seqCresc[j] += sequenza[i] //ci interessa la somma degli elementi della sottosequenza
		} else {
			seqCresc = append(seqCresc, sequenza[i])
			j += 1
		}
	}

	return seqCresc
}

func main() {
	seq := readInteri()
	seqCresc := sottoSeqCresc(seq)
	fmt.Println(seqCresc)
}
