package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func push(l *linkedList, val int) {
	node := newNode(val)
	node.next = l.head
	l.head = node
}

func pop(l *linkedList) int {
	eliminato := l.head.item
	l.head = l.head.next
	return eliminato
}

func valuta(espressione string) int {
	var l linkedList
	
  ...
  push(&l, n)
	...  
  push(&l, pop(&l)+pop(&l))
	...

	return l.head.item
}


/* Note:

In questo programma, la pila è:

1. definita parzialmente come struttura di dati astratta:
  - sono definite delle funzioni "pop", "push", ...
  - ma non è definito un tipo "pila", perciò le funzioni anziché avere come argomento una "pila" hanno come argomento un puntatore a "linkedList"
 (si confronti il programma con il programma pila-lista-con-nome.go)

2. implementata mediante una lista concatenata:
  - la funzione push effettua un inserimento in testa
  - la funzione pop effettua una rimozione in testa

Il fatto che non esista il tipo "pila" comporta che non possa essere dichiarata una variabile di questo tipo.
Ad esempio, nella funzione "valuta" si dichiara una variabile l di tipo linkedList (riga 36)

Tale variabile, tuttavia, viene modificata solo mediante l'uso delle funzioni "pop", "push",
quindi di fatto è usata come se fosse una struttura dati astratta.


Possibili miglioramenti:
- mancano altre funzioni tipiche della pila (isEmpty, top)
- se riceve una pila vuota (cosa che non succede in questo specifico programma, se l'input è ben formato), 
la funzione "pop" dà un panic error alla riga 31. 


*/
