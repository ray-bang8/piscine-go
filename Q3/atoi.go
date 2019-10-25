package piscine

func Atoi(s string) int {
	nb := S(s)
	meow := []rune(s)
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
		} else if meow[i] == '+' && i == 0 {
			plus++
		} else if meow[i] == '-' && i == 0 {
			minus++

		} else {
			return 0
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
