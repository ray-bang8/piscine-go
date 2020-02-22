

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

