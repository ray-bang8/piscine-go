func DOOP() {
	func modulo(a, b int) int {
		return a % b
	}
	
	func divide(a, b int) int {
		if b == 0 {
			return 0
		}
		return a / b
	}
	
	func substract(a, b int) int {
		return a - b
	}
	
	func add(a, b int) int {
		return a + b
	}
	
	func multiply(a, b int) int {
		return a * b
	}
	
	func doop(a int, b int, op string) (int, string) {
		var res int
		switch op {
		case "+":
			res = add(a, b)
	
		case "-":
			res = substract(a, b)
	
		case "*":
			res = multiply(a, b)
	
		case "/":
			if b == 0 {
				return 0, "No division by 0"
			} else {
				res = divide(a, b)
			}
		case "%":
			if b == 0 {
				return 0, "No modulo by 0"
			} else {
				res = modulo(a, b)
			}
		}
		return res, ""
	}
	
	func main() {
	
		args := os.Args[1:]
	
		if len(args) == 0 {
			return
		} else if len(args) > 3 {
			fmt.Println(0)
		} else {
			a, Aerr := strconv.Atoi(args[0])
			if Aerr != nil {
				fmt.Println(0)
			} else {
				op := args[1]
				b, Berr := strconv.Atoi(args[2])
				if Berr != nil {
					fmt.Println(0)
				} else {
					res, err := doop(a, b, op)
					if len(err) > 0 {
						fmt.Println(err)
					} else {
						if a > 0 && b > 0 && res < 0 {
							fmt.Println(0)
						} else if a < 0 && b < 0 && res > 0 {
							fmt.Println(0)
						} else if a < 0 && b > 0 {
							fmt.Println(res)
						} else {
							fmt.Println(res)
						}
					}
				}
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
		} else if str[i] == '+' && i == 0 {
			plus++
		} else if str[i] == '-' && i == 0 {
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