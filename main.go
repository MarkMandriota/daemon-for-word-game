package main

import (
	"github.com/atotto/clipboard"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var list = make(map[rune][]string)

func init() {
	if len(os.Args) < 2 {
		panic("No input file specified.")
	}

	words, _ := ioutil.ReadFile(os.Args[1])
	for _, word := range strings.Fields(string(words)) {
		list[[]rune(word)[0]] = append(list[[]rune(word)[0]], word)
	}
}

func main() {
	var word string
	for word != "exorcise" {
		for i := len([]rune(word)) - 1; i >= 0; i-- {
			if v, has := list[[]rune(strings.ToLower(word))[i]]; has {
				rand.Seed(time.Now().UnixNano())
				clipboard.WriteAll(v[rand.Intn(len(v)-1)])

				break
			}
		}

		<-time.After(450 * time.Millisecond)
		word, _ = clipboard.ReadAll()
	}
}