//LETTERE DI ALGOLANDIA: grafo orientato che ha un insieme di vertici che rappresentato le lettere dell'alfabeto e l'insieme degli archi è l'insieme delle regole: c'è un arco da X a Y se X può essere seguito da Y
// -> guardo appunti su ipad

package main

import (
	"fmt"
	"sort"
)

type Grafo map[rune][]rune

// Legge n regole XY, costruisce grafo e lo restiuisce
func read() (g Grafo) {
	var n int
	fmt.Scan(&n)
	g = make(Grafo)

	archi := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&archi[i])
	}
	for _, v := range archi {
		a := v[0]
		b := v[1]
		g[rune(a)] = append(g[rune(a)], rune(b))
	}
	return g
}

// stampa tutte le lettere che possono seguire k rispettando regole
func followers(k rune, g Grafo) {
	fmt.Println(string(g[k]))
}

// stampa in ordine alfabetico tutte le lettere che possono seguire k rispettando regole
func max(g Grafo) {
	var max int = 0
	var lettere []rune
	for _, v := range g {
		if len(v) > max {
			max = len(v)
		}
	}
	for k, v := range g {
		if len(v) == max {
			lettere = append(lettere, k)
		}
	}

	sort.Slice(lettere, func(i, j int) bool {
		return lettere[i] < lettere[j]
	})

	for _, v := range lettere {
		fmt.Print(string(v))
	}
	fmt.Println()
}

// stampa tutte le lettere che possono essere contenute in una parola che inizia con k rispettando le regole
func parola(g Grafo, v rune, aux map[rune]bool) {
	fmt.Print(string(v))
	aux[v] = true
	for _, v2 := range g[v] {
		if aux[v2] != true {
			parola(g, v2, aux)
		}
	}
}

func main() {
	var aux map[rune]bool = make(map[rune]bool)
	g := read()
	//followers('A', g)
	//(g)
	parola(g, 'H', aux)
}
