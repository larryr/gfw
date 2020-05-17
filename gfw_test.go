package gfw

import (
	"os"
	"testing"
)

func TestProcess(t *testing.T) {

	f, e := os.Open("test/samp/ex_warning_overlap.gfw")
	if e != nil {
		t.FailNow()
	}
	e = Process(f, os.Stdout)
	if e != nil {
		t.Log(e)
		t.FailNow()
	}
}
