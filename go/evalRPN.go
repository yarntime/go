func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "*":
			first := stack[len(stack)-1]
			second := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, first*second)
		case "/":
			first := stack[len(stack)-1]
			second := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if first != 0 {
				i := second / first
				i = int(i)
				stack = append(stack, i)
			} else {
				stack = append(stack, 0)
			}
		case "+":
			first := stack[len(stack)-1]
			second := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, first+second)
		case "-":
			first := stack[len(stack)-1]
			second := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, second-first)
		default:
			val, _ := strconv.Atoi(tokens[i])
			stack = append(stack, val)
		}
	}
	return stack[0]
}