package main

import (
	"fmt"
	"larryr/gfw/lex"
	"os"
)

func main() {
	lexer := lex.NewLexer("gfw-lex", os.Stdin)

	for {
		t := lexer.NextToken()
		fmt.Printf("token: %v\n", t)
		if t.Typ == lex.TokEOF {
			break
		}
	}
}
