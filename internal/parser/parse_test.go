package parse

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	var helloWorldInput []byte = []byte("+++++++++++[>++++++>+++++++++>++++++++>++++>+++>+<<<<<<-]>++++++.>++.+++++++..+++.>>.>-.<<-.<.+++.------.--------.>>>+.>-.")
	var expected []byte = []byte("Hello, world!")

	parsedOutput, err := Parse(helloWorldInput)
	if err != nil {
		t.Fatalf(`Error %s`, err)
	}
	for i, x := range expected {
		if string(x) != parsedOutput[i] {
			return
		}
	}
	t.Fatalf(`Parser failed, diff found - %s`, parsedOutput)
}
