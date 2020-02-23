


func Itoa(nbr int) string {
	result := ""
	t := 1

	if nbr < 0 {
		result += "-"
		t = -1
	}
	if nbr != 0 {
		q := (nbr / 10) * t
		if q != 0 {
			Itoa(q)
		}
		d := ((nbr % 10) * t) + '0'
		result += string(rune(d))
	} else {
		result += "0"
	}

	return result
}
____________________________
func Itoa(n int) string { // better
	res := ""
	t := 1
	if n < 0 {
		res += "-"
		t = -1
	}
	n = n * t

	for n > 0 {
		res = res + string(n%10+48)
		n /= 10
	}
	return res
}
__________________________________________
func Itoa(n int) string {
	count := 0
	for n > 0 {
		count++
		n = n/10
	}
	a := make([]rune, count)
	i := 0
	for n > 0 {
		a[i] = n%10
		n = n/10
	}
	for 
}