package run

import (
	"fmt"
	"os"

	"../lexer"
)

const maxCell = 30000

// Run from Tokens
func Run(ts []lexer.Token) {
	cells := make([]byte, maxCell)
	pointer := 0
	for i := 0; i < len(ts); i++ {
	START:
		t := ts[i]
		if !isInRange(pointer) {
			continue
		}

		switch t.Kind {
		case '>':
			pointer++
		case '<':
			pointer--
		case '+':
			cells[pointer]++
		case '-':
			cells[pointer]--
		case '.':
			fmt.Print(string(cells[pointer]))
		case ',':
			fmt.Scan(&cells[pointer])
		case '[':
			if cells[pointer] == 0 {
				for li := i; li < len(ts); li++ {
					if ts[li].Kind == ']' {
						i = li + 1
						goto START
					}
				}
				fmt.Fprintf(os.Stderr, `ERROR: 対応する "]" が見つかりませんでした %#v (%d %d)\n`, t.Kind, t.Line, t.Column)
			}
		case ']':
			if cells[pointer] != 0 {
				for li := i; li > 0; li-- {
					if ts[li].Kind == '[' {
						i = li + 1
						goto START
					}
				}
				fmt.Fprintf(os.Stderr, `ERROR: 対応する "]" が見つかりませんでした %#v (%d %d)\n`, t.Kind, t.Line, t.Column)
			}
		default:
			// fmt.Fprintf(os.Stderr, "WARNING: INVALID TOKEN %#v (%d %d) SKIPPED\n", t.Kind, t.Line, t.Column) // 不正トークン用
		}
	}
}

func isInRange(i int) bool {
	return 0 <= i && i < maxCell
}
