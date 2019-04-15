func trap(height []int) int {
	total := len(height)

	var ans float64
	for i := 1; i < total-1; i++ {
		var maxLeft, maxRight float64

		for j := i; j >= 0; j-- {
			maxLeft = math.Max(maxLeft, float64(height[j]))
		}

		for j := i; j < total; j++ {
			maxRight = math.Max(maxRight, float64(height[j]))
		}

		ans += math.Min(maxLeft, maxRight) - float64(height[i])

	}
	return int(ans)
}
