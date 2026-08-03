package main

import (
	"fmt"
)
import "errors"

//creating map

var mp = make(map[int]Student)

type Student struct {
	name string
	age  int
	mark int
}

//for get All Student

func (student Student) getDetails() map[int]Student {
	if len(mp) == 0 {
		return nil
	}
	return mp
}

// get student by id
func (student Student) getStudentByID(id int) (Student, error) {
	stu, exist := mp[id]

	if exist != true {
		return Student{}, errors.New("student does not found")
	}
	return stu, nil
}

// add student
func (student Student) addStudent(id int, name string, age int, mark int) (string, error) {
	if _, exist := mp[id]; exist == true {
		return "", errors.New("These Student Already exist")
	}

	mp[id] = Student{name, age, mark}

	return "Student Added Successfully", nil
}

// delete student
func (Student Student) removeStudent(id int) (string, error) {
	_, exist := mp[id]

	if exist != true {
		return "", errors.New("student not exist to delete")
	}

	delete(mp, id)
	return "student deleted successfully", nil
}
func main() {
	var choice int
	student := Student{}
	for {
		fmt.Println("Enter your choice : ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			return
		}
		if choice == 1 {
			mpp := student.getDetails()
			for value := range mpp {
				fmt.Printf("ID : %v \nName : %v\nAge : %d\nMark : %v\n", value, mpp[value].name, mpp[value].age, mpp[value].mark)
			}
		}

		if choice == 2 {
			var id int
			var name string
			var age int
			var mark int
			fmt.Print("Enter Student ID : \n")
			_, err = fmt.Scan(&id)
			fmt.Print("Enter Student Name : \n")
			_, err = fmt.Scan(&name)
			fmt.Print("Enter Age : \n")
			_, err = fmt.Scan(&age)
			fmt.Print("Enter Mark : \n")
			_, err = fmt.Scan(&mark)
			addStudent, err := student.addStudent(id, name, age, mark)
			if err == nil {
				fmt.Println(addStudent)
			} else {
				fmt.Println(err)
			}
		}

		if choice == 0 {
			break
		}
	}

}
