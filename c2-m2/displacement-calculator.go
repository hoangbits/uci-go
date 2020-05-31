package main

import "fmt"

func main() {
	fn := GenDisplaceFn(10, 2, 0)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}

// GenDisplaceFn forming a function with initial value except time of displace formulation
func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return 1.0/2*acceleration*time*time + initialVelocity*time + initialDisplacement
	}
	return fn
}
