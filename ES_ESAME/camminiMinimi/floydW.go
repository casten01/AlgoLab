//IMPLEMENTAZIONE DELL'ALGORITMO DI FLOYD e WARSHALL: cammino minimo tra tutte le coppie di vertici
//ALGORITMO di FLOYD-WARSHALL: cammino minimo tra tutte le coppie di vertici, il grafo non deve contenere cicli di peso negativo
/*
- Implementazione grafo: matrice dei pesi
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	pesi    [][]int //matrice dei pesi, in [i,j] = peso arco da i a j
	vertici int
}

const MAX int = math.MaxInt32

func floydWarshall(g Graph) (D [][]int, P [][]int) {
	P = make([][]int, g.vertici)
	D = make([][]int, g.vertici)

	//Per ogni coppia di vertici, calcola la minima lunghezza di un cammino che li congiunge senza passare per altri vertici
	/*
		D[i,j]  = w(v_i, vj) se i != j
					  = 0 se i == j
					  = INF altrimenti
	*/
	for i := 0; i < g.vertici; i++ {
		D[i] = make([]int, g.vertici)
		P[i] = make([]int, g.vertici)
		for j := 0; j < g.vertici; j++ {
			if g.pesi[i][j] != 0 {
				D[i][j] = g.pesi[i][j]
			} else if i == j {
				D[i][j] = 0
			} else {
				D[i][j] = MAX
			}
			P[i][j] = 0
		}
	}

	//Per ogni coppia di vertici, calcola la minima lunghezza di un cammino che li congiunge, i cui vertici intermedi appartengono ai vertici del grafo (indice <= k). Nella matrice P si tiene traccia del massimo indice dei vertici intermedi lungo tale cammino minimo
	// D[i][j] (k) = min {D[i][j](k), D[i][k](k-1) + D[k][j](k-1) }
	for k := 1; k <= g.vertici; k++ {
		for i := 0; i < g.vertici; i++ {
			for j := 0; j < g.vertici; j++ {
				if D[i][k-1]+D[k-1][j] < D[i][j] {
					D[i][j] = D[i][k-1] + D[k-1][j]
					P[i][j] = k
				}
			}
		}
	}

	return D, P
}

// FUNZIONI AUSILIARIE UTILI:

/*crea un grafo implementato con una matrice dei pesi, riempie la matrice da un file */
func createMatrice(n int) (g Graph) {
	g.vertici = n
	g.pesi = readFileMatrice(n)

	return g
}

/*legge una matrice da un file data la sua dimensione*/
func readFileMatrice(n int) [][]int {
	mat := make([][]int, n)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var i int = 0
	for scanner.Scan() {
		mat[i] = make([]int, n)
		line := strings.Split(scanner.Text(), " ")
		for j := 0; j < len(line); j++ {
			mat[i][j], err = strconv.Atoi(line[j])
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	return mat
}

/*stampa la matrice in input, sostiuisce i max con 0*/
func printMat(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			if mat[i][j] == MAX {
				fmt.Print(" 0")
			} else {
				fmt.Print(mat[i][j], " ")
			}
		}
		fmt.Println()
	}
}

/*ricostruisce i cammini minimi tra tutte le coppie di vertici*/
func rebuiltPath(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		fmt.Printf("Cammini minimi da [%d]: \n", i+1)
		for j := 0; j < len(mat); j++ {
			fmt.Printf("- a [%d]: ", j+1)
			var path []int
			if mat[i][j] == 0 && i != j {
				fmt.Printf("direttamente collegati")
			} else if i == j {
				fmt.Printf("stesso vertice")
			} else {
				fmt.Printf("passando per ")
				var k int = mat[i][j]
				for k != 0 {
					path = append(path, k)
					k = mat[k-1][j]
				}
				for c := len(path) - 1; c >= 0; c-- {
					fmt.Printf("[%d] ", path[c])
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func main() {
	g := createMatrice(6)
	printMat(g.pesi)
	_, P := floydWarshall(g)
	// fmt.Println("------------------------------")
	// printMat(D)
	fmt.Println("------------------------------")
	printMat(P)
	rebuiltPath(P)

}
