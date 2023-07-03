// 1.1 RAPPRESENTAZIONE di GRAFI -> SLICE DI LISTE DI ADIACENZA
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// nodo linkedList
type node struct {
	item int
	next *node
}

// testa linkedList
type linkedList struct {
	head *node
}

// grafo
type graph struct {
	n         int //numero di vertici
	adiacenti []linkedList
}

// Creo nuovo nodo nella lista
func newNode(val int) *node {
	return &node{val, nil}
}

// inserimento in testa nella lista di adiacenza (lista concatenata)
func insert(list *linkedList, val int) {
	node := newNode(val)
	node.next = list.head
	list.head = node
}

// Funzione che restituisce l’indirizzo di un nuovo grafo con n nodi.
func nuovoGrafo(n int) *graph {
	var graph graph
	graph.n = n
	graph.adiacenti = make([]linkedList, n)

	return &graph
}

//Funzione per leggere un grafo da standard-input, formato:
/*
n
v1 v2
v1 v3
v2 v...
*/
func readGrapgSI() *graph {
	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	var istruzioni []string
	for scanner.Scan() {
		istruzioni = append(istruzioni, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	n, _ := strconv.Atoi(istruzioni[0])

	graph := nuovoGrafo(n)

	for i := 1; i < len(istruzioni)/2-1; i++ {
		val, _ := strconv.Atoi(istruzioni[i+1])
		val2, _ := strconv.Atoi(istruzioni[i])

		insert(&graph.adiacenti[val2], val)
		i++
	}

	return graph
}

func printList(l *linkedList) {
	p := l.head
	for p != nil {
		fmt.Print(p.item)
		if p.next != nil {
			fmt.Print(" ")
		}
		p = p.next
	}
	fmt.Println()
}

func stampaGrafo(g *graph) {
	fmt.Printf("Il grafo ha %d nodi\n", g.n)
	for i, lista := range g.adiacenti {
		fmt.Print(i, ": ")
		printList(&lista)
	}
}

func main() {

	graph := readGrapgSI()
	stampaGrafo(graph)
}
