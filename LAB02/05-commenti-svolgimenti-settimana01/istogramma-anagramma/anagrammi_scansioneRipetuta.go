/* DIFETTI (funzionamento):
Nel tentativo di eliminare da parola2 le lettere trovate man mano,
si ha un out-of-bound quando si arriva alla fine di parola 2.
Inoltre parola2 viene distrutta nel processo.

DIFETTO (complessità):
Si scansiona parola2 ripetutamente. Meglio annotare con una struttura dati di supporto (mappa/slice)
il numero di occorrenze delle lettere di ciascuna parola e confrontare le mappe/slice delle due parole.
*/

package main

import "fmt"

func main() {
	var parola1, parola2, oldparola2 string
	fmt.Scan(&parola1)
	fmt.Scan(&parola2)
	oldparola2 = parola2
	if len(parola1) != len(parola2) {
		fmt.Println(parola1, " non è anagramma di ", parola2)
		return
	}
	for _, lettera := range parola1 {
		var found bool
		for i, letteracheck := range parola2 {
			if lettera == letteracheck {
				parola2 = parola2[:i] + parola2[i+1:]
				found = true
			}
		}
		if found == false {
			fmt.Println(parola1, " non è anagramma di ", parola2)
			return
		}
	}
	fmt.Println(parola1, " è anagramma di ", oldparola2)
}
