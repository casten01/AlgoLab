//ORDINAMENTO PER INSERIMENTO: Scrivete una funzione che legga da standard input una sequenza di interi distinti terminati da 0, memorizzandoli in un vettore ordinato (valutate se è più opportuno usare un array o una slice): ogni volta che viene letto un nuovo intero, il vettore viene scorso fino a trovare l’esatta collocazione del numero, quindi si crea lo spazio per il nuovo numero spostando in avanti i numeri successivi già memorizzati. Questo algoritmo è utile per riempire un vettore mantenendolo ordinato ad ogni passo.

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

//scambia due elementi di indice i e j e restituisce la stessa slice in input
func switchAB(seq []int, i, j int) []int {
	var tmp int = seq[j]
	seq[j] = seq[i]
	seq[i] = tmp

	return seq
}

//ordina con insertionSort la slice passata in input
func ordina(seq []int) []int {
	seq = readInteri()

	for i := 1; i < len(seq); i++ {
		for j := i - 1; j >= 0; j-- {
			if seq[i] < seq[j] {
				seq = switchAB(seq, i, j)
				i = j
			}
		}
	}
	return seq
}

func main() {
	var seq []int
	fmt.Println(ordina(seq))
}
