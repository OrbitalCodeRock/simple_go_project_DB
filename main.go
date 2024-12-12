// package declaration
package main

// Imports (Modules and/or Packages)
import (
	"fmt"
	"strings"
)

// Entry point
func main() {
	var userAnswer string

	fmt.Println("Hello... ")

	// Collect user input up to the next \n (newline) character.
	// & Gives us the memory address of userAnswer.
	fmt.Scanln(&userAnswer)

	switch strings.ToLower(userAnswer) {
	case "world":
		HandleRightAnswer(userAnswer)
	case "world!":
		HandleRightAnswer(userAnswer)
	case "world.":
		HandleRightAnswer(userAnswer)
	default:
		HandleWrongAnswer(userAnswer)
	}
}

// Function for handling an incorrect answer.
func HandleWrongAnswer(userAnswer string) {
	fmt.Printf("Hello %.5s! That's... not quite what I was looking for!", userAnswer)
}

// Function for handling a correct answer.
func HandleRightAnswer(userAnswer string) {
	fmt.Printf("Hello %.5s! That's correct!", userAnswer)
}
