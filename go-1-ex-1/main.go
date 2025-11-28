package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	var firstName string = "Jana"
	var lastName string = "Slavnic"
	var dayOfBirth int = 23
	var monthOfBirth int = 11
	var yearOfBirth int = 2008
	var numberOfSiblings uint = 0
	var heightInMeters float32 = 1.72

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
