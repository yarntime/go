func compress(chars []byte) int {
	cur := 1
	count := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i - 1] {
			count++
		} else  {
			if count != 1 {
				c := getChar(count)
				for j := 0; j < len(c); j++ {
					chars[cur] = c[j]
					cur++
				}
			}
			chars[cur] = chars[i]
			cur++
			count = 1
		}
	}
	if count > 1 {
		c := getChar(count)
		for j := 0; j < len(c); j++ {
			chars[cur] = c[j]
			cur++
		}
	}
	return cur
}

func getChar(a int) []byte {
	result := make([]byte, 0)
	for a != 0 {
		result = append([]byte{byte('0' + a % 10)}, result...)
		a /= 10
	}
	return result
}