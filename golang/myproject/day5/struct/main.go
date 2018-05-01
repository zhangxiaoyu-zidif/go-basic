package main

import (
	"fmt"
)

type Student struct {
	name  string
	sex   string
	grade int
	class int
	score int
	Next  *Student
}

func (stu *Student) NewStudent(name1 string, sex1 string, grade1 int, class1 int, score1 int) {
	stu.name = name1
	stu.class = class1
	stu.grade = grade1
	stu.score = score1
	stu.sex = sex1
}

func main() {
	var stu1 Student
	stu1.name = "xiao mi"
	stu1.class = 1
	stu1.grade = 1
	stu1.score = 100
	stu1.sex = "female"

	var stu2 Student
	stu1.Next = &stu2
	stu2.name = "huawei"

	fmt.Println(stu1)
	fmt.Println(stu1.Next.name)

	var stu3 = new(Student)
	stu3.name = "zte"
	stu2.Next = stu3

	fmt.Println(stu2)
	fmt.Println(stu2.Next.name)

	var stu4 = &Student{}
	stu4.name = "Lenove"
	stu3.Next = stu4
	fmt.Println(stu3)
	fmt.Println(stu3.Next.name)

	var stu5 Student
	stu5.NewStudent("xiaoming", "male", 1, 1, 90)
	fmt.Println(stu5)
}
