//2. VISITE DI GRAFI: rappresentiamo il grafo come lista dei vertici

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type grafo map[string][]string

//-> l'insieme dei vertici è definito implicitamente dall'insieme delle chiavi della mappa

func readGrafo() grafo {
	g := make(map[string][]string)

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		info := strings.Split(scanner.Text(), " ")
		if info[1] != "null" {
			g[info[0]] = append(g[info[0]], info[1])
			continue
		}
		g[info[0]] = make([]string, 0)
	}
	return g
}

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

// attraversamento in profondità
func dfs1(g grafo, v string, aux map[string]bool) {
	fmt.Println(v)
	aux[v] = true
	for _, v2 := range g[v] {
		if aux[v2] != true {
			dfs1(g, v2, aux)
		}
	}
}

// attraversamento in ampiezza
func bfs1(g grafo, v string, aux map[string]bool) {
	coda := []string{v}
	aux[v] = true

	for len(coda) > 0 {
		v := coda[0]
		coda = coda[1:]
		fmt.Println("\t", v)

		for _, v2 := range g[v] {
			if !aux[v2] {
				coda = append(coda, v2)
				aux[v2] = true
			}
		}
	}
}

func main() {
	g := readGrafo()
	//printGrafo(g)
	aux1 := make(map[string]bool, len(g))
	aux2 := make(map[string]bool, len(g))
	fmt.Println("Visista in profondità: ")
	dfs1(g, "marco", aux1)
	fmt.Println("Visita in ampiezza: ")
	bfs1(g, "marco", aux2)
}
