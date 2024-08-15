package parse

import (
	"fmt"
	"strings"
)

type loopState struct {
	active bool
	value  int
}

func Sanitize(content []byte) ([]byte, error) {
	// Returns a brainfuck only chars string (minus comments basically)
	var out []byte
	var brainfuckChars string = "><.-+[]"
	for _, c := range content {
		if strings.Contains(brainfuckChars, string(c)) {
			out = append(out, c)
		}
	}
	return out, nil
}

func Parse(content []byte) ([]string, error) {
	content, err := Sanitize(content)
	if err != nil {
		panic("fuck")
	}

	fmt.Println(string(content))

	// Collects outputs
	var out []string

	// TODO: Do not use an array
	// var memory []int = make([]int, 0)
	var memory [64]int
	var ptr int = 0

	// Flow control
	var ls loopState
	ls.active = false
	ls.value = 0

	// Continue counter
	var cc int

	for i := 0; i < len(content); i++ {
		// TODO: , operator
		if content[i] == '+' {
			memory[ptr] = memory[ptr] + 1
		} else if content[i] == '-' {
			if memory[ptr] >= 0 {
				memory[ptr] = memory[ptr] - 1
			}
		} else if content[i] == '>' {
			ptr = ptr + 1
		} else if content[i] == '<' {
			ptr = ptr - 1
		} else if content[i] == '.' {
			out = append(out, string(rune(memory[ptr])))
		} else if content[i] == '[' {
			if memory[ptr] == 0 {
				ls.active = false
				ls.value = 0
			} else {
				ls.active = false
				ls.value = i
			}
		} else if content[i] == ']' {
			if memory[ptr] != 0 {
				ls.active = true
				i = ls.value
			}
		} else {
			cc = cc + 1
			continue
		}
		// DEBUG:
		fmt.Printf("%v (skipped %v): %v\n", i, cc, memory)
	}
	return out, nil
}
