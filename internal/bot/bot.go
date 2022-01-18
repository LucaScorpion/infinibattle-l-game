package bot

import (
	"bufio"
	"io"
	"os"
	"strings"
	"time"
)

type Bot struct {
	rw           *bufio.ReadWriter
	currentState stateFn
}

func NewBot() *Bot {
	return &Bot{
		rw: bufio.NewReadWriter(
			bufio.NewReader(os.Stdin),
			bufio.NewWriter(os.Stdout),
		),
		currentState: botStarting,
	}
}

func (bot *Bot) Start() {
	for bot.currentState != nil {
		bot.currentState = bot.currentState(bot)
	}
}

func (bot *Bot) expectLine(expected string) {
	bot.printComment("Waiting for: " + expected)
	line := bot.readLine()
	if line != expected {
		panic("Expected " + expected + " but got " + line)
	}
}

func (bot *Bot) readLine() string {
	bytes, err := bot.rw.ReadBytes('\n')
	line := string(bytes)

	// Keep reading until the actual end of line.
	for err == io.EOF {
		bot.printComment("Re-reading stdin")
		time.Sleep(10 * time.Millisecond)
		bytes, err = bot.rw.ReadBytes('\n')
		line += string(bytes)
	}

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(line)
}

func (bot *Bot) writeLine(s string) {
	if _, err := bot.rw.WriteString(s + "\n"); err != nil {
		panic(err)
	}
	if err := bot.rw.Flush(); err != nil {
		panic(err)
	}
}

func (bot *Bot) printComment(s string) {
	bot.writeLine("# " + s)
}
