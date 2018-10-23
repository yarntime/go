func isSubsequence(s string, t string) bool {
    tLast := len(t) - 1
    sLast := len(s) - 1
    for tLast >= 0 && sLast >= 0 {
        for tLast >= 0 && s[sLast] != t[tLast] {
            tLast--
        }
        if tLast < 0 {
            break
        }
        tLast--
        sLast--
    }
    return sLast < 0
}