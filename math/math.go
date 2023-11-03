package math
import "fmt"
// Add is our function that sums two integers
func Add(x, y int) (res int) {
	fmt.Println("this is local module")
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) (res int) {
	return x - y
}
