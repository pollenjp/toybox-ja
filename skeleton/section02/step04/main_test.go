package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"testing"

	path "github.com/pollenjp/sandbox-go/2022-08-03-140501/sample/path"
)

func TestMain(t *testing.T) {
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		fmt.Printf("Called from %s, line #%d, func: %v\n",
			file, line, runtime.FuncForPC(pc).Name())
	}

	fpath := path.Path{Filepath: file}
	t.Logf("%s\n", fpath)

	re := regexp.MustCompile("(?P<stem>.*)_test")
	match := re.FindAllStringSubmatch(fpath.Stem(), 1)
	groupNames := re.SubexpNames()
	var groupIdx int
	for i, name := range groupNames {
		if name == "stem" {
			groupIdx = i
			break
		}
	}

	originPath := fpath.Parent().Join(match[0][groupIdx] + ".go")
	t.Logf("%s\n", originPath)

	cmdList := []string{"go", "run", originPath.String()}
	subProc := exec.Command(cmdList[0], cmdList[1:]...)

	input := strings.Join(
		[]string{
			"1",
			"2",
		},
		"\n",
	)
	t.Logf("input:\n%s", input)
	subProc.Stdin = strings.NewReader(input)

	output, err := subProc.Output()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	defer subProc.Wait()

	outputLines := strings.Split(string(output), "\n")
	t.Logf("output:\n%s", string(output))

	// fmt.Println(outputLines)
	lastLine := outputLines[len(outputLines)-2]

	if lastLine != "回答>正解!" {
		t.Errorf("Wanted: %v, Got: %v", input, string(output))
	}
}
