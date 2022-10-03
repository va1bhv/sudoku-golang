package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	board [9][9]int
}

type Pos struct {
	x int
	y int
}

func draw_board(board Board) {
	for i, row := range board.board {
		for j, value := range row {
			fmt.Printf("%d ", value)
			if j == 2 || j == 5 {
				fmt.Print(" | ")
			}
		}
		if i == 2 || i == 5 {
			fmt.Print("\n- - -    - - -    - - -")
		}
		fmt.Println()

	}
}

func get_empty_pos(board Board) Pos {
	for i, row := range board.board {
		for j, value := range row {
			if value == 0 {
				return Pos{i, j}
			}
		}
	}
	return Pos{-1, -1}
}

func is_valid_value(board Board, value int, pos Pos) bool {
	for i := 0; i < 9; i++ {
		// Check within row
		if value == board.board[pos.x][i] && i != pos.x {
			return false
		}
		// Check within column
		if value == board.board[i][pos.y] && i != pos.y {
			return false
		}
	}
	// Check within the box
	box_x := pos.x / 3
	box_y := pos.y / 3
	for i := box_x * 3; i < box_x*3+3; i++ {
		for j := box_y * 3; j < box_y*3+3; j++ {
			if value == board.board[i][j] && i != pos.x && j != pos.y {
				return false
			}
		}
	}
	return true
}

func solve(board *Board) bool {
	empty_pos := get_empty_pos(*board)
	if empty_pos.x == -1 {
		return true
	}
	for i := 1; i <= 9; i++ {
		if is_valid_value(*board, i, empty_pos) {
			board.board[empty_pos.x][empty_pos.y] = i
			if solve(board) {
				return true
			}
			board.board[empty_pos.x][empty_pos.y] = 0
		}
	}
	return false
}

func file_reader() string {
	// First element in os.Args is always the program name,
	// So we need at least 2 arguments to have a file name argument.
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	// Data is the file content, we can use it
	return string(data)
}

func get_board(board_str string) Board {
	var board Board
	var err error
	for i, line := range strings.Split(board_str, "\n") {
		line = strings.Replace(line, "\r", "", 1)
		for j, value := range strings.Split(line, ", ") {
			board.board[i][j], err = strconv.Atoi(value)
			check(err)
		}
	}
	return board
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func board_to_String(board Board) string {
	var board_str []string
	for _, row := range board.board {
		for _, value := range row {
			board_str = append(board_str, strconv.Itoa(value), ", ")
		}
		board_str = append(board_str, "\n")
	}
	return strings.Join(board_str, "")
}

func write_to_file(board Board, filename string) {
	file, err := os.Create(filename)
	data := []byte(board_to_String(board))
	check(err)
	defer file.Close()
	num_bytes, err := file.Write(data)
	check(err)
	fmt.Printf("Wrote %d bytes.\n", num_bytes)
}

func main() {
	board := get_board(file_reader())
	fmt.Println("Unsolved board:")
	draw_board(board)
	solve(&board)
	fmt.Println("\nSolved board:")
	draw_board(board)
	ans := "N"
	fmt.Println("\nPrint to file?(y/N)")
	fmt.Scanln(&ans)
	if ans == "y" {
		write_to_file(board, "Solved_"+os.Args[1])
	}
}
