func longestPalindrome(s string) string {
    if len(s) <= 1 {
        return s
    }
    start, max := 0, 1
    dp := [1001][1001]bool{}
    for i := 0; i < len(s); i++ {
        dp[i][i] = true
        if i > 0 && s[i] == s[i - 1] {
            dp[i - 1][i] = true
            max = 2
            start = i - 1
        }
    }
    for reLen := 3; reLen <= len(s); reLen++ {
        for st := 0; st <= len(s) - reLen; st++ {
            en := st + reLen - 1
            if s[st] == s[en] && dp[st+1][en-1] {
                dp[st][en] =  true
                max = reLen
                start = st
            }
        }
    }
    return s[start:start+max]
}