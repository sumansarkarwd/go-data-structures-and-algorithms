package main

import "fmt"

type contactInfo struct {
	email       string
	phoneNumber int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// suman := person{firstName: "Suman", lastName: "Sarkar"}
	// fmt.Println(suman)

	// var suman person

	// suman.firstName = "Suman"
	// suman.lastName = "Sarkar"

	suman := person{
		firstName: "Suman",
		lastName:  "Sarkar",
		contactInfo: contactInfo{
			email:       "sumansarkarwd@gmail.com",
			phoneNumber: 8276809245,
		},
	}

	// sumanPointer := &suman
	// sumanPointer.updateFirstName("Susmita")
	suman.updateFirstName("Susmita")
	suman.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println("")
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
