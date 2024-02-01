package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	var wordLine, symbolNumber, bytesNumber, wordSingle int
	// Напишите код здесь

	// Подсказка. Чтобы обойти все символы в строке используйте цикл for _, v := range value
	// Помните вам нужно вывести количество букв, а не символов.
	// Используйте unicode.IsSpace() для определения пробела
	// Используйте unicode.IsLetter() для определения буквы

	//	text = "Помните вам нужно вывести количество букв, а не символов!"

	for _, v := range text {
		//		fmt.Printf("Symbol: %c\n", v)
		if unicode.IsLetter(v) {
			symbolNumber++
			wordSingle++
		} else if (!unicode.IsLetter(v) || unicode.IsSpace(v)) && wordSingle > 0 {
			wordLine++
			wordSingle = 0
		}
		bytesNumber++
	}

	if wordSingle > 0 {
		wordLine++

	}
	fmt.Printf("Количество слов: %d\n", wordLine)
	fmt.Printf("Количество букв: %d\n", symbolNumber)
	fmt.Printf("Количество байт: %d\n", bytesNumber)
	fmt.Printf("Количество байт: %d\n", len(text))
}
