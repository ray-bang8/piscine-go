func ROT13() {
	arg := os.Args[1:]
	if len(arg) == 1 {
		massiv := []rune(arg[0])
		for i := 0; i < len(massiv); i++ {
			if massiv[i] >= 'a' && massiv[i] <= 'z' {
				if massiv[i] > 'm' {
					massiv[i] = massiv[i] - 13
				} else {
					massiv[i] = massiv[i] + 13
				}
			} else if massiv[i] >= 'A' && massiv[i] <= 'Z' {
				if massiv[i] > 'M' {
					massiv[i] = massiv[i] - 13
				} else {
					massiv[i] = massiv[i] + 13
				}
			}
		}
		for _, v := range massiv {
			z01.PrintRune(v)
		}

	}
	z01.PrintRune('\n')
}