package main

import "testing"

func TestGetFilter(t *testing.T) {

	got := getFilter()
	want := "filter"

	if want != got {
		t.Errorf("got %v want %v", got, want)
	}
}
