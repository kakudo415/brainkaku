package parser

import (
	"../lexer"
)

// Node of AST
type Node struct {
	Token    lexer.Token
	Children []*Node
}

type status struct {
	ts  []lexer.Token
	pos int
}

// Parser (make AST)
func Parser(ts []lexer.Token) *Node {
	s := new(status)
	s.ts = ts
	root := s.parse()
	root.Token = lexer.Token{Kind: lexer.ROOT}
	return root
}

func (s *status) parse() *Node {
	node := new(Node)
	for s.pos < len(s.ts) {
		switch c := s.current(); c.Kind {
		case lexer.EOF:
			return node
		case lexer.INVALID:
			// WARNING...?
			s.pos++
		case lexer.LOOPS:
			s.pos++
			loop := s.parse()
			loop.Token = lexer.Token{Kind: lexer.LOOP}
			node.Children = append(node.Children, loop)
		case lexer.LOOPE:
			s.pos++
			return node
		default:
			node.Children = append(node.Children, &Node{Token: c})
			s.pos++
		}
	}
	return node
}

func (s *status) current() lexer.Token {
	if 0 <= s.pos && s.pos < len(s.ts) {
		return s.ts[s.pos]
	}
	return lexer.Token{Kind: lexer.EOF}
}
