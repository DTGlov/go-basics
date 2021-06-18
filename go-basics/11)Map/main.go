package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":  9.78,
		"pie":   7.99,
		"salad": 7.98,
		"bun":   9,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])
	fmt.Println(menu["bun"])

	//LOOPING MAPS
	for k,v := range menu{
		fmt.Println(k,"-",v)
	}

	//ints as key types
	phonebook := map[int]string{
		1234313:"Dave",
		9786756:"Adrian",
		8564233:"Asher",
	}
	phonebook[1234313]="Geoffrey"
	fmt.Println(phonebook)
}