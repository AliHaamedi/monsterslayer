package interaction

import "fmt"

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

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

func PrintRoundStatistics(rd *RoundData) {
	if rd.Action == "ATTACK" {
		fmt.Printf("Player attack monster for %v damage.\n", rd.PlayerAttackDmg)
	} else if rd.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player perform a strong attack against monster for %v damage.\n", rd.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed for %v .\n", rd.PlayerHealValue)
	}
	fmt.Printf("Monster attack player for %v damage.\n", rd.MonsterAttackDmg)

	fmt.Printf("Player Health: %v\n", rd.PlayerHealth)
	fmt.Printf("Monster Health: %v\n", rd.MonsterHealth)
}

func DeclareWinner(winner string) {
	fmt.Println("-------------------------")
	fmt.Println("GAME OVER!")
	fmt.Println("-------------------------")
	fmt.Printf("%v won!\n", winner)
}
