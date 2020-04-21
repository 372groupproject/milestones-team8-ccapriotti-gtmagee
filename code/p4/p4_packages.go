package main

// import course package
import "./course"

func main() {
	// case for exported functions matters
	c := course.New("Comparative Programming Languages", "372", "Christian Collberg", "CS")
	c.Print()
}
