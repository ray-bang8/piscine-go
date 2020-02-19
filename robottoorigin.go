func robottoorigin() {
	args := os.Args[1:]
	var UD int
	var RL int
	var f string
	for _, v := range args[0] {
		if v == 'U' {
			UD++
		} else if v == 'D' {
			UD--
		} else if v == 'R' {
			RL++
		} else if v == 'L' {
			RL--
		}

		if UD == 0 && RL == 0 {
			f = "true"
		} else {
			f = "false"
		}
	}

	for _, v := range f {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}