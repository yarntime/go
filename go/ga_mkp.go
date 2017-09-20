package ga

import (
	"fmt"
	"math/rand"
	"sort"
)

type Constraints struct {
	BagLimit int
	Weights  []int
}

type Problem struct {
	Profits     []int
	Constraints []Constraints
}

type Solution []int
type Probabilites []float64

type Interval struct {
	from, to float64
}

type Generic struct {
	Population           []Solution
	problem              Problem
	PopulationSize       int
	mutateProbability    float64
	crossoverProbability float64
	rnd                  *rand.Rand
}

func CreateInitial(ps, ss int, p Problem, m, c float64, r *rand.Rand) *Generic {
	g := &Generic{
		PopulationSize:       ps,
		problem:              p,
		mutateProbability:    m,
		crossoverProbability: c,
		rnd:                  r,
	}
	g.Population = make([]Solution, 0)
	g.PopulationSize = ps
	var packed int

	// create g.rndomly a valid population
	for i := 0; i < ps; {
		solution := make([]int, 0)
		for j := 0; j < ss; j++ {
			packed = g.rnd.Intn(2)
			solution = append(solution, packed)
		}

		g.makeSolutionValid(solution)

		if g.isNotDouble(solution, g.Population) {
			g.Population = append(g.Population, solution)
			i++
		}
	}
	return g
}

func (g *Generic) AnalyzeBestSolution() {
	best := g.getFittestSolution()
	quality := g.evaluate(best)

	fmt.Printf("total best solution with profit %d is %d \n", quality, best)
}

func (g *Generic) LogCurrentBest(generation int) {
	best := g.getFittestSolution()
	quality := g.evaluate(best)
	fmt.Printf("generation: %d, best quality %d\n", generation, quality)
}

func (g *Generic) GenerateOffspringPopulation() {

	offspringsLimit := int(OffspringsProportion * float64(g.PopulationSize))

	offspringPopulation := make([]Solution, 0)

	probabilites := g.computeProbabilites()

	for len(offspringPopulation) < offspringsLimit {

		parents := g.getParents(probabilites)
		children := g.getOffsprings(parents)

		offspringPopulation = g.insertChildren(children, offspringPopulation)
	}

	fitnesses := g.getFitnessValuesOfCurrentPopulation()

	g.Population = g.replaceWorst(fitnesses, offspringPopulation)
}

func (g *Generic) getFitnessValuesOfCurrentPopulation() []int {
	var population = g.Population
	var result = make([]int, 0)

	for i := 0; i < len(population); i++ {
		result = append(result, g.getFitness(population[i]))
	}

	return result
}

func (g *Generic) insertChildren(children, offspringPopulation []Solution) []Solution {
	for i := 0; i < len(children); i++ {
		if g.isValidAndNotDouble(children[i], offspringPopulation) {
			if DEBUG {
				fmt.Printf("children[%d] %v\n, valid children\n", i, children[i])
			}
			offspringPopulation = append(offspringPopulation, children[i])
		}
	}
	return offspringPopulation
}

func (g *Generic) evaluate(s Solution) int {
	var totalProfit = 0

	for i := 0; i < len(s); i++ {
		if s[i] == 1 {
			totalProfit += g.problem.Profits[i]
		}
	}

	var contraintsMet = g.checkConstraints(s)

	if !contraintsMet {
		return -1
	}
	return totalProfit
}

func (g *Generic) checkConstraints(s Solution) bool {
	var totalWeight, bagLimit int
	var constraint Constraints
	var constraintMet bool

	constraintMet = true

	for c := 0; c < len(g.problem.Constraints); c++ {
		constraint = g.problem.Constraints[c]
		bagLimit = g.problem.Constraints[c].BagLimit
		totalWeight = g.getWeight(s, constraint)
		if totalWeight > bagLimit {
			constraintMet = false
			break
		}
	}

	return constraintMet
}

func (g *Generic) getWeight(s []int, c Constraints) int {
	var weight = 0

	for i := 0; i < len(s); i++ {
		if s[i] == 1 {
			weight += c.Weights[i]
		}
	}

	return weight
}

func (g *Generic) makeSolutionValid(s Solution) {
	var r int
	for g.evaluate(s) < 0 {
		r = g.rnd.Intn(len(s))
		s[r] = 0
	}
}

func (g *Generic) isNotDouble(s Solution, p []Solution) bool {
	var isDouble = false
	var isSame bool

	if len(p) == 0 {
		return !isDouble
	}

	for i := 0; i < len(p); i++ {
		isSame = true
		for j := 0; j < len(s); j++ {
			if p[i][j] != s[j] {
				isSame = false
				break
			}
		}
		if isSame {
			isDouble = true
			break
		}
	}

	return !isDouble
}

