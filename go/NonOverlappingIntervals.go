/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func eraseOverlapIntervals(intervals []Interval) int {
    end := -1 << 31
    res:=0
    if len(intervals) == 1 {
        return 0
    }
    
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].End < intervals[j].End
    })
    
    for _,v := range intervals {
        if v.Start >= end {
            end = v.End
        } else {
            res++
        }
    }
    return res
}