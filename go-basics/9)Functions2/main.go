package main

import (
	"fmt"
	"strings"
)

//FUNCTIONS THAT RETURN MULTIPLE VALUES

func getInitials(n string) (string,string){
//returning two strings
s := strings.ToUpper(n)
names := strings.Split(s," ")
var initials []string

for _,v := range names{
	initials = append(initials,v[:1])
}
if len(initials) > 1{
	return initials[0],initials[1]
}
return initials[0],"_"
}

func main()  {
	fn1,sn1 :=getInitials("tifa lockhart")
	fmt.Println(fn1,sn1)
	
}