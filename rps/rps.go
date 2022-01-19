package rps

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	ROCK     = 0 // beats scissors. (scissors +1) % 3 = 0
	PAPER    = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS = 2 // beats paper. (paper + 1) % 3 = 2
)

func PlayRound(playerValue int) (int, string, string) {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}

	if playerValue == computerValue {
		fmt.Println("It's a draw")
		// *** decrement i by 1, since we're repeating the round
		i--
	} else if playerValue == -1 {
		fmt.Println("Invalid choice!")
		i--
	} else if playerValue == (computerValue+1)%3 {
		playerScore = playerWins(playerScore)
	} else {
		computerScore = computerWins(computerScore)
	}

	return 0, "", ""
}
