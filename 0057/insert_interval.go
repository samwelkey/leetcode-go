/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func insert(intervals []Interval, newInterval Interval) []Interval {
    intervals = append(intervals,newInterval)
    return merge(intervals)
}


func merge(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start < intervals[j].Start {
			return true
		}
		if intervals[i].Start == intervals[j].Start && intervals[i].End <= intervals[j].End {
			return true
		}
		return false
	})
	var ans []Interval
	ans = append(ans, intervals[0])

	var i = 1

	for i < len(intervals) {

		last := &ans[len(ans)-1]
		if last.End < intervals[i].Start {
			ans = append(ans, intervals[i])
			i++
			continue
		}

		if last.End >= intervals[i].Start && last.End < intervals[i].End {
			last.End = intervals[i].End
		}
		i++
	}
	return ans
}
