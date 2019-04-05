
func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] >= nums[high] {
			index := binarySearch(nums, target, low, mid)
			if index != -1 {
				return index
			}
			low = mid + 1
		} else {
			index := binarySearch(nums, target, mid, high)
			if index != -1 {
				return index
			}
			high = mid - 1
		}
	}
	return -1
}

func binarySearch(nums []int, target, low, high int) int {
	for low <= high {
		mid := low + (high-low)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			high = mid - 1
		} else{
			low = mid + 1
		}
	}
	return -1
}
