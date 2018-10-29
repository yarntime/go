func groupAnagrams(strs []string) [][]string {
    resultMap := make(map[string][]string)
    for i := 0; i < len(strs); i++ {
        s := SortString(strs[i])
        if _, ok := resultMap[s]; !ok {
            resultMap[s] = make([]string, 0)
        }
        resultMap[s] = append(resultMap[s], strs[i])
    }
    result := make([][]string, 0)
    
    for _, v := range resultMap {
        result = append(result, v)
    }
    return result
}

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}