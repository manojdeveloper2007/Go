package student

type Student struct {
	name string
	age  int
	mark int
}

func (s Student) GetStudent() Student {
	stu := Student{"Manoj", 19, 97}
	return stu
}
