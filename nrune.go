package piscine

func NRune(s string, n int) rune {
	cox := StrLen(s)
	shmox := []rune(s)
	if n > 0 && n <= cox {
		return shmox[n-1]
	}
	return 0
}

func StrLen(str string) int {
	a := 0
	for range str {
		a++
	}
	return a
}
