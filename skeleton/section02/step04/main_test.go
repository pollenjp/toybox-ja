package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	// TODO: file名から _test を覗く処理をしたい
	cmdList := []string{"go", "run", "main.go"}
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
