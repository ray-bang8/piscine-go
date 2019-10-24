package piscine

func IsUpper(str string) bool {
	red := StrLen(str)
	yellow := []rune(str)
	for i := 0; i < red; i++ {
		if yellow[i] >= 'A' && yellow[i] <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true

}

func StrLen(str string) int {
	j := 0
	for range str {
		j++
	}
	return j
}
