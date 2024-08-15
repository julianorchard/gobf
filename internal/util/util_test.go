package util

import "testing"

var SimpleTestFilePath string = "../../test/hello-world.bf"

func TestFileRead(t *testing.T) {
	fileContent, err := ReadFile(SimpleTestFilePath)
	if len(fileContent) <= 0 || err != nil {
		t.Fatalf(`ReadFile(%s) = empty file`, SimpleTestFilePath)
	}
}
