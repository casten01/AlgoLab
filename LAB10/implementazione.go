package main

import "fmt"

type Nodo struct {
	val      int
	visitato bool
}

type Arco struct {
	peso int
	part *Nodo
	arr  *Nodo
}

type Grafo struct {
	nodi map[int]*Nodo
	adj  map[*Nodo][]*Arco
}

func arcoPesoMin(archi []*Arco) *Arco {
	min := archi[0]
	for i := 1; i < len(archi); i++ {
		if archi[i].peso < min.peso {
			min = archi[i]
		}
	}
	return min
}

func algoritmo(grafo *Grafo, partenza int, arrivo int) int {
	n := grafo.nodi[partenza]
	vicini := grafo.adj[n]
	count := 0
	for n.val != arrivo {
		arcoMin := arcoPesoMin(vicini)
		if !arcoMin.arr.visitato {
			count++
			arcoMin.arr.visitato = true
			n = arcoMin.arr
			vicini = grafo.adj[n]
		} else {
			return -1
		}
	}
	return count
}

func main() {
	var n, m, h, s int
	var part, arr, peso int
	g := new(Grafo)
	g.nodi = make(map[int]*Nodo)
	g.adj = make(map[*Nodo][]*Arco)
	fmt.Println("Inserisci N M H S")
	fmt.Scanf("%d %d %d %d", &n, &m, &h, &s)
	for i := 1; i <= n; i++ {
		g.nodi[i] = new(Nodo)
		g.nodi[i].visitato = false
		g.nodi[i].val = i
		g.adj[g.nodi[i]] = make([]*Arco, 0)
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d %d", &part, &arr, &peso)
		arco := new(Arco)
		arco.peso = peso
		arco.part = g.nodi[part]
		arco.arr = g.nodi[arr]
		g.adj[g.nodi[part]] = append(g.adj[g.nodi[part]], arco)
		g.adj[g.nodi[arr]] = append(g.adj[g.nodi[arr]], arco)
	}
	fmt.Println(algoritmo(g, h, s))
}
