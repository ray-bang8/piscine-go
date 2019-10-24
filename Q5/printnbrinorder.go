package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else {
		s := []int{}
		s = fromNumToArr(n, s)
		s = SortInteger(s)
		for _, i := range s {
			z01.PrintRune(rune(i) + 48)
		}
	}

}
func fromNumToArr(n int, s []int) []int {
	if n == 0 {
		return s
	} else {
		s = append(s, n%10)
		s = fromNumToArr((n / 10), s)
	}
	return s
}
func SortInteger(chucha []int) []int {
	pshan := 0
	for range chucha {
		pshan++
	}
	for i := 0; i < pshan; i++ {
		for j := i; j < pshan; j++ {
			if chucha[i] > chucha[j] {
				swap(chucha, i, j)
			}
		}
	}
	return chucha
}
func swap(chucha []int, i, j int) {
	tmp := chucha[i]
	chucha[i] = chucha[j]
	chucha[j] = tmp
}

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}
