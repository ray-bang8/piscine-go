
func ItoaBase(value, base int) string {
	res := ""
	for value > 0 {
		res = string(value%base+48) + res
		value /= base
	}
	return res

}

