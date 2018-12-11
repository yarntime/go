type Pair struct {
    Key rune
    Count int
}
func frequencySort(s string) string {
    counts := make(map[rune]*Pair)
    for _, c := range s {
        if _, ok := counts[c]; !ok {
            counts[c] = &Pair{
                Key: c,
                Count: 0,
            }
        }
        p := counts[c]
        p.Count++
    }
    values := make([]*Pair, 0)
    for _, v := range counts {
        values = append(values, v)
    }
    sort.Slice(values, func(i, j int) bool {
        return values[i].Count > values[j].Count
    })
    result := make([]rune, 0)
    for _, value := range values {
        for i := 0; i < value.Count; i++ {
            result = append(result, value.Key)
        }
    }
    return string(result)
}