package markdown_test

import (
	"os"
	"testing"
)

func TestPWD(t *testing.T) {
	workingDir, err := os.Getwd()
	if err != nil {
		t.Fail()
	}
	t.Log(workingDir)
}
