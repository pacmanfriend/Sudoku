package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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

func (s *sudoku) transposing() {
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			if i != j {
				s[i][j], s[j][i] = s[j][i], s[i][j]
			}
		}
	}
}

func (s *sudoku) getRandomArea() int {
	rand.Seed(time.Now().UnixNano())

	var area int

	switch rand.Intn(2) {
	case 0:
		area = 0
	case 1:
		area = 3
	default:
		area = 6
	}

	return area
}

func (s *sudoku) swapRowsSmall() {
	area := s.getRandomArea()

	row1, row2 := rand.Intn(3)+area, rand.Intn(3)+area

	for i := 0; i < size; i++ {
		s[row1][i], s[row2][i] = s[row2][i], s[row1][i]
	}
}

func (s *sudoku) swapColumnsSmall() {
	area := s.getRandomArea()

	col1, col2 := rand.Intn(3)+area, rand.Intn(3)+area

	for i := 0; i < size; i++ {
		s[i][col1], s[i][col2] = s[i][col2], s[i][col1]
	}
}

func (s *sudoku) swapRowsArea() {
	area1, area2 := s.getRandomArea(), s.getRandomArea()

	if area1 == area2 {
		return
	}

}

func (s *sudoku) swapColumnsArea() {

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
			fmt.Printf("%-3v", s[i][j])
		}

		fmt.Println()
	}

	fmt.Println("Please, set the cell: ")
}

func main() {
	newSudoku := generateSudoku()

	newSudoku.print()

	//newSudoku.transposing()
	newSudoku.swapRowsSmall()

	newSudoku.print()
}
