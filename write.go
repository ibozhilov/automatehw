package automatehw

import (
	//"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"syscall"
)

//Generate a personalized string in .tex format for the student
func Write(txt, fn, gr, spec, name, info, date string) string {
	txt = WriteFN(txt, fn)
	txt = WriteDate(txt, date)
	txt = WriteGr(txt, gr)
	txt = WriteName(txt, name)
	txt = WriteSpec(txt, spec)
	txt = WriteInfo(txt, info)
	txt = WriteRandomMatrix(txt, fn)
	txt = WriteRandomMatrix(txt, fn)
	txt = WriteRandomLinearSystem(txt, fn)
	txt = WriteVector(txt, fn)
	txt = GenerateRandomMatrix4x4(txt, fn)
	txt = GenerateRandomMatrixOfOperator(txt, fn)
	return txt
}

//Write the FN of the student in the document
func WriteFN(txt, fn string) string {
	re := regexp.MustCompile(`<\*[Ff][Nn]\*>`)
	return re.ReplaceAllString(txt, fn)
}

//Write the Group number of the student in the document
func WriteGr(txt, gr string) string {
	re := regexp.MustCompile(`<\*[Gg][Rr]\*>`)
	return re.ReplaceAllString(txt, gr)
}

//Write the course, which the student is taking
func WriteSpec(txt, spec string) string {
	re := regexp.MustCompile(`<\*[Ss][Pp][Ee][Cc]\*>`)
	return re.ReplaceAllString(txt, spec)
}

//Write the name of the student in the document
func WriteName(txt, name string) string {
	re := regexp.MustCompile(`<\*[Nn][Aa][Mm][Ee]\*>`)
	return re.ReplaceAllString(txt, name)
}

//Write info about the homework or the test
func WriteInfo(txt, info string) string {
	re := regexp.MustCompile(`<\*[Ii][Nn][Ff][Oo]\*>`)
	return re.ReplaceAllString(txt, info)
}

//Write the date, that will be printed on the document
func WriteDate(txt, date string) string {
	re := regexp.MustCompile(`<\*[Dd][Aa][Tt][Ee]\*>`)
	return re.ReplaceAllString(txt, date)
}

//Generates a random matrix from the FN of the student and write it down in the document.
//Maxima is used for the generation of the matrix.
func WriteRandomMatrix(txt, fn string) string {
	re := regexp.MustCompile(`\\pmatrix\{.+\}`)
	f, err := os.Create("/home/ivan/gocode/src/fmi/automatehw/" + fn + ".mac")
	defer syscall.Unlink(f.Name()) // Think whether it is neccesary to delete the file
	if err != nil {
		panic(err)
	}
	//generate maxima script for the current student
	generate_random_matrix := "A: [fn_0,fn_1,fn_2,fn_3,fn_4]$\nM: matrix(A)$\nfor i: 1 thru 4 do\nM: addrow(M, random_permutation(A))$\ntex1(M);"
	generate_random_matrix = strings.Replace(generate_random_matrix, "fn_0", string(fn[0]), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "fn_1", string(fn[1]), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "fn_2", string(fn[2]), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "fn_3", string(fn[3]), -1)
	generate_random_matrix = strings.Replace(generate_random_matrix, "fn_4", string(fn[4]), -1)

	f.Write([]byte(generate_random_matrix))
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
	txt = strings.Replace(txt, `<*random_matrix*>`, matrix_latex, 1)
	return txt
}

//
func WriteRandomLinearSystem(txt, fn string) string {
	txt = strings.Replace(txt, `<*random_linear_system*>`, SystemOfEquationsToTex(GenerateHomogenousEquationsFromFN(fn, GenerateVariables("x", len(fn)), 3, 7)), -1)
	return txt
}

func WriteVectors(txt, fn string) string {
	newstr := "\\left.\n\\begin{split}\na&=&"
	vectors := GenerateVectorsFromFN(fn, 3)
	newstr = newstr + vectors[0] + "\\\\\nb&=&" + vectors[1] + "\\\\\nc&=&" + vectors[2] + "\n\\end{split}\n\\right."
	txt = strings.Replace(txt, `<*random_vectors*>`, newstr, 1)
	return txt
}

func WriteVector(txt, fn string) string {
	vector := GenerateVectors(fn, 4, 3, false)
	txt = strings.Replace(txt, `<*random_vector*>`, VectorToString(vector[0]), 1)
	txt = strings.Replace(txt, `<*random_vector*>`, VectorToString(vector[1]), 1)
	txt = strings.Replace(txt, `<*random_vector*>`, VectorToString(vector[2]), 1)
	return txt
}

func VectorToString(vector []int) string {
	svector := "("
	for i, value := range vector {
		if i == 0 {
			svector = svector + strconv.Itoa(value)
		} else {
			svector = svector + "," + strconv.Itoa(value)
		}
	}
	svector = svector + ")"
	return svector
}
