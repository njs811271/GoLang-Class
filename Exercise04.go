package main

import (
	"fmt"
)

type myType int

var x myType

func main() {

	fmt.Println(x)        //Print out the variable x
	fmt.Printf("%T\n", x) //print out the type of variable
	x = 42                //reasigned value
	fmt.Println(x)        //output the value

}
