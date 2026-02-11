package main

import (
	"fmt"
	//"io"
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
	// fcontent := strings.Fields(content)
	// result := strings.Join(fcontent, " ")

	for i := 0
	// fmt.Println(fcontent)
	output := argu[1]
	err = os.WriteFile(output, []byte(result), 0644)
	if err!= nil {
		
		log.Fatal(err)
	}


}