func letterCasePermutation(S string) []string {
    result := make([]string, 0)
    result = append(result, "")
    for i := 0; i < len(S); i++ {
        c := byte(0)
        if S[i] >= 'a' && S[i] <= 'z' {
            c = S[i]
        } else if S[i] >= 'A' && S[i] <= 'Z' {
            c = S[i] - 'A' + 'a'
        } else {
            for j := 0; j < len(result); j++ {
                result[j] = result[j] + string(S[i])
            }
            continue
        }
        tmp := make([]string, 0)
        for j := 0; j < len(result); j++ {
            tmp = append(tmp, result[j] + string(c))
            tmp = append(tmp, result[j] + string(c - 'a' + 'A'))
        }
        result = tmp
    }
    return result
}