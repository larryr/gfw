package main

import (
	"flag"
	"fmt"
	"larryr/gfw/lex"
	"os"
)

var debug bool

func main() {
	flag.BoolVar(&debug, "d", false, "debug output")
	flag.Parse()

	lex.Debug(debug)

	lexer, err := lex.NewLexer("gfw-lex", os.Stdin)
	if err != nil {
		fmt.Printf("input error: %v\n", err)
		return
	}
	for {
		t := lexer.NextToken()
		fmt.Printf("token: %v\n", t)
		if t.Typ == lex.TokEOF {
			break
		}
		if t.Typ == lex.TokError {
			break
		}
	}
}
