func LASTWORD() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		var len, count int
		var word string
		for _, j := range arg {
			if j != ' ' && j != '\n' && j != '\t' {
				word += string(j)
			} else if (j == ' ' || j == '\n' || j == '\t') && word != "" {
				word, len = "", len+1
			}
		}
		if word != "" {
			len++ //last word
		}

		word = ""
		answer := make([]string, len+1)
		for _, j := range arg {
			if j != ' ' && j != '\n' && j != '\t' {
				word += string(j)
			} else if (j == ' ' || j == '\n' || j == '\t') && word != "" {
				answer[count] = word
				word, count = "", count+1
			}
		}
		if word != "" {
			answer[count] = word //last word
		}
		for _, j := range answer[len-1] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	} else {
		z01.PrintRune(10)
	}	
}
___________________
func main() {
	argString := os.Args[1]
	var (
		indexTo, indexFrom int
		slice              string
	)
	lenArgString := len(argString) - 1
	for i := lenArgString; i >= 0; i-- {
		if indexTo == 0 && argString[i] != ' ' {
			indexTo = i + 1
		} else if indexTo > 0 && argString[i] == ' ' {
			indexFrom = i + 1
			break
		}
	}

	slice = arg[index2:index1]
	for _, v := range slice {
		z01.PrintRune(v)
	}

	slice = argString[indexFrom:indexTo]
	for _, v := range slice {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}