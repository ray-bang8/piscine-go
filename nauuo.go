package main

func Nauuo(a, b, c int) string {
	str := ""
	if a > (b + c) {
		str = "+"
	} else if b > (a + c) {
		str = "-"
	} else if a == b && c == 0 {
		str = "0"
	} else {
		str = "?"
	}

	return str
}
