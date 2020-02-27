package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if ArrayStrLength(os.Args) == 1 {
		PrintStr("\n")
		return
	}

	// arr := make([]byte, 2048)
	var arr [2048]byte
	p := &arr[0]
	arg := []rune(os.Args[1])

	i := 0
	j := 0

	for j < 2048 {
		for i < ArrayRuneLength(arg) {
			switch arg[i] {
			case '>':
				j++
				p = &arr[j]
			case '<':
				j--
				p = &arr[j]
			case '+':
				*p++
			case '-':
				*p--
			case '.':
				PrintStr(string(*p))
			case '[':
				cont := 0

				if *p == 0 {

					for cont >= 0 {
						i++
						if arg[i] == ']' {
							cont--
						} else if arg[i] == '[' {
							cont++
						}
					}
				}
			case ']':
				cont := 0
				if *p != 0 {
					for cont >= 0 {
						i--
						if arg[i] == '[' {
							cont--
						} else if arg[i] == ']' {
							cont++
						}
					}
				}
			}

			i++
		}
		j = 2048
	}
}
func ArrayStrLength(str []string) int {

	length := 0

	for range str {
		length++
	}

	return length
}
func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
	// z01.PrintRune('\n')
}
func ArrayRuneLength(str []rune) int {

	length := 0

	for range str {
		length++
	}

	return length
}
