package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"./codegen"
	"./lexer"
	"./run"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, `USAGE: bk <command "run" or "make"> [file]`)
	}
	input, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s を開けませんでした", os.Args[2])
	}

	ts := lexer.Lexer(input)

	switch os.Args[1] {
	case "run":
		run.Run(ts)
	case "make":
		fmt.Print(codegen.Gen(ts))
	}
}
