// ESERCIZIO TEMA D'ESAME GENNAIO 2023
package main

import "fmt"

type node struct {
	val  int
	next *node
}

type lista struct {
	head *node
	tail *node
}

func insert(p *lista, val int) {
	new := &node{val, nil}
	p.tail.next = new
	p.tail = new
}

func insertCirc(p *lista, val int) {
	new := &node{val, nil}
	new.next = p.head
	p.tail.next = new
	p.tail = new
}

func printList(l lista) {
	p := l.head
	for p != nil {
		fmt.Print(p.val, " ")
		p = p.next
	}
	fmt.Println()
}

func printListCirc(l lista) {
	p := l.head
	for p != l.tail.next {
		fmt.Print(p.val, " ")
		p = p.next
	}
	fmt.Println()
}

func f(p *node, k int) int {
	var a int
	if p == nil {
		return 0
	}
	a = 1 + f(p.next, k)
	if a == k {
		fmt.Println(p.val)
	}

	return a
}

//Qualunque elemento della lista punti p, stampa elementi a partire da zero
func stampaDaZero(p *node) {
	var supp *node = p
	for {
		if supp.val == 0 {
			break
		}
		supp = supp.next
	}
	for {
		fmt.Printf("%d ", supp.val)
		supp = supp.next
		if supp.val == 0 {
			break
		}
	}
}

func main() {
	var l lista
	l.head = &node{3, nil}
	l.tail = l.head
	insertCirc(&l, 2)
	insertCirc(&l, 0)
	insertCirc(&l, 1)
	insertCirc(&l, 7)
	//printListCirc(l)
	stampaDaZero(l.head.next)
}

// 3 -> 2 -> 0 -> 1 -> 7
/*
1. Cosa stampa la chiamata di f(p,k) con k==1, p punta a 3
	Stampa 7, perchè all'ultima chiamata ricorsiva p punta all'ultimo elemento della lista e a = 1 + 0 e quindi uguale a k che non cambia mai in tutta l'esecuzione della funzione

2. Cosa stampa la chiamata di f(p,k) con k==5, p punta a 3
	Stampa 3, perchè p punta a 3 quando a == k che non cambia mai in tutta l'esecuzione della funzione

3.Cosa stampa la chiamata di f(p,k) con k==10, p punta a 3
	Nulla

4.Ricevendo il puntatore alla testa di una lista e un intero k la funzione f stampa il k-esimo elemento della lista partendo dalla coda contando da 1 a n e restituisce il numero n di elementi della coda

5. La complessità della funzione è
*/
