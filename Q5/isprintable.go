package piscine

func IsPrintable(str string) bool {
	a := StrLen(str)
	b := []rune(str)
	for i := 0; i < a; i++ {
		if b[i] >= 32 && b[i] <= 126 {
			continue
		} else {
			return false
		}
	}
	return true
}

func StrLen(str string) int {
	c := 0
	for range str {
		c++
	}
	return c
}
