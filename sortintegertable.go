package piscine

func SortIntegerTable(s []int) {

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if s[i] > s[j] {
				swap(s, i, j)
			}
		}
	}
}
func swap(s []int, i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}
