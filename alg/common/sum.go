package common

// sumTwoNum 求数组中两数之和等于 target 的两个值.
func sumTwoNum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ { // 每两个比较一下.
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
