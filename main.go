package main //
import "fmt"

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

	for i := 0; i < 8; i++ { //For is the only loop in Go. i has a scope only within the loop. i cannot be declared with var.
		fmt.Println(i)
	}
	// fmt.Println(i) will throw an error as i is not defined outside the loop.
	fmt.Println(add(5, 6))

	count := 0
	for count < 5 {
		fmt.Println("count", count)
		count++
	}
}
func add(a int, b int) int {
	return a + b
}
