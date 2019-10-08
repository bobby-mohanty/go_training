package main

import (
	"fmt"
	"math"
	"reflect"
)

type myInt int

func main()  {
	var f float64 = math.Pi
	var i myInt = 50
	v1 := reflect.ValueOf(f)
	v2 := reflect.ValueOf(i)
	fmt.Println("Type of variable 'f' : ", reflect.TypeOf(f))
	fmt.Println("Type of variable 'i' : ", reflect.TypeOf(i))
	fmt.Println("Reflect ValueOf params")
	fmt.Println("\nValueOf for variable 'f' : \nType :", v1.Type(), "\nKind :", v1.Kind(), "\nValue :", v1.Float(),
		"\nSettable :", v1.CanSet())
	fmt.Println("\nValueOf for variable 'f' : \nType :", v2.Type(), "\nKind :", v2.Kind(), "\nValue :", v2.Int(),
		"\nSettable :", v2.CanSet())
	fmt.Println("ValueOf on pointers")
	v3 := reflect.ValueOf(&f)
	v4 := reflect.ValueOf(&i)
	fmt.Println("\nValueOf for variable '&f' : \nType :", v3.Type(), "\nKind :", v3.Kind(), "\nValue :", v3.Elem(),
		"\nSettable :", v3.Elem().CanSet())
	fmt.Println("\nValueOf for variable '&i' : \nType :", v4.Type(), "\nKind :", v4.Kind(), "\nValue :", v4.Elem(),
		"\nSettable :", v4.Elem().CanSet())
}