package main

import "testing"

func TestTempInF(t *testing.T) {
	actual := tempInF(273.15)
	expected := 32.0
	if actual != expected {
		t.Error("Expected temp in F to eql %v, was %v", actual, expected)
	}
}
