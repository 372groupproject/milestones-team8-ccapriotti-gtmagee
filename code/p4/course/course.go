package course

import "fmt"

type course struct {
	name       string
	number     string
	instructor string
	department string
}

func New(name string, number string, instructor string, department string) course {
	c := course{name, number, instructor, department}
	return c
}

func (c course) Print() {
	fmt.Printf("%s%s: %s - taught by %s\n", c.department, c.number, c.name, c.instructor)
}
