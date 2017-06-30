package automatehw

import (
	//"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

//Fill in the data of the student and details about the test/homework
//into a *.tex file
func WriteData(txt, fn, gr, name *string, spec, info, date string) {
	WriteFN(txt, fn)
	WriteDate(txt, &date)
	WriteGr(txt, gr)
	WriteName(txt, name)
	WriteSpec(txt, &spec)
	WriteInfo(txt, &info)
}

//Write the FN of the student in the document
func WriteFN(txt, fn *string) {
	re := regexp.MustCompile(`(?i)<\*fn\*>`)
	*txt = re.ReplaceAllString(*txt, *fn)
}

//Write the Group number of the student in the document
func WriteGr(txt, gr *string) {
	re := regexp.MustCompile(`(?i)<\*gr\*>`)
	*txt = re.ReplaceAllString(*txt, *gr)
}

//Write the course, which the student is taking
func WriteSpec(txt, spec *string) {
	re := regexp.MustCompile(`(?i)<\*spec\*>`)
	*txt = re.ReplaceAllString(*txt, *spec)
}

//Write the name of the student in the document
func WriteName(txt, name *string) {
	re := regexp.MustCompile(`(?i)<\*name\*>`)
	*txt = re.ReplaceAllString(*txt, *name)
}

//Write info about the homework or the test
func WriteInfo(txt, info *string) {
	re := regexp.MustCompile(`(?i)<\*info\*>`)
	*txt = re.ReplaceAllString(*txt, *info)
}

//Write the date, that will be printed on the document
func WriteDate(txt, date *string) {
	re := regexp.MustCompile(`(?i)<\*date\*>`)
	*txt = re.ReplaceAllString(*txt, *date)
}

//Replace the coefficient in the file depending on it role. If it is a
// coefficient infront of a variable then use (c), if it is the leading
// coefficient the use (lc), if it is a free coefficient then use (fc)
func WriteCoeffsToFile(paramValue int, paramSign string, txt *string) {
	if paramValue > 1 {
		*txt = strings.Replace(*txt, "<*(c)"+paramSign+"*>", "+"+strconv.Itoa(paramValue), -1)
		*txt = strings.Replace(*txt, "<*(lc)"+paramSign+"*>", strconv.Itoa(paramValue), -1)
		*txt = strings.Replace(*txt, "<*"+paramSign+"*>", strconv.Itoa(paramValue), -1)
		*txt = strings.Replace(*txt, "<*(fc)"+paramSign+"*>", "+"+strconv.Itoa(paramValue), -1)
	} else if paramValue == 1 {
		*txt = strings.Replace(*txt, "<*(c)"+paramSign+"*>", "+", -1)
		*txt = strings.Replace(*txt, "<*(lc)"+paramSign+"*>", "", -1)
		*txt = strings.Replace(*txt, "<*"+paramSign+"*>", "1", -1)
		*txt = strings.Replace(*txt, "<*(fc)"+paramSign+"*>", "+1", -1)
	} else if paramValue == -1 {
		*txt = strings.Replace(*txt, "<*(c)"+paramSign+"*>", "-", -1)
		*txt = strings.Replace(*txt, "<*(lc)"+paramSign+"*>", "-", -1)
		*txt = strings.Replace(*txt, "<*"+paramSign+"*>", "-1", -1)
		*txt = strings.Replace(*txt, "<*(fc)"+paramSign+"*>", "-1", -1)
	} else {
		*txt = strings.Replace(*txt, "<*(c)"+paramSign+"*>", strconv.Itoa(paramValue), -1)
		*txt = strings.Replace(*txt, "<*(lc)"+paramSign+"*>", strconv.Itoa(paramValue), -1)
		*txt = strings.Replace(*txt, "<*"+paramSign+"*>", strconv.Itoa(paramValue), -1)
		*txt = strings.Replace(*txt, "<*(fc)"+paramSign+"*>", strconv.Itoa(paramValue), -1)
	}

}
