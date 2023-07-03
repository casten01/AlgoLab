package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Edge struct {
	U, V   int     // Identificatori dei punti
	Weight float64 // Distanza tra i punti
}

func Prim(points []Point) []Edge {
	n := len(points)

	// Calcolo delle distanze tra i punti
	distances := make([][]float64, n)
	for i := 0; i < n; i++ {
		distances[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				distances[i][j] = math.Sqrt((points[j].X-points[i].X)*(points[j].X-points[i].X) + (points[j].Y-points[i].Y)*(points[j].Y-points[i].Y))
			}
		}
	}
	//->evito di salvarmele tutte e le calcolo man mano sotto

	// Inizializzazione
	visited := make([]bool, n)
	visited[0] = true

	minimumSpanningTree := make([]Edge, 0)
	var iterazioni int = 0

	// Costruzione dell'albero ricoprente minimo
	for i := 1; i < n; i++ {
		minWeight := math.Inf(1)
		u, v := 0, 0

		// Trova l'arco con il peso minimo tra i punti visitati e non visitati
		for j := 0; j < n; j++ {
			if visited[j] {
				for k := 0; k < n; k++ {
					iterazioni++
					if !visited[k] && distances[j][k] < minWeight {
						minWeight = distances[j][k]
						u, v = j, k
					}
				}
			}
		}

		visited[v] = true
		minimumSpanningTree = append(minimumSpanningTree, Edge{U: u, V: v, Weight: minWeight})
	}
	fmt.Println(iterazioni)
	return minimumSpanningTree
}

func main() {
	// Esempio di utilizzo
	points := []Point{
		{X: 0, Y: 0},
		{X: 1, Y: 2},
		{X: 3, Y: 1},
		{X: 4, Y: 3},
		{X: 2, Y: 5},
	}

	minimumSpanningTree := Prim(points)

	// Stampa degli archi dell'albero ricoprente minimo
	for _, edge := range minimumSpanningTree {
		fmt.Printf("Edge: %d - %d, Weight: %.2f\n", edge.U, edge.V, edge.Weight)
	}
}
