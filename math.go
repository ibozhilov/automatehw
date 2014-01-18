package automatehw

//import "fmt"

//Check whether two vectors are perpendiculat
func IsPerpendicular(vector1, vector2 []int) bool {
	if len(vector1) != len(vector2) {
		panic("The two vectors have different dimension")
	}
	var product int
	for i, value := range vector1 {
		product = product + vector2[i]*value
	}
	if product == 0 {
		return true
	}
	return false
}

//Check whether the two vectors are linear dependent
func IsLinearIndependent(vector1, vector2 []int) bool {
	//fmt.Println("hey")
	if len(vector1) != len(vector2) {
		panic("The two vectors have different dimension")
	}
	var coef float32
	coef = float32(vector1[0]) / float32(vector2[0])
	for i, value := range vector1 {
		if coef != float32(value)/float32(vector2[i]) {
			return true
		}
	}
	return false
}
