package piscine

func TrimAtoi(s string) int {
	nb := S(s)
	meow := []rune(s)
	s = ""
	final := 0
	plus := 0
	minus := 0
	num := 0
	for i := 0; i < nb; i++ {
		if meow[i] >= '0' && meow[i] <= '9' {
			num = 0
			for j := '0'; j < meow[i]; j++ {
				num++
			}
			final = final*10 + num
		} else if meow[i] == '+' {
			plus++
		} else if meow[i] == '-' {
			minus++

		} else {
		}
	}
	if minus == 1 {
		final = -final
	}
	return final
}

func S(s string) int {
	a := 0
	for range s {
		a++
	}
	return a
}
