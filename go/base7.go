func convertToBase7(num int) string {
    if num == 0 {
        return "0"
    }
    flag := false
    if num < 0 {
        flag = true
        num = -num
    }
    result := make([]byte, 0)
    for num != 0 {
        q := num % 7
        num /= 7
        result = append([]byte{byte('0' + q)}, result...)
    }
    if flag {
        result = append([]byte{'-'}, result...)
    } 
    return string(result)
}