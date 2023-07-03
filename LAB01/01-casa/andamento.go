/* ANDAMENTO: data da standard input una serie di interi positivi terminata da 0, stampare ’+’ ogni volta che il
nuovo valore è maggiore o uguale al precedente e ’-’ altrimenti.
*/

package main

import "fmt"

func main() {
	sequenza := make([]int, 1)
	var n, i int
	fmt.Scan(&n)
	sequenza[0] = n

	for n != 0 {
		i += 1
		fmt.Scan(&n)
		sequenza = append(sequenza, n)
		if sequenza[i] >= sequenza[i-1] {
			fmt.Println("+")
		} else {
			fmt.Println("-")
		}
	}
}
