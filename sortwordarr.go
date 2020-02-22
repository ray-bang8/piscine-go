

func SortWordArr(str []string) {

	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] > str[j] {
				str[i], str[j] = str[j], str[i]
			}
		}
	}

}
