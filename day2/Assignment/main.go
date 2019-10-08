package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

var idCounter = 0

type salary struct {
	basic         float64
	hra           float64
	lta           float64
	cityAllowance float64
}

type contact struct {
	mob       string
	city      string
	emergency string
}

type employee struct {
	id          int
	fname       string
	lname       string
	dob         string
	rating      string
	leavesTaken int
	totalLeaves int
	friends     []int
	salary
	contact
}

type employees map[int]employee

func showOptions() {
	const options = "\n1. Create employee\n" +
		"2. Show employee details\n" +
		"3. Update employee details\n" +
		"4. Delete employee\n" +
		"5. Show friends\n" +
		"6. Calculate appraisal\n" +
		"7. Show emergency details\n" +
		"8. Delete friend\n" +
		"0. Exit\n\n" +
		"Enter a valid option: "
	fmt.Println(options)
}

func removeDupliactes(s []int) []int {
	var result []int
	tmp := make(map[int]bool)
	for _, v := range s {
		if _, ok := tmp[v]; !ok {
			result = append(result, v)
		}
		tmp[v] = true
	}
	return result
}

func (e *employees) populateFriends() {
	for key, emp := range *e {
		for _, v := range emp.friends {
			if frnd, ok := (*e)[v]; ok {
				frnd.friends = append(frnd.friends, key)
				frnd.friends = removeDupliactes(frnd.friends)
				(*e)[v] = frnd
			}
		}
		fmt.Println("Inside populate:", emp.friends)
	}
}

func (e *employees) createAndUpdateEmployee(id int) error {
	var emp employee
	if id == 0 {
		idCounter++
		id = idCounter
		emp.id = id
	} else {
		emp = (*e)[id]
	}
	fmt.Println("Enter First Name")
	_, err := fmt.Scan(&emp.fname)
	if err != nil {
		return err
	}
	fmt.Println("Enter Last Name")
	_, err = fmt.Scan(&emp.lname)
	if err != nil {
		return err
	}
	fmt.Println("Enter D.O.B")
	_, err = fmt.Scan(&emp.dob)
	if err != nil {
		return err
	}
	fmt.Println("Enter Rating")
	_, err = fmt.Scan(&emp.rating)
	if err != nil {
		return err
	}
	fmt.Println("Enter leaves taken")
	_, err = fmt.Scan(&emp.leavesTaken)
	if err != nil {
		return err
	}
	fmt.Println("Enter total leaves")
	_, err = fmt.Scan(&emp.totalLeaves)
	if err != nil {
		return err
	}
	fmt.Println("Enter list of friends")
	var input string
	_, err = fmt.Scan(&input)
	if err != nil {
		return err
	}
	for _, v := range strings.Split(input, ",") {
		if n, e := strconv.Atoi(v); e == nil {
			emp.friends = append(emp.friends, n)
		}
	}
	emp.friends = removeDupliactes(emp.friends)
	fmt.Println(emp.friends)
	fmt.Println("\nSALARY DETAILS\n Enter basic pay")
	_, err = fmt.Scan(&emp.salary.basic)
	if err != nil {
		return err
	}
	fmt.Println("Enter HRA")
	_, err = fmt.Scan(&emp.salary.hra)
	if err != nil {
		return err
	}
	fmt.Println("Enter LTA")
	_, err = fmt.Scan(&emp.salary.lta)
	if err != nil {
		return err
	}
	fmt.Println("Enter City Allowance")
	_, err = fmt.Scan(&emp.salary.cityAllowance)
	if err != nil {
		return err
	}
	fmt.Println("\nCONTACT DETAILS\nEnter Contact no.")
	_, err = fmt.Scan(&emp.contact.mob)
	if err != nil {
		return err
	}
	fmt.Println("Enter City")
	_, err = fmt.Scan(&emp.contact.city)
	if err != nil {
		return err
	}
	fmt.Println("Enter emergency contact no.")
	_, err = fmt.Scan(&emp.contact.emergency)
	if err != nil {
		return err
	}
	if err == nil {
		(*e)[id] = emp
	}
	return err
}

func (e employees) displayEmpDetails(id int) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"ID", "NAME", "D.O.B", "RATING", "LEAVE_BALANCE", "CONTACT", "CITY", "EMERGENCY", "SALARY"})
	for idx, v := range e {
		if id == 0 {
			leaveDeduction := 0.0
			if v.leavesTaken > v.totalLeaves {
				leaveDeduction = float64(v.leavesTaken-v.totalLeaves) * (v.basic * 0.02)
			}
			tw.AppendRow(table.Row{idx, v.fname+" "+v.lname, string(v.dob), string(v.rating),
				float64(v.totalLeaves-v.leavesTaken), v.contact.mob, v.contact.city, v.contact.emergency,
				float64(v.salary.basic+v.salary.hra+v.salary.lta+v.salary.cityAllowance-leaveDeduction)})
		} else if id == idx {
			leaveDeduction := 0.0
			if v.leavesTaken > v.totalLeaves {
				leaveDeduction = float64(v.leavesTaken-v.totalLeaves) * (v.basic * 0.02)
			}
			tw.AppendRow(table.Row{idx, v.fname+" "+v.lname, string(v.dob), string(v.rating),
				float64(v.totalLeaves-v.leavesTaken), v.contact.mob, v.contact.city, v.contact.emergency,
				float64(v.salary.basic+v.salary.hra+v.salary.lta+v.salary.cityAllowance-leaveDeduction)})
		}
	}
	fmt.Println(tw.Render())
}

