package main

import "fmt"

type FullName struct {
		First string
    	Last  string
}

type BirthDate struct {
	Day
	Month
	Year
}

type Profile struct {
	// TODO:
    FullName
    BirthDate
	NumberOfSiblings byte
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		FullName: FullName{
		First: "Jana",
		Last: "Slavnic",
		},
    BirthDate: BirthDate{
    Day : 23,
    Month : 11,
    Year : 2008,
    },
		NumberOfSiblings: 0,
		ZodiacSign:       ' ',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
    me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
