package main

import (
	"github.com/alihaamedi/monsterslayer/actions"
	"github.com/alihaamedi/monsterslayer/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	startGame()

	winner := ""

	for winner == "" {
		winner = executeRound()
	}

	finishGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++

	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerAttackDmg int
	var playerHealValue int
	var monsterAttackDmg int

	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		playerAttackDmg = actions.AttackMonster(true)
	}

	monsterAttackDmg = actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthAmount()

	roundData := interaction.RoundData{
		Action:           userChoice,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
		PlayerAttackDmg:  playerAttackDmg,
		PlayerHealValue:  playerHealValue,
		MonsterAttackDmg: monsterAttackDmg,
	}

	interaction.PrintRoundStatistics(&roundData)
	gameRounds = append(gameRounds, roundData)
	if monsterHealth <= 0 {
		return "player"
	} else if playerHealth <= 0 {
		return "monster"
	}

	return ""
}

func finishGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
