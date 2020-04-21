package student

type Student struct {
	Name  string
	Grade int
}

func New(name string, grade int) Student {
	s := Student{name, grade}
	return s
}
