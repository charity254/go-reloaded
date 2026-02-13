package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	argu := os.Args[1:]

	if len(argu) < 2 {
		fmt.Println("You need to pass two file names")
		return
	}

	input := argu[0]
	r, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	content := string(r)
	tokens := tokenize(content)
	tokens = cap(lower(upper(bin(hex(tokens)))))

	
	// for _, char := range tokens {
	// 	fmt.Printf("%s\n", char)
	// }
	output := argu[1]
	result := strings.Join(tokens, " ") 

	err = os.WriteFile(output, []byte(result), 0644)
	if err!= nil {
		log.Fatal(err)
	}
}