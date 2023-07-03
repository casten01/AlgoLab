package main

import "fmt"

type listNode struct {
	item int
	next *listNode
}

type linkedList struct {
	head *listNode
}

func newNode(val int) *listNode {
	return &listNode{val, nil}
}

func addNewNode(l *linkedList, val int) {
	node := newNode(val)
	node.item = val
	node.next = l.head
	l.head = node
}

func main() {
	var l *linkedList

	addNewNode(l, 3)
	printList(l)
}

func printList(l *linkedList) {
	fmt.Println("stampo lista")
	p := l.head
	for p != nil {
		fmt.Print(p.item, " ")
		p = p.next
	}
	fmt.Println()
}

/* CORREZIONE:
In riga 28 è stata dichiarata una variabile puntatore a nestedList, ma non è stato allocato lo spazio per la struttura nestedList.
In particolare questo vuol dire che l ha valore “nil" e quando si cerca di prendere il suo campo “head” nella funzione addNewNode (riga 23)
si ha un panic error (di fatto stiamo cercando di prendere il campo “head” di… nessuna struttura: il puntatore non punta a nulla!)

Per risolvere, bisogna allocare una struttura nestedList subito dopo la dichiarazione, ad esempio così:
l = &linkedList{nil}
(ovvero allocando lo spazio per una linkedList con campo head a nil)

Un’alternativa (forse più diretta) sarebbe stato dichiarare l come nestedList:
	var l linkedList
	l.head = nil
e poi passarne l’indirizzo alla funzione addNewNode:
	addNewNode(&l, val)
*/
