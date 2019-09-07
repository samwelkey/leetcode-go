func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	left := make([]int, 0)
	right := make([]int, len(nums))

	for i, num := range nums {
		if i%k == 0 {
			left = append(left, num)
			continue
		}
		if left[i-1] > num {
			left = append(left, left[i-1])
		} else {
			left = append(left, num)
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i%k == k-1 {
			right[i] = nums[i]
			continue
		}
		if i+1 < len(nums) {
			if nums[i] < right[i+1] {
				right[i] = right[i+1]
			} else {
				right[i] = nums[i]
			}
		} else {
			right[i] = nums[i]
		}
	}

	res := make([]int, 0)
	for i := 0; i+k-1 < len(nums); i++ {
		if right[i] > left[i+k-1] {
			res = append(res, right[i])
		} else {
			res = append(res, left[i+k-1])
		}
	}
	return res
}
