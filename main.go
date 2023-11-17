package main

import "github.com/alihaamedi/monster/interaction"

var currentRound = 0

func main() {
	startGame()

	winner := ""

	for winner == "" {
		winner = executeRound()
	}

	finishGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++

	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)

	return ""
}

func finishGame() {}
