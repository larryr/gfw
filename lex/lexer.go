//Copyright (C) 2020 Larry Rau. all rights reserved
package lex

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

/*
Lexer for the gfw policy language.
Core lexer driver.

(pretty much lifted from a Rob Pike presentation)
*/

type Lexer struct {
	name   string //dbg aid
	in     *bufio.Scanner
	input  string     //input stream of runes
	start  int        //start of current token
	pos    int        //current position in rune-stream
	width  int        //current width of in-process rune run
	state  stateFn    //state machine element
	tokens chan Token //place to emit new tokens
}

type stateFn func(*Lexer) stateFn

// create a lexer to tokenize the input string
func NewLexer(name string, in io.Reader) *Lexer {
	scan := bufio.NewScanner(in)
	l := &Lexer{
		name:   name,
		in:     scan,
		state:  lexText,
		tokens: make(chan Token, 2),
	}
	return l
}

// return the next item
// read item token from channel and return; or change state if no token.
func (lex *Lexer) NextToken() Token {
	for {
		select {
		case item := <-lex.tokens:
			return item
		default:
			lex.state = lex.state(lex)
		}
	}
	panic("not reached")
}

// emit the new item to the channel of tokens
func (lex *Lexer) emit(it TokenType) {
	lex.tokens <- Token{it, lex.input[lex.start:lex.pos]}
	lex.start = lex.pos
}

// get the next rune
func (lex *Lexer) next() rune {
	if lex.pos >= len(lex.input) {
		lex.width = 0
		// try to fill buffer
		return -1 //eof
	}
	var r rune
	r, lex.width = utf8.DecodeRuneInString(lex.input[lex.pos:])
	lex.pos += lex.width
	// fill buffer if empty
	if lex.pos >= len(lex.input) {
		s := lex.nextLine()
		if s != "" {
			lex.input = lex.input + s
		}
	}
	return r //a rune
}

// consume next rune; if in valid set
func (lex *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, lex.next()) >= 0 {
		return true
	}
	lex.backup()
	return false
}

// consume a run of runes
func (lex *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, lex.next()) >= 0 {
	}
	lex.backup()
}

// ignore text up to current position
func (lex *Lexer) ignore() {
	lex.start = lex.pos
}

// move back in stream
func (lex *Lexer) backup() {
	lex.pos -= lex.width
}

// return the next rune without advancing
func (lex *Lexer) peek() rune {
	r := lex.next()
	lex.backup()
	return r
}

func (lex *Lexer) nextLine() string {
	if lex.in.Scan() {
		return lex.in.Text() + "\n" //add '\n' back
	}
	if err := lex.in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return ""
}
