package gfw

import (
	"os"
	"testing"
)

func TestLexer(t *testing.T) {

	f, e := os.Open("test/samp/ex_warning_overlap.gfw")
	if e != nil {
		t.FailNow()
	}
	// read all of file into a string

	e = Process(f, os.Stdout)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}
}
