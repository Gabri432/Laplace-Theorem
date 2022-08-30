# Teorema di Laplace
Un semplice programma scritto in GO che permetterà di calcolare il determinante di una matrice quadrata usando il metodo di Laplace.
Per maggiori dettagli `https://en.wikipedia.org/wiki/Laplace_expansion`.

## Funzionalità
- Calcolerà il determinante di matrici di qualunque dimensione.

## Struttura del Progetto
- `Laplace Theorem` (cartella principale), dove sono presenti i seguenti file:
  - `laplace.go`, dove ci sono i principali file della libreria,
  - `laplace_test.go`, contenente i file per il testing,
  - `types.go`, contenente tutti i tipi custom,
  - `readme.md` e `license`.


## Come usare
### Clona il progetto
```
git clone https://github.com/Gabri432/Laplace-Theorem.git
```
### Esegui i test
```
go test [-v] [-run TestFunctionName]
```
### Esegui il codice
```
\...\laplace-theorem> go run .
```

### Esempio
- Esempio 1: Creazione di una nuova matrice dal codice
```
package main

import (
	"fmt"
	laplace "github.com/Gabri432/LaplaceExpansion"
)

func main() {

	row1 := laplace.NewRow([]float64{1, 2, 3, 6})
	row2 := laplace.NewRow([]float64{5, 7, 3, 0})
	row3 := laplace.NewRow([]float64{1, 11, 1, 4})
	row4 := laplace.NewRow([]float64{9, 2, 8, 4})
	matrix := laplace.NewMatrix([]laplace.MatRow{row1, row2, row3, row4})
	fmt.Println(matrix.LaplaceDet())

}

=== Output ===

 120
```

- Esempio 2: Creazione di una nuova matrice usando la funzione "MatrixFromTerminal()"
```

func main() {
  matrix := lapalce.MatrixFromTerminal()
  matrix.Start()
}

=== Output ===

\...\laplace-theorem> go run .
Insert Input:  <<< It will appear a description about how to use the code.
 1) each line creates a row of the Matrix.
 2) each number of the line must be separated by a semicolon;      
 3) once you have finished write END to start running the program.
Example:
 3;2
 1;5
 END    <<< End of description
1;3;3;1 <<< User starts inserting here.
4;2;9;1
5;3;1;7
9;5;6;8
END   <<< User writes 'END' to start calculating the matrix determinant.

Result: -224      <<< Output of the calculation.
```
- Esempio 3: Inserire un input errato mentre si usa "MatrixFromTerminal()" e "(Matrix).Start()".
```
1;3;3;1
4;2;9
5;3;1;7
9;5;6;8
END                       <<< User scrive 'END' per iniziare il calcolo del determinante.
This isn't a square matrix    <<< Messaggio di errore dal codice.
 Retry.
```


## Note ulteriori

- Il programma dovrebbe funzionare correttamente. Tuttavia puoi testarlo usando questo sito web:
```
https://matrixcalc.org/det.html
```
- L' "Espansione di Laplace" non è molto efficiente per matrici grandi. Infatti la sua complessità temporale è O(n!).
- Per cui il codice potrebbe eseguirsi lentamente a seconda della grandezza della matrice.
