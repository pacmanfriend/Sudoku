package main

import "errors"

const size = 9

type sudoku [size][size]int8

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

func main() {

}
