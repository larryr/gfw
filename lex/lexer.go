//Copyright (C) 2020 Larry Rau. all rights reserved
package lex

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"unicode/utf8"
)

/*
Lexer for the gfw policy language.
Core lexer driver.

(pretty much lifted from a Rob Pike presentation)
*/

type Lexer struct {
	name   string     //dbg aid
	input  string     //input stream of runes
	start  int        //start of current token
	pos    int        //current position in rune-stream
	width  int        //current width of in-process rune run
	state  stateFn    //state machine element
	tokens chan Token //place to emit new tokens
	scope  []stateFn  //stack of scopes for returning to prev statefn
}

type stateFn func(*Lexer) stateFn

// create a lexer to tokenize the input string
func NewLexer(name string, in io.Reader) (*Lexer, error) {
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}
	l := &Lexer{
		name:   name,
		input:  string(buf),
		state:  lexOuter,
		tokens: make(chan Token, 2),
	}
	return l, nil
}

//debugging
var debug bool = false

func Debug(d bool) {
	debug = d
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

	var r rune
	if lex.pos >= len(lex.input)-1 {
		lex.width = 0
		r = symEOF
	} else {
		r, lex.width = utf8.DecodeRuneInString(lex.input[lex.pos:])
		lex.pos += lex.width
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

// consume a run of runes while rune is valid
func (lex *Lexer) acceptRun(valid string) {
	//keep consuming until a non-valid char is found
	for strings.IndexRune(valid, lex.next()) >= 0 {
	}
	lex.backup()
}

// consume a run of runes until match rune found
func (lex *Lexer) acceptRunUntil(match string) {
	//keep consuming until the invalid is found
	for strings.IndexRune(match, lex.next()) < 0 {
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

// consume whitespace by accepting then ignoring
// if none find silently return
func (lex *Lexer) consumeWhitespace() {
	s := lex.pos
	lex.acceptRun(runWhitespace)
	if lex.pos > s {
		lex.ignore() //ignore if changed
	}
}

// consume whitespace by accepting then ignoring
// if none find silently return
func (lex *Lexer) consumeNonNLwhite() {
	s := lex.pos
	lex.acceptRun(runNonNLwhite)
	if lex.pos > s {
		lex.ignore() //ignore if changed
	}
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (lex *Lexer) errorf(format string, args ...interface{}) stateFn {
	lex.tokens <- Token{TokError, fmt.Sprintf(format, args...)}
	return nil
}

// return last scope state
func (lex *Lexer) popStateFn() stateFn {
	l := len(lex.scope)
	if l == 0 {
		return nil
	}
	fn := lex.scope[l-1]        //last scope
	lex.scope = lex.scope[:l-1] //pop it
	return fn
}

// push a scope state
func (lex *Lexer) pushStateFn(fn stateFn) {
	lex.scope = append(lex.scope, fn)
}

// flush the scope stack
func (lex *Lexer) flushStateFn() {
	lex.scope = lex.scope[:0]
}
