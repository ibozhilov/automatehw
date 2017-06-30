package automatehw

import (
	"io/ioutil"
	"regexp"
	"strings"
)

type Student struct {
	Name  string
	Fn    string
	Group string
}

//Generate array of Students from inputfile
func Read(filepath string) []Student {
	fnTemplate := regexp.MustCompile(`\d{5,6}`)
	nameTemplate := regexp.MustCompile(`[А-я]+ [А-я]+ [А-я]+`)
	groupTemplate := regexp.MustCompile(`Група (\d)`)
	bytelist, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	csv := string(bytelist)
	var line, fn, name, group string
	var studentList []Student
	for _, value := range csv {
		if string(value) != "\n" {
			line = line + string(value)
			continue
		}
		fn = fnTemplate.FindString(line)
		name = strings.Replace(nameTemplate.FindString(line), ",", " ", -1)
		group = groupTemplate.FindStringSubmatch(line)[1]
		studentList = append(studentList, Student{name, fn, group})
		line = ""
	}
	return studentList
}
