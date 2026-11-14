package game

import (
	"fmt"
	"math/rand"
)

var DifficultLevel [3]string = [3]string{"Easy", "Medium", "Hard"}
var CountOfChances [3]int = [3]int{10, 5, 3}
var EndOfGame bool = false

func RandomNumber() int {
	return rand.Intn(101)
}

func HelloMessage() {
	var helloMessage = "Welcome to the Number Guessing Game! \n I'm thinking of a number between 1 and 100. \n You have 5 chances to guess the correct number."
	fmt.Println(helloMessage)
	fmt.Println(" ")
}
func ChooseDifLevel() uint8 {
	fmt.Print("Enter your choice: ")
	var difficultyLevel uint8
	fmt.Scan(&difficultyLevel)
	return difficultyLevel
}
func SelectDifLevelMesage() {
	var chooseMessage = "Please select the difficulty level: \n 1. Easy (10 chances) \n 2. Medium (5 chances) \n 3. Hard (3 chances)"
	fmt.Println(chooseMessage)
}
func StartTheGame(dufficultylevel uint8) {
	fmt.Printf("Great! You have selected the %s difficulty level.\n Let's start the game! \n", DifficultLevel[dufficultylevel-1])
}
func ChoosingNumber() int {
	var chosenNumber int
	fmt.Print("Enter your guess: ")
	fmt.Scan(&chosenNumber)
	return chosenNumber
}
func Help(chosenNumber, randomNumber, step int) {
	if chosenNumber == randomNumber {
		fmt.Printf("Congratulations! You guessed the correct number in %d attempts.", step)
		EndOfGame = true
		return
	} else if chosenNumber > randomNumber {
		fmt.Printf("Incorrect! The number is less than %d. \n", chosenNumber)
		return
	} else {
		fmt.Printf("Incorrect! The number is greater than %d. \n", chosenNumber)
		return
	}
}
