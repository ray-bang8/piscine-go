package piscine

func Join(s []string, m string) string {
	nb := M(s)
	var str string
	for i, v := range s {
		str = str + v
		if i != nb-1 {
			str += m
		}
	}
	return str
}

func M(s []string) int {
	m := 0
	for range s {
		m++
	}
	return m
}
