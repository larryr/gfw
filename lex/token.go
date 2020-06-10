//Copyright (C) 2020 Larry Rau. all rights reserved
package lex

import "fmt"

/*
TokenType Identifiers
*/
type Token struct {
	Typ TokenType //type id
	Val string
}

type TokenType int

func (item *Token) String() string {
	switch item.Typ {
	case TokError:
		return item.Val
	case TokEOF:
		return "EOF"
	case TokEOL:
		return "EOL"
	}
	if len(item.Val) > 10 {
		return fmt.Sprintf("%. 10q...", item.Val)
	}
	return fmt.Sprintf("%q", item.Val)
}

const (
	TokError TokenType = iota //lexing error
	TokEOF
	TokEOL
	TokIDENTIFIER
	//keywords
	TokOPTION
	TokINTERFACE
	TokALIAS
	TokFIREWALL
	TokPOLICY
	TokLOCAL
	TokTCP
	TokUDP
	TokICMP
	TokCUSTOM
	TokFORMULA
	//fw operators
	TokALLOW   // >
	TokTWALLOW // <>
	TokDROP    // /
	TokREJECT  // //
	// symbols
	TokSTAR   //"*"
	TokAT     //"@"
	TokLBRACK // [
	TokRBRACK // ]
	TokDOT    // .
	TokPIPE   // |
	// special values
	TokNETADDRESS  // octet.octet.octet.octet '/' subnet
	TokHOSTADDRESS // octet.octet.octet.octet
	TokPORT        // ':'<number>
	TokOCTET       // <digit> [<digit>] [<digit>]
	TokINT         // <number>
	TokSUBNETMASK  // <digit> [<digit>]
	TokCOMMENT     // '#' [<text>]\n
)
