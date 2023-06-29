package main

import "testing"

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1)
	if err != nil {
		t.Error("got an error: ", err.Error(), " which we shouldnt have")
	}
}
