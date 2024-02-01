package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[1]

	var fileName, fileExt string
	// Напишите код, который выведет следующее
	// filename: <name>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex
	// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
	// ) Они могут помочь для проверки решения

	s := `\`

	if string(os.PathSeparator) == "/" {
		s = "/"
	} else if s == string(os.PathSeparator) {
		fmt.Println("something")
	}

	fmt.Println(s)

	fileName = filePth[strings.LastIndex(filePth, s)+1:]
	if strings.LastIndex(filePth, ".") > 0 {

		fileExt = fileName[strings.LastIndex(fileName, ".")+1:]
		fileName = fileName[:strings.LastIndex(fileName, ".")]
	}

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
