package bot

import (
	"bufio"
	"os"
	"strings"
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
		currentState: botStartingState,
	}
}

func (bot *Bot) Start() {
	for bot.currentState != nil {
		bot.currentState = bot.currentState(bot)
	}
}

func (bot *Bot) readLine() string {
	bytes, err := bot.rw.ReadBytes('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(bytes))
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
