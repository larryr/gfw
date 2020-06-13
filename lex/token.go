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
	TokDEFAULTS
	TokLOCAL
	TokTCP
	TokUDP
	TokICMP
	TokNATIVE
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
	TokHYPH   // -
	TokPIPE   // |
	TokLCURLY // {
	TokRCURLY // }
	TokSLIST  // ( start list
	TokELIST  // ) end list
	// special values
	TokNETADDR    // octet.octet.octet.octet '/' subnet
	TokHOSTADDR   // octet.octet.octet.octet
	TokPORT       // ':'<number>
	TokOCTET      // <digit> [<digit>] [<digit>]
	TokINT        // <number>
	TokSUBNETMASK // <digit> [<digit>]
	TokCOMMENT    // '#' [<text>]\n
	TokENDSTMT
)

func (tok Token) String() string {

	if len(tok.Val) > 20 {
		return fmt.Sprintf("%v: %.20q...", tok.Typ, tok.Val)
	}
	return fmt.Sprintf("%v: %q", tok.Typ, tok.Val)
}

func (tt TokenType) String() string {
	switch tt {
	case TokError:
		return "TokError"
	case TokEOF:
		return "TokEOF"
	case TokEOL:
		return "TokEOL"
	case TokIDENTIFIER:
		return "TokIDENTIFIER"
	case TokOPTION:
		return "TokOPTION"
	case TokINTERFACE:
		return "TokINTERFACE"
	case TokALIAS:
		return "TokALIAS"
	case TokFIREWALL:
		return "TokFIREWALL"
	case TokDEFAULTS:
		return "TokDEFAULTS"
	case TokLOCAL:
		return "TokLOCAL"
	case TokTCP:
		return "TokTCP"
	case TokUDP:
		return "TokUDP"
	case TokICMP:
		return "TokICMP"
	case TokNATIVE:
		return "TokNATIVE"
	case TokFORMULA:
		return "TokFORMULA"
	case TokALLOW:
		return "TokALLOW"
	case TokTWALLOW:
		return "TokTWALLOW"
	case TokDROP:
		return "TokDROP"
	case TokREJECT:
		return "TokREJECT"
	case TokSTAR:
		return "TokSTAR"
	case TokAT:
		return "TokAT"
	case TokLBRACK:
		return "TokLBRACK"
	case TokRBRACK:
		return "TokRBRACK"
	case TokDOT:
		return "TokDOT"
	case TokHYPH:
		return "TokHYPH"
	case TokPIPE:
		return "TokPIPE"
	case TokLCURLY:
		return "TokLCURLY"
	case TokRCURLY:
		return "TokRCURLY"
	case TokSLIST:
		return "TokSLIST"
	case TokELIST:
		return "TokELIST"
	case TokNETADDR:
		return "TokNETADDR"
	case TokHOSTADDR:
		return "TokHOSTADDR"
	case TokPORT:
		return "TokPORT"
	case TokOCTET:
		return "TokOCTET"
	case TokINT:
		return "TokINT"
	case TokSUBNETMASK:
		return "TokSUBNETMASK"
	case TokCOMMENT:
		return "TokCOMMENT"
	case TokENDSTMT:
		return "TokENDSTMT"
	}
	return "BAD"
}
