package main

import (
	"strconv"
	"strings"
)

func hex (str  []string) []string {
	for i := 0; i < len(str); i++ {
		if str[i] == "(hex)" && i > 0 {
			value, err := strconv.ParseInt(str[i-1], 16, 64)
			 if err != nil {
				continue
			 }
			 str[i-1] = strconv.Itoa(int(value))
			 str = append(str[:i], str[i+1:]...)
			 i--
		}
	}
	return str
}

func bin (str []string) []string {
	for i := 0; i < len(str); i++ {
		if str[i] == "(bin)" && i > 0 {
			value, err := strconv.ParseInt(str[i-1], 2, 64)

			if err != nil {
				continue
			}
			str[i-1] = strconv.Itoa(int(value))
			str = append(str[:i], str[i+1:]... )
			i--
		}
	}
	return str
}

func upper (str []string) []string {
	for i := 0; i < len(str); i++ {
		if str[i] == "(up)" && i > 0 {
			value := strings.ToUpper(str[i-1])
			str[i-1] = value
			str = append(str[:i], str[i+1:]... )
			i--
		}
		
	}
	return str
}

func lower (str []string) []string {
	for i := 0; i < len(str); i++ {
		if str[i] == "(low)" && i > 0 {
			value := strings.ToLower(str[i-1])
			str[i-1] = value
			str = append(str[:i], str[i+1:]... )
			i--
		}
		
	}
	return str
}

func cap (str []string) []string {
	for i := 0; i < len(str); i++ {
		if str[i] == "(cap)" && i > 0 {
			word := str[i-1]
			if len(word) == 0 {
				continue
			} 
			if len(word) == 1 {
				word = strings.ToUpper(word)
				str[i-1] = word
			}
			fLet := strings.ToUpper(string(word[0]))
			rWord := strings.ToLower(string(word[1:]))
			word = fLet+rWord
			str[i-1] = word

			str = append(str[:i], str[i+1:]... )
			i--
		}
	}
	return str
}