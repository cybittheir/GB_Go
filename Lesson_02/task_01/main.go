package main

import (
	"fmt"
	"math/rand"
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

}
