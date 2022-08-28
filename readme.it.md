# Teorema di Laplace
Un semplice programma scritto in GO che permetterà di calcolare il determinante di una matrice quadrata usando il metodo di Laplace.
Per maggiori dettagli `https://en.wikipedia.org/wiki/Laplace_expansion`.

## Funzionalità
- It will calculate square matrices of any size.

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
```
\...\laplace-theorem> go run .
Insert Input:  <<< Apparirà una descrizione su come inserire la matrice.
 1) each line creates a row of the Matrix.
 2) each number of the line must be separated by a semicolon;      
 3) once you have finished write END to start running the program.
Example:
 3;2
 1;5
 END    <<< Fine della descrizione
        <<< User comnicia a scrivere da qui.
```
- Esempio di input
```
1;3;3;1
4;2;9;1
5;3;1;7
9;5;6;8
END   <<< User scrive 'END' per iniziare il calcolo del determinante della matrice.

Result: -224
```
- Esempio in caso di input errato
```
1;3;3;1
4;2;9
5;3;1;7
9;5;6;8
END                       <<< User scrive 'END' per iniziare il calcolo del determinante della matrice.
This isn't a square matrix    <<< Messaggio di errore dal programma.
 Retry.
```

## Note ulteriori

- Il programma dovrebbe funzionare correttamente. Tuttavia puoi testarlo usando questo sito web:
```
https://matrixcalc.org/det.html
```
- L' "Espansione di Laplace" non è molto efficiente per matrici grandi. Infatti la sua complessità temporale è O(n!).
- Per cui il codice potrebbe eseguirsi lentamente a seconda della grandezza della matrice.
