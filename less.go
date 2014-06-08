package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	Margin = 2
)

var (
	FrameSize = 20
)

func main() {
	name := os.Args[1]
	showLineNums := "F"

	if len(os.Args) == 3 {
		showLineNums = os.Args[2]
	}

	contents, err := ioutil.ReadFile(name)

	if err != nil {
		log.Fatal("Error reading from file " + name)
	}

	lines := fileToLines(contents)

	l := makeLens(name, lines, showLineNums)
	l.listen()
}

func makeLens(name string, lines []string, showLineNumbers string) Lens {
	l := Lens{}

	l.showLineNumbers = showLineNumbers
	l.setBuffer(lines)
	l.margin = Margin
	l.top = 0
	l.fileName = name

	return l
}

func fileToLines(contents []byte) []string {
	return strings.Split(string(contents), "\n")
}
