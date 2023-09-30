import (
"fmt"
"math"
)
func findGreaterNumber(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func main() {
	num1 := 10
    num2 := 20

    result := findGreaterNumber(num1, num2)
    fmt.Printf("The greater number between %d and %d is: %d\n", num1, num2, result)
// Function 1: Add two integers and return the result
add := func(x, y int) int {
return (x + y)
}


// Function 2: Subtract two integers and return the result
subtract := func(x, y int) int {
return x - y
}

// Function 3: Multiply two integers and return the result
multiply := func(x, y int) int {
return x * y
}

// Function 4: Divide two floats and return the result
divide := func(x, y float64) float64 {
if y == 0 {
panic("Division by zero is not allowed.")
}
return x / y
}

// Function 5: Calculate the square root of a float and return it
squareRoot := func(x float64) float64 {
if x < 0 {
panic("Square root of a negative number is not defined.")
}
return math.Sqrt(x)
}

// Function 6: Print a message to the console
printMessage := func(message string) {
fmt.Println(message)
}

// Example usage of the functions
fmt.Println("1 + 2 =", add(1, 2))
fmt.Println("5 - 3 =", subtract(5, 3))
fmt.Println("4 * 6 =", multiply(4, 6))
fmt.Println("10 / 2 =", divide(10, 2))
fmt.Println("Square root of 25 =", squareRoot(25))
printMessage("Hello, Go!")
}