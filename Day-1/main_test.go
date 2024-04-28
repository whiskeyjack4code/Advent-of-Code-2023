package main

import (
	"os"
	"testing"
)

func TestHelpSantaGodDamnIt(t *testing.T) {
	fileName := "inputTest.txt"
	file, err := os.Open(fileName)

	if err != nil {
		t.Errorf("could not open file: %s, got error %s", fileName, err)
	}

	res := HelpSantaGodDamnIt(file)
	expected := 455

	if res != expected {
		t.Errorf("expected %d but got %d", expected, res)
	}
}
