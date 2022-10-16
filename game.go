package main

import "fmt"

func main() {

	//yazdırma
	fmt.Println("Welcome to my quiz game.")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello %v, welcome to the game. \n ", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play.")
	}

	if age < 10 {
		fmt.Println("No, you can't play.")
	}

	score := 0         //score point starting
	numQuestions := 2 //number of questions

	fmt.Println("continue")

	//question 1
	fmt.Printf("Which is better, RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)
	true || false -> true
	//answer 1
	if answer+" "+answer2 == "RTX 3090" {
		fmt.Println("Correct!")
		score++ //give a score point
	} else if answer+" "+answer2 == "rtx 3090" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	//question 2
	fmt.Printf("How many cores does have the Ryzen 5 2600 have? ")
	var cores int
	fmt.Scan(&cores)
	//answer 2
	if cores == 6 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("you scored %v out of %v.\n", score, numQuestions)
	percent := (float64(score) / float64(numQuestions) * 100
	fmt.Printf("you scored: %v%%. ", percent)
}
