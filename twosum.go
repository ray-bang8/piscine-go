
func TwoSum(nums []int, target int) []int {

	a := []int{0, 0}
	for i, v := range nums {
		for j, g := range nums {
			if v+g == target && i != j {
				a[0] = j
				a[1] = i
				return reverse(a)
			}
		}
	}
	return nil

}

func reverse(a []int) []int {
	b := []int{0, 0}
	b[0] = a[1]
	b[1] = a[0]
	return b
}

// func main() {
// 	case1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	out := TwoSum(case1, 6)
// 	fmt.Println(out)
// }
