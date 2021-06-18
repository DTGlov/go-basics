package main

import "fmt"

nameSix  := "theyname" //shorthand cannot be used outside function

func main() {

	//STRINGS
	//strings must be double quotes
	var nameOne string = "myname" //explicit type
	var nameTwo = "yourname"  //inferred type
	var nameThree string // setting the value for future use
	
	fmt.Println(nameOne,nameTwo,nameThree)

	nameOne = "hisname" // updating nameOne
	nameFour := "hername" //shorthand for defining variable, only used to initialize the variable


	//INTEGERS
	var myNumber int = 45
	var hisNumber = 12
	herNumber := 69

	fmt.Println(myNumber,hisNumber,herNumber)

	//BITS & MEMORY
	var numOne int8 = 25 //int8 represents range of numbers , visit https://golang.org/pkg/builtin for more info
	var numTwo int8 = -129
	var numThree uint = 45 //unit means the value has to be +tive

	//FLOATS
	var scoreOne float32 = -1.5
	var scoreTwo float64 = 324332535454266768.56 //float64 has higher precision
	scoreThree := 1.5 //inferred as float64
}