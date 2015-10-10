package data

import (
	"testing"
)

func TestGetSentences(t *testing.T) {
	testInput := "Hello everybody. My name is Joel. What's your name, if I may ask?"
	sentences := GetSentences(testInput)

	if len(sentences) != 3 {
		t.Fatal("Sentences length was incorrect. Got %s", len(sentences))
	}
}
