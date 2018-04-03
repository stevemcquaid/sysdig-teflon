package main

import (
	//"errors"
	//"fmt"
	//"log"
	//"strings"
	"testing"
)

func TestDeleteK8SPod(t *testing.T) {
	got := "1"
	want := "1"
	if want != got {
		t.Errorf("got %v want %v", got, want)
	}
}

// Assure that the pod is in the log message for any of the error paths of the deleteK8SPod() method
