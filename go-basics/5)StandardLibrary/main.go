package main

import (
	"fmt"
	"sort"
	"strings"
)


func main() {
	greeting := "hello there friends"

	//STRINGS
	fmt.Println(strings.Contains(greeting,"hello"))
	fmt.Println(strings.ReplaceAll(greeting,"e","i"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting,"r"))
	fmt.Println(strings.Split(greeting,""))

	//SORT
	ages := []int{12,12,3,4,5,4,66,75,76}
	
	sort.Ints(ages) //sorts a slice of integers ,modifies the original slice
	fmt.Println(ages)

	index := sort.SearchInts(ages,66)
	fmt.Println(index)

	nomen := []string{"mada","moshi","falana","jui"}

	sort.Strings(nomen) //sorts a slice of strings
	fmt.Println(nomen)

	fmt.Println(sort.SearchStrings(nomen,"falana"))
}