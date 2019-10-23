package piscine

func LastRune(s string) rune {

	nb := StrLen(s)
	cat := []rune(s)
	return cat[nb-1]
}
func StrLen(s string) int {
	a := 0
	for range s {
		a++
	}
	return a
}
