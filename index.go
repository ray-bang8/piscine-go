package piscine

func Index(s string, toFind string) int {
	newstr := []rune(s)
	nb := StrLen(s)
	strFind := []rune(toFind)
	nbFind := StrLen(toFind)
	match := 0
	index := 0
	if nbFind == 0 {
		return 0
	}
	for i := 0; i < nbFind; i++ {
		for j := index; j < nb; j++ {
			match = 0
			if strFind[i] == newstr[j] {
				match = 1
				index = j
				break
			}

		}
		if match == 0 {
			return -1
		}
	}
	return index - nbFind + 1
}
