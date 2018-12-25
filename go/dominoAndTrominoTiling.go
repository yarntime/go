func numTilings(N int) int {
    a, b, c := 1, 2, 5
    mod := 1000000007
    if N == 1 {
        return 1
    }
    if N == 2 {
        return 2
    }
    if N == 3 {
        return 5
    }
    for i := 4; i <= N; i++ {
        cur := (c * 2 % mod + a) % mod
        a = b
        b = c
        c = cur
    }
    return c
}