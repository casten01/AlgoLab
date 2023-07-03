//IMPLEMENTAZIONE DI LISTE CONCATENATE SEMPLICI: scrivete un programma con l’implementazione di una lista concatenata semplice, seguendo i lucidi presentati a lezione. Definite i tipi listNode e linkedList e scrivete queste funzioni:
/*
• newNode, che crea un nuovo nodo di lista;
• addNewNode, che inserisce un nuovo nodo in testa alla lista;
• printList, che stampa una lista;
• searchList, che cerca un elemento in una lista;
• removeItem, che cancella un item da una lista
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

//nodo della lista
type listNode struct {
	item int
	next *listNode
}

//testa della coda
type linkedList struct {
	head *listNode
}

//crea un nuovo nodo per la lista
func newNode(val int) *listNode {
	return &listNode{val, nil}
}

//aggiunge un nuovo nodo in testa alla lista
func addNewNode(l *linkedList, val int) {
	node := newNode(val)
	node.next = l.head
	l.head = node
}

//stampa lista
func printList(l linkedList) {
	p := l.head
	for p != nil {
		fmt.Printf("[%d] ", p.item)
		p = p.next
	}
	fmt.Println()
}

//stampa inverso
func printReverse(l linkedList) {
	p := l.head
	len := countElements(l)
	var listaArr []listNode

	for p != nil {
		listaArr = append(listaArr, *p)
		p = p.next
	}

	for i := len - 1; i >= 0; i-- {
		fmt.Printf("[%d] ", listaArr[i].item)
	}
	fmt.Println()
}

//conta elementi lista
func countElements(l linkedList) int {
	var counter int
	p := l.head
	for p != nil {
		counter++
		p = p.next
	}
	return counter
}

//cerca un elemento in una lista
func searchList(l linkedList, val int) bool {
	p := l.head
	for p != nil {
		if p.item == val {
			return true
		}
		p = p.next
	}
	return false
}

//cancella un elemento dalla lista
func removeItem(l linkedList, val int) {
	var curr, prev *listNode = l.head, nil
	for curr != nil {
		if curr.item == val {
			if prev == nil {
				l.head = curr.next
			} else {
				prev.next = curr.next
			}
		}

		prev = curr
		curr = curr.next
	}
}

//cancella tutti gli elementi della lista
func removeAll(l linkedList) {
	curr := l.head
	for curr != nil {
		fmt.Println("eliminato valore: ", curr.item)
		removeItem(l, curr.item)
		curr = curr.next
	}
}

//legge istruzioni da file
func readFromFile() []string {
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

	return istruzioni
}

//Per testare funzioni usiamo una tabella di istruzioni che ci permettono di lavorare con la coda:
/*
+ n -> Se n non appartiene all’insieme lo inserisce, altrimenti non fa nulla
- n -> Se n appartiene all’insieme lo elimina, altrimenti non fa nulla
? n -> Stampa un messaggio che dichiara se n appartiene all’insieme
c   -> Stampa il numero di elementi dell’insieme
p   -> Stampa gli elementi dell’insieme (nell’ordine in cui compaiono nella lista)
o   -> Stampa gli elementi dell’insieme nell’ordine inverso
d   -> Cancella tutti gli elementi dell’insieme
f   -> Termina l’esecuzione
*/
func main() {
	p := newNode(9)
	list := new(linkedList)
	list.head = p
	istruzioni := readFromFile()

	for i, istr := range istruzioni {
		switch istr {
		case "+":
			val, _ := strconv.Atoi(istruzioni[i+1])
			if searchList(*list, val) {
				fmt.Println("Elemento già inserito")
			} else {
				addNewNode(list, val)
			}
		case "-":
			val, _ := strconv.Atoi(istruzioni[i+1])
			if searchList(*list, val) {
				removeItem(*list, val)
			} else {
				fmt.Println("Elemento da eliminare non presente")
			}
		case "?":
			val, _ := strconv.Atoi(istruzioni[i+1])
			if searchList(*list, val) {
				fmt.Println("Elemento appartiene all'insieme")
			} else {
				fmt.Println("Elemento non appartiene all'insieme")
			}
		case "c":
			fmt.Printf("Numero di elementi: %d \n", countElements(*list))
		case "p":
			printList(*list)
		case "o":
			printReverse(*list)
		case "d":
			removeAll(*list)
		case "f":
			break
		default:
			continue
		}

	}
}
