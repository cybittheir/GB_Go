package main

import (
	"fmt"
	"math/rand"
)

func filter(s []int, n int) []int {
	var s0 = make([]int, 0, len(s))
	for _, v := range s {
		if v%n == 0 && v > 0 {
			s0 = append(s0, v)
		}
	}

	return s0
}

func insert(s []int, p int, n int) []int {
	var s0 = make([]int, 0, len(s))
	var s1 = make([]int, 0, len(s))
	var s2 = make([]int, 0, len(s))
	s0 = s[0:p]
	s1 = append(s1, n)
	s1 = append(s1, s[p:]...)
	s2 = append(s0, s1...)
	return s2
}

func main() {
	l := 7
	cap := 10
	var s = make([]int, l, cap)

	for i, _ := range s {
		s[i] = rand.Intn(30)
	}

	fmt.Println(s)

	n := rand.Intn(l-1) + 1
	s1 := filter(s, n)

	fmt.Printf("%d -> %v\n", n, s1)

	p := rand.Intn(l-1) + 2

	v := rand.Intn(50)

	s2 := insert(s, p, v)

	fmt.Printf("%d -> [%d] >> %v\n", v, p, s2)

}
