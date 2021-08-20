package main

import (
	"fmt"
)

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	fn := func(t float64) float64 {
		displacement := (0.5 * a * (t * t)) + (v * t) + s
		return displacement
	}
	return fn
}

func main() {
	var a, v, si, t float64
	fmt.Println("Enter value of acceleration")
	fmt.Scan(&a)
	fmt.Println("Enter value of initial Velocity")
	fmt.Scan(&v)
	fmt.Println("Enter value of initial distance")
	fmt.Scan(&si)
	fmt.Println("Enter value of time")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v, si)
	fmt.Println(fn(t))
}
