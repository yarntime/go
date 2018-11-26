type State struct {
    current int;
    step int;
}

type StateQueue []State

func numBusesToDestination(routes [][]int, S int, T int) int {
    stopToRoute := map[int]map[int]bool{}
    for routeIndex := 0; routeIndex < len(routes); routeIndex++ {
        for _, stop := range routes[routeIndex] {
            if _, ok := stopToRoute[stop]; !ok {
                stopToRoute[stop] = map[int]bool{}
            }
            stopToRoute[stop][routeIndex] = true
        }
    }
    
    q := StateQueue{State{S, 0}}
    visitedStops := map[int]bool{}
    
    for len(q) != 0 {
        front := q[0]
        q = q[1:]
        
        if front.current == T {
            return front.step
        }
        if visitedStops[front.current] {
            continue
        }
        
        visitedStops[front.current] = true
        for route, _ := range stopToRoute[front.current] {
            fmt.Println(routes[route])
            for _, stop := range routes[route] {
                q = append(q, State{stop, front.step + 1})
            }
            routes[route] = []int{}
        }
    }
    return -1
}