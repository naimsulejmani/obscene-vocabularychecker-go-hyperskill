package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	tabooWordsSet := map[string]struct{}{}

	var fileName string

	fmt.Scanln(&fileName)

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		tabooWordsSet[strings.ToLower(scanner.Text())] = struct{}{}
	}

	for {
		var saySomething string

		fmt.Scanln(&saySomething)
		if strings.EqualFold(saySomething, "exit") {
			break
		}

		words := strings.Split(saySomething, " ")
		for _, word := range words {
			word = strings.Trim(strings.ToLower(saySomething), ".,!?'\"[]()%#")
			if _, ok := tabooWordsSet[word]; ok {
				fmt.Println(strings.Repeat("*", len(saySomething)))
			} else {
				fmt.Println(saySomething)
			}
		}

	}
	fmt.Println("Bye!")
}
