func merge(intervals []Interval) []Interval {
    
    result := make([]Interval, 0)
    
    sort.Slice(intervals, func (i, j int) bool { return intervals[i].Start < intervals[j].Start})
    
    for _, interval := range intervals {
        if len(result) == 0 || result[len(result)-1].End < interval.Start {
            result = append(result, interval)
        } else {
            curEnd := result[len(result)-1].End
            result[len(result)-1].End = max(curEnd, interval.End)
        }
    }
    
    return result 
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}