package main

import "fmt"

func main()  {
	age := "25"
	name := "Ben"

	// printf (formatted strings) %_ = formatt specifier
	fmt.Printf("my age is %v and my name is %v \n ", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f points! \n", 225.55) // for floats
	fmt.Printf("you scored %0.1f points! \n", 225.556) //%0.1f = to 1dp and vice versa

	// sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n ", age, name)
	fmt.Println("the saved string is:", str)
}