# Laplace-Theorem
A simple program in go that will be able to calculate the matrix determinant by using the Laplace Theorem.
You can check the details on `https://en.wikipedia.org/wiki/Laplace_expansion`.

## Features
- It will calculate square matrices from 2 to 4 rows of size.

## Project Structure
- `Laplace Theorem` (main folder), where the following files are located:
  - `laplace.go`, where you find the main functions of the library,
  - `laplace_test.go`, containing all test functions of the library,
  - `types.go`, containing all the custom types used,
  - `readme.md` and `license`.


## How to use it
### Clone the project
```
git clone https://github.com/Gabri432/Laplace-Theorem.git
```
### Run the tests
```
go test [-v] [-run TestFunctionName]
```

## Notes
- The project does test all the functions for the calculations, however one test is not passed. This one has to test if the determinant of a matrix-6x6 is correct. So if anyone can help me with the code he/she will be really appreciated.
- It doesn't always work with matrix-5x5, as example:
```
1;2;3;4;5
2;3;4;5;1
3;4;5;1;2
4;5;1;2;3
5;1;2;3;4

>>> 25 (it should be 1875)
```
- The project should work properly for matrices-3x3, 4x4 and 5x5. However you can test by using the following website:
```
https://matrixcalc.org/det.html
```
