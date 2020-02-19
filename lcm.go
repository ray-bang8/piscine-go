func Lcm(first, second int) int {
	n1 := first
	n2 := second
	var res int
	i := 1
	for i >= 1 {
		temp := i * n2
		if temp%n1 == 0 && temp%n2 == 0 {
			res = temp
			break
		}
		i++

	}
	return res

}