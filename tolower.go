package piscine

func ToLower(s string) string {
	geralt := []rune(s)
	for i := 0; i < len(s); i++ {
		if geralt[i] >= 'A' && geralt[i] <= 'Z' {
			geralt[i] = geralt[i] + 32
		}
	}
	return string(geralt)
}
