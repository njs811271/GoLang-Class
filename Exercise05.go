package main

import (
	"fmt"
)

//Custom Type
type myType int

var x myType
var y int

func main() {

	fmt.Println(x)        //Print out the variable x
	fmt.Printf("%T\n", x) //print out the type of variable
	x = 42                //reasigned value
	fmt.Println(x)        //output the value

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
