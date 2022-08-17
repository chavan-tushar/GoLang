package main

import "fmt"
import "reflect"

var (
	myFirstName, myLastName = "Tushar", "Chavan"
	myAge                   = calcAge(1990)
)

func calcAge(birthYear int) int {
	return 2022 - birthYear

}

func main() {
	fmt.Println("Hello World, my First Name is", myFirstName, ",my last name is", myLastName, ". I am ", myAge, " years old.")
	fmt.Println("Type of firstname is ", reflect.TypeOf(myFirstName), ". Type of age is ", reflect.TypeOf(myAge))
}
