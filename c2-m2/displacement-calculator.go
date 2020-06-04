package main

import "fmt"

//fn := GenDisplaceFn(10, 2, 1)
// fmt.Println(fn(3))
// fmt.Println(fn(5))

func main() {
	var d, v, s0, time float64
	fmt.Printf("Enter intial acceleration:\n")
	fmt.Scan(&d)
	fmt.Printf("Enter initial Velocity:\n")
	fmt.Scan(&v)
	fmt.Printf("Enter initial Displacement:\n")
	fmt.Scan(&s0)
	fn := GenDisplaceFn(d, v, s0)
	fmt.Printf("Enter Time:\n")
	fmt.Scan(&time)
	fmt.Println("Displacement value:", fn(time))
}

// GenDisplaceFn forming a function with initial value except time of displace formulation
func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return 1.0/2*acceleration*time*time + initialVelocity*time + initialDisplacement
	}
	return fn
}
