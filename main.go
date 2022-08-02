package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	size     = 9
	mixCount = 9999
)

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

	s.mix()

	return s
}

func (s *sudoku) selectDifficultyLevel(level int) *sudoku {
	n := 0.0

	switch level {
	case 1:
		n = 0.5
	case 2:
		n = 0.65
	case 3:
		n = 0.8
	}

	s2 := s

	cellsCount := int(81 * n)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < cellsCount; i++ {
		s2[rand.Intn(9)][rand.Intn(9)] = 0
	}

	return s2
}

func (s *sudoku) mix() {
	i := 0

	for i < mixCount {
		s.transposing()
		s.swapRowsSmall()
		s.swapColumnsSmall()
		s.transposing()

		s.swapRowsArea()
		s.transposing()
		s.swapColumnsArea()
		s.transposing()

		i++
	}
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

func (s *sudoku) getRandomArea(a int) int {
	var area int

	switch a {
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
	rand.Seed(time.Now().UnixNano())

	area := s.getRandomArea(rand.Intn(3))

	row1, row2 := rand.Intn(3)+area, rand.Intn(3)+area

	for i := 0; i < size; i++ {
		s[row1][i], s[row2][i] = s[row2][i], s[row1][i]
	}
}

func (s *sudoku) swapColumnsSmall() {
	rand.Seed(time.Now().UnixNano())

	area := s.getRandomArea(rand.Intn(3))

	col1, col2 := rand.Intn(3)+area, rand.Intn(3)+area

	for i := 0; i < size; i++ {
		s[i][col1], s[i][col2] = s[i][col2], s[i][col1]
	}
}

func (s *sudoku) swapRowsArea() {
	rand.Seed(time.Now().UnixNano())

	area1, area2 := s.getRandomArea(rand.Intn(3)), s.getRandomArea(rand.Intn(3))

	if area1 == area2 {
		return
	}

	j := 0
	for j < 3 {
		for i := 0; i < size; i++ {
			s[area1+j][i], s[area2+j][i] = s[area2+j][i], s[area1+j][i]
		}

		j++
	}
}

func (s *sudoku) swapColumnsArea() {
	rand.Seed(time.Now().UnixNano())

	area1, area2 := s.getRandomArea(rand.Intn(3)), s.getRandomArea(rand.Intn(3))

	if area1 == area2 {
		return
	}

	j := 0
	for j < 3 {
		for i := 0; i < size; i++ {
			s[i][area1+j], s[i][area2+j] = s[i][area2+j], s[i][area1+j]
		}

		j++
	}
}

func (s *sudoku) setCell(row, col, dig int8) error {
	if row < 0 || row > size-1 {
		return errors.New("row out of boundary")
	}

	if col < 0 || col > size-1 {
		return errors.New("column out of boundary")
	}

	s[row][col] = dig
	return nil
}

func (s *sudoku) clearCell(row, col int8) error {
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
			if s[i][j] == 0 {
				fmt.Printf("%-3v", "-")
			} else {
				fmt.Printf("%-3v", s[i][j])
			}
		}

		fmt.Println()
	}

	fmt.Println("Please, set the cell: ")
}

func main() {
	newSudoku := generateSudoku()
	newSudoku.print()

	game := newSudoku.selectDifficultyLevel(3)
	game.print()
}
