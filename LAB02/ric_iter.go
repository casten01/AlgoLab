//4) RICORSIONE e ITERAZIONE

package main

import (
	"fmt"
	"log"
	"time"
)

func f_rec(n int) uint64 {
	// defer timeTrack(time.Now(), "f_rec")
	if n == 1 || n == 2 {
		//fmt.Println("caso base")
		return 1
	}
	//fmt.Println("caso intermedio")

	return f_rec(n-1) + f_rec(n-2) //
}

//===========================

func f_iter1(n int) uint64 {
	// defer timeTrack(time.Now(), "f_iter1")
	var f uint64
	var f1, f2 uint64 = 1, 1

	if n == 2 || n == 1 {
		return 1
	}

	for n >= 3 {
		n--
		f = f1 + f2
		f1 = f2
		f2 = f
	}
	return f
}

func f_iter2(n int) uint64 {
	// defer timeTrack(time.Now(), "f_iter2")
	var f uint64
	var f1, f2 uint64 = 1, 1

	if n == 2 || n == 1 {
		return 1
	}

	//con n-1 funzione uguale alla prima
	for i := 2; i <= n; i++ {
		f = f1 + f2
		f1 = f2
		f2 = f
	}
	return f
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(f_rec(n))
	fmt.Println(f_iter1(n))
	fmt.Println(f_iter2(n))
	fmt.Println(f_riter(1, 1, 10))
}

//Senza eseguire la funzione al computer, rispondete alle seguenti domande:
//1. Cosa restituisce la funzione f_rec(7)? 13
/*
       f6         +        f5
   f5   +   f4    +    f4   +     f3
f4 + f3 + f3 + f2 + f3 + f2 + f2 + f1
f3 + f2 + f2 + f1 + f2 + f1 + 1 + f2 + f1 + 1 + 1 + 1
f2 + f1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1 + 1
= 13
*/
//2. Perché n è dichiarato come intero mentre il valore restituito è di tipo uint64?
/*
non può essere negativo e può essere molto grosso a fronte di un input piccolo
*/
//3. Riassumete a parole cosa restituisce la funzione se riceve come argomento un intero positivo n maggiore di 0.
/*
	sequenza di fibonacci
*/

//Considerate ora le due funzioni f_iter1 e f_iter2

//4. Considerando solo il valore restituito, le due funzioni sono equivalenti? (ovvero: restituiscono sempre lo stesso valore?) NO
//5. Le due funzioni sono equivalenti alla funzione f_rec? la prima si la seconda no
//6. Modificate (se necessario) le funzioni f_iter1 e f_iter2 in modo che risultino essere equivalenti a f_rec.   -> n-1 nel secondo ciclo
//7. Stimate il numero di operazioni che si svolgono durante l’esecuzione di f_rec, f_iter1 e f_iter2: sono paragonabili?

//Considerate infine la seguente funzione ricorsiva f_riter

func f_riter(a, b uint64, n int) uint64 {
	// defer timeTrack(time.Now(), "f_siter")
	if n == 2 {
		return a
	}
	if n == 1 {
		return b
	}

	return f_riter(a+b, a, n-1)
}

//8. Convincetevi che questa funzione può essere usata per calcolare f_rec. In particolare: con quali argomenti devo invocare f_riter per ottenere il valore restituito da f_rec(n)?
// a e b devono essere uguali a 1, mentre n è lo stesso n di prima
//9. Rappresentate graficamente lo schema delle chiamate ricorsive definiti dall’invocazione f_rec(7) e dalla chiamata equivalente del tipo f_riter(..., ..., ...).
//10. Considerate il numero di chiamate ricorsive effettute da f_rec(n) e dalla chiamata equivalente del tipo f_riter(..., ..., ...). Sono paragonabili?

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %dms", name, elapsed.Nanoseconds()/1000)
}
