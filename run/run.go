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
		fmt.Scan(&s.cells[s.pointer])
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
