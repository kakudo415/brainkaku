package run

import (
	"fmt"

	"../lexer"
	"../parser"
)

const cellMax = 30000

type status struct {
	cells   []byte
	pointer int
	buffer  string
}

// Run command
func Run(root *parser.Node) {
	s := new(status)
	s.cells = make([]byte, cellMax)
	s.run(root)
}

func (s *status) run(node *parser.Node) {
	if !isInRange(s.pointer) {
		return
	}
	switch t := node.Token; t.Kind {
	case lexer.ROOT:
		for _, child := range node.Children {
			s.run(child)
		}
	case lexer.NEXT:
		s.pointer++
	case lexer.PREV:
		s.pointer--
	case lexer.INCR:
		s.cells[s.pointer]++
	case lexer.DECR:
		s.cells[s.pointer]--
	case lexer.WRITE:
		fmt.Print(string(s.cells[s.pointer]))
	case lexer.READ:
		if len(s.buffer) == 0 {
			var buffer string
			fmt.Scanln(&buffer)
			s.buffer += buffer
		}
		if len(s.buffer) > 0 {
			rb := []rune(s.buffer)
			s.cells[s.pointer] = byte(rb[0])
			if len(s.buffer) > 1 {
				s.buffer = string(rb[1:])
			} else {
				s.buffer = ""
			}
		}
	case lexer.LOOP:
		for s.cells[s.pointer] != 0 {
			for _, child := range node.Children {
				s.run(child)
			}
		}
	default:
		panic("UNKNOWN TOKEN WHEN PARSING")
	}
}

func isInRange(i int) bool {
	return 0 <= i && i < cellMax
}
