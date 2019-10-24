package piscine

func Capitalize(s string) string {
	newstr := []rune(s)
	nb := StrLen(s)
	if newstr[0] >= 'a' && newstr[0] <= 'z' {
		newstr[0] = newstr[0] - 32
	}
	for i := 1; i < nb; i++ {
		if newstr[i] >= 'a' && newstr[i] <= 'z' && (newstr[i-1] < 'a' && newstr[i-1] > 'Z' || newstr[i-1] < 'A' && newstr[i-1] > '9' || newstr[i-1] < '0' || newstr[i-1] > 'z') {
			newstr[i] = newstr[i] - 32
		}
		if newstr[i] >= 'A' && newstr[i] <= 'Z' && (newstr[i-1] >= '0' && newstr[i-1] <= '9' || newstr[i-1] >= 'a' && newstr[i-1] <= 'z' || newstr[i-1] >= 'A' && newstr[i-1] <= 'Z') {
			newstr[i] = newstr[i] + 32
		}
	}
	return string(newstr)
}

func StrLen(s string) int {
	a := 0
	for range s {
		a++
	}
	return a
}
