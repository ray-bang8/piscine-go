func BRACKEts() {
	if len(os.Args) == 1 {
		z01.PrintRune('\n')
		return
	}

	for i := 1; i < len(os.Args); i++ {
		arr := []rune(os.Args[i])
		var stack []rune

		if os.Args[i] == "" {
			PrintStr("OK")
			z01.PrintRune('\n')
		} else {
			for j := 0; j < StrLen(os.Args[i]); j++ {
				if arr[j] == '(' || arr[j] == '[' || arr[j] == '{' {
					stack = append(stack, arr[j])
				} else {
					switch arr[j] {
					case ')':
						if stack[len(stack)-1] == '(' {
							stack = stack[0 : len(stack)-1]

						}
					case ']':
						if stack[len(stack)-1] == '[' {
							stack = stack[0 : len(stack)-1]

						}
					case '}':
						if stack[len(stack)-1] == '{' {
							stack = stack[0 : len(stack)-1]
						}
					}
				}
			}

			if len(stack) == 0 {
				PrintStr("OK")
				z01.PrintRune('\n')
			} else {
				PrintStr("Error")
				z01.PrintRune('\n')
			}

		}
	}
}
func PrintStr(str string) {

	stringi := []rune(str)

	for _, letter := range stringi {

		z01.PrintRune(letter)
	}

}
func StrLen(str string) int {
	m := 0
	for range str {
		m++
	}
	return m
}
_____________________-
29)func ListSize(l *List) int {

	n := l.Head
	size := 0
	for n != nil {
		size++
		n = n.Next
	}
	return size
}