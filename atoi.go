package student

// import "fmt"

func Atoi(s string) int {
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
