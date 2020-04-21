package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./student"
)

var students = []student.Student{}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("\nCommands: \n")
		fmt.Println("  enter: enter a student's information ")
		fmt.Println("  average: get the average grade")
		fmt.Println("  highest: get the highest grade")
		fmt.Println("  quit: exit the program \n")

		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "enter":
			enter()
		case "average":
			average()
		case "highest":
			highest()
		case "quit":
			os.Exit(0)
		}
	}
}

func enter() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Student's name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Printf("Student's grade: ")
	scanner.Scan()
	grade, _ := strconv.Atoi(scanner.Text())

	s := student.New(name, grade)
	students = append(students, s)
}

func average() {
	average := 0
	i := 0

	for _, student := range students {
		average += student.Grade
		i++
	}

	fmt.Printf("\nAverage: %d\n", average/i)
}

func highest() {
	var highest student.Student

	for _, student := range students {
		if student.Grade > highest.Grade {
			highest = student
		}
	}

	fmt.Printf("\nHighest: %s - %d\n", highest.Name, highest.Grade)
}
