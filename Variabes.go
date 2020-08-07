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

	fmt.Println("Basic Variable output\n===============")
	fmt.Println(a)
	fmt.Println("\n")

	//Run the next Lesson
	varType()

	varCreateNewType()
}

//determine what a variable type is
func varType() {

	var thisInt = 42
	var stringVar = "Shaken, not stirred"

	fmt.Println("Variables and Types\n===============")
	fmt.Println("Variable value: ", thisInt)       //Print with line return at the end
	fmt.Printf("%T\n", thisInt)                    //Print format
	fmt.Print("Variable value: ", stringVar, "\n") //Print at face value with no additional formatting.
	//Added newline manually
	fmt.Printf("%T\n", stringVar)
	fmt.Printf("\n")

	//You can use Sprint to print to a string which is helpful when formatting to a variable
	y := 5
	S := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Print(S, "\n")
	//https://godoc.org/fmt formatting options

	//You can use Fprint to print to a file
}

func varCreateNewType() {

	//creating a new type of variable rather than int,string,bool,etc
	type hotDog int
	var newType hotDog
	//Int variable
	oldType := 42
	newType = 5

	fmt.Print("Creating a Custom Type\n", "=================\n")
	fmt.Printf("%T\n", oldType)
	fmt.Printf("%T\n", newType)

	//Converting one type to another type
	oldType = int(newType)
	fmt.Println(oldType)
	fmt.Printf("%T\n", newType)

}
