package main

import (
<<<<<<< HEAD
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Напишите функцию фильтрации слайса. Отфильтруйте слайс arr1 так, чтобы он содержал только не четные числа
// То есть, например arr1 = [0, 2, 3, 1, 5, 4] после фильтрации [3, 1, 5]
// Не используйте готовые функции из пакета slices
func filter(arr1 []int) []int {
	var r []int
	for _, v := range arr1 {
		if v > 0 && v%2 != 0 {
			r = append(r, v)
		}
	}
	return r
}

// Напишите функцию вставки элемента в слайс на позицию
// То есть, например arr1 = [0, 2, 3, 1, 5, 4] pos = 1, value = 4. Результат [0, 4, 2, 3, 1, 5, 4]
// Не используйте готовые функции из пакета slices
func insert(pos, value int, arr1 []int) []int {

	return append(arr1[:pos], append([]int{value}, arr1[pos:]...)...)
}

func main() {
	// Этот код нужен для ввода массива из стандартного ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите элементы массива через пробел: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Разбиваем строку на массив строк
	strValues := strings.Split(input, " ")

	// Преобразуем строки в числа и заполняем массив
	var arr []int
	for _, str := range strValues {
		val, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		arr = append(arr, val)
	}

	var pos int
	fmt.Println("Введите позицию для вставки")
	if _, err := fmt.Fscan(reader, &pos); err != nil {
		log.Fatal("Неправильное значение")
	}

	if pos < 0 || pos > len(arr)-1 {
		log.Fatal("Позиция вставки должна входить в диапазон индексов введенного слайса")
	}

	var value int
	fmt.Println("Введите значение для вставки")
	if _, err := fmt.Fscan(os.Stdin, &value); err != nil {
		log.Fatal("Неправильное значение")
	}

	// Скопируйте слайс arr в слайс arr1.
	// Это можно сделать через функцию Append или через функцию copy

	arr1 := make([]int, len(arr))
	copy(arr1, arr)

	fmt.Println(arr1)
	arr1 = insert(pos, value, arr1)
	fmt.Println(arr1)

	arr1 = filter(arr1)

	fmt.Println(arr1)
=======
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

	n := rand.Intn(l-1) + 2
	s1 := filter(s, n)

	fmt.Printf("%d -> %v\n", n, s1)

	p := rand.Intn(l-1) + 2

	v := rand.Intn(50)

	s2 := insert(s, p, v)

	fmt.Printf("%d -> [%d] >> %v\n", v, p, s2)

>>>>>>> 99eb12bcde0205d1e37a23fddcb2ee0834dbea8d
}
