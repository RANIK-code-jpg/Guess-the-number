package game

import "fmt"

func Session() {
	var endOfGame bool = false
	HelloMessage()
	SelectDifLevelMesage()
	chosenLevel := ChooseDifLevel()
	StartTheGame(chosenLevel)
	randomNumber := RandomNumber()
	for i := 0; i < int(CountOfChances[chosenLevel-1]) && endOfGame != true; i++ {
		currentNumber := ChoosingNumber()
		Help(currentNumber, randomNumber, i+1, &endOfGame)
	}
	if endOfGame == false {
		fmt.Println("You lost! Random number was", randomNumber)
		return
	}
	return
}
func NewGame() int {
	var pressToContinue int
	fmt.Println("Press 0 to end the game or 1 to play again")
	fmt.Scan(&pressToContinue)
	for {
		if pressToContinue != 0 && pressToContinue != 1 {
			fmt.Println("Chose 0 to end the game or 1 to play again")
			fmt.Scan(&pressToContinue)
		} else {
			break
		}
	}
	return pressToContinue
}

func AllGameSession() {
	for {
		Session()
		passed := NewGame()
		if passed == 0 {
			fmt.Println("Good Bye!")
			break
		}
	}
}
