package main

import "fmt"

func Add(x int, y int) int {
	return x + y
}

func Subtract(x, y int) (int, int) {
	return x - y, y - x
}

// Variadic Functions - Only and should be last
func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Here f has function value
func Mult(f func (fact int) (int, int)) {
	s1, s2 := f(5)
	fmt.Println("x5\t", s1, s2)

	s1, s2 = f(10)
	fmt.Println("x10\t", s1, s2)
}

func main() {
	x, y := 20, 10
	result := Add(x, y)
	fmt.Println(x, "+", y, "=", result)
	result1, result2 := Subtract(x ,y)
	fmt.Println(x, "-", y, "=", result1, "\n", y, "-", x, "=",  result2)

	total := Sum(1, 2, 3, 4)
	fmt.Println("The Sum is:", total)

	total = Sum(5, 6)
	fmt.Println("The Sum is:", total)

	total = Sum()
	fmt.Println("The Sum is:", total)

	// Providing a Slice as an argument
	nums := []int{1, 2, 3, 4, 5}
	total = Sum(nums...)
	fmt.Println("The Sum is:", total)

	a, b := 2, 5
	fn := func(fact int) (int, int) {
		s1, s2 := a * fact, b * fact
		a = 2*a // Value changes
		return s1, s2
	}
	fmt.Println("a =", a, "\tb =", b)
	Mult(fn)
	fmt.Println("a =", a, "\tb =", b)
	b = 2*b
	Mult(fn)
	fmt.Println("a =", a, "\tb =", b)
}
