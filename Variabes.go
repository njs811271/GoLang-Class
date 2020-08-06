package main

import "fmt"

//DECLARE there is a variable with the IDENTIFIER "Z" and that the VARIABLE with the IDENTIFIER "Z"
//is of TYPE int. When no value is specified the variable inistializes to 0 for ints, 0.0 for float, false for bool
//"" for strings
var z int

//DECLARE the variable "y"
//ASSIGN the value 43
var y = 43

func main() {
	//Use _ as a empty variable. Allows you to receive data you will never use, or debug a non-defined variable
	_, a := "Hello and Tes", 45 //Comment Test
	fmt.Println(a)

	fmt.Println("TESTING new code")

}
