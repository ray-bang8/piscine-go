package piscine

func IsAlpha(str string) bool {
	newstr := []rune(str)
	a := StrLen(str)
	for i := 0; i < a; i++ {
		if (newstr[i] >= 'a' && newstr[i] <= 'z') || (newstr[i] >= 'A' && newstr[i] <= 'Z') || (newstr[i] >= '0' && newstr[i] <= '9') {
			continue
		} else {
			return false
		}
	}
	return true
}

func StrLen(str string) int {
	b := 0
	for range str {
		b++
	}
	return b
}
