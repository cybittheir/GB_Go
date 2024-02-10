package main

import (
	"fmt"
	"log"
	"os"
<<<<<<< HEAD
=======
	"regexp"
>>>>>>> 99eb12bcde0205d1e37a23fddcb2ee0834dbea8d
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[1]

<<<<<<< HEAD
=======
	//	filePth := "d:\\php7\\extras\\ssl\\openssl"

>>>>>>> 99eb12bcde0205d1e37a23fddcb2ee0834dbea8d
	var fileName, fileExt string
	// Напишите код, который выведет следующее
	// filename: <name>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex
	// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
	// ) Они могут помочь для проверки решения
<<<<<<< HEAD

	s := `\`

	if string(os.PathSeparator) == "/" {
		s = "/"
	} else if s == string(os.PathSeparator) {
		fmt.Println("something")
	}

	fmt.Println(s)

	fileName = filePth[strings.LastIndex(filePth, s)+1:]
	if strings.LastIndex(filePth, ".") > 0 {
=======
	//==============================

	// проверка на правильность написания пути к файлу для Windows (обязательное экранирование обратного слеша)

	winSep := `\`

	if winSep == string(os.PathSeparator) {
		if regexp.MustCompile(`:`).MatchString(filePth) {
			checkPth(`:\\`, filePth)
		} else if filePth[0:1] == string('\\') {
			checkPth(`\\`, filePth[1:])
			checkPth(`\\`, filePth[2:]) // на случай если в начале пути слеши экранируются, а дальше нет. Хотя это конечно не уберегает от ошибки в одном месте
		}
	}

	// сначала отделяем имя файла, потом, если есть, расширение

	fileName = filePth[strings.LastIndex(filePth, string(os.PathSeparator))+1:]

	if strings.LastIndex(fileName, ".") >= 0 {
>>>>>>> 99eb12bcde0205d1e37a23fddcb2ee0834dbea8d

		fileExt = fileName[strings.LastIndex(fileName, ".")+1:]
		fileName = fileName[:strings.LastIndex(fileName, ".")]
	}

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
<<<<<<< HEAD
=======

func checkPth(r, s string) bool {

	if !(regexp.MustCompile(r).MatchString(s)) {
		log.Fatal("Ошибка ввода. Требуется разделять названия папок двойным обратным слешем - \\\\")
	}
	return true
}
>>>>>>> 99eb12bcde0205d1e37a23fddcb2ee0834dbea8d
