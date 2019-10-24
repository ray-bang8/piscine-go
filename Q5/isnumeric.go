package piscine

func IsNumeric(str string) bool {

	hole := []rune(str)
	Batyr := StrLen(str)
	for i := 0; i < Batyr; i++ {
		if hole[i] >= '0' && hole[i] <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}

func StrLen(str string) int {
	a := 0
	for range str {
		a++
	}
	return a
}
