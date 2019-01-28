type Record struct {
	Objects      []int
	Value        int
	LeftCapacity int
}

func NewRecord() *Record {
	return &Record{
		Objects:      make([]int, 0),
		Value:        0,
		LeftCapacity: 0,
	}
}

func CopyRecord(r *Record) *Record {
	newRecord := NewRecord()
	newRecord.Objects = make([]int, len(r.Objects))
	copy(newRecord.Objects, r.Objects)
	newRecord.Value = r.Value
	newRecord.LeftCapacity = r.LeftCapacity
	return newRecord
}

func (af *ArtificialFish) Adjust1(bag int, objectID int) {
	unusedObjects := make([]int, 0)
	for i := objectID + 1; i < OBJECT_NUM; i++ {
		if af.Object[i] == UNSELECTED && af.Capacity[bag] >= AllGoods[i].Weight {
			unusedObjects = append(unusedObjects, i)
		}
	}

	dp := make([]*Record, 0)
	init := NewRecord()
	init.LeftCapacity = af.Capacity[bag]
	dp = append(dp, init)
	for i := 0; i < len(unusedObjects); i++ {
		newDp := make([]*Record, 0)
		for j := 0; j < len(dp); j++ {
			if dp[j].LeftCapacity > AllGoods[unusedObjects[i]].Weight {
				newRecord := CopyRecord(dp[j])
				newRecord.Objects = append(newRecord.Objects, unusedObjects[i])
				newRecord.Value += AllGoods[unusedObjects[i]].Value
				newRecord.LeftCapacity -= AllGoods[unusedObjects[i]].Weight
				newDp = append(newDp, newRecord)
			}
			newDp = append(newDp, dp[j])
		}
		dp = newDp
		sort.Slice(dp, func(i, j int) bool {
			return dp[i].Value > dp[j].Value
		})
		if len(dp) > 100 {
			dp = dp[0:100]
		}
	}
	for i := 0; i < len(dp[0].Objects); i++ {
		af.Object[dp[0].Objects[i]] = bag
		af.Capacity[bag] -= AllGoods[dp[0].Objects[i]].Weight
	}
}
