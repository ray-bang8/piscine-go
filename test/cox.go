package basicatoi

func BasicAtoi(s string) int {
	res := 0
	for _, val := range s {
		a := 0
		for i := '1'; i <= val; i++ {
			a++
		}
		res = res*10 + a
	}
	return res
}