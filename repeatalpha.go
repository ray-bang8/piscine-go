func REPEATALPHA() {
	arg := os.Args
	count := 0
	if len(arg) == 2 {

		arr := []rune(arg[1])
		for i := 0; i < len(arr); i++ {

			if arr[i] >= 'a' && arr[i] <= 'z' || arr[i] >= 'A' && arr[i] <= 'Z' {

				if arr[i] >= 'a' && arr[i] <= 'z' {
					count = int(arr[i]-'a') + 1
				}

				if arr[i] >= 'A' && arr[i] <= 'Z' {
					count = int(arr[i]-'A') + 1
				}
				for ; count > 0; count-- {
					z01.PrintRune(arr[i])
				}
			}
		}
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('\n')
	}
}