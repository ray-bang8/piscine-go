package piscine

func IsAlpha(str string) bool {
	newstr := []rune(str)
	for i := 0; i < len(str); i++ {
		if (newstr[i] >= 'a' && newstr[i] <= 'z') || (newstr[i] >= 'A' && newstr[i] <= 'Z') || (newstr[i] >= '0' && newstr[i] <= '9') {
			continue
		} else {
			return false
		}
	}
	return true
}
