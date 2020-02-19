func LastRune(s string) rune {
	m := 0
	ms := []rune(s)
	for range s {

		m++
	}
	return ms[m-1]

}_