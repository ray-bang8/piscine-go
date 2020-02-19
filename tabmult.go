func tabmult() {
    if len(os.Args[1:]) != 1 {
        z01.PrintRune('\n')
        return
    } else {
        a, _ := strconv.Atoi(os.Args[1])
        b := os.Args[1]
        if a < 0 {
            z01.PrintRune('\n')
            return
        }
        for i := 1; i <= 9; i++ {
            res := i * a
            z01.PrintRune(rune(i + 48))
            z01.PrintRune(' ')
            z01.PrintRune('x')
            z01.PrintRune(' ')
            for _, v := range b {
                z01.PrintRune(v)
            }
            z01.PrintRune(' ')
            z01.PrintRune('=')
            z01.PrintRune(' ')
            str := strconv.Itoa(res)
            for _, v := range str {
                z01.PrintRune(v)
            }
            z01.PrintRune('\n')

        }
    }
}
func Atoi(s string) int {
	len := 0
	for range s {
		len++
	}
	plus := 0
	minus := 0
	final := 0
	count := 0
	str := []rune(s)
	for i := 0; i < len; i++ {
		if str[i] >= '0' && str[i] <= '9' {
			count = 0
			for j := '0'; j < str[i]; j++ {
				count++
			}
			final = final*10 + count
		} else if str[i] == '-' && str[i] == 0 {
			minus++
		} else if str[i] == '+' && str[i] == 0 {
			plus++
		} else {
			return 0
		}

	}
	if minus == 1 {
		final = -final
	}
	return final
}