package main

import (
	"testing"
)

func sample(t testing.T) bool {

	data := 1
	data2 := 2

	if data != data2 {
		t.Error("hey done")
		return false
	}
	return true
}
