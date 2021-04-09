func twoSum(nums []int, target int) []int {
	for i1, v1 := range nums {
		for i2, v2 := range nums {
			if i1 == i2 {
				continue
			} else if (v1 + v2) == target {
				return []int{i1, i2}
			}
		}
	}
	return []int{}
}
