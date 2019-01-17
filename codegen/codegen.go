package codegen

import (
	"../parser"
)

// Gen -erate assembly code
func Gen(root *parser.Node) string {
	var code string
	code += ".intel_syntax noprefix\n"
	code += ".global main\n"
	code += "main:\n"
	return code
}
