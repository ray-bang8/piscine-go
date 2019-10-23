package piscine

func ToUpper(s string) string {
	cox := []rune(s)
	for i := 0; i < len(s); i++ {
		if cox[i] >= 'a' && cox[i] <= 'z' {
			cox[i] = cox[i] - 32
		}
	}
	return string(cox)
}
