package ltcode

func maxProduct(nums []int) int {
	minimum, maximum, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			minimum, maximum = maximum, minimum
		}
		maximum = max(nums[i], nums[i]*maximum)
		minimum = min(nums[i], nums[i]*minimum) //lhcerr: max
		res = max(res, maximum)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
