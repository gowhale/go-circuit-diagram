// Package main contains code to do with ensuring coverage is over 80%
package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const (
	minPercentCov = 80.0

	coverageStringNotFound = -1.0
	firstItemIndex         = 1
	floatByteSize          = 64
	emptySliceLen          = 0
	lenOfPercentChar       = 1
	indexOfEmptyLine       = 1
)

var excludedPkgs = map[string]bool{
	"golang-repo-template":           true,
	"golang-repo-template/pkg/fruit": true,
}

func main() {
	if err := run(&execute{}); err != nil {
		log.Fatalln(err)
	}
	log.Println("Tests=PASS Coverage=PASS")
}

//go:generate go run github.com/vektra/mockery/cmd/mockery -name executer -inpkg --filename executer_mock.go
type executer interface {
	runGoTest() (string, error)
	covertOutputToCoverage(termOutput string) ([]testLine, error)
	validateTestOutput(tl []testLine, o string) error
}

func run(e executer) error {
	output, err := e.runGoTest()
	if err != nil {
		log.Println(output)
		return err
	}

	tl, err := e.covertOutputToCoverage(output)
	if err != nil {
		return err
	}

	return e.validateTestOutput(tl, output)
}

var execCommand = exec.Command

type execute struct{}

func (*execute) runGoTest() (string, error) {
	cmd := execCommand("go", "test", "./...", "--cover")
	output, err := cmd.CombinedOutput()
	termOutput := string(output)
	return termOutput, err
}

type testLine struct {
	pkgName   string
	coverage  float64
	coverLine bool
}

func getCoverage(line string) (testLine, error) {
	if !strings.Contains(line, "go: downloading") {
		pkgName := strings.Fields(line)[firstItemIndex]
		if _, ok := excludedPkgs[pkgName]; !ok {
			coverageIndex := strings.Index(line, "coverage: ")
			if coverageIndex != coverageStringNotFound {
				lineFields := strings.Fields(line[coverageIndex:])
				pkgPercentStr := lineFields[firstItemIndex][:len(lineFields[firstItemIndex])-lenOfPercentChar]
				pkgPercentFloat, err := strconv.ParseFloat(pkgPercentStr, floatByteSize)
				if err != nil {
					return testLine{}, err
				}
				log.Println(pkgPercentStr)
				return testLine{pkgName: pkgName, coverage: pkgPercentFloat, coverLine: true}, nil
			}
			return testLine{pkgName: pkgName, coverage: coverageStringNotFound, coverLine: true}, nil
		}
	}
	return testLine{coverLine: false}, nil
}

func (*execute) covertOutputToCoverage(termOutput string) ([]testLine, error) {
	testStruct := []testLine{}
	lines := strings.Split(termOutput, "\n")
	for _, line := range lines[:len(lines)-indexOfEmptyLine] {
		tl, err := getCoverage(line)
		if err != nil {
			return nil, err
		}
		if tl.coverLine {
			testStruct = append(testStruct, tl)
		}
	}

	return testStruct, nil
}

func (*execute) validateTestOutput(tl []testLine, o string) error {
	invalidOutputs := []string{}
	for _, line := range tl {
		switch {
		case !line.coverLine:
			invalidOutputs = append(invalidOutputs, fmt.Sprintf("pkg=%s is missing tests", line.pkgName))
		case line.coverage < minPercentCov:
			invalidOutputs = append(invalidOutputs, fmt.Sprintf("pkg=%s cov=%f under the %f%% minimum line coverage", line.pkgName, line.coverage, minPercentCov))
		}
	}
	if len(invalidOutputs) == emptySliceLen {
		return nil
	}
	log.Println(o)
	log.Println("###############################")
	log.Println("invalid pkg's:")
	for i, invalid := range invalidOutputs {
		log.Printf("id=%d problem=%s", i, invalid)
	}
	log.Println("###############################")
	return fmt.Errorf("the following pkgs are not valid: %+v", invalidOutputs)
}
