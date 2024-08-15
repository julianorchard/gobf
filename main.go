package main

import (
	"fmt"
	"strings"

	parser "github.com/julianorchard/bfgo/internal/parser"
	"github.com/julianorchard/bfgo/internal/util"
)

func main() {
	file, err := util.ReadFile("./test/hello-world-verbose.bf")
	if err != nil {
		panic(err)
	}
	output, err := parser.Parse(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Join(output, ""))
}
