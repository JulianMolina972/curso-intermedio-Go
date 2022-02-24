package main

/*Definir una serie de propiedades y una serie de metodos propias de esa clase
Un objeto se instancie a partir de esa clase va poder acceder al repertorio de propiedades y metodos.*/
import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

//Forma numero 4 con función
func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	//Forma numero 1
	e := Employee{}
	fmt.Printf("%v\n", e)

	//Forma numero 2
	e2 := Employee{
		id:       1,
		name:     "Name",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)
	//Forma numero 3
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 1
	e3.name = "Name"
	fmt.Printf("%v\n", *e3)

	//Forma numero 4
	e4 := NewEmployee(1, "Name 2", true)
	fmt.Printf("%v\n", *e4)
}
