package automatehw

//import "fmt"

//Check whether two vectors are perpendiculat
func IsPerpendicular(vector1, vector2 []int) bool {
	if ScalarProductOfVectors(vector1, vector2) == 0 {
		return true
	}
	return false
}

//Check whether the two vectors are linear independent
func IsLinearIndependent(vector1, vector2 []int) bool {
	checkVectorsLength(vector1, vector2)
	var coef float32
	coef = float32(vector1[0]) / float32(vector2[0])
	for i, value := range vector1 {
		if coef != float32(value)/float32(vector2[i]) {
			return true
		}
	}
	return false
}

//Returns a perpendicular vector to a given one, by simply switching the places of the coordinates of the vector two by two
// and changeing the sign
func GetPerpendicular(vector1 []int) []int {
	var vector []int
	dim := len(vector1)
	if dim == 0 {
		panic("The vector is empty")
	} else if dim == 1 {
		return []int{0}
	} else if dim%2 == 0 {
		for i := range vector {
			if i%2 == 0 {
				vector = append(vector, (-1)*vector1[i+1])
			} else {
				vector = append(vector, vector1[i-1])
			}
		}
	} else {
		for i := 0; i < dim-1; i++ {
			if i%2 == 0 {
				vector = append(vector, (-1)*vector1[i+1])
			} else {
				vector = append(vector, vector1[i-1])
			}
		}
		vector = append(vector, 0)
	}
	return vector
}

//Returns a perpendicular vector to a 4 dimensional one, through one of the three methods
func GetPerpendicular4(vector1 []int, n int) []int {
	if len(vector1) != 4 {
		panic("The length of the vector is greater or less than 4!")
	}
	switch {
	case n == 0:
		return []int{vector1[1], (-1) * vector1[0], vector1[3], (-1) * vector1[2]}
	case n == 1:
		return []int{vector1[2], (-1) * vector1[3], (-1) * vector1[0], vector1[1]}
	case n == 2:
		return []int{(-1) * vector1[3], vector1[2], (-1) * vector1[1], vector1[0]}
	}
	return []int{}
}

//Calculates the scalar product of two vectors
func ScalarProduct(vector1, vector2 []int) int {
	var product int
	checkVectorsLength(vector1, vector2)
	for i := range vector1 {
		product = product + vector1[i]*vector2[i]
	}
	return product
}

//Return the sum of two vectors
func sumVectors(vector1, vector2 []int) []int {
	var sum []int
	checkVectorsLength(vector1, vector2)
	for i := range vector1 {
		sum = append(sum, vector1[i]+vector2[i])
	}
	return sum
}

//Check whether the two vectors have the same lenght. If not raise an exception
func checkVectorsLength(vector1, vector2 []int) {
	if len(vector2) != len(vector1) {
		panic("Could not sum the vectors! The two vectors has different dimension!")
	}
}
