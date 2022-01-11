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
		currentState: awaitBotStart,
	}
}

func (bot *Bot) Start() {
	for bot.currentState != nil {
		bot.currentState = bot.currentState(bot)
	}
}

func (bot *Bot) expectCommand(expected string) string {
	bot.printComment("Waiting for command: " + expected)
	cmd, arg := bot.readCommand()
	if cmd != expected {
		panic("Expected command " + expected + " but got " + cmd)
	}
	return arg
}

func (bot *Bot) readCommand() (string, string) {
	parts := strings.SplitN(bot.readLine(), " ", 2)

	arg := ""
	if len(parts) > 1 {
		arg = strings.TrimSpace(parts[1])
	}

	return parts[0], arg
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
