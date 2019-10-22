package piscine

func AlphaCount(str string) int {
	b := 0
	for _, babe := range str {
		if (babe >= 'a' && babe <= 'z') || (babe >= 'A' && babe <= 'Z') {
			b++
		}
	}
	return b
}
