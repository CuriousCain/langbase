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
	inputText := "Hello my name is Joel"
	controlLength := 4

	pairs := GetPairs(inputText)

	if len(pairs) != controlLength {
		t.Fatal("Pairs length does not match. Got: ", len(pairs))
	}

	testHead := pairs[1].Head
	testTail := pairs[1].Tail

	if testHead != "my" {
		t.Fatal("Test Head does not match")
	}

	if testTail != "name" {
		t.Fatal("Test Tail does not match")
	}
}
