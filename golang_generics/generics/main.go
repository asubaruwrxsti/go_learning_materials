package main

import "fmt"

func main() {
	// Initialize a map for the integer values
	m1 := map[string]int64{
		"first": 1,
		"second": 2,
	}

	// Initialize a map for the float values
	m2 := map[string]float64{
		"pi": 3.14,
		"e": 2.718,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(m1), 
		SumFloats(m2),
	)

	fmt.Printf("Generic Sum: %v and %v\n",
		SumIntsOrFloats[string, int64](m1), // [string, int64] is the type parameters for SumIntsOrFloats
		SumIntsOrFloats[string, float64](m2), // [string, float64] is the type parameters for SumIntsOrFloats
	)

	fmt.Printf("Generic Sum, inferred: %v and %v\n",
		SumIntsOrFloats(m1), // [string, int64] is the type parameters for SumIntsOrFloats
		SumIntsOrFloats(m2), // [string, float64] is the type parameters for SumIntsOrFloats
	)
}

// SumInts adds together the values of m
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
func SumIntsOrFloats[K comparable, V int64 | float64] (m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Declare a SumIntsOrFloats function with two type parameters (inside the square brackets), K and V, and one argument that uses the type parameters, m of type map[K]V. The function returns a value of type V.

// The type parameters are declared as comparable, which means that they can be compared with == and !=. The type parameters are used in the function signature to declare the type of the argument m and the return type of the function.

// Specify for the V type parameter a constraint that is a union of two types: int64 and float64. Using | specifies a union of the two types, meaning that this constraint allows either type. Either type will be permitted by the compiler as an argument in the calling code.

// Specify that the m argument is of type map[K]V, where K and V are the types already specified for the type parameters. Note that we know map[K]V is a valid map type because K is a comparable type. If we hadn’t declared K comparable, the compiler would reject the reference to map[K]V.