package lexer

import (
	"unicode/utf8"
)

// Token Type
type Token struct {
	Kind rune

	Line   int
	Column int
}

type status struct {
	src []byte

	pos    int
	line   int
	column int
}

// Lexer func (Tokenizer)
func Lexer(input []byte) []Token {
	var ts []Token

	var s = new(status)
	s.src = input
	s.line = 1
	s.column = 1

	for s.current() > 0 {
		var t Token
		switch s.current() {
		case '>':
			t.Kind = '>'
			t.Line, t.Column = s.line, s.column
		case '<':
			t.Kind = '<'
			t.Line, t.Column = s.line, s.column
		case '+':
			t.Kind = '+'
			t.Line, t.Column = s.line, s.column
		case '-':
			t.Kind = '-'
			t.Line, t.Column = s.line, s.column
		case '.':
			t.Kind = '.'
			t.Line, t.Column = s.line, s.column
		case ',':
			t.Kind = ','
			t.Line, t.Column = s.line, s.column
		case '[':
			t.Kind = '['
			t.Line, t.Column = s.line, s.column
		case ']':
			t.Kind = ']'
			t.Line, t.Column = s.line, s.column
		default: // INVALID TOKEN
			t.Kind = '?'
			t.Line, t.Column = s.line, s.column
		}

		if s.current() == '\n' {
			s.next()
			s.line++
			s.column = 1
		} else {
			s.next()
			s.column++
		}

		ts = append(ts, t)
	}

	return ts
}

func (s *status) current() rune {
	if s.pos >= len(s.src) {
		return -1
	}
	r, _ := utf8.DecodeRune(s.src[s.pos:])
	if r == utf8.RuneError {
		return -1
	}
	return r
}

func (s *status) next() {
	size := utf8.RuneLen(s.current())
	if size < 0 {
		return
	}
	s.pos += size
}
