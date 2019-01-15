package codegen

import (
	"fmt"
	"os"

	"../lexer"
)

// Gen -erate assembly code
func Gen(ts []lexer.Token) string {
	var code string
	code += ".intel_syntax noprefix\n"
	code += ".global main\n"
	code += "main:\n"
	for _, t := range ts {
		switch t.Kind {
		case '>':
		case '<':
		case '+':
		case '-':
		case '.':
		case ',':
		case '[':
		case ']':
		default:
			fmt.Fprintf(os.Stderr, "WARNING: INVALID TOKEN %#v (%d %d)\n", t.Kind, t.Line, t.Column)
		}
	}
	return code
}
