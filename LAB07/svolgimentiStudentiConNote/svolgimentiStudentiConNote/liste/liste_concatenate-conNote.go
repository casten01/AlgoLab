/*
Scrivete un programma con l’implementazione di una lista concatenata semplice,
seguendo i lucidi presentati a lezione. Definite i tipi listNode e linkedList e
scrivete queste funzioni:
• newNode, che crea un nuovo nodo di lista;
• addNewNode, che inserisce un nuovo nodo in testa alla lista;
• printList, che stampa una lista;
• searchList, che cerca un elemento in una lista;
• removeItem, che cancella un item da una lista.
Per testare le vostre funzioni scrivete una funzione main che gestisca un insieme
dinamico (che variano nel tempo) di interi. Il main deve leggere da standard input
una sequenza di istruzioni secondo il formato nella tabella qui sotto, dove n è un
intero. I vari elementi sulla riga sono separati da uno o più spazi. Quando una riga
è letta, viene eseguita l’operazione associata; le operazioni di stampa sono
effettuate sullo standard output, e ogni operazione deve iniziare su una nuova riga.
*/
package main

import (
	"fmt"
	"strconv"
)

type listNode struct {
	item int
	next *listNode
}
type linkedList struct {
	head *listNode
}

func main() {
	var (
		ok  bool
		l   linkedList
		s   string
		nrs []*listNode
	)
	/*
		+ n Se n non appartiene all’insieme lo inserisce, altrimenti non fa nulla
		- n Se n appartiene all’insieme lo elimina, altrimenti non fa nulla
		? n Stampa un messaggio che dichiara se n appartiene all’insieme
		c Stampa il numero di elementi dell’insieme
		p Stampa gli elementi dell’insieme (nell’ordine in cui compaiono nella lista)
		o Stampa gli elementi dell’insieme nell’ordine inverso
		d Cancella tutti gli elementi dell’insieme
		f Termina l’esecuzione
	*/
	for {
		fmt.Scan(&s)
		switch s[0] {
		case '+':
			n, _ := strconv.Atoi(s[2:])
			tf, _ := searchList(l, n)
			if tf == true {
				addNewNode(l, n)
			}
		case '-':
			n, _ := strconv.Atoi(s[2:])
			tf, _ := searchList(l, n)
			if tf == true {
				removeItem(&l, n)
			}
		case '?':
			n, _ := strconv.Atoi(s[2:])
			tf, _ := searchList(l, n)
			fmt.Print("La lista ")
			if tf == false {
				fmt.Print("non ")
			}
			fmt.Println("contiene", n)
		case 'c':
			p := l.head
			i := 0
			for p != nil {
				i++
				p = p.next
			}
			fmt.Println(i)
		case 'p':
			printList(l)
		case 'o':
			p := l.head
			for p != nil {
				nrs = append(nrs, p)
				p = p.next
			}
			for j := len(nrs); j >= 0; j-- {
				fmt.Print(nrs[j])
			}
		case 'd':
			p := l.head
			for p != nil {
				p = nil
				p = p.next
			}
		case 'f':
			ok = true
		}
		if ok == true {
			break
		}
	}
}
func newNode(n int) *listNode {
	return &listNode{n, nil}
}
func addNewNode(l linkedList, n int) linkedList {
	node := newNode(n)
	node.next = l.head
	l.head = node
	return l
}
func printList(l linkedList) {
	p := l.head
	for p != nil {
		fmt.Print(p.item, " ", '\n')
		p = p.next
	}
}
func searchList(l linkedList, n int) (bool, *listNode) {
	tf := false
	p := l.head
	for p != nil {
		if p.item == n {
			tf = true
		}
		p = p.next
	}
	return tf, p
}
func removeItem(l *linkedList, n int) bool {
	var curr, prev *listNode = l.head, nil
	tf := false
	for curr != nil {
		if curr.item == n {
			if prev == nil {
				l.head = l.head.next
			} else {
				prev.next = curr.next
			}
			tf = true
		}
		prev = curr
		curr = curr.next
	}
	return tf
}

/* CORREZIONE

L’esercizio è impostato correttamente, ed è tutto sommato corretto. Tuttavia eseguendolo sembra del tutto sbagliato.
In realtà, sembra più sbagliato di quanto non sembri, innanzitutto a causa di errori nella lettura dell’input.

Consideriamo al riga 50: ricordiamo che, quando si legge una stringa da stdin, la lettura si interrompe dopo il primo spazio.
Dunque, la stringa “+ 7” verrebbe spezzata se c’è uno spazio tra + e 7.
Invece, la stringa “+7” senza spazi, verrebbe letta interamente.
In questo secondo caso però, il numero inizia alla posizione 1 della stringa, quindi la riga
            n, _ := strconv.Atoi(s[2:])
andrebbe sostituita con
            n, _ := strconv.Atoi(s[1:])
(Lo stesso errore si trova nella lettura dei comandi “-n” e “?n")
Se invece si vuole leggere l’input con lo spazio dopo il primo simbolo '+', '-', '?' (come specificato nella tabella dell'esercizio),
allora meglio leggere l’input riga per riga con uno scanner.

A valle di questo ci sono altri errori:

1- l’inserimento va fatto solo quando l’elemento non c’è, ovvero quando la searchList restituisce false (non true, come nella tua riga 55

2- la funzione addNewNode restituisce la lista modificata, quindi quando invochi addNewNode devi assegnare il valore restituito, altrimenti la lista della funzione chiamante non viene modificata (vedi lucidi annotati, slide 9/18). Cioè, la riga 56
                addNewNode(l, n)
Va sostituita con
                 l = addNewNode(l, n)

3- nella stampa al contrario, c’è un problema di off-by-one. Devi cominciare dalla posizione con indice  len(nrs) - 1 invece che da quella di len(nrs)

4- nella stampa di riga 117
	fmt.Print(p.item, " ", '\n’)
il carattere ‘\n’ viene stampato come codice numerico, non come stringa, quindi su standard output viene stampato un “10” dopo ogni elemento della lista. Bisogna usare le doppie virgolette:
	fmt.Print(p.item, " ", "\n”)

5- searchList: la ricerca potrebbe essere interrotta dopo aver trovato il nodo con l'item cercato.
Il secondo valore restituito dovrebbe essere l'indirizzo del nodo con l'item cercato, invece la funzione restituisce sempre nil.
Bisogna aggiungere un return dentro l'if, con p come secondo valore restituito:


func searchList(l linkedList, n int) (bool, *listNode) {
	p := l.head
	for p != nil {
		if p.item == n {
			return true, p
		}
		p = p.next
	}
	return false, nil
}

NB: potrebbero esserci altri errori/difetti non riportati in queste note.


*/
