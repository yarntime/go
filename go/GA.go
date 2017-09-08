package main

import (
	"fmt"
	"math/rand"

	. "async/fish"
	gago "github.com/MaxHalford/gago"
	"os"
)

type Solution []int

func (X Solution) Evaluate() float64 {
	result := 0.0
	for i := 0; i < OBJECT_NUM; i++ {
		if X[i] != 0 {
			result += float64(Value[i])
		}
	}

	return -result
}

func (X Solution) Mutate(rng *rand.Rand) {
	n := rng.Intn(OBJECT_NUM/2) + 1
	gago.MutPermuteInt(X, n, rng)
	X.Repair(rng)
}

func (X Solution) Crossover(Y gago.Genome, rng *rand.Rand) (gago.Genome, gago.Genome) {
	var o1, o2 = gago.CrossGNXInt(X, Y.(Solution), 10, rng)
	var s1, s2 = Solution(o1), Solution(o2)
	s1.Repair(rng)
	s2.Repair(rng)
	return s1, s2
}

func (X Solution) Clone() gago.Genome {
	var Y = make(Solution, len(X))
	copy(Y, X)
	return Y
}

func (X Solution) IsValidSolution() bool {
	var usedCapacity = make([]int, BAG_NUM)
	for i := 0; i < OBJECT_NUM; i++ {
		if X[i] != 0 {
			for j := 0; j < BAG_NUM; j++ {
				usedCapacity[j] += Dim[j][i]
			}
		}
	}

	for i := 0; i < BAG_NUM; i++ {
		if usedCapacity[i] > Capacity_Limit[i] {
			return false
		}
	}

	return true
}

func (X Solution) Repair(rng *rand.Rand) {
	for !X.IsValidSolution() {
		for pos := rng.Intn(OBJECT_NUM); pos < OBJECT_NUM; pos++ {
			if X[pos] != 0 {
				X[pos] = 0
				break
			}
		}
	}
}

func Make(rng *rand.Rand) gago.Genome {
	var solution = make(Solution, OBJECT_NUM)

	for i := 0; i < OBJECT_NUM; i++ {
		solution[i] = rng.Intn(2)
	}

	solution.Repair(rng)

	return gago.Genome(solution)
}

func main() {

	fmt.Printf("starting\n")

	err := InitGoods("FILE")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var ga = gago.Generational(Make)

	ga.Initialize()

	//fmt.Printf("Best fitness at generation 0: %f\n", ga.Best.Fitness)
	checkResult(ga.Best.Genome.(Solution))

	for i := 1; i < 20000; i++ {
		fmt.Printf("generation %d, ", i)
		ga.Enhance()
		checkResult(ga.Best.Genome.(Solution))
	}
}

func checkResult(s Solution) {
	var value float64
	var capacities = make([]int, BAG_NUM)
	var error = 0

	for i := 0; i < OBJECT_NUM; i++ {
		if s[i] != 0 {
			for j := 0; j < BAG_NUM; j++ {
				capacities[j] += Dim[j][i]
			}
			value += float64(Value[i])
		}
	}

	for i := 0; i < BAG_NUM; i++ {
		if capacities[i] > Capacity_Limit[i] {
			fmt.Printf("bag is %v\n", i)
			fmt.Printf("limit is %v, current is %v\n", Capacity_Limit[i], capacities[i])
			fmt.Printf("invalid result.\n")
			error++
		}
	}

	if error != 0 {
		fmt.Printf("invalid result, there are %d bags error\n", error)
	} else {
		fmt.Printf("valid result. fitness is %f\n", value)
	}
}
