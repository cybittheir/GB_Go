package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	t := strings.TrimSpace(text)

	if len(t) == 0 {
<<<<<<< HEAD
		log.Fatal("Нет ни одного значимого символа. Повторите ввод предложения.")
	}

	m := make(map[rune]int)
	var l int
	var x int
	for _, v := range t {
		if unicode.IsLetter(v) {
			m[unicode.ToLower(v)]++
			l++
		}
		x++
=======
		log.Fatal("Нет ни одного символа. Повторите ввод.")
	}

	m := make(map[rune]int)

	for _, v := range t {
		if unicode.IsLetter(v) {
			m[unicode.ToLower(v)]++
		}
>>>>>>> 99eb12bcde0205d1e37a23fddcb2ee0834dbea8d
	}

	z := len(m)

	if z == 0 {
<<<<<<< HEAD
		log.Fatal("Не обнаружено ни одной буквы. Повторите ввод предложения.")
=======
		log.Fatal("Не обнаружено ни одной буквы. Повторите ввод.")
>>>>>>> 99eb12bcde0205d1e37a23fddcb2ee0834dbea8d
	}

	var p float32
	for i, s := range m {
<<<<<<< HEAD
		p = float32(s) / float32(l)
		fmt.Printf("%c - %d; %0.2f / %0.1f", i, s, p, 100*p)
		fmt.Println("%")
	}
	fmt.Println("===================")
	fmt.Printf("Всего букв: %d; уникальных букв: %d;\nВсего символов: %d\n", l, z, x)
=======
		p = float32(s) / float32(z)
		fmt.Printf("%c - %d; %0.2f / %0.1f", i, s, p, 100*p)
		fmt.Println("%")
	}
>>>>>>> 99eb12bcde0205d1e37a23fddcb2ee0834dbea8d
}
