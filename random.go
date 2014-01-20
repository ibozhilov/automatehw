package automatehw

import (
	//"fmt"
	"math"
	"math/rand"
	"strconv"
)

//Generate random equation initialized by the fn f the student
func RandomHomogenousEquations(fn string, variables []string, numberOfEquations int, maxnumber int) []string {
	num, _ := strconv.ParseInt(fn, 10, 0)
	source := rand.NewSource(num)
	random := rand.New(source)
	var homogenousEquations []string
	var homogenousEquation string
	var coefficient, rnum int
	for count := 0; count < numberOfEquations; count++ {
		homogenousEquation = ""
		for i, value := range variables {
			rnum = 0
			for rnum == 0 {
				rnum = random.Intn(maxnumber)
			}
			coefficient = rnum * int(math.Pow(-1, float64(random.Int())))
			if i == 0 {
				homogenousEquation = homogenousEquation + strconv.Itoa(coefficient) + "*" + string(value)
			} else {
				if coefficient < 0 {
					homogenousEquation = homogenousEquation + strconv.Itoa(coefficient) + "*" + string(value)
				} else {
					homogenousEquation = homogenousEquation + "+" + strconv.Itoa(coefficient) + "*" + string(value)
				}
			}
		}
		homogenousEquation = homogenousEquation + "=0"
		homogenousEquations = append(homogenousEquations, homogenousEquation)
	}
	return homogenousEquations
}

// Generate homogenous equations with coeficients the figures of the FN
func GenerateHomogenousEquationsFromFN(fn string, variables []string, numberOfEquations int, maxnumber int) []string {
	num, _ := strconv.ParseInt(fn, 10, 0)
	source := rand.NewSource(num)
	random := rand.New(source)
	var homogenousEquations, permutation []string
	var homogenousEquation string
	var coefficient, rnum int
	for count := 0; count < numberOfEquations; count++ {
		homogenousEquation = ""
		permutation = RandomPermutationOfFN(fn, count*13)
		for i, value := range variables {
			rnum, _ = strconv.Atoi(permutation[i])
			for rnum == 0 {
				rnum = random.Intn(maxnumber)
			}
			coefficient = rnum * int(math.Pow(-1, float64(random.Int())))
			if i == 0 {
				homogenousEquation = homogenousEquation + strconv.Itoa(coefficient) + "*" + string(value)
			} else {
				if coefficient < 0 {
					homogenousEquation = homogenousEquation + strconv.Itoa(coefficient) + "*" + string(value)
				} else {
					homogenousEquation = homogenousEquation + "+" + strconv.Itoa(coefficient) + "*" + string(value)
				}
			}
		}
		homogenousEquation = homogenousEquation + "=0"
		homogenousEquations = append(homogenousEquations, homogenousEquation)
	}
	return homogenousEquations
}

//Generate random permutation of the figures in the FN
func RandomPermutationOfFN(fn string, n int) []string {
	num, _ := strconv.ParseInt(fn, 10, 0)
	source := rand.NewSource(num * int64(n))
	random := rand.New(source)
	intPerm := random.Perm(len(fn))
	var permutation []string
	for _, value := range intPerm {
		permutation = append(permutation, string(fn[value]))
	}
	return permutation
}

//Generate random vectors
func GenerateVectors(fn string, dim int, n int, positive bool) [][]int {
	var vectors [][]int
	num, _ := strconv.ParseInt(fn, 10, 0)
	source := rand.NewSource(num)
	random := rand.New(source)
	var vector []int
	if positive == true {
		for i := 0; i < n; i++ {
			vector = []int{}
			if i > 0 {
				for IsLinearIndependent(vectors[0], vector) {
					for t := 0; t < dim; t++ {
						value := random.Intn(4) + 1
						vector = append(vector, value)
					}
				}
			} else {
				for t := 0; t < dim; t++ {
					value := random.Intn(4) + 1
					vector = append(vector, value)
				}
			}
			vectors = append(vectors, vector)
		}
	} else {
		for i := 0; i < n; i++ {
			vector = []int{0, 0, 0, 0}
			if i > 0 {
				//fmt.Println(IsLinearIndependent(vectors[0], vector))
				for !IsLinearIndependent(vectors[0], vector) {
					vector = []int{}
					for t := 0; t < dim; t++ {
						value := (random.Intn(4) + 1) * int(math.Pow(-1, float64(random.Intn(13))))
						vector = append(vector, value)
					}
				}
			} else {
				vector = []int{}
				for t := 0; t < dim; t++ {
					value := random.Intn(4) + 1
					vector = append(vector, value)
				}
			}
			vectors = append(vectors, vector)
			//fmt.Println(vectors)
		}
	}
	return vectors
}
