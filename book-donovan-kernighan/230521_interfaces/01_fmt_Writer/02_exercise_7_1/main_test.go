package main

import (
	"testing"
)

func TestWrite(t *testing.T) {
	aText := "Using the ideas from ByteCounter, implement counters for words and for lines. You will find bufio.ScanWords useful."
	var c WordCounter = 0
	c.Write([]byte(aText))
	expected := 17
	result := c
	if expected != int(result) {
		t.Error("Expected", expected, "Got", result)
	}
}
