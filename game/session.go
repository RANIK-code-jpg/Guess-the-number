package game

import "fmt"

func Session() {
	HelloMessage()
	SelectDifLevelMesage()
	chosenLevel := ChooseDifLevel()
	StartTheGame(chosenLevel)
	randomNumber := RandomNumber()
	for i := 0; i < int(chosenLevel); i++ {
		currentNumber := ChoosingNumber()
		Help(currentNumber, randomNumber, i+1)
	}
	if EndOfGame == false {
		fmt.Println("You lost! Randim number was", randomNumber)
		return
	}
	return
}