func (e *employees) deleteEmployee() {
	var id int
	fmt.Print("Enter the ID of the existing employee: ")
	_, _ = fmt.Scan(&id)
	if _, ok := (*e)[id]; !ok {
		fmt.Printf("Employee with ID : %d doesn't exist", id)
		return
	} else {
		delete(*e, id)
	}
}

func (e *employees) showFriends() {
	var id int
	fmt.Print("Enter the ID of the existing employee: ")
	_, _ = fmt.Scan(&id)
	if _, ok := (*e)[id]; !ok {
		fmt.Printf("Employee with ID : %d doesn't exist", id)
		return
	} else {
		tw := table.NewWriter()
		tw.AppendHeader(table.Row{"ID", "NAME"})
		tw.SetIndexColumn(1)
		for idx, v := range (*e)[id].friends {
			if emp, ok := (*e)[v]; ok && emp.id != id {
				tw.AppendRow(table.Row{idx, emp.fname + " " + emp.lname})
			}
		}
		fmt.Println(tw.Render())
	}
}

func (e *employees) deleteFriends() {
	var f1, f2 int
	fmt.Print("Enter the IDs of the employees (e.g. emp1 emp2): ")
	_, _ = fmt.Scan(&f1, &f2)
	f1 = int(f1)
	f2 = int(f2)
	if _, ok := (*e)[f1]; !ok {
		fmt.Printf("Employee with ID : %d doesn't exist", f1)
		return
	}
	if _, ok := (*e)[f2]; !ok {
		fmt.Printf("Employee with ID : %d doesn't exist", f2)
		return
	}
	(*e)[f1].friends[sort.SearchInts((*e)[f1].friends, f2)] = 0
	(*e)[f2].friends[sort.SearchInts((*e)[f2].friends, f1)] = 0
}

func (e *employees) calculateAppraisal() {
	var id int
	fmt.Print("Enter the ID of the existing employee: ")
	_, _ = fmt.Scan(&id)
	if emp, ok := (*e)[id]; !ok {
		fmt.Printf("Employee with ID : %d doesn't exist", id)
		return
	} else {
		leaveDeduction := 0.0
		if emp.leavesTaken > emp.totalLeaves {
			leaveDeduction = float64(emp.leavesTaken-emp.totalLeaves) * (emp.basic * 0.02)
		}
		if strings.ToLower(emp.rating) == "a" {
			emp.basic += emp.basic * 0.30
		} else if strings.ToLower(emp.rating) == "b" {
			emp.basic += emp.basic * 0.10
		} else if strings.ToLower(emp.rating) == "c" {
			emp.basic += emp.basic * 0.05
		}
		salary := emp.basic + emp.hra + emp.cityAllowance + emp.lta - leaveDeduction
		fmt.Printf("Salary after appraisal for Employee %d: %0.2f\n", id, salary)
	}
}

func (e *employees) showEmergencyContact() {
	var id int
	fmt.Print("Enter the ID of the existing employee: ")
	_, _ = fmt.Scan(&id)
	if emp, ok := (*e)[id]; !ok {
		fmt.Printf("Employee with ID : %d doesn't exist", id)
		return
	} else {
		fmt.Printf("Emergency contact for Employee %d: %s\n", id, emp.emergency)
	}
}

func main() {
	var option int
	emp := employees{}
	cont := false
	for !cont {
		showOptions()
		_, _ = fmt.Scan(&option)
		switch option {
		case 0:
			cont = true
		case 1:
			if err := emp.createAndUpdateEmployee(0); err != nil {
				log.Fatal(err)
			}
			emp.populateFriends()
		case 2:
			var id int
			fmt.Print("Enter the ID of the existing employee: (Enter 0 to show all employees)")
			_, _ = fmt.Scan(&id)
			if _, ok := emp[id]; !ok && id != 0 {
				fmt.Printf("Employee with ID : %d doesn't exist", id)
				break
			} else {
				emp.displayEmpDetails(id)
			}
		case 3:
			var id int
			fmt.Print("Enter the ID of the existing employee: ")
			_, _ = fmt.Scan(&id)
			if _, ok := emp[id]; !ok {
				fmt.Printf("Employee with ID : %d doesn't exist", id)
				break
			} else if err := emp.createAndUpdateEmployee(id); err != nil {
				log.Fatal(err)
			}
		case 4:
			emp.deleteEmployee()
		case 5:
			emp.showFriends()
		case 6:
			emp.calculateAppraisal()
		case 7:
			emp.showEmergencyContact()
		case 8:
			emp.deleteFriends()
		default:
			fmt.Println("Invalid option. Please enter again")
		}
	}
}
