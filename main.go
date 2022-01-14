package main

import "fmt"

func main() {

	fmt.Println("=============== STRINGS ===============")
	// strings
	var nameOne string = "nkduy"
	var nameTwo = "Duy"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Mario"
	nameThree = "Peach"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"
	fmt.Println(nameFour)


	fmt.Println("=============== NUMBERS ===============")
	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory https://pkg.go.dev/builtin
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint16 = 256

	// float
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 883759237594363564868547567.7
	scoreThree := 1.5

}
