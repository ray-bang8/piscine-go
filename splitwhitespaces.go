package piscine

func SplitWhiteSpaces(str string) []string {
	var len, count int
	for i, j := range str {
		if (j == '\n' || j == '\t' || j == ' ') && (i != 0 && str[i-1] != '\n' && str[i-1] != '\t' && str[i-1] != ' ') {
			len++
		}
	}

	var word string
	answer := make([]string, len+1)
	for _, j := range str {
		if j != ' ' && j != '\n' && j != '\t' {
			word += string(j)
		} else if (j == ' ' || j == '\n' || j == '\t') && word != "" {
			answer[count] = word
			word, count = "", count+1
		}
	}
	answer[count] = word //last word
	return answer
}
