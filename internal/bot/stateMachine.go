package bot

type stateFn func(*Bot) stateFn

func botStartingState(bot *Bot) stateFn {
	bot.printComment("Waiting for bot-start")
	if l := bot.readLine(); l != "bot-start" {
		panic("Expected bot-start but got: " + l)
	}
	return botStartedState
}

func botStartedState(bot *Bot) stateFn {
	bot.printComment("Waiting for game-init")
	if l := bot.readLine(); l != "game-init" {
		panic("Expected game-init but got: " + l)
	}
	return nil
}
