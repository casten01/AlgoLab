package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Insert html document (one line): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	doc := scanner.Text()
	wellFormed, errPos, remainingTags := checkTags(doc)
	if wellFormed {
		fmt.Println("The document is well formed")
	} else {
		fmt.Println("The document is NOT well formed:")
		if len(remainingTags) > 0 {
			fmt.Println("\tTags not closed:", remainingTags)
		} else {
			fmt.Println("\tError position:", errPos)
		}
	}
}

func checkTags(doc string) (bool, int, []string) {
	var stack stack = stack{nil}
	tags := strings.Split(doc, " ")
	for i, t := range tags {
		if !strings.Contains(t, "/") {
			stack.push(t)
		} else if strings.Contains(t, "/") {
			openingTag := stack.pop()
			if openingTag[1:] != t[2:] {
				return false, i + 1, nil
			}
		}
	}
	if stack.isEmpty() {
		return true, -1, nil
	} else {
		var remainingTags = make([]string, 0)
		for !stack.isEmpty() {
			remainingTags = append(remainingTags, stack.pop())
		}
		for i, j := 0, len(remainingTags)-1; i < j; i, j = i+1, j-1 {
			remainingTags[i], remainingTags[j] = remainingTags[j], remainingTags[i]
		}
		return false, -1, remainingTags
	}
}

// IMPLEMENTAZIONE PILA TRAMITE LISTA //

type listNode struct {
	next *listNode
	item string
}

type stack struct {
	head *listNode
}

func newNode(item string) *listNode {
	return &listNode{nil, item}
}

func (list *stack) push(item string) {
	newNode := newNode(item)
	newNode.next = list.head
	list.head = newNode
}

func (list *stack) pop() string {
	node := list.head
	list.head = node.next
	return node.item
}

func (list stack) isEmpty() bool {
	if list.head == nil {
		return true
	}
	return false
}

func (list stack) print() {
	fmt.Print("[ ")
	var node *listNode = list.head
	for node != nil {
		fmt.Print(node.item, " ")
		node = node.next
	}
	fmt.Println("]")
}

/*



/* Note:

In questo programma, la pila è:

1. definita come struttura di dati astratta:
  - è definito un tipo "stack"
  - sono definiti dei metodi generali "pop", "push", ...


2. usata come struttura di dati astratta:
  - nel main si dichiara una variabile "stack" di tipo "stack"
  - la variabile "stack" viene modificata solo mediante l'uso dei metodi "pop", "push"

3. implementata mediante una lista concatenata:
  - il metodo push effettua un inserimento in testa
  - il metodo pop effettua una rimozione in testa


Difetto importante: se riceve una pila vuota, il metodo "pop" dà un panic error alla riga 78
(vedi nota nel programma pila-lista-con-nome.go)
Ad esempio su input: "</a>"

Ci sono due approcci per ovviare a questo problema:
- o si specifica in un commento la pre-condizione al metodo "po" (la pila non deve essere vuota) -> questo significa che la funzione chiamante deve prima fare questo controllo,
- o si aggiunge un controllo prima della riga 35 usando il metodo "isEmpty"
In quest'ultimo caso si può arricchire la funzione con una gestione più articolata dell'errore (ad esempio restituendo un secondo valore per segnalare l'errore)



Possibili miglioramenti:
- si potrebbe completare la struttura di dati astratta con un'altra funzione tipica della pila (top)
- (riga 32) lo '/' se c'è, è sempre in posizione 1, non serve usare la Contains

*/