func (g *Generic) isValidAndNotDouble(s Solution, t []Solution) bool {
	var isValid = g.evaluate(s) >= 0
	var isNotDouble = g.isNotDouble(s, t)
	return isValid && isNotDouble
}

func (g *Generic) getFitness(s Solution) int {
	return g.evaluate(s)
}

func (g *Generic) getFittestSolution() Solution {
	highestFitness := -1
	var fittestSolution Solution
	var fittest int
	for i := 0; i < len(g.Population); i++ {
		fittest = g.getFitness(g.Population[i])
		if fittest > highestFitness {
			highestFitness = fittest
			fittestSolution = g.Population[i]
		}
	}

	return fittestSolution
}

func (g *Generic) computeProbabilites() Probabilites {

	probabilites := make([]float64, 0)
	fitnesses := make([]int, 0)
	totalFitness := 0

	for i := 0; i < len(g.Population); i++ {
		fitness := g.getFitness(g.Population[i])
		totalFitness += fitness
		fitnesses = append(fitnesses, fitness)
	}

	for i := 0; i < len(g.Population); i++ {
		p := float64(fitnesses[i]) / float64(totalFitness)
		probabilites = append(probabilites, p)
	}

	return probabilites
}

func (g *Generic) getParents(p Probabilites) []Solution {
	parents := make([]Solution, 0)
	intervals := g.createProbabilityIntervals(p)

	for len(parents) < 2 {
		r := g.rnd.Float64()

		parent := g.getSolutionByProbability(intervals, r)

		if g.isNotDouble(parent, parents) {
			parents = append(parents, parent)
		}
	}

	return parents
}

func (g *Generic) createProbabilityIntervals(p Probabilites) []Interval {
	var intervals []Interval
	var interval Interval

	for i := 0; i < len(p); i++ {
		if i == 0 {
			interval.from = 0.0
		} else {
			interval.from = intervals[i-1].to
		}
		interval.to = interval.from + p[i]

		intervals = append(intervals, interval)
	}

	return intervals
}

func (g *Generic) isProbable(p float64) bool {
	return g.rnd.Float64() < p
}

func (g *Generic) getSolutionByProbability(intervals []Interval, r float64) Solution {
	var solution Solution
	for i := 0; i < len(intervals); i++ {
		if r >= intervals[i].from && r < intervals[i].to {
			solution = make([]int, len(g.Population[i]))
			copy(solution, g.Population[i])
			break
		}
	}

	return solution
}

func (g *Generic) getOffsprings(parents []Solution) []Solution {

	offsprings := make([]Solution, 0)
	var parentIndex int

	c1 := parents[0]
	c2 := parents[1]

	if g.isProbable(g.crossoverProbability) {
		for i := 0; i < len(parents[0]); i++ {
			if g.isProbable(0.5) {
				parentIndex = 0
			} else {
				parentIndex = 1
			}

			c1[i] = parents[parentIndex][i]
			c2[i] = parents[1-parentIndex][i]
		}

		c1 = g.mutate(c1)
		c2 = g.mutate(c2)
	}

	offsprings = append(offsprings, c1)
	offsprings = append(offsprings, c2)

	return offsprings
}

func (g *Generic) mutate(s Solution) Solution {
	if g.isProbable(g.mutateProbability) {
		var pos = g.rnd.Intn(len(s))
		s[pos] = 1 - s[pos]
	}
	return s
}

func (g *Generic) replaceWorst(fitnesses []int, offsprings []Solution) []Solution {
	population := g.Population
	mapping := make([]Map, 0)
	sanitizedOffsprings := g.copyLimited(offsprings, len(population))
	for i := 0; i < len(fitnesses); i++ {
		mapping = append(mapping, Map{
			index: i,
			val:   fitnesses[i],
		})
	}

	sort.Sort(MapSlice(mapping))

	if DEBUG {
		fmt.Printf("len sanitizedOffsprings %d\n", len(sanitizedOffsprings))
	}
	cur := 0
	for i := 0; i < len(mapping); i++ {
		if cur < len(sanitizedOffsprings)-1 {
			cur++
			if DEBUG {
				fmt.Printf("len population %d, k %d, cur %d \n", len(population), mapping[i].index, cur)
			}
			population[mapping[i].index] = sanitizedOffsprings[cur]
		} else {
			break
		}
	}

	return population
}

func (g *Generic) copyLimited(sourcePopulation []Solution, max int) []Solution {

	if len(sourcePopulation) <= max {
		return sourcePopulation
	}

	result := make([]Solution, max)

	for i := 0; i < max; i++ {
		result[i] = sourcePopulation[i]
	}

	return result
}
