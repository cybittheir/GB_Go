package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) bool {
	// Напишите код, проверяющий является ли число простым здесь
	//
	//
	for s := 2; s < n/2; s++ {
		if n%s == 0 {
			fmt.Println("деление на", s, "без остатка")
			return false
		}
	}
	return true
}

func main() {
	// В данном примере мы сознатель опускаем проверку типов.
	// Если передать в качестве ввода значение типа строка или булево значение, то программа аварийно завершится
	// Пакет log предоставляет нам базовые функции логирования.
	// log.Fatal печатает на экране переданный текст ошибки и звершается
	var value int
	fmt.Println("Введите число")
	if _, err := fmt.Fscan(os.Stdin, &value); err != nil {
		log.Fatal(err)
	}

	if isPrime(value) {
		fmt.Println("Число", value, "простое")
	} else {
		fmt.Println("Число", value, "не простое")

	}
}
