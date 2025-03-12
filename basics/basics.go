package basics //

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @brief Importing necessary packages.
 *
 * The `import` keyword is used to include the necessary packages in the Go program.
 * Packages are used to organize and reuse code. The `fmt` package is imported here
 * to provide formatted I/O functions, such as printing to the console.
 */

func Run_tutorial() { // Only uppercase functions are exported.
	fmt.Println("Hello, World!")

	var name string = "John Doe"
	var age int = 25
	fmt.Printf("Name: %s\nAge: %d\n", name, age)

	batsman := "David Warner"
	runs := 34
	fmt.Printf("Batsmen: %s\nRuns: %d\n", batsman, runs)

	if runs >= 50 {
		fmt.Println("Half-century!")
	} else { // else should be added in the same line as the closing brace
		fmt.Println("Not a half-century!")
	}

	var wg sync.WaitGroup     //WaitGroup is used to wait for a collection of goroutines to finish.
	for i := 0; i < 50; i++ { //For is the only loop in Go. i has a scope only within the loop. i cannot be declared with var.
		wg.Add(1)
		go printNumbers(i, &wg) // go keyword is used to create a new goroutine. // &wg is the address of the variable wg.
	}
	wg.Wait() // Wait blocks until the WaitGroup counter is zero.
	// fmt.Println(i) will throw an error as i is not defined outside the loop.
	fmt.Println(add(5, 6))

	p := Person{"John Doe", 25}
	fmt.Printf("Name: %s\nAge: %d\n", p.Name, p.Age)

}

func add(a int, b int) int { // lowercase functions are not exported.
	return a + b
}

type Person struct { //Structs are used to define custom data types.
	Name string
	Age  int
}

func printNumbers(n int, wg *sync.WaitGroup) { // *sync.WaitGroup is a pointer to the WaitGroup variable.
	defer wg.Done()
	time.Sleep(1 * time.Second) // Fun fact: time.Sleep(time.Duration(n) * time.Second) is called sleep sort.
	fmt.Println(n)
}
