package main

import (
	"errors"
	"fmt"
)

const size = 9

type sudoku [size][size]int8

type sudokuCheck [size][size]bool

func generateSudoku() sudoku {
	s := sudoku{}

	for i := 0; i < size; i++ {
		s[0][i] = int8(i + 1)
	}

	for i := 1; i < size; i++ {
		start := (s[i-1][0] + 3) % 9

		if i == 3 {
			start++
		}

		if i == 6 {
			start++
		}

		for j := 0; j < size; j++ {
			cell := (start + int8(j)) % 9

			if cell == 0 {
				cell = 9
			}

			s[i][j] = cell
		}
	}

	return s
}

func (s *sudoku) set(row, col, dig int8) error {
	if row < 0 || row > size-1 {
		return errors.New("row out of boundary")
	}

	if col < 0 || col > size-1 {
		return errors.New("column out of boundary")
	}

	s[row][col] = dig
	return nil
}

func (s *sudoku) clear(row, col int8) error {
	if row < 0 || row > size-1 {
		return errors.New("row out of boundary")
	}

	if col < 0 || col > size-1 {
		return errors.New("column out of boundary")
	}

	s[row][col] = 0
	return nil
}

func (s *sudoku) print() {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(s[i][j])
		}

		fmt.Println()
	}

	fmt.Println("Please, set the cell: ")
}

func main() {
	newSudoku := generateSudoku()

	newSudoku.print()
}
