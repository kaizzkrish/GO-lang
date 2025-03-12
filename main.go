package main //
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	var name string = "John Doe"
	var age int = 25
	var name1 string = "Jane Doe"
	var age1 int = 30
	fmt.Printf("Name:%s\nAge:%d\nname1:%s\nAge1:%d\n", name1, age, name, age1)

	batsman := "David Warner"
	runs := 151
	fmt.Printf("Batsmen: %s\nRuns: %d\n", batsman, runs)

	if runs >= 50 && runs < 100 { // && should be used instead of &
		fmt.Println("Half-century!")
	} else if runs >= 100 { // elif should be used instead of else if
		fmt.Println("Century!")
	} else { // else should be added in the same line as the closing brace
		fmt.Println("Not a half-century!")
	}

	var wg sync.WaitGroup    //WaitGroup is used to wait for a collection of goroutines to finish.
	for i := 0; i < 5; i++ { //For is the only loop in Go. i has a scope only within the loop. i cannot be declared with var.
		wg.Add(1)
		go printNumbers(i, &wg) // go keyword is used to create a new goroutine. // &wg is the address of the variable wg.
	}
	wg.Wait() // Wait blocks until the WaitGroup counter is zero.
	// fmt.Println(i) will throw an error as i is not defined outside the loop.
	fmt.Println(add(5, 6))

	type Person struct { //Structs are used to define custom data types.
		Name      string
		Age       string
		MarriedAt int
	}

	p := Person{"John Doe", "25", 25}
	fmt.Printf("Name: %s\nAge: %s\nMArried at: %d\n", p.Name, p.Age, p.MarriedAt)

	count := 0
	for count < 5 {
		fmt.Println("count", count)
		count++
	}
}
func add(a int, b int) int {
	return a + b
}

func printNumbers(n int, wg *sync.WaitGroup) { // *sync.WaitGroup is a pointer to the WaitGroup variable.
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println(n)
}
