func PrintMemory(arr [10]int) {
	index := 0
	for i := 0; i < len(arr); i++ {
		index++
		if index == 5 || index == 9 {
			fmt.Println()
		}

		if arr[i] != 0 {
			// PrintNbrBase(arr[i],"0123456789ABCDEF")
			dec2hexa(arr[i])
			fmt.Print("00 0000 ")
		} else {
			fmt.Println("0000 0000")
		}

	}

}

func dec2hexa(n int) {
	var arr [100]rune
	i := 0
	for n != 0 {
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
}