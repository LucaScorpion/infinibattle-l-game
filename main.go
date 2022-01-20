package main

import (
	"fmt"
	"infinibattle-l-game/internal/bot"
	"infinibattle-l-game/internal/lgame"
	"math"
)

func main() {
	bot.NewBot().Start()
}

/*==== RESEARCH ====*/

type stateWithSubStates struct {
	state     lgame.GameState
	subStates []lgame.GameState
}

// This method is for research purposes.
func main2() {
	emptyBoard := lgame.GameState{
		PlayerTurn: 0,
		Players: [2]lgame.Player{
			{
				Piece: lgame.LPiece{
					{-1, -1},
					{-1, -1},
					{-1, -1},
					{-1, -1},
				},
				Score: 0,
			},
			{
				Piece: lgame.LPiece{
					{-1, -1},
					{-1, -1},
					{-1, -1},
					{-1, -1},
				},
				Score: 0,
			},
		},
		Neutrals: [2]lgame.NeutralPiece{
			{-1, -1},
			{-1, -1},
		},
	}

	placements := lgame.GetLShapeMoves(lgame.DefaultSettings(), emptyBoard)
	fmt.Printf("There are %d ways to place an L on an empty board.\n", len(placements))

	var scoringStates []*stateWithSubStates
	for _, p := range placements {
		if p.Players[0].Score > 0 {
			scoringStates = append(scoringStates, &stateWithSubStates{
				state:     p,
				subStates: []lgame.GameState{},
			})
		}
	}
	fmt.Printf("Of those, %d ways score a point.\n", len(scoringStates))

	boardLength := lgame.DefaultSettings().BoardWidth * lgame.DefaultSettings().BoardHeight
	totalSubStates := 0
	for _, s := range scoringStates {
		for pos := 0; pos < boardLength-1; pos++ {
			neutralOne := lgame.NeutralPiece{
				X: pos % lgame.DefaultSettings().BoardWidth,
				Y: pos / lgame.DefaultSettings().BoardWidth,
			}

			occ := lgame.GetOccupation(s.state)
			if _, ok := occ[lgame.Coordinate(neutralOne)]; ok {
				continue
			}

			oneNeutralPlaced := s.state
			oneNeutralPlaced.Neutrals[0] = neutralOne

			for pos2 := pos + 1; pos2 < boardLength; pos2++ {
				neutralTwo := lgame.NeutralPiece{
					X: pos2 % lgame.DefaultSettings().BoardWidth,
					Y: pos2 / lgame.DefaultSettings().BoardWidth,
				}

				if _, ok := occ[lgame.Coordinate(neutralTwo)]; ok {
					continue
				}

				newState := oneNeutralPlaced
				newState.Neutrals[1] = neutralTwo
				s.subStates = append(s.subStates, newState)
				totalSubStates++
			}
		}
	}
	fmt.Printf("Placing 2 neutral pieces gives a total of %d possible boards (%d ways per scoring state).\n", totalSubStates, totalSubStates/len(scoringStates))

	totalPossibilities := 0
	minScoringPlacements := math.MaxInt
	maxScoringPlacements := 0
	var singleScoringOption []lgame.GameState
	for _, scoring := range scoringStates {
		for _, sub := range scoring.subStates {
			fullStates := lgame.GetLShapeMoves(lgame.DefaultSettings(), sub)
			totalPossibilities += len(fullStates)

			scoring := 0
			for _, full := range fullStates {
				if full.Players[1].Score == 1 {
					scoring++
				}
			}
			if scoring < minScoringPlacements {
				minScoringPlacements = scoring
			}
			if scoring > maxScoringPlacements {
				maxScoringPlacements = scoring
			}

			if scoring == 1 {
				singleScoringOption = append(singleScoringOption, sub)
			}
		}
	}
	fmt.Printf("Placing the second L gives a total of %d possible boards.\n", totalPossibilities)
	fmt.Printf("For these boards, each arrangement allows the opponent to make between %d and %d scoring moves (inclusive).\n", minScoringPlacements, maxScoringPlacements)
	fmt.Printf("There are %d arrangements (%d if you remove rotations and reflections) where the opponent has only a single scoring move:\n", len(singleScoringOption), len(singleScoringOption)/8)

	for i := 0; i < len(singleScoringOption)/8; i++ {
		fmt.Println(lgame.DrawState(lgame.DefaultSettings(), singleScoringOption[i]))
	}
}
