func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var newSlice []int
	newSlice = append(nums1, nums2...)

	if len(newSlice) == 0 {
		return float64(0)
	}

	sort.Ints(newSlice)
	length := len(newSlice)
	if length%2 == 0 {
		return float64(newSlice[length/2-1] + newSlice[length/2]) / 2
	} else {
		return float64(newSlice[length/2])
	}
}
