package pointers

import "fmt"

func Pointers() {
	var car string = "Mercedes"
	mycar := &car
	fmt.Printf("My Car is %s and it is stored at %p\n", *mycar, mycar)
	var intPointer = new(int)
	fmt.Println(intPointer, *intPointer)
	*intPointer = 1
	fmt.Println(*intPointer)

}
