package main

import "fmt"

func main() {
	age := 45

	fmt.Println(age<45)
	fmt.Println(age<15)
	fmt.Println(age>50)

	if age < 30{
		fmt.Println("yes it is")
	}else{
		fmt.Println("No it isnt")
	}
	mane := []string{"jac","mk","kors","lolo","poko"}
	for index,value := range mane{
	if index == 2{
		fmt.Println("continuing at pos",index)
		continue
	}
	if index > 2 {
		fmt.Println("breaking at pos",index)
		break
	}
	fmt.Printf("the value at pos %v is %v \n",index,value)
	}
}