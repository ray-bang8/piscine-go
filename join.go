package piscine

func Join(strs []string, sep string) string {
	a := ""
	for _, b := range strs {
		a += b + sep
	}
	return a
}
