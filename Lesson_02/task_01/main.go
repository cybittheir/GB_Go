package main

import (
	"fmt"
	"math/rand"
<<<<<<< HEAD
	"time"
)

func main() {

	var arr [100]int
	max := 20
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)
	for i := range arr {
		value := rnd.Intn(max)
		arr[i] = value
	}
	fmt.Println(arr)

	// Создайте слайс нужного размера
	slice := make([]int, max)

	for _, v := range arr {
		slice[v]++
	}

	// Заполните слайс количеством встречаемых чисел

	fmt.Println(slice)
=======
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
>>>>>>> 99eb12bcde0205d1e37a23fddcb2ee0834dbea8d

}
