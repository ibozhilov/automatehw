package automatehw

import (
	"math"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"syscall"
)

//Generate array with variables
func GenerateVariables(variable string, n int) []string {
	var variables []string
	for i := 1; i <= n; i++ {
		variables = append(variables, variable+"_"+strconv.Itoa(i))
	}
	return variables
}

//Generates LaTeX code for homogenous system of linear equations from array of equations
func SystemOfEquationsToTex(homogenousSystem []string) string {
	latexcode := "V:\\left|\n\\begin{split}\n"
	re := regexp.MustCompile(`(\A|-|\+)[1-9]\*`)
	for _, equation := range homogenousSystem {
		for i, value := range re.FindAllString(equation, -1) {
			if i == 0 {
				equation = strings.Replace(equation, value, strings.Replace(strings.Replace(value, "*", "&", -1), "1", "", -1), -1)
			} else {
				equation = strings.Replace(equation, value, "&"+strings.Replace(strings.Replace(strings.Replace(value, "*", "", -1), "1", "", -1), "-", "-&", 1), 1)
			}
		}
		equation = strings.Replace(equation, "+", "+&", -1)
		equation = strings.Replace(equation, "=", "&=&", 1)
		latexcode = latexcode + equation + "\\\\\n"
	}
	latexcode = latexcode + "\\end{split}\n\\right."
	return latexcode
}

//Generate random permutation of the figures of the fn represented as vectors
func GenerateVectorsFromFN(fn string, n int) []string {
	var vectors, permutation []string
	for i := 0; i < n; i++ {
		permutation = RandomPermutationOfFN(fn, i*8)
		vector := "("
		for t, value := range permutation {
			if t == 0 {
				vector = vector + value
			} else {
				vector = vector + "," + value
			}
		}
		vector = vector + ")"
		vectors = append(vectors, vector)
	}
	return vectors
}

//Takes the LaTeX code from Maxima for a matrix and returns it for LaTeX compiling
func MaximaMatrixToLaTeX(txt, fn string) string {
	re := regexp.MustCompile(`\\pmatrix\{.+\}`)
	f, err := os.Create("/home/ivan/gocode/src/fmi/automatehw/" + fn + ".mac")
	defer syscall.Unlink(f.Name()) // Think whether it is neccesary to delete the file
	if err != nil {
		panic(err)
	}
	f.Write([]byte(txt))
	f.Close()
	cmd := exec.Command(`maxima`, `--very-quiet`, `-b`, f.Name())
	b, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	matrix_latex := re.FindString(string(b))
	matrix_latex = strings.Replace(matrix_latex, "}", `\end{pmatrix}`, 1)
	matrix_latex = strings.Replace(matrix_latex, "{", "}", 1)
	matrix_latex = strings.Replace(matrix_latex, `\`, `\begin{`, 1)
	return matrix_latex
}

//Generate random 3 by 3 matrix with two rational and one irrational number as eigenvalue
func GenerateRandomMatrix3x3(txt, fn string) string {
	num, _ := strconv.ParseInt(fn, 10, 0)
	source := rand.NewSource(num)
	random := rand.New(source)

	//generate maxima script for the current student
	generate_random_matrix := `A: matrix([1,а,б],[0,1,в],[0,0,1])$
	B: matrix([1,0,0],[0,-1,0],[0,0,1])$
	C: matrix([г,д,1],[е,1,0],[1,0,0])$
	D: C.B.A$
	E: matrix([0,2,0],[1,0,0],[0,0,з])$
	M: D.E.invert(D)$
	tex1(M);`
	generate_random_matrix = strings.Replace(generate_random_matrix, "а", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "б", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "в", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "г", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "д", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "е", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "ж", strconv.Itoa(random.Intn(3)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "з", strconv.Itoa(random.Intn(3)+1), -1)

	//replace the new generated matrix in the main LaTeX code
	txt = strings.Replace(txt, `<*random_matrix1*>`, MaximaMatrixToLaTeX(generate_random_matrix, fn), 1)
	return txt
}

//Generate random 4 by 4 matrix with two rational and one irrational number as eigenvalue
func GenerateRandomMatrix4x4(txt, fn string) string {
	num, _ := strconv.ParseInt(fn, 10, 0)
	source := rand.NewSource(num)
	random := rand.New(source)

	/* B: matrix([1,0,0,0],[0,-1,0,0],[0,0,1,0],[0,0,0,-1])$
	C: matrix([ж,з,и,1],[й,к,1,0],[л,1,0,0],[1,0,0,0])$
	D: C.B.A$ */
	//generate maxima script for the current student
	generate_random_matrix := `T: 1/2*matrix([-1,1,1,1],[1,1,1,-1],[-1,1,-1,1],[1,1,-1,1])$
	D: matrix([0,0,0,0],[0,0,0,0],[0,0,м,0],[0,0,0,н])$
	M: T.D.invert(T)$
	M: M;
	tex1(M);`
	generate_random_matrix = strings.Replace(generate_random_matrix, "а", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "б", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "в", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "г", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "д", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "е", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "ж", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "з", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "и", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "й", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "к", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "л", strconv.Itoa(random.Intn(2)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "м", strconv.Itoa((random.Intn(5)+1)*int(math.Pow(-1, float64(random.Intn(13))))), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "н", strconv.Itoa((random.Intn(5)+1)*int(math.Pow(-1, float64(random.Intn(13))))), -1)

	//replace the new generated matrix in the main LaTeX code
	txt = strings.Replace(txt, `<*random_matrix1*>`, MaximaMatrixToLaTeX(generate_random_matrix, fn), 1)
	return txt
}

func GenerateRandomMatrixOfOperator(txt, fn string) string {
	num, _ := strconv.ParseInt(fn, 10, 0)
	source := rand.NewSource(num)
	random := rand.New(source)

	//generate maxima script for the current student
	generate_random_matrix := `A: [а,б,в,г]$
	B: [д,е,ж,з]$
	M: matrix(A,и*A+й*B,B,к*A+л*B)$
	tex1(M);`
	generate_random_matrix = strings.Replace(generate_random_matrix, "а", strconv.Itoa(random.Intn(3)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "б", strconv.Itoa(random.Intn(3)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "в", strconv.Itoa(random.Intn(3)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "г", strconv.Itoa(random.Intn(3)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "д", strconv.Itoa(random.Intn(3)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "е", strconv.Itoa(random.Intn(3)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "ж", strconv.Itoa(random.Intn(3)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "з", strconv.Itoa(random.Intn(3)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "и", strconv.Itoa(random.Intn(3)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "й", strconv.Itoa(random.Intn(3)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "к", strconv.Itoa(random.Intn(3)+1), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "л", strconv.Itoa(random.Intn(3)+1), -1)

	//write the new generated matrix into the main LaTeX code
	txt = strings.Replace(txt, `<*random_matrix2*>`, MaximaMatrixToLaTeX(generate_random_matrix, fn), 1)
	return txt
}
