package main

import "fmt"

func main() {
	//ARRAYS ---> Arrays must have a fix length
	var ages [3]int = [3]int{1, 2, 3}
	var myages  = [3]int{1, 2, 3}
	name := [4]string{"luka","trevor","juan"}

	name[1]= "Ivanovic"

	fmt.Println(ages,myages,len(ages))
	fmt.Println(name,len(name))

	//SLICES
	var scores = []int{1,2,3,4,5}
	scores[2] = 5

	scoresTwo :=append(scores,45)

	fmt.Println(scores,len(scores))
	fmt.Println(scoresTwo)

	//SLICE RANGING

	rangeOne := name[1:3]
	rangeTwo := name[:2]

	fmt.Println(rangeOne,rangeTwo)
}