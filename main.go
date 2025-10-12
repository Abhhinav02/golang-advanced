package main

import (
	"fmt"
	"reflect"
)

// methods( ) with 'reflect'

type Greeter struct{}

// Method with two parameters
func (g Greeter) Greet(fname, lname string) string {
	return "Hello " + fname + " " + lname
}

func main() {
	g := Greeter{}
	t := reflect.TypeOf(g)
	v := reflect.ValueOf(g)

	fmt.Println("Type:", t)

	// ✅ Correct way: iterate through methods properly
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("Method %d: %s\n", i, method.Name)
	}

	// ✅ Directly invoke the "Greet" method by name
	m := v.MethodByName("Greet")
	if !m.IsValid() {
		fmt.Println("❌ Method not found")
		return
	}

	// Call the method with arguments using reflection
	results := m.Call([]reflect.Value{
		reflect.ValueOf("Skyy"),
		reflect.ValueOf("Banerjee"),
	})

	// ✅ Extract and print the returned string
	fmt.Println("Greet result:", results[0].String())
}


// O/P
// $ go run main.go
// Type: main.Greeter
// Method 0: Greet
// Greet result: Hello Skyy Banerjee
