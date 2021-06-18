package main

import "fmt"

func updateName(x *string) string{
	*x= "wedge" //deference to get the value
     return *x
}

func main() {
	name := "tifa"

    //updateName(name)
	fmt.Println("memory address of name is:",&name)

	m := &name //pointer address to value name==tifa
	fmt.Println("the pointer value of &name:",*m)// dereference to get the value

    fmt.Println(name,&name)
	newName:=updateName(m) //passing in the pointer
	fmt.Println(newName,&name)
}