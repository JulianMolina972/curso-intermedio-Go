package main

import "fmt"

/*Interfaces
-Diferentes lenguajes de programación utilizan sintaxis explícita para decir
que una clase implementa una interfaz.
-Go lo hace de manera implícita lo que permite la reutilización de codigo y el polimorfismo.
*/

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (ftEmployee FullTimeEmployee) GetMessage() string {
	return "Full Time Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporaryEmployee) GetMessage() string {
	return "Temporary Employee"
}

type PrintInfo interface {
	GetMessage() string
}

func GetMessage(p PrintInfo) {
	fmt.Println(p.GetMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 22
	ftEmployee.id = 5
	fmt.Printf("%v\n", ftEmployee)
	tEmployee := TemporaryEmployee{}
	GetMessage(tEmployee)
	GetMessage(ftEmployee)
	//GetMessage(ftEmployee)
}
