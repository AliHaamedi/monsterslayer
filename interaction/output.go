package interaction

import "fmt"

func PrintGreeting() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting new game...")
	fmt.Println("Good luck!")
}

func ShowAvailableActions(specialRoundAvailable bool) {

	fmt.Println("Please choose your action")

	fmt.Println("-------------------------")
	fmt.Println("(1) Attack monster")
	fmt.Println("(2) Heal")

	if specialRoundAvailable {
		fmt.Println("(3) Special attack")
	}

}
