package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialRoundAvailable bool) string {
	for {
		playerChoice, _ := getPlayerInput()

		if playerChoice == "1" {
			return "ATTACK"
		} else if playerChoice == "2" {
			return "HEAL"
		} else if playerChoice == "3" && specialRoundAvailable {
			return "SPECIAL_ATTACK"
		}
		fmt.Println("fetching the user input failed, please try again...")
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Your choice: ")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput, nil
}
