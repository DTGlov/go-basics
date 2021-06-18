package main

import "fmt"

func updateName(x string) string{
	x ="wedge"
     return	x
	
}
func updateMenu(y map[string]float64)  map[string]float64{
	y["coffee"]=56.99
	return y

}

func main() {
	//group A types --> strings,ints,bools,floats,arrays,structs, a copy of the name is altered
	//name := "tifa"

	// names  := updateName(name)
	// fmt.Println(names,&names)
	// fmt.Println(name,&name)

	//group b types --> slices,map,functions ,the menu is altered directly
	menu := map[string]float64{
		"pie":45.98,
		"cheese cake":56.23,
		"Yoghurt":34.56,
	}
	//newMenu :=updateMenu(menu)
	for k,v :=range menu {
		fmt.Println(k,"-",&v)
	}
	// for k,v :=range newMenu {
	// 	fmt.Println(k,"-",&v)
	// }
	// fmt.Println(menu,&menu)
	// fmt.Println(newMenu,&newMenu)
}