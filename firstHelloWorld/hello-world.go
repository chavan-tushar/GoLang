package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	myFirstName, myLastName = "Tushar", "Chavan"
	myAge                   = calcAge(1990)
	years                   = "10"
	ageAfterYears           = calcAgeAfterYears(myAge, years)
)

func calcAgeAfterYears(ageToday int, years string) int {
	yearsInt, err := strconv.Atoi(years)
	if err == nil {
		return yearsInt + ageToday
	}
	return 0
}

func calcAge(birthYear int) int {
	return 2022 - birthYear

}

func main() {
	myFavSport := "Cricket"

	//Declaring myFavSportPtr as a pointer
	//&myFavSport will return the memory address of myFavSport
	//we are storing that memory address in a pointer variable called myFavSportPtr
	var myFavSportPtr *string = &myFavSport

	fmt.Println("Hello World, my First Name is", myFirstName, ",my last name is", myLastName, ". I am ", myAge, " years old.")
	fmt.Println("Type of firstname is ", reflect.TypeOf(myFirstName), ". Type of age is ", reflect.TypeOf(myAge))
	fmt.Println("I will be", ageAfterYears, "years old after", years, "years.")
	fmt.Println("memory address of myFavSport is ", &myFavSport)
	//we have stored memory address of myFavSport variable in myFavSportPtr variable.
	//*myFavSportPtr will return the value stored in memory address at which the pointer myFavSportPtr is pointing
	fmt.Println("Memory address of myFavSport is", myFavSportPtr, "and value stored is ", *myFavSportPtr)

}
