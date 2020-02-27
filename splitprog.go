

func split(str, charset string) []string {
	count := 1
	n := 0
	if charset == "" {
		empty := make([]string, len(str))
		i := 0
		for v := range str {
			empty[i] = string(str[v])
			i++
		}
		return empty

	}
	for i := 0; i <= len(str)-len(charset); i++ {
		if str[i:i+len(charset)] == charset {
			i = i + len(charset) - 1
			count++
		}
	}
	output := make([]string, count)
	countEl := 0
	for i := 0; i < len(str); i++ {

		if count == countEl+1 {
			output[countEl] = str[n:]
			break
		}
		if str[i:i+len(charset)] == charset {
			output[countEl] = str[n:i]
			i = i + len(charset) - 1
			n = i + 1
			countEl++
		}
	}
	return output
}

// func SPLITMAIN() {
// 	args := os.Args[1:]
// 	str := args[0]
// 	charset := args[1]
// 	if len(args) == 2 {
// 		fmt.Println(split(str, charset))
// 	} else {
// 		fmt.Println()
// 	}
// }
