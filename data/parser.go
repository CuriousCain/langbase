package data

import "strings"

type WordPair struct {
	head string
	tail string
}

func GetSentences(inputText string) []string {
	untrimmedSentences := strings.FieldsFunc(inputText, separateString) //TODO Recognize sentence type
	return trimSpaces(untrimmedSentences)
}

func GetPairs(inputText string) []WordPair {
	stringPairs := strings.Split(inputText, " ")
	pairsLength := len(stringPairs)

	var pair WordPair
	var pairs []WordPair

	for i := 0; i < pairsLength-1; i++ {
		pair.head = stringPairs[i]
		pair.tail = stringPairs[1+i]

		pairs = append(pairs, pair)
	}

	return pairs
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
