//Copyright (C) 2020 Larry Rau. all rights reserved
package lex

/*
Define all the symbols and keywords and mapping to tokens.

These are ony meant for the internal lexer.
*/

const (
	runWhitespace  = " \n\t\v\f\r"
	runNonNLwhite  = " \t\v\f" //non \n \r whitepsace
	runDigits      = "0123456789"
	runAlpha       = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	runAlphanum    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" + runDigits
	runAlphanumsym = runAlphanum + "_"
	runIPH         = runDigits + "."  //host address
	runIPN         = runDigits + "./" //network address range
	runSYMS        = "*@[].:|#"
	runEOL         = "\r\n"            //end of lines
	runLineTail    = runEOL + "|#"     //several lines can end with comment or formula
	runSTMTSTART   = runAlphanum + "*" //potential start of statements

	// special symbols
	symSTAR   = '*'
	symAT     = '@'
	symLBRACK = '['
	symRBRACK = ']'
	symDOT    = '.'
	symCOLON  = ':'
	symPIPE   = '|'
	symHYPH   = '-'
	symHASH   = '#'
	symLCURLY = '{'
	symRCURLY = '}'
	symLPAREN = '('
	symRPAREN = ')'
	symCOMA   = ','
	symEOF    = -1
)

// keywords -- well known identifiers
var sections = map[string]TokenType{
	"options":    TokOPTION,
	"interfaces": TokINTERFACE,
	"aliases":    TokALIAS,
	"firewall":   TokFIREWALL,
	"defaults":   TokDEFAULTS,
	"native":     TokNATIVE,
}

var keywords = map[string]TokenType{
	"local": TokLOCAL,
	"tcp":   TokTCP,
	"udp":   TokUDP,
	"icmp":  TokICMP,
}

var allkeywords = map[string]TokenType{
	"options":    TokOPTION,
	"interfaces": TokINTERFACE,
	"aliases":    TokALIAS,
	"firewall":   TokFIREWALL,
	"defaults":   TokDEFAULTS,
	"native":     TokNATIVE,
	"local":      TokLOCAL,
	"tcp":        TokTCP,
	"udp":        TokUDP,
	"icmp":       TokICMP,
}

// firewall operators -- symbols for fw statements
var fwops = map[string]TokenType{
	">":  TokALLOW,
	"<>": TokTWALLOW,
	"/":  TokDROP,
	"//": TokREJECT,
}
