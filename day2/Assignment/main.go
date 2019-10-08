package main

import "fmt"

type salary struct {
	basic         float64
	hra           float64
	lta           float64
	cityAllowance float64
	medical       float64
}


type contact struct {
	mob     int
	address string
}

type employee struct {
	id            int
	fname         string
	lname         string
	dob           string
	rating        string
	leavesTaken   float64
	leavesBalance float64
	salary
	contact
}

func showOptions() {
	options := "1. Create employee\n2. Show employee details\n3. Update employee details\n4. Delete employee\n5. Show friends\n6. Calculate appraisal\n7. Show emergency details\n8. Delete friend"
	fmt.Println(options)
}

func main() {
	showOptions()
}
