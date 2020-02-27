

func AtoiBase(s string, base string) int {
	noValid := ""

	if base == "0123456789" {

		return At(s)
	}

	if len(base) < 2 {
		noValid = "NV"
	} else {
		for i := 0; i < len(base)-1; i++ {
			for j := i + 1; j < len(base); j++ {
				if base[i] == base[j] ||
					base[i] == '+' || base[i] == '-' ||
					base[j] == '+' || base[j] == '-' {
					noValid = "NV"
				}
			}
		}
	}

	if noValid == "NV" {
		return 0
	}

	return PrintNumberAtoiBase(s, base)

}
func PrintNumberAtoiBase(nbr string, base string) int {

	arrayBase := []rune(base)
	arrayNbr := []rune(nbr)
	result := 0

	for i := 0; i < len(nbr); i++ {
		for j := 0; j < len(base); j++ {

			if arrayNbr[i] == arrayBase[j] {

				result += j * (Power(len(base), len(nbr)-i-1))
			}
		}
	}

	return result

}

//Power returns the result of n^m
func Power(n int, m int) int {
	helper := n
	if m == 0 {
		n = 1
	} else {
		for i := 1; i < m; i++ {
			n *= helper
		}
	}
	return n
}
func At(s string) int {
	min := 0
	plu := 0
	sum := 0
	for i, v := range s {
		if s[i] >= '0' && s[i] <= '9' {
			num := 0
			for j := '0'; j < v; j++ {
				num++
			}
			sum = sum*10 + num
		} else if s[i] == '-' && i == 0 {
			min++
		} else if s[i] == '+' && i == 0 {
			plu++
		} else {
			return 0
		}
	}
	if min == 1 {
		sum = -sum
	}
	if min != 1 && sum < 0 {
		return 0
	}
	if min == 1 && sum > 0 {
		return 0
	}
	return sum

}
