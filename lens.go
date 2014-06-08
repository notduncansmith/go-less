package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"strconv"
)

type Lens struct {
	fileName        string
	buffer          []string
	top             int
	bottom          int
	margin          int
	showLineNumbers string
}

const (
	defaultColor = termbox.ColorDefault
)

func (l *Lens) setBuffer(lines []string) {
	newBuf := make([]string, 0, len(lines))

	for i, v := range lines {
		newBuf = append(newBuf, l.getMargin(i)+v)
	}

	l.buffer = newBuf
}

func (l *Lens) len() int {
	return len(l.buffer)
}

func (l *Lens) size() int {
	return l.top - l.bottom
}

func (l *Lens) printToBuffer() {
	termbox.Clear(defaultColor, defaultColor)

	screen := make([]string, 0, l.len()+1)

	screen = append(screen, l.buffer[l.top:l.bottom]...)
	screen = append(screen, ": ")

	for lineIndex, line := range screen {
		x := 0

		for _, chr := range line {
			termbox.SetCell(x, lineIndex, chr, defaultColor, defaultColor)
			x++
		}
	}
}

func (l *Lens) down() {
	if l.bottom < l.len() {
		l.bottom = l.bottom + 1
		l.top = l.top + 1
	}

	l.printToBuffer()
	termbox.Flush()
}

func (l *Lens) up() {
	if l.top > 0 {
		l.bottom = l.bottom - 1
		l.top = l.top - 1
	}

	l.printToBuffer()
	termbox.Flush()
}

func (l *Lens) getMargin(lineNumber int) string {
	if l.showLineNumbers == "N" {
		return marginWithLineNumbers(lineNumber)
	}

	return marginWithoutLineNumbers()
}

func marginWithoutLineNumbers() string {
	m := "  "

	for len(m) < Margin {
		m = " " + m
	}

	return m
}

func marginWithLineNumbers(lineNumber int) string {
	m := strconv.Itoa(lineNumber)

	for len(m) < Margin {
		m = " " + m
	}

	return m + "  "
}

func (l *Lens) listen() {
	err := termbox.Init()

	if err != nil {
		panic(err)
	}

	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	_, h := termbox.Size()
	l.bottom = h - 1

	for {
		l.printToBuffer()
		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Ch {
			case 'j':
				l.down()
			case 'k':
				l.up()
			case 'q':
				termbox.Close()
				os.Exit(0)
			default:
				fmt.Println("Yup, that's stuff all right")
			}
		}
	}
}
