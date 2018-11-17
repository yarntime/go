func toLowerCase(str string) string {
    result := make([]byte, 0)
    for i := 0; i < len(str); i++ {
        if str[i] >= 'A' && str[i] <= 'Z' {
            result = append(result, str[i] - 'A' + 'a')
        } else {
            result = append(result, str[i])
        }
    }
    return string(result)
}