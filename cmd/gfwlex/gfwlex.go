package main

import (
	"flag"
	"fmt"
	"larryr/gfw/lex"
	"os"
)

var debug bool
var tokcnt, stmtcnt, formcnt, addrcnt, identcnt int

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
		if debug {
			fmt.Printf("token: %v\n", t)
		}

		tokcnt++
		countTokTyps(t)

		if t.Typ == lex.TokEOF {
			break
		}
		if t.Typ == lex.TokError {
			fmt.Printf("error:%v\n", t)
			break
		}
	}

	//summarize
	fmt.Printf("statements:\t%v\n", stmtcnt)
	fmt.Printf("tokens:  \t%v\n", tokcnt)
	fmt.Printf("idents:  \t%v\n", identcnt)
	fmt.Printf("addrs:   \t%v\n", addrcnt)
	fmt.Printf("formulas:\t%v\n", formcnt)
}

func countTokTyps(t lex.Token) {
	switch t.Typ {
	case lex.TokENDSTMT:
		stmtcnt++
	case lex.TokFORMULA:
		formcnt++
	case lex.TokIDENTIFIER:
		identcnt++
	case lex.TokHOSTADDR:
		addrcnt++
	case lex.TokNETADDR:
		addrcnt++
	}
}
