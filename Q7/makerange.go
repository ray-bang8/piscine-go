package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	answer := make([]int, max-min)
	for i := range make {
		make[i] = i + min
	}
	return answer
}
