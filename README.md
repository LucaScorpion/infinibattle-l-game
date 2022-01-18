# Infinibattle L-game bot

## Development

```shell
# Run the tests.
go test ./...

# Format the code.
go fmt ./...

# Run the bot.
go run ./main.go
```

## Bot State Diagram

```mermaid
stateDiagram-v2

[*] --> BotInit
note left of BotInit: Output\nbot-start

BotInit --> GameInitialised: game-init\n[turnState]
GameInitialised --> AwaitingTurn: game-start
AwaitingTurn --> TurnInitialised: turn-init\n[turnState]
TurnInitialised --> TurnStarted: turn-start
TurnStarted --> AwaitingTurn

AwaitingTurn --> [*]: throw / game-end
AwaitingTurn --> AwaitingTurn: sleep

note right of TurnStarted: Output\n[PlacePiecesCommand]\nturn-end

```
