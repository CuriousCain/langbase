package data

import (
	"testing"
)

func TestGetSentences(t *testing.T) {
	controlSentence := "My name is Joel"
	controlLength := 4

	testInput := "Hello everybody. My name is Joel. What's your name, if I may ask? Oh that's great!"
	sentences := GetSentences(testInput)

	if len(sentences) != controlLength {
		for _, s := range sentences {
			t.Log(s)
		}
		t.Fatal("Sentences length was incorrect. Got", len(sentences))
	}

	if sentences[1] != controlSentence {
		t.Fatal("Sentence does not match. Expected:", controlSentence, "but got:", sentences[1])
	}
}

func TestGetPairs(t *testing.T) {

}
