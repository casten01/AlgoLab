package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	Source int
	Target int
	Weight int
}

type Graph struct {
	Vertices int
	Edges    []Edge
}

type Subset struct {
	Parent int
	Rank   int
}

// Funzione per trovare il set radice di un elemento
func find(subsets []Subset, i int) int {
	if subsets[i].Parent != i {
		subsets[i].Parent = find(subsets, subsets[i].Parent)
	}
	return subsets[i].Parent
}

// Funzione per unire due set in base al rango
func union(subsets []Subset, x, y int) {
	rootX := find(subsets, x)
	rootY := find(subsets, y)

	if subsets[rootX].Rank < subsets[rootY].Rank {
		subsets[rootX].Parent = rootY
	} else if subsets[rootX].Rank > subsets[rootY].Rank {
		subsets[rootY].Parent = rootX
	} else {
		subsets[rootY].Parent = rootX
		subsets[rootX].Rank++
	}
}

// Algoritmo di Kruskal per trovare l'albero di copertura minimo
func KruskalMST(graph Graph) []Edge {
	// Ordina gli archi in ordine crescente di peso
	sort.Slice(graph.Edges, func(i, j int) bool {
		return graph.Edges[i].Weight < graph.Edges[j].Weight
	})

	mst := make([]Edge, 0)
	subsets := make([]Subset, graph.Vertices)

	// Inizializza i set di ogni vertice
	for i := 0; i < graph.Vertices; i++ {
		subsets[i].Parent = i
		subsets[i].Rank = 0
	}

	// Aggiungi archi all'albero di copertura minimo finché non si raggiunge il numero di vertici - 1
	for _, edge := range graph.Edges {
		rootSource := find(subsets, edge.Source)
		rootTarget := find(subsets, edge.Target)

		// Se gli estremi dell'arco appartengono a set diversi, uniscili e aggiungi l'arco all'albero di copertura minimo
		if rootSource != rootTarget {
			union(subsets, rootSource, rootTarget)
			mst = append(mst, edge)
		}
	}

	return mst
}

func main() {
	// Creazione del grafo di esempio
	graph := Graph{
		Vertices: 4,
		Edges: []Edge{
			{Source: 0, Target: 1, Weight: 10},
			{Source: 0, Target: 2, Weight: 6},
			{Source: 0, Target: 3, Weight: 5},
			{Source: 1, Target: 3, Weight: 15},
			{Source: 2, Target: 3, Weight: 4},
		},
	}

	// Trova l'albero di copertura minimo
	mst := KruskalMST(graph)

	// Stampa gli archi dell'albero di copertura minimo
	for _, edge := range mst {
		fmt.Printf("Arco: %d - %d, Peso: %d\n", edge.Source, edge.Target, edge.Weight)
	}
}
