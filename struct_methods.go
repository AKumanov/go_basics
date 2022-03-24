package main

import "fmt"

func studentFactory(firstName string, lastName string, grades []int, age int) Student {
	student := Student{firstName, lastName, grades, age}
	return student
}

type Student struct {
	firstName string
	lastName  string
	grades    []int
	age       int
}

// created a method to show the age
func (s Student) getAge() int {
	return s.age
}

func (s Student) calculateAverageGrade() float32 {
	var average float32 = 0
	// for i := 0; i < len(grades); i++ {
	// 	average += s.grades[i]
	// }
	// return average / len(grades)

	// another way to iterate over the grades array
	for _, grade := range s.grades {
		average += float32(grade)
	}
	return float32(average / float32(len(s.grades)))
}

func (s *Student) setAge(age int) {
	s.age = age
}

// created a method to show the full name of the student
func (s Student) getFullName() string {
	return s.firstName + " " + s.lastName
}

func main() {
	fmt.Println("Struct Methods in go..")
	s1 := Student{"Alex", "Kumanov", []int{70, 90, 80, 85, 83, 77, 66, 12, 10}, 0}
	age := s1.getAge()
	fmt.Println(age)
	s1.setAge(30)
	fmt.Println(s1.age)
	avgGrade := s1.calculateAverageGrade()
	fmt.Printf("The average grade of %v is: %.2f\n", s1.getFullName(), avgGrade)
	newStudent := studentFactory("Ivan", "Testov", []int{10, 20, 30}, 30)
	fmt.Println("Average grades", newStudent.calculateAverageGrade())

}
