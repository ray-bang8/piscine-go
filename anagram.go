import "fmt"

func Anagram(str1, str2 string) bool {

	if CountLen(str1) != CountLen(str2) {
		return false
	}

	newarg := make([]rune, CountLen(str2))
	j := 0
	for _, two := range str2 {
		if two != ' ' {
			newarg[j] = two
			j++
		}
	}

	count := 0
	for _, one := range str1 {
		for j, two := range newarg {
			if one == two {
				count++
				newarg[j] = '#'
				break
			}
		}
	}
	fmt.Println(string(newarg))
	return CountLen(str2) == count

}

func CountLen(str string) int {

	counter := 0

	for _, x := range str {
		if x >= 'a' && x <= 'z' {
			counter++
		}
	}
	return counter
}