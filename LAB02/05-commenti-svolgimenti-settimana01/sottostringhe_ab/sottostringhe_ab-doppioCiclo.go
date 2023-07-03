/* DIFETTO: ogni volta che trovo una 'a', scandisco l'intero suffisso per cercare tutte le 'b'
La complessità in tempo è quadratica: per ogni posizione della stringa, scandisco tutta la parte rimanente di stringa
(si consideri ad es la stringa aaaaaaab)
*/

package main

import "fmt"

func main() {
	var s string
	var counter int
	fmt.Scanf("%s", &s)

	for k, v := range s {
		if v == 'a' {
			for i := k + 1; i < len(s); i++ {
				if s[i] == 'b' {
					counter++
				}
			}
		}
	}
	fmt.Println(counter)
}
