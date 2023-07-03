/* Costruisce mappa delle occorrenze di ciascuna parola, poi le confronta: l'idea è buona.

Si può migliolare l'implementazione:
- meglio evitare numeri magici (64, 91), sostituire con 'a', 'A'
- meglio una mappa che ha per chiavi direttamente delle rune, invece che delle stringhe
*/

package main

import (
	. "fmt"
	"strings"
)

func main() {

	var parola1, parola2 string
	Print("prima parola: ")
	Scan(&parola1)
	Println()
	Print("seconda parola: ")
	Scan(&parola2)
	parola1 = strings.ToUpper(parola1)
	parola2 = strings.ToUpper(parola2)
	Println(parola1, parola2)
	var mappa1 map[string]int = map[string]int{}
	var mappa2 map[string]int = map[string]int{}

	for _, v := range parola1 {
		vv := string(v)
		if v > 64 && v < 91 {
			mappa1[vv]++
		}
	}
	for _, v := range parola2 {
		vv := string(v)
		if v > 64 && v < 91 {
			mappa2[vv]++
		}
	}
	Println(mappa1, mappa2)

	if len(mappa1) != len(mappa2) {
		Println("NON sono anal")
		return
	}

	for key1, _ := range mappa1 {
		if mappa1[key1] != mappa2[key1] {
			Println("NON sono anal")
			return
		}
	}
	Println("ANAL")
}
