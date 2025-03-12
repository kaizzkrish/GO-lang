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
}
