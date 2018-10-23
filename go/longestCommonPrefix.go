func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    result := ""
    for i := 0; i < (1 << 31); i++ {
        if strs[0][i:] == "" {
			return result
		}
        target := strs[0][i]
        for j := 1; j < len(strs); j++ {
            if strs[j][i:] == "" {
                return result
            }
            if strs[j][i] != target {
                return result
            }
        }
        result = result + string(target)
    }
    return result
}