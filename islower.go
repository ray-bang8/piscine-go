package piscine

func IsLower(str string) bool {

	boss := []rune(str)
	shake := StrLen(str)
	for i := 0; i < shake; i++ {
		if boss[i] >= 'a' && boss[i] <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}

func StrLen(str string) int {
	me := 0
	for range str {
		me++
	}
	return me
}
