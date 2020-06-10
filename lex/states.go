//Copyright (C) 2020 Larry Rau. all rights reserved
package lex

import "unicode"

const (
	runWhitespace  = " \n\t\v\f\r"
	runDigits      = "0123456789"
	runAlpha       = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	runAlphanum    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" + runDigits
	runAlphanumsym = runAlphanum + "_"
	runIPH         = runDigits + "."  //host address
	runIPN         = runDigits + "./" //network address range
	runSYMS        = "*@[].:|#"

	// special symbols
	symSTAR   = "*"
	symAT     = "@"
	symLBRACK = "["
	symRBRACK = "]"
	symDOT    = "."
	symCOLON  = ":"
	symPIPE   = "|"
	symHASH   = '#'
)

// keywords -- well known identifiers
var sections = map[string]TokenType{
	"OPTIONS":    TokOPTION,
	"INTERFACES": TokINTERFACE,
	"ALIASES":    TokALIAS,
	"FIREWALL":   TokFIREWALL,
	"POLICIES":   TokPOLICY,
	"CUSTOM":     TokCUSTOM,
}

var keywords = map[string]TokenType{
	"local": TokLOCAL,
	"tcp":   TokTCP,
	"udp":   TokUDP,
	"icmp":  TokICMP,
}

// firewall operators -- symbols for fw statements
var fwops = map[string]TokenType{
	">":  TokALLOW,
	"<>": TokTWALLOW,
	"/":  TokDROP,
	"//": TokREJECT,
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

// lexer states

// start scanning for policy statements
func lexStatements(lex *Lexer) stateFn {

	switch r := lex.next(); {

	case isAlphaNumeric(r):
		lex.backup()
		return lexIdentOrIPS

	case r = symHASH:
		return lexComment
	}

}

// lexComment scans a comment. The comment marker is known to be present.
func lexComment(lex *lexer) stateFn {
	lex.pos++ 
	i := strings.Index(l.input[l.pos:], rightComment)
	if i < 0 {
		return l.errorf("unclosed comment")
	}
	l.pos += Pos(i + len(rightComment))
	if !strings.HasPrefix(l.input[l.pos:], l.rightDelim) {
		return l.errorf("comment ends before closing delimiter")

	}
	l.pos += Pos(len(l.rightDelim))
	l.ignore()
	return lexText
}
