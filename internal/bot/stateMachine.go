package bot

import (
	"fmt"
	"strconv"
	"time"
)

type stateFn func(*Bot) stateFn

func awaitBotStart(bot *Bot) stateFn {
	bot.expectCommand("bot-start")
	return awaitGameInit
}

func awaitGameInit(bot *Bot) stateFn {
	bot.expectCommand("game-init")
	return awaitGameStart
}

func awaitGameStart(bot *Bot) stateFn {
	bot.expectCommand("game-start")
	return awaitTurn
}

func awaitTurn(bot *Bot) stateFn {
	bot.printComment("Waiting for command: turn-init, sleep, or throw")
	cmd, arg := bot.readCommand()

	switch cmd {
	case "sleep":
		if seconds, err := strconv.Atoi(arg); err == nil {
			bot.printComment(fmt.Sprintf("Sleeping for %d seconds", seconds))
			time.Sleep(time.Duration(seconds) * time.Second)
		}
		return awaitTurn
	case "throw":
		return nil
	case "turn-init":
		return awaitTurnStart
	default:
		panic("Invalid command while awaiting turn: " + cmd)
	}
}

func awaitTurnStart(bot *Bot) stateFn {
	bot.expectCommand("turn-start")
	// TODO
	bot.writeLine("turn-end")
	return awaitTurn
}
