package main

import (
	"strconv"
	"strings"
	"unicode"
)
func isPunctuation (s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, char := range s {
		if !unicode.IsPunct(char) {
			return false
		}
	}
	return true
}

func hex (str  []string) []string {
	for i := 0; i < len(str); i++ {
		if len(str[i]) > 5 && str[i][:5] == "(hex," && str[i][len(str[i])-1] == ')' {
			num := str[i][5:len(str[i])-1]
			n, err := strconv.Atoi(strings.TrimSpace(num))
			if err != nil {
				continue
			}
			for j := 1; j <= n && i-j >= 0; j++ {
				value, err := strconv.ParseInt(str[i-j], 16, 64)
				if err != nil {
					continue
				}
				str[i-j]= strconv.Itoa(int(value))
			}
			str = append(str[:i], str[i+1:]... )
				i--
		}else if str[i] == "(hex)" && i > 0 {
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
		if len(str[i]) > 5 && str[i][:5] == "(bin," && str[i][len(str[i])-1] == ')' {
			num := str[i][5:len(str[i])-1]
			n, err := strconv.Atoi(strings.TrimSpace(num))
			if err != nil {
				continue
			}
			for j := 1; j <= n && i-j >= 0; j++ {
				value, err := strconv.ParseInt(str[i-j], 2, 64)
				if err != nil {
					continue
				}
				str[i-j] = strconv.Itoa(int(value))
			}
			str = append(str[:i], str[i+1:]... )
			i--
		} else if str[i] == "(bin)" && i > 0 {
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
		if len(str[i]) > 4 && str[i][:4] == "(up," && str[i][len(str[i])-1] == ')' {
			num := str[i][4:len(str[i])-1]
			n, err := strconv.Atoi(strings.TrimSpace(num))
			if err != nil {
				continue
			}
			for j := 1; j <= n && i-j >= 0; j++ {
				value := strings.ToUpper(str[i-j])
				str[i-j] = value
			}
			str = append(str[:i], str[i+1:]... )
			i--
		}else if str[i] == "(up)" && i > 0 {
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
		if len(str[i]) > 5 && str[i][:5] == "(low," && str[i][len(str[i])-1] == ')' {
			num := str[i][5:len(str[i])-1]
			n, err := strconv.Atoi(strings.TrimSpace(num))
			if err != nil {
				continue
			}
			for j := 1; j <= n && i-j >= 0; j++ {
				value := strings.ToLower(str[i-j])
				str[i-j] = value
			}
			str = append(str[:i], str[i+1:]... )
			i--
		}else if str[i] == "(low)" && i > 0 {
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
		if len(str[i]) > 5 && str[i][:5] == "(cap," && str[i][len(str[i])-1] == ')' {
			num := str[i][5:len(str[i])-1]
			n, err := strconv.Atoi(strings.TrimSpace(num))
			if err != nil {
				continue
			}
			for j := 1; j <= n && i-j >= 0; j++ {
				word := str[i-j]
				if len(word) == 0 {
					continue
				}
				if len(word) == 1 {
					word = strings.ToUpper(word)
					str[i-j] = word
				} else{
					fLet := strings.ToUpper(string(word[0]))
					rWord := strings.ToLower(string(word[1:]))
					word = fLet+rWord
					str[i-j] = word
				}
			}
			str= append(str[:i], str[i+1:]... )
			i--
		}else if str[i] == "(cap)" && i > 0 {
			word := str[i-1]
			if len(word) == 0 {
				continue
			} 
			if len(word) == 1 {
				word = strings.ToUpper(word)
				str[i-1] = word
			} else{
				fLet := strings.ToUpper(string(word[0]))
				rWord := strings.ToLower(string(word[1:]))
				word = fLet+rWord
				str[i-1] = word
			}

			str = append(str[:i], str[i+1:]... )
			i--
		}
	}
	return str
}

func punctuation(words []string) string {
	if len(words) == 0 {
		return ""
	}
	var result strings.Builder

	insideQuotes := false

	for i:=0; i<len(words); i++ {
		currentWord := words[i]

		isQuote := currentWord == "'"
		if i > 0 {
			prevWord := words[i-1]
			prevIsQuote := prevWord == "'"

			curIsPunc := isPunctuation(currentWord)

			addSpace := true

			if prevIsQuote && insideQuotes {
				addSpace = false
			} else if isQuote && insideQuotes {
				addSpace = false
			} else if curIsPunc && !isQuote {
				addSpace = false
			}
			if addSpace {
				result.WriteString(" ")
			}
		}
		result.WriteString(currentWord)

		if isQuote {
			insideQuotes = !insideQuotes
		}
	}
	return result.String()
}

func article(str []string) []string {
	if len(str) == 0 {
		return str
	}
	for i :=0; i < len(str); i++ {
		curWord := strings.ToLower(str[i])

		if curWord == "a" {
			if i+1 < len(str) {
				nextWord := str[i+1]
				if len(nextWord) > 0 {
					fChar := strings.ToLower(string(nextWord[0]))

					if fChar == "a" || fChar == "e" || fChar == "i" || fChar == "o" || fChar == "u" || fChar == "h" {
						if str[i] == "A" {
							str[i] = "An"
						} else {
							str[i] = "an"
						}
					}
				}
			}
		}
	}
	return str
}