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

	//	filePth := "d:\\php7\\extras\\ssl\\openssl"

	var fileName, fileExt string
	// Напишите код, который выведет следующее
	// filename: <name>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex
	// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
	// ) Они могут помочь для проверки решения

	fileName = filePth[strings.LastIndex(filePth, "\\")+1:]
	if strings.LastIndex(filePth, ".") > 0 {

		fileExt = fileName[strings.LastIndex(fileName, ".")+1:]
		fileName = fileName[:strings.LastIndex(fileName, ".")]
	}

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
