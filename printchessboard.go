func printchessboard() {
	if len(os.Args[1:]) != 2 {
		PrintErr()
		return
	} else {
		col, e := strconv.Atoi(os.Args[1])
		line, er := strconv.Atoi(os.Args[2])
		if e != nil || er != nil {
			PrintErr()
			return
		}

		for l := 1; l <= line; l++ {
			for c := 1; c <= col; c++ {
				if (isOdd(c) && isOdd(l)) || (!isOdd(c) && !isOdd(l)) {
					z01.PrintRune('#')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}

	}
}
func isOdd(n int) bool {
	if n%2 == 0 {
		return false
	}
	return true
}