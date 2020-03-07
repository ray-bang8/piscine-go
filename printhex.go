
func PrintHex() {
	arg := os.Args
	if len(arg) == 2 {
		n := Atoi(arg[1])
		if n == 0 {
			z01.PrintRune(0)
		}
		var arr [70]rune
		i := 0
		for n != 0 {
			if n < 0 {
				z01.PrintRune('-')
				n = n / -1
			}
			temp := n % 16
			if temp < 10 {
				arr[i] = rune(temp + 48)
				i++
			} else {
				arr[i] = rune(temp + 87)
				i++
			}
			n = n / 16
		}
		for j := i - 1; j >= 0; j-- {
			z01.PrintRune(arr[j])
		}
		z01.PrintRune('\n')

	} else {
		z01.PrintRune('\n')
	}

}
func Atoi(s string) int {
	len := 0
	for range s {
		len++
	}
	meow := []rune(s)
	final := 0
	plus := 0
	minus := 0
	num := 0
	for i := 0; i < len; i++ {
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
____________________________________-
func PRINTHEX() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	ans := toHex(n)

	for _, c := range ans {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func toHex(n int) string {

	if n == 0 {
		return "0"
	}

	base := "0123456789abcdef"

	res := ""

	for n != 0 {
		res = string(base[n%len(base)]) + res
		n /= len(base)
	}

	return res
}