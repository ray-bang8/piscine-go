package main

import (
	"fmt"
	"os"
)

func main() {
	matrix := os.Args[1:]
	if !checkForValidity(matrix) {
		return
	}
	var m [9][9]int
	m = toArray(matrix)

	if !isSudokuSolved(m) {
		fmt.Println("Error")
	}
}

func isSudokuSolved(m [9][9]int) bool {

	row := 0
	col := 0
	isFilled := true

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if m[i][j] == 0 {
				row = i
				col = j
				isFilled = false
				break
			}
		}
		if !isFilled {
			break
		}
	}
	if isFilled {
		for i := 0; i < 9; i++ {
			fmt.Println(m[i])
		}

		return true
	}

	for n := 1; n <= 9; n++ {
		if isAllowed(m, row, col, n) {
			m[row][col] = n
			if isSudokuSolved(m) {
				return true
			}
			m[row][col] = 0
		}
	}
	return false
}

func isAllowed(m [9][9]int, row int, col int, num int) bool {
	for i := 0; i < 9; i++ {
		if m[row][i] == num {
			return false
		}
	}

	for j := 0; j < 9; j++ {
		if m[j][col] == num {
			return false
		}
	}

	sqrt := 3
	rowstart := row - row%sqrt
	colstart := col - col%sqrt

	for i := rowstart; i < rowstart+sqrt; i++ {
		for j := colstart; j < colstart+sqrt; j++ {
			if m[i][j] == num {
				return false
			}
		}
	}
	return true
}

func toArray(m []string) [9][9]int {
	var arr [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if m[i][j] == '.' {
				arr[i][j] = 0
			} else {
				arr[i][j] = runeToInt(rune(m[i][j]))
			}
		}
	}

	return arr
}

func runeToInt(r rune) int {
	c := 0
	for i := '0'; i < r; i++ {
		c++
	}
	return c
}

func intToRune(n int) rune {
	r := '0'
	for i := 0; i < n; i++ {
		r++
	}
	return r
}

func isFilled(s []string) bool {

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s[i][j] == '.' {
				return false
			}
		}
	}
	return true
}

func checkForValidity(s []string) bool {
	if LengOfArray(s) != 9 {
		fmt.Println("Error")
		return false
	}
	for _, word := range s {
		if LengOfString(word) != 9 {
			fmt.Println("Error")
			return false
		}
		for _, letter := range word {
			if letter != '.' && (letter < '1' || letter > '9') || letter == '"' {
				fmt.Println("Error")
				return false
			}
		}
	}
	return true
}

func LengOfArray(s []string) int {
	c := 0
	for range s {
		c++
	}
	return c
}
func LengOfString(s string) int {
	c := 0
	for range s {
		c++
	}
	return c
}
