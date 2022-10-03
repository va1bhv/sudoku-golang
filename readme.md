# Sudoku board solver
## Implemented in Golang
---
###Usage
It accepts a (9, 9) board in a text document, prints the unsolved and solved boards to stdout and then asks if the solved board should be saved.

```
go run solver.go board.txt // To run directly from go
go build -o solver solver.go; ./solver board.txt
```

The board in the txt file should be in the following format:
```
0, 1, 5, 0, 0, 4, 0, 0, 0
6, 0, 0, 1, 0, 0, 2, 0, 0
0, 0, 2, 6, 9, 8, 0, 5, 0
2, 5, 0, 3, 0, 1, 0, 0, 0
0, 0, 0, 0, 0, 6, 3, 0, 2
4, 0, 0, 0, 0, 0, 0, 6, 1
8, 2, 4, 5, 1, 0, 0, 3, 0
5, 0, 0, 4, 0, 0, 0, 1, 0
0, 0, 1, 0, 0, 3, 4, 0, 0
```

The output for the above board is:
```
Unsolved board:
0 1 5  | 0 0 4  | 0 0 0
6 0 0  | 1 0 0  | 2 0 0
0 0 2  | 6 9 8  | 0 5 0
- - -    - - -    - - -
2 5 0  | 3 0 1  | 0 0 0
0 0 0  | 0 0 6  | 3 0 2
4 0 0  | 0 0 0  | 0 6 1
- - -    - - -    - - -
8 2 4  | 5 1 0  | 0 3 0
5 0 0  | 4 0 0  | 0 1 0
0 0 1  | 0 0 3  | 4 0 0

Solved board:
3 1 5  | 2 3 4  | 6 7 8
6 4 8  | 1 7 5  | 2 9 3
3 7 2  | 6 9 8  | 1 5 4
- - -    - - -    - - -
2 5 6  | 3 4 1  | 7 8 9
1 8 7  | 9 5 6  | 3 4 2
4 3 9  | 7 8 2  | 5 6 1
- - -    - - -    - - -
8 2 4  | 5 1 7  | 9 3 6
5 6 3  | 4 2 9  | 8 1 7
7 9 1  | 8 6 3  | 4 2 5

Print to file?(y/N)
y
Wrote 252 bytes.
```

The the solved board is saved with the same name as the unsolved board, with `Solved_` prefixed. 
The solved output is as follows:
```
3, 1, 5, 2, 3, 4, 6, 7, 8, 
6, 4, 8, 1, 7, 5, 2, 9, 3, 
3, 7, 2, 6, 9, 8, 1, 5, 4, 
2, 5, 6, 3, 4, 1, 7, 8, 9, 
1, 8, 7, 9, 5, 6, 3, 4, 2, 
4, 3, 9, 7, 8, 2, 5, 6, 1, 
8, 2, 4, 5, 1, 7, 9, 3, 6, 
5, 6, 3, 4, 2, 9, 8, 1, 7, 
7, 9, 1, 8, 6, 3, 4, 2, 5, 
```