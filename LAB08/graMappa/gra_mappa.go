//1.2 Struttura per i vertici e mappa per la relazione di adiacenza
/*
Grafo (orientato) in cui ogni vertice rappresenta un utente di una rete sociale (tipo Twitter). Ogni vertice ha un nome, un’età, e una serie di hobby (rappresentiamoli come come slice di stringhe: ogni stringa descrive un hobby).
C’è un arco dal vertice A al vertice B se l’utente rappresentato dal vertice A segue l’utente rappresentato dal vertice B.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Struttura che rappresenta un utente di una rete sociale
type item struct {
	nome    string
	eta     int
	hobbies []string
}

// vertice con chiave usata nella mappa
type vertice struct {
	val item
	key string //coincide con il nome
}

// struttura complessiva che rappresenta il grafo
type grafo struct {
	vertici   map[string]*vertice   //Insieme di vertici, facile accesso
	adiacenti map[string][]*vertice //Insieme dei vertici adiacenti
}

// Crea un nuovo grafo inizializzando le due mappe
func graphNew() *grafo {
	var g grafo
	g.vertici = make(map[string]*vertice)
	g.adiacenti = make(map[string][]*vertice)
	return &g
}

// Crea e restituisce un vertice date le sue info
func createVertice(nome string, eta int, hobbies []string) *vertice {
	var v vertice
	v.val.nome = nome
	v.val.eta = eta
	v.val.hobbies = hobbies
	v.key = nome

	return &v
}

// Inserisce un vertice nel grafo, direttamente come puntatore
func insertVertice(v *vertice, g *grafo) {
	g.vertici[v.key] = v
}

// Legge la serie di vertici che rappresentano gli utenti
func readGrapgSI(g *grafo) {
	// open file
	f, err := os.Open("inputMap.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords) //parola per parola

	var infoNodo []string
	for scanner.Scan() {
		infoNodo = append(infoNodo, scanner.Text())
	}
	var i int = 0

	//nome eta hobbie1,hobbie2,hobbie3
	for i < len(infoNodo) {
		nome := infoNodo[i]                          //nome
		eta, _ := strconv.Atoi(infoNodo[i+1])        //eta
		hobbies := strings.Split(infoNodo[i+2], ",") //hobbies
		v := createVertice(nome, eta, hobbies)
		insertVertice(v, g)
		i += 3 //passo all'utente dopo, potevo farlo riga per riga
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	defer f.Close()
}

// Legge legami follow nel formato: nome seguito1 seguito2
func readFollow(g *grafo) {
	// open file
	f, err := os.Open("inputMapAmici.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var infoAmici []string
	var i int = 0

	for scanner.Scan() {
		infoAmici = append(infoAmici, scanner.Text())
		tmp := strings.Split(infoAmici[i], " ")
		tmp2 := strings.Split(tmp[1], ",")
		var seguiti []*vertice
		for _, v := range tmp2 {
			k := g.vertici[v]
			seguiti = append(seguiti, k)
		}
		key := tmp[0]
		g.adiacenti[key] = seguiti

		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	defer f.Close()

}

// stampa i vertici e le info degli utenti
func printGrafo(g *grafo) {
	for _, v := range g.vertici {
		fmt.Printf("nome: %s \neta: %d \nhobbies: ", v.val.nome, v.val.eta)
		for i := 0; i < len(v.val.hobbies); i++ {
			fmt.Printf("%s ", v.val.hobbies[i])
		}
		fmt.Printf("\nSegue: ")

		seguiti := g.adiacenti[v.key]
		for i := 0; i < len(seguiti); i++ {
			nome := seguiti[i].val.nome
			fmt.Printf("%s ", nome)
		}

		fmt.Println()
		fmt.Println("---------------")
	}
}

// Scrivete una funzione che data una stringa A stampi gli hobby dell’utente di nome A e l’elenco di tutti gli hobby delle persone che seguono A.
func showHobbies(s string, g *grafo) {
	fmt.Printf("Gli hobbies di %s sono: ", s)
	v := g.vertici[s]
	for i := 0; i < len(v.val.hobbies); i++ {
		fmt.Printf("%s ", v.val.hobbies[i])
	}

	fmt.Printf("\nGli hobbies delle persone che seguono %s sono:\n", s)
	for k, v := range g.adiacenti {
		for _, name := range v {
			if name.val.nome == s {
				fmt.Println(g.vertici[k].val.hobbies)
			}

		}
	}
}

// Scrivete una funzione che data una stringa A stampi gli hobby dell’utente di nome A e l’elenco di tutti gli hobby delle persone seguite da A. Quale è la più complessa tra questa e l’operazione precedente?
func showHobbies2(s string, g *grafo) {
	fmt.Printf("Gli hobbies di %s sono: ", s)
	v := g.vertici[s]
	for i := 0; i < len(v.val.hobbies); i++ {
		fmt.Printf("%s ", v.val.hobbies[i])
	}

	fmt.Printf("\nGli hobbies delle persone seguite da %s sono:\n", s)
	for _, v := range g.adiacenti[s] {
		fmt.Println(g.vertici[v.key].val.hobbies)
	}
}

func main() {
	grafo := graphNew()

	readGrapgSI(grafo)
	readFollow(grafo)

	//printGrafo(grafo)

	showHobbies2("Chiara", grafo)

}
