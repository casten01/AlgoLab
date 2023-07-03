/* Il programma usa una mappa, chiamata diz, per tenere traccia di quante volte nella riga compare ogni lettera.
Questa è una buona idea, ma la realizzazione presenta dei difetti.

DIFETTO 1.
Il difetto più grave è nel secondo ciclo, che produce la stampa dell'istogramma.
Per ogni lettera i dell'alfabeto si usa la mappa per recuperare il numeri di ripetizioni della lettera.
Basterebbe accedere all'elemento di chiave i nella mappa (costo costante):

for i := 'a'; i <= 'z'; i++ {
	if diz[byte(i)] > 0 {
		fmt.Println(string(i), " ", strings.Repeat("*", diz[byte(i)]))
	}
}

Invece si itera sulla mappa alla ricerca dell'elemento di chiave i
Asintoticamente la complessità non cambia, poiché la dimensione della mappa è costante (26 elementi)
e non dipende dalla lunghezza della stringa in input;
tuttavia si ripete INUTILMENTE per 26 volte la ricerca nella mappa,
che invece consentirebbe una ricerca in tempo costante.

NB: la conversione da rune a byte, e/o viceversa, non aiuta... forse era meglio una mappa con chiavi di tipo runa.


DIFETTO 2.
Anche il primo ciclo, usato per popolare la mappa, potrebbe essere scritto meglio.
La mappa viene popolata esaminando una parola alla volta - in effetti questo non è particolarmente utile:
bastava un solo ciclo invece dei due cicli annidati:

for i := 0; i < len(s); i++ { // per ogni posizione della riga
	unicode.ToLower(rune(s[i]))
		if s[k] >= 97 && s[k] <= 122 {
			diz[s[k]]++
		}
	}
}

Inoltre, in questo caso sembra più comodo usare il range, dato che l'indice di posizione non ci interessa.
Infine, l'uso di "numeri magici" come 97 o 122 è sconsigliato: meglio usare 'a', 'A'.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	s1 := strings.Split(s, " ")

	diz := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		parola := s1[i]
		for k := 0; k < len(parola); k++ {
			unicode.ToLower(rune(parola[k]))
			if parola[k] >= 97 && parola[k] <= 122 { //
				diz[parola[k]]++
			}
		}
	}

	for i := 'a'; i <= 'z'; i++ {
		for k, v := range diz {
			if rune(k) == i {
				fmt.Println(string(k), " ", strings.Repeat("*", v))
			}
		}
	}
}
