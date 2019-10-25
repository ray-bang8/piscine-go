package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	Arr1 := make([]int, max-min)
	for i := range Arr1 {
		Arr1[i] = i + min
	}
	return Arr1
}
