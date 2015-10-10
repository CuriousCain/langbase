package data

import "strings"

func GetSentences(inputText string) []string {
	untrimmedSentences := strings.FieldsFunc(inputText, separateString) //TODO Recognize sentence type
	return trimSpaces(untrimmedSentences)
}

func separateString(c rune) bool {
	return c == '.' || c == '!' || c == '?'
}

func trimSpaces(inputSentences []string) []string {
	var sentences []string

	for _, sentence := range inputSentences {
		sentences = append(sentences, strings.TrimSpace(sentence))
	}

	return sentences
}
