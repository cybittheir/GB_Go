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
		log.Fatal("Нет ни одного символа. Повторите ввод.")
	}

	m := make(map[rune]int)

	for _, v := range t {
		if unicode.IsLetter(v) {
			m[unicode.ToLower(v)]++
		}
	}

	z := len(m)

	if z == 0 {
		log.Fatal("Не обнаружено ни одной буквы. Повторите ввод.")
	}

	var p float32
	for i, s := range m {
		p = float32(s) / float32(z)
		fmt.Printf("%c - %d; %0.2f / %0.1f", i, s, p, 100*p)
		fmt.Println("%")
	}
}
