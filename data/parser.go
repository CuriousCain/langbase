package data

import "strings"

func GetSentences(inputText string) []string {
	sentences := strings.Split(inputText, ".") //TODO: Don't only split on period

	return sentences
}
