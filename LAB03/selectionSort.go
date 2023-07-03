//ORDINAMENTO PER SELEZIONE: Scrivete una funzione che riceva una slice di interi e la ordini usando l’algoritmo SelectionSort: alla fine dell’esecuzione, la slice originaria passata come argomento dovrà risultare ordinata. Può essere utile restituire la stessa slice (ordinata), ad esempio per poterla passare come argomento ad altre funzioni, come in fmt.Println(selectionSort(v)).

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

//VERSIONE ITERATIVA: La funzione selectionSortIter( int a[]) ripete la seguente operazioni tante volte quanto è lunga la slice da ordinare: per ogni prefisso di lunghezza n (con n inizialmente pari alla lunghezza della slice) cerca nel prefisso l’elemento massimo e lo scambia con quello nell’ultima posizione del prefisso.

func selectionSortIter(seq []int) []int {
	for i := 0; i < len(seq); i++ {
		n := len(seq) - i
		var max int = 0
		for j := 0; j < n; j++ {
			if seq[j] > seq[max] {
				max = j
			}
		}
		seq = switchAB(seq, max, n-1)
	}

	return seq
}

//VERSIONE RICORSIVA: la funzione selectionSortRec deve funzionare come segue: innanzitutto cerca nel vettore l’elemento massimo e lo sposta nell’ultima posizione; poi richiama se stessa ricorsivamente per ordinare i primi n-1 elementi del vettore. La base della ricorsione è data dai vettori di lunghezza 0 o 1, che sono sempre ordinati.
func selectionSortRec(seq []int, n int) []int {
	if n == 0 || n == 1 {
		return seq
	} else {
		var max int
		for i := 0; i < n; i++ {
			if seq[i] > seq[max] {
				max = i
			}
			seq = switchAB(seq, max, n-1)
		}
		selectionSortRec(seq, n-1)
	}
	return seq
}

func main() {
	seq := readInteri()
	fmt.Println(selectionSortIter(seq))
	fmt.Println(selectionSortRec(seq, len(seq)))
}
