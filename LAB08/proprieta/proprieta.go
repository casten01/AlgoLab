//3. PROPRIETA DEL GRAFO: da implementare con rappresentazioni diverse

package main

import (
	"fmt"
	"math/rand"
)

type grafo map[string][]string

// Genera un grafo casuale a partire dalla probabilità p (0<=p<=1) di inserire un arco(genero un numero casuale e se maggiore di p inserisco arco)
func gen(p float32, nomi []string) grafo {
	var g grafo = make(map[string][]string)

	for _, v := range nomi {
		for _, b := range nomi {
			if b != v {
				n := rand.Float32()
				if n < p {
					g[v] = append(g[v], b)
				}
			}
		}
	}

	return g
}

// Calcola il grado del vertice v. Si ricorda che il grado di un vertice è definito come il numero di vertici ad esso adiacenti
func degree(s string, g grafo) (grado int) {
	return len(g[s])
}

// Testa l’esistenza di un cammino semplice che collega i vertici v e w. Si ricorda che un cammino si dice semplice quando attraversa ogni vertice al più una volta.
func path(v string, w string, g grafo) bool {
	for _, v2 := range g[v] {
		if v2 == w {
			return true
		}
	}
	aux := make(map[string]bool)
	return dfs1(g, v, w, aux)
}

// attraversamento in profondità modificato
func dfs1(g grafo, v string, w string, aux map[string]bool) bool {
	if v == w {
		return true
	}
	aux[v] = true
	for _, v2 := range g[v] {
		if v2 == w {
			return true
		}
		if !aux[v2] {
			if dfs1(g, v2, w, aux) {
				return true
			}
		}
	}
	return false
}

//4) Funzione che conta il numero di componenti connesse di un grafo (non orientato). Si ricorda che si chiama componente connessa di un grafo ogni insieme massimale di vertici connessi tra loro da un cammino.

//5) cc (v) stampa l’elenco dei vertici della componente connessa contenente v;

//6) span (v) calcola uno spanning tree con radice v e lo stampa nella rappresentazione "a sommario".Si ricorda che si definisce spanning tree (in italiano, albero di copertura) un albero che ha per nodi tutti e soli i vertici del grafo. Osservate che per ottenere uno spanning tree con radice v è sufficiente eseguire una visita della componente connessa contenente v stampando ad ogni passo l’arco attraversato. Che tipo di visita si deve eseguire per avere la garanzia di ottenere uno spanning tree di altezza minimale?

//7) twocolor testa se il grafo è bicolorabile. Un grafo si dice bicolorabile quando è possibile assegnare ad ogni vertice del grafo uno dei due colori bianco o nero in modo che due nodi vicini abbiano sempre colori diversi. Quando un grafo è bicolorabile, si dice anche che è bipartito. Ad esempio, il grafo nell’illustrazione è bipartito (basta colorare i vertici vi di bianco e i vertici ui di nero).

// -> Osservate che per verificare questa proprietà del grafo è sufficiente eseguire una visita in profondità, assegnando colori alternati ai vertici che si visitano man mano.

//8) oddcycles testa se il grafo contiene cicli di lunghezza dispari. Si ricorda che un ciclo è un cammino che parte e finisce nello stesso vertice. Prima di implementare questa operazione, osservate quale relazione c’è tra questa proprietà e la precedente!

// stampa il grafo
func printGrafo(g grafo) {
	for k, seguiti := range g {
		if len(seguiti) > 0 {
			fmt.Printf("%s segue: ", k)
			for _, v := range seguiti {
				fmt.Printf("%s ", v)
			}
			fmt.Println()
			continue
		}
		fmt.Printf("%s non segue nessuno \n", k)
	}
}

func main() {
	var nomi []string = []string{"marco", "viola", "massi", "pablo", "gabriele", "chiara"}
	g := gen(0.3, nomi)
	printGrafo(g)

	if path("marco", "gabriele", g) {
		fmt.Println("Esiste un cammino!")
	} else {
		fmt.Println("Non esiste un cammino!")
	}
}
