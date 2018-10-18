var direct = map[byte][2]byte{
    '0':{'9','1'},
    '1':{'0','2'},
    '2':{'1','3'},
    '3':{'2','4'},
    '4':{'3','5'},
    '5':{'4','6'},
    '6':{'5','7'},
    '7':{'6','8'},
    '8':{'7','9'},
    '9':{'8','0'}}
func openLock(deadends []string, target string) int {
	processed := make(map[string]bool)
	dead := make(map[string]bool)
	for _, str := range deadends {
		dead[str] = true
	}
    if _, ok := dead["0000"]; ok {
        return -1
    }
	steps := 0
	queue := make([]string, 0)
	queue = append(queue, "0000")
	processed["0000"] = true
	for len(queue) != 0 {
		size := len(queue)
		for k := 0; k < size; k++ {
			now := queue[0]
			queue = queue[1:]
			if now == target {
				return steps
			}
			for i := 0; i < len(now); i++ {
				for j := 0; j < 2; j++ {
					tmp := now[0:i] + string(direct[now[i]][j]) + now[i + 1:]
					if _, ok := dead[tmp]; ok {
						continue
					}
					if _, ok := processed[tmp]; ok {
						continue
					} else {
						processed[tmp] = true
					}
					queue = append(queue, tmp)
				}
			}
		}
        steps++
	}
	return -1
}