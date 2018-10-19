type Record struct {
    Pos int
    PreStep int
}
var direct = [3]int{-1, 0, 1}
func canCross(stones []int) bool {
    if len(stones) == 0 || stones[1] != 1 {
        return false
    }
    target := stones[len(stones) - 1]
    validPos := make(map[int]bool)
    isVisited := make(map[Record]bool)
    for _, stone := range stones {
        validPos[stone] = true
    }
    isVisited[Record{Pos:0, PreStep: 0}] = true
    isVisited[Record{Pos:1, PreStep: 1}] = true
    queue := make([]*Record, 0)
    queue = append(queue, &Record{Pos:1, PreStep: 1})
    for len(queue) != 0 {
        p := queue[0]
        queue = queue[1:]
        pos := p.Pos
        PreStep := p.PreStep
        for i := 0; i < 3; i++ {
            s := PreStep + direct[i]
            newPos := pos + s
            if newPos == target {
                return true
            }
            r := &Record{Pos: newPos, PreStep: PreStep + direct[i]}
            if newPos < 0 {
                continue
            }
            if _, ok := isVisited[*r]; ok {
                continue
            }
            if _, ok := validPos[newPos]; !ok {
                continue
            }
            isVisited[*r] =  true
            queue = append(queue, r)
        }
    }
    return false
}