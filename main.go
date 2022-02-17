package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int //variable explicita
	x = 8
	y := 7 //variable implicita

	fmt.Println(x)
	fmt.Println(y)

	myValue, err := strconv.ParseInt("7", 0, 64) // convertir string a enteros(int)
	if err != nil {                              //manejo de errores
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}
	//Los mapas son estructuras de llave valor
	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println(m["Key"])
	//range sirve para hacer iteraciones
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	//Agregamos con append un nuevo elemento al slice
	s = append(s, 16)
	//recorremos nuevamente el slice
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	//c := make(chan int)
	//go doSomething(c)
	//<-c

	g := 25
	fmt.Println(g)
	h := &g //le asigno la direccion del dato g
	f := *h //puntero *h     g y f son numeros enteros y h es el puntero
	fmt.Println(h)
	fmt.Println(f)
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
