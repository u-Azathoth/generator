// Package generator select random name from words dictionary
// You can use this for creating new files for example:
// mkdir $(generator) etc.
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/GeertJohan/go.rice"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getWordsDictionary() []string {
	templateBox, err := rice.FindBox("assets")
	check(err)

	data, err := templateBox.String("words.txt")
	check(err)

	return strings.Split(string(data), " ")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	words := getWordsDictionary()
	index := rand.Intn(len(words))

	fmt.Println(strings.ToLower(words[index]))
}
