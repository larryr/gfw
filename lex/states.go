//Copyright (C) 2020 Larry Rau. all rights reserved
package lex

import (
	"fmt"
	"net"
	"strings"
	"unicode"
)

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

// isStmtStart returns true if the rune could start a statement
func isStmtStart(r rune) bool {
	if strings.IndexRune(runSTMTSTART, r) >= 0 {
		return true
	}
	return false
}

// lexer states

// start scanning for policy statement sections
// eat all whitespace and comments looking for a section identifier
func lexOuter(lex *Lexer) stateFn {

	// consume whitespace
	lex.acceptRun(runWhitespace)
	lex.ignore()

	switch r := lex.next(); {
	case r == symEOF:
		lex.emit(TokEOF)
		return nil
	case isAlphaNumeric(r):
		lex.backup()
		return lexSectionHeader

	case r == symHASH:
		// starting a comment; consume it.
		lex.pushStateFn(lexOuter)
		return lexComment
	}
	lex.flushStateFn()

	return lexOuter
}

// lexSection will scan the header of a section.
// Each section is delimited by '{' '}'
func lexSectionHeader(lex *Lexer) stateFn {
	// find an identifier
	lex.acceptRun(runAlphanumsym)
	ident := lex.input[lex.start:lex.pos]
	section, ok := sections[ident]
	if !ok {
		return lex.errorf("expected section name, got: %v", ident)
	}
	lex.emit(section)

	lex.consumeWhitespace()

	// expect open curly
	if lex.next() != symLCURLY {
		return lex.errorf("expected left brace, got:%v", lex.next())
	}
	lex.emit(TokLCURLY)

	lex.consumeWhitespace()

	if section == TokNATIVE {
		// we don't need/want to tokenize the statements of native; just pass along
		return lexSectionNative
	}
	return lexSection
}

// lexSection scans a non-native section which is a set of statements
// statements can only start with a small set of symbols
func lexSection(lex *Lexer) stateFn {

	lex.consumeWhitespace()

	switch r := lex.peek(); {
	case isStmtStart(r):
		return lexStatement
	case r == symHASH:
		lex.pushStateFn(lexSection)
		return lexComment
	case r == symRCURLY:
		lex.pos++
		lex.emit(TokRCURLY)
		return lexOuter
	}
	return lex.errorf("stopping here")
}

func lexStatement(lex *Lexer) stateFn {

	lex.consumeNonNLwhite()

	var r rune
	switch r = lex.peek(); {
	case unicode.IsLetter(r):
		lex.pushStateFn(lexStatement)
		return lexIdentifier
	case unicode.IsDigit(r):
		lex.pushStateFn(lexStatement)
		return lexNumbers
	case r == symCOLON:
		return lexPort
	case r == symHYPH:
		lex.pos++
		lex.emit(TokHYPH)
		return lexStatement
	case r == symLPAREN:
		lex.pos++
		lex.emit(TokSLIST)
		return lexList
	case isEndOfLine(r):
		lex.pos++
		lex.emit(TokENDSTMT)
		return lexSection
	case r == symHASH:
		lex.emit(TokENDSTMT)
		return lexSection
	}
	return lex.errorf("unexpected char:%v", r)
}

// lexIdentifier will consume a valid identifier; first char is present
func lexIdentifier(lex *Lexer) stateFn {
	lex.acceptRun(runAlphanumsym)
	word := lex.input[lex.start:lex.pos]
	tok, ok := allkeywords[word]
	if ok {
		lex.emit(tok)
	} else {
		lex.emit(TokIDENTIFIER)
	}
	return lex.popStateFn()
}

// lexNumbers scans host or net addresses or address ranges
func lexNumbers(lex *Lexer) stateFn {
	lex.acceptRun(runIPN)
	numb := lex.input[lex.start:lex.pos]
	if strings.IndexByte(numb, '/') >= 0 {
		//looks like a network
		_, ipnet, err := net.ParseCIDR(numb)
		if err != nil {
			return lex.errorf("bad network: %v", err)
		}
		lex.tokens <- Token{TokNETADDR, fmt.Sprintf("%v", ipnet)}
		lex.start = lex.pos
	} else {
		ip := net.ParseIP(numb)
		if ip == nil {
			return lex.errorf("bad host address: %s", numb)
		}
		lex.emit(TokHOSTADDR)
	}
	r := lex.peek()
	if !unicode.IsSpace(r) && r != symHYPH {
		return lex.errorf("expected space after addr")
	}
	return lex.popStateFn()
}

// lexList will scan a list of ident|addrs e.g. (anident,10.1.0.1)
func lexList(lex *Lexer) stateFn {
	//look for identfiers or numbers
	lex.consumeNonNLwhite()

	var r rune
	switch r := lex.peek(); {
	case unicode.IsLetter(r):
		lex.pushStateFn(lexList)
		return lexIdentifier
	case unicode.IsDigit(r):
		lex.pushStateFn(lexList)
		return lexNumbers
	case r == symCOMA:
		lex.pos++
		return lexList
	case r == symRPAREN:
		lex.pos++
		lex.emit(TokELIST)
		return lexStatement
	}
	return lex.errorf("in list unexpected char:%v", r)
}

// lexPort will scan a port number -- ':'<int>
func lexPort(lex *Lexer) stateFn {
	lex.pos++ //eat ':'
	lex.consumeWhitespace()
	lex.acceptRun(runDigits)
	lex.emit(TokPORT)
	if !unicode.IsSpace(lex.peek()) {
		return lex.errorf("expected space after addr")
	}
	return lexSection
}

// lexSectionNative will scan the native section as lines
//
func lexSectionNative(lex *Lexer) stateFn {

	lex.consumeWhitespace()

	switch r := lex.next(); {
	case r == symHASH:
		lex.pushStateFn(lexSectionNative)
		return lexComment
	case isEndOfLine(r):
		lex.emit(TokENDSTMT)
		return lexSectionNative
	case r == symRCURLY:
		lex.emit(TokRCURLY)
		return lexOuter
	default:
		// scan and emit each line as a separate Formula token
		//lex.consumeWhitespace()
		lex.acceptRunUntil(runLineTail)
		if lex.start < lex.pos {
			lex.emit(TokFORMULA)
			return lexSectionNative
		}
	}
	return lex.errorf("unexpected input")
}

// lexComment scans a comment. The comment marker is known to be present.
func lexComment(lex *Lexer) stateFn {
	lex.pos++                  //eat comment marker
	lex.acceptRunUntil(runEOL) //read until eol
	lex.pos++                  //eat the EOL
	lex.ignore()               //eat the comment
	return lex.popStateFn()    //return to prev state
}
