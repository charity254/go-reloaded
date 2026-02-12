package main

import (
	"strings"
	"unicode"

)
func tokenize (str string) []string {
	var tokens []string
	var word strings.Builder

	characters := []rune(str) //convert the str into a slice of runes

	for i := 0; i < len(characters); i++ {
		char := characters[i]
	
	if unicode.IsLetter(char) ||unicode.IsNumber(char) {
		word .WriteRune(char)
	} else if unicode.IsSpace(char) { //the end of a word
		if word.Len() > 0 {
			tokens = append(tokens, word.String())
			word.Reset()
		}
	} else if char == '(' { //save any word being build
		if word.Len() > 0 {
			tokens = append(tokens, word.String())
			word.Reset()
		}
		for j := i +1; j < len(characters); j++ { //finding the closing bracket
			if characters[j] == ')' {
				bracket := string(characters[i : j+1]) //
				tokens = append(tokens, bracket)
				// move past the brackets section
				i = j
				break
			}
		}
	} else if unicode.IsPunct(char) {
		if word.Len() > 0 {
			tokens = append(tokens, word.String())
			word.Reset()
		}
		tokens = append(tokens, string(char))
	} else {
		word.WriteRune(char)
		}
	}
	if word.Len() > 0 {
		tokens = append(tokens, word.String())
	}
	return tokens
}