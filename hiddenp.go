


func NIDENP() {
	if len(os.Args) == 3 {
		answer := ""
		s1 := []rune(os.Args[1])
		s2 := []rune(os.Args[2])
		j := 0

		for i := 0; i < len(s1); i++ {

			for ; j < len(s2); j++ {
				if s1[i] == s2[j] {
					answer += string(s1[i])
					j++
					break
				}
			}
		}
		if answer == string(s1) {
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
		}
		z01.PrintRune(10)
	} else {
		z01.PrintRune(10)
	}
}
__________________________________
func HiddenP(s1, s2 string) bool {

	for _, c := range s1 {
		x := indexOfRune(s2, c)

		if x < 0 {
			return false
		}

		// fmt.Println(s2)

		s2 = s2[x+1:]
	}

	return true
}

func indexOfRune(s string, r rune) int {

	for i, c := range s {
		if c == r {
			return i
		}
	}

	return -1
}