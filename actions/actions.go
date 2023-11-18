package actions

import (
	"math/rand"
	"time"
)

var newSource = rand.NewSource(time.Now().UnixNano())
var customRand = rand.New(newSource)

var currentMonsterHealth = 100
var currentPlayerHealth = 100

func AttackMonster(isSpecialAttack bool) {
	minAttackValue := 5
	maxAttackValue := 10

	if isSpecialAttack {
		minAttackValue = 10
		maxAttackValue = 20
	}
	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)

	currentMonsterHealth -= dmgValue
}

func HealPlayer() {
	minHealValue := 10
	maxHealValue := 20

	healValue := generateRandBetween(minHealValue, maxHealValue)

	healthDiff := 100 - currentPlayerHealth

	if healthDiff >= healValue {
		currentMonsterHealth += healValue

	} else {
		currentPlayerHealth = 100
	}

}

func AttackPlayer() {
	minAttackValue := 9
	maxAttackValue := 13

	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)

	currentPlayerHealth -= dmgValue
}

func generateRandBetween(min int, max int) int {
	return customRand.Intn(max-min) + min
}
