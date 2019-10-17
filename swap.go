package piscine

func Swap(a *int, b *int) {
	var x int = *a
	*a = *b
	*b = x
}
