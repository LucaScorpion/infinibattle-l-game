package bot

import (
	"fmt"
	"infinibattle-l-game/internal/lgame"
	"infinibattle-l-game/internal/parser"
	"time"
)

type stateFn func(*Bot) stateFn

func awaitBotStart(bot *Bot) stateFn {
	bot.printComment("Waiting for: bot-start")
	cmd := bot.readLine()

	switch cmd {
	case "q":
		// Quick start for debugging purposes.
		return awaitTurnStart
	case "bot-start":
		return awaitGameInit
	default:
		panic("Invalid command while starting bot: " + cmd)
	}
}

func awaitGameInit(bot *Bot) stateFn {
	bot.expectLine("game-init")
	return awaitGameStart
}

func awaitGameStart(bot *Bot) stateFn {
	bot.expectLine("game-start")
	return awaitTurn
}

func awaitTurn(bot *Bot) stateFn {
	bot.printComment("Waiting for: turn-init, sleep, or throw")
	cmd := bot.readLine()

	switch cmd {
	case "sleep":
		bot.printComment("Sleep is not currently implemented")
		//if seconds, err := strconv.Atoi(arg); err == nil {
		//	bot.printComment(fmt.Sprintf("Sleeping for %d seconds", seconds))
		//	time.Sleep(time.Duration(seconds) * time.Second)
		//}
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
	bot.expectLine("turn-start")
	state := parser.ParseGameState(bot.readLine())
	startTime := time.Now()
	lgame.GetNextState(state)
	thinkTime := time.Now().Sub(startTime).Seconds()
	bot.printComment(fmt.Sprintf("Thinking took %.2f seconds", thinkTime))
	bot.writeLine("turn-end")
	return awaitTurn
}
