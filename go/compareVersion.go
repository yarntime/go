func compareVersion(version1 string, version2 string) int {
    v1 := strings.Split(version1, ".")
    v2 := strings.Split(version2, ".")
    var i int
    for i = 0; i < len(v1) && i < len(v2); i++ {
        i1, _ := strconv.Atoi(v1[i])
        i2, _ := strconv.Atoi(v2[i])
        if i1 > i2 {
            return 1
        } else if i1 < i2 {
            return -1
        }
    }
    for i < len(v1) {
        i1, _ := strconv.Atoi(v1[i])
        if i1 != 0 {
            return 1
        }
        i++
    }
    for i < len(v2) {
        i2, _ := strconv.Atoi(v2[i])
        if i2 != 0 {
            return -1
        }
        i++
    }
    return 0
}