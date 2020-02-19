func TwoSum(nums []int, target int) []int {

	a := []int{0, 0}
	for i, v := range nums {
		for j, g := range nums {
			if v+g == target && i != j {
				a[0] = j
				a[1] = i
				return a
			}
		}
	}
	return nil

}