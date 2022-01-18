package bot

import (
	"fmt"
	"infinibattle-l-game/internal/lgame"
	"infinibattle-l-game/internal/parser"
	"time"
)

type stateFn func(*Bot) stateFn

func botStarting(bot *Bot) stateFn {
	bot.writeLine("bot-start")
	return initGame
}

func initGame(bot *Bot) stateFn {
	bot.expectLine("game-init")

	bot.printComment("Waiting for: game-start")
	for bot.readLine() != "game-start" {
		// Ignore any game init input.
	}

	return awaitTurn
}

func awaitTurn(bot *Bot) stateFn {
	bot.printComment("Awaiting turn or command")
	cmd := bot.readLine()

	switch cmd {
	case "turn-init":
		return initTurn
	case "sleep":
		bot.printComment("Sleeping for 1 second")
		time.Sleep(1 * time.Second)
		return awaitTurn
	case "game-end":
		fallthrough
	case "throw":
		bot.printComment("gg")
		return nil
	default:
		panic("Invalid command while awaiting turn: " + cmd)
	}
}

func initTurn(bot *Bot) stateFn {
	// Parse the turn state.
	bot.printComment("Initialising turn")
	state := parser.ParseGameState(bot.readLine())

	// Get the next state, benchmark.
	bot.printComment("Starting ")
	startTime := time.Now()
	nextState := getNextState(lgame.DefaultSettings(), state)
	thinkTime := time.Now().Sub(startTime).Seconds()
	bot.printComment(fmt.Sprintf("Thinking took %.3f seconds", thinkTime))

	for bot.readLine() != "turn-start" {
		// Ignore any other turn input.
	}

	bot.writeLine(parser.GetMoveOutput(nextState))
	bot.writeLine("turn-end")
	return awaitTurn
}
