package main

import "fmt"

func main(){

	//Print
	fmt.Print("hello ") //.Print does not add a newline
	fmt.Print("World! \n")
	fmt.Print("new line \n")

	//Println
	age := 45
	name := "Pablo"
	fmt.Println("Come on") //Prints a newline
	fmt.Println("Come People")
	fmt.Println("my age is",age,"and my name is",name)

	//Printf --->Formatted Strings
	fmt.Printf("my age is %v and my name is %v \n",age,name)
	fmt.Printf("my age is %q and my name is %q \n",age,name) //places a quote around the variable , must be used on string variable
	fmt.Printf("age is of type %T \n",age) //outputs the type of the variable
	fmt.Printf("you scored %0.1f points! \n",45.7) //outputs floats

	//Sprintf (save formatted strings)
	newName :=fmt.Sprintf("my name is %v",name) // saves the string in a variable

	fmt.Println(newName)
	
}