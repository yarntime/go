type Tire struct {
    IsPrefix  bool
    Next [26]*Tire
}

func replaceWords(dict []string, sentence string) string {
    root := &Tire{IsPrefix: false}
    for _, str := range dict {
        if str == "" {
            continue
        }
        cur := root
        for i := 0; i < len(str); i++ {
            if cur.Next[str[i] - 'a'] == nil {
                cur.Next[str[i] - 'a'] = &Tire{IsPrefix: false}
            }
            cur = cur.Next[str[i] - 'a']
        }
        cur.IsPrefix = true
    }
    result := ""
    strs := strings.Split(sentence, " ")
    for i := 0; i < len(strs); i++ {
        cur := root
        pos := -1
        for j := 0; j < len(strs[i]); j++ {
            if cur.IsPrefix {
                pos = j
                break
            } else if cur.Next[strs[i][j] - 'a'] != nil {
                cur = cur.Next[strs[i][j] - 'a']
            } else if cur.Next[strs[i][j] - 'a'] == nil {
                break
            }
        }
        if pos == -1 {
            result += strs[i] + " "
        } else {
            result += strs[i][0:pos] + " "
        }
    }
    return result[0:len(result) - 1]
}