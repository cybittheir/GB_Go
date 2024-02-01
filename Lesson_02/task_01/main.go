package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a [100]int

	l := 20

	t := make([]int, l)

	t1 := rand.Perm(l)
	t2 := rand.Perm(l)
	t3 := rand.Perm(l)
	t4 := rand.Perm(l)
	t5 := rand.Perm(l)

	for x := 0; x < 5; x++ {
		switch x {
		case 0:
			copy(t, t1)
		case 1:
			copy(t, t2)
		case 2:
			copy(t, t3)
		case 3:
			copy(t, t4)
		case 4:
			copy(t, t5)
		}

		for i, v := range t {
			z := rand.Intn(20)
			if z == v {
				a[i+l*x] = 0
			} else {
				a[i+l*x] = z
			}
		}
	}

	fmt.Println(a)

	r := make([]int, l)

	for _, v := range a {
		r[v] += 1
	}

	fmt.Println(r)

}
