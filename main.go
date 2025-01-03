package main

import "fmt"

func main()  {

	// strings
	var nameOne string = "Ben"

	fmt.Println(nameOne)

	//ints
	var ageOne int = 25
	ageTwo := 40

	fmt.Println(ageOne, ageTwo)

	//bits & memory
	var numOne int8 = 25
	var numTwo uint16 = 200 //uint - cant generate a negative number
	
	fmt.Println(numOne, numTwo)

	//float - decimal numbers
	var scoreOne float32 = 20.22
	var scoreTwo float64 = 362746234782653627337.9 // we will bw using float64

	fmt.Println(scoreOne, scoreTwo)
}