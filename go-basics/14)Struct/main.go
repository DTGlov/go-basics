package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//helper function

func getInput(prompt string,r *bufio.Reader) (string,error){
    fmt.Print(prompt)
	input,err := r.ReadString('\n') 

	return strings.TrimSpace(input), err

}

func createBill() bill { //responsible for getting user inputs
	reader := bufio.NewReader(os.Stdin) //Stdin refers to the terminal


	name, _ := getInput("Create a new bill name: ",reader)

	b := newBill(name)
	fmt.Println("Created the bill:",b.name)

	return b
}

func promptOptions(b bill){
reader := bufio.NewReader(os.Stdin)

opt, _ := getInput("Choose option(a -add item,s - save bill, t - add tip): ",reader)
	switch opt {
	case "a":
		name,_ :=getInput("Item name:",reader)
		price,_ := getInput("item price",reader)

		p,err := strconv.ParseFloat(price,64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name,p)

		fmt.Println("item added - ",name,price)
		promptOptions(b)
	case "t":
		tip,_ := getInput("Enter tip amount ($): ",reader)
		t,err := strconv.ParseFloat(tip,64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added - ",tip)
	fmt.Println(tip)
	case "s":
		b.save()
	fmt.Println("you saved the file -",b.name)
	default:
	fmt.Println("that was not a valid option...")
	promptOptions(b)
	
}

}

func main() {

mybill := createBill()
promptOptions(mybill)


}