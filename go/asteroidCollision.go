func asteroidCollision(asteroids []int) []int {
    stack := make([]int, 0)
    for i := 0; i < len(asteroids); i++ {
        if asteroids[i] < 0 {
            explode := false
            for len(stack) != 0 {
                if stack[len(stack) - 1] < 0 {
                    break
                }
                if -asteroids[i] > stack[len(stack) - 1] {
                    stack = stack[0:len(stack) - 1]
                } else if -asteroids[i] == stack[len(stack) - 1] {
                    explode = true
                    stack = stack[0:len(stack) - 1]
                    break
                } else if -asteroids[i] < stack[len(stack) - 1] {
                    explode = true
                    break
                }
            }
            if !explode {
                stack = append(stack, asteroids[i])
            }
        } else {
            stack = append(stack, asteroids[i])
        }
    }
    return stack
}