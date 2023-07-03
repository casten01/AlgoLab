//LISTA CON PUNTATORE A TAIL: Considerate ora una lista concatenata di interi, dotata di due riferimenti al primo e all’ultimo elemento della lista, come descritto dal tipo linkedListwithTail nella porzione di codice qui sotto. Quando la lista è vuota, sia head che tail sono nil. La funzione newNode alloca lo spazio per un nuovo nodo e ne inizializza il valore con l’argomento passato.

package main

import "fmt"

type listNode struct {
	item int
	next *listNode
}

type linkedListwithTail struct {
	head, tail *listNode
}

func newNode(val int) *listNode {
	return &listNode{val, nil}
}

//inserimento alla fine
func addNewNodeAtEnd(l *linkedListwithTail, val int) {
	if l.tail == nil {
		l.tail = newNode(val)
		l.head = l.tail
	} else {
		l.tail.next = newNode(val)
		l.tail = l.tail.next
	}
}

//confronto tra lista semplice e lista con tail
func main() {
	var list linkedListwithTail
	addNewNodeAtEnd(&list, 5)
	addNewNodeAtEnd(&list, 8)
	addNewNodeAtEnd(&list, 3)

	curr := list.head
	for curr != nil {
		fmt.Println("[]->", curr.item)
		curr = curr.next
	}

	fmt.Printf("testa-> %d, coda-> %d \n", list.head.item, list.tail.item)
}

//confronto con lista senza riferimento alla fine, per quali operazioni si ha un tempo migliore:
/*
A) restituisci 1 se la lista contiene un dato elemento
-> stesso costo
B) cancella l’ultimo elemento della lista
-> con la lista normale O(n) con la lista con coda O(1)
C) aggiungi un elemento all’inizio della lista
-> stesso costo
D) aggiungi un elemento alla fine della lista
-> con la lista normale O(n) con la lista con coda O(1)
*/
