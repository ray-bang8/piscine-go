package piscine

func ConcatParams(args []string) string {
	result := ""
	for i, c := range args {
		if i == 0 {
			result = c
			continue
		}
		result = result + "\n" + c
	}
	return result
}
