type RandomizedSet struct {
    m map[int]int
    n []int
}

func Constructor() RandomizedSet {
    return RandomizedSet{m: make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.m[val]; ok {
        return false
    }
    this.m[val] = len(this.n)
    this.n = append(this.n, val)
    return true
}

func (this *RandomizedSet) Remove(val int) bool {
    if idx, ok := this.m[val]; ok {
        this.n[idx] = this.n[len(this.n) - 1]
        this.m[this.n[idx]] = idx    
        this.n = this.n[:len(this.n) - 1]
        delete(this.m, val)
        return true
    }
    return false
}

func (this *RandomizedSet) GetRandom() int {
    if len(this.n) == 0 {
        return 0
    }
    return this.n[rand.Intn(len(this.n))]
}