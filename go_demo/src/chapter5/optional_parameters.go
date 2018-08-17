package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i <= 4; i++ {
		a, b, c := PythagoreanTriple(i, i+1)
		s1 := Heron(a, b, c)
		s2 := Heron(PythagoreanTriple(i, i+1))
		fmt.Printf("△1==%10f==△2== %10f\n", s1, s2)

	}
}

func Heron(a, b, c int) float64 {
	α, β, γ := float64(a), float64(b), float64(c)
	s := (α + β + γ) / 2
	return math.Sqrt(s * (s - α) * (s - β) * (γ))
}
func PythagoreanTriple(m, n int) (a, b, c int) {
	if m < n {
		m, n = n, m
	}
	return (m * m) - (n * n), (2 * m * n), (m * m) + (n * n)
}
