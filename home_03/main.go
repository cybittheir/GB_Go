package main

/*
Напишите программу, которая будет хранить ваши url. На основании созданного шаблона допишите код, который позволяет
добавлять новые ссылки, удалять и выводить список.
Для решения задачи используйте структуры. Обязательные поля структуры должны быть дата добавления, имя ссылки,
теги для ссылки через запятую и сам url.

* Задача со звездочкой. Реализовать хранение списка url в файле, подгрузка списка урлов из файла и удаление соответственно.
Можно переделать в этом плане формат и вместо отслеживания нажатий клавиатуры сделать через классические аругменты.

** Задача с двумя звездочками, реализовать полностекстовый поиск по имени ссылки и по тегам.
Можно реализовать сначала только по тегам так как это проще и после этого попробовать взяться за полнотекстовый поиск.
Желательно посмотреть подходы в интернете как можно подобное реализовать.
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type Tags struct {
	Tag []string
}

type Record struct {
	Date time.Time
	Url  string
	Desc string
	Tags
}

func (t Tags) GetTags() []string {
	return t.Tag
}

func (r Record) GetUrl() string {
	return r.Url
}

func (d Record) GetDesc() string {
	return d.Desc
}

func (d Record) GetTime() time.Time {
	return d.Date
}

func main() {

	urlMap := make(map[string]Record)

	fmt.Println("Программа для добавления url в список")
	fmt.Println("a - добавить запись")
	fmt.Println("l - показать список")
	fmt.Println("r - удалить запись")
	fmt.Println("Для выхода из приложения нажмите Esc")

	defer func() {
		// Завершаем работу с клавиатурой при выходе из функции
		fmt.Println("Bye.")
		_ = keyboard.Close()
	}()

OuterLoop:
	for {
		// Подключаем отслеживание нажатия клавиш
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}

		fmt.Print("Выберите режим (a/r/l/ESC): ")
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", string(char))

		switch char {
		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			// Добавление нового url в список хранения
			fmt.Println("\nВведите новую запись в формате <url описание_ссылки теги,через,запятую>:")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("\nОшибка. Требуется ввести аргументы в формате: url описание_ссылки теги,через,запятую")
				continue OuterLoop
			}

			// Напишите свой код здесь
			desc := strings.Replace(args[1], "_", " ", -1)
			tags := strings.Fields(strings.Replace(args[2], ",", " ", -1))
			url := strings.Replace(args[0], "https://", "", 1)
			url = strings.Replace(url, "http://", "", 1)

			r := Record{time.Now(), url, desc, Tags{tags}}
			if urlMap[url].Url == url {
				fmt.Println("\nОшибка. Этот URL уже записан")
				continue OuterLoop
			}
			urlMap[url] = r

			//			records = append(records, r)
			fmt.Println("Запись добавлена")

		case 'l':
			// Вывод списка добавленных url. Выведите количество добавленных url и список с данными url
			// Вывод в формате
			// Имя: <Описание>
			// URL: <url>
			// Теги: <Теги>
			// Дата: <дата>

			// Напишите свой код здесь
			if len(urlMap) < 1 {
				fmt.Println("\nОшибка. Нет ни одной записи")
				continue OuterLoop
			}
			fmt.Println("--=== Список URL ===--")
			for _, r := range urlMap {
				fmt.Printf("Имя: %s;\n", r.Desc)
				fmt.Printf("URL: %s;\n", r.Url)
				fmt.Printf("Теги: %s;\n", r.Tag)
				fmt.Printf("Дата: %s;\n", r.Date.Format("01/02/2006 15:04"))
				fmt.Println("--=== ===--")
			}

		case 'r':
			if len(urlMap) < 1 {
				fmt.Println("\nОшибка. Нет ни одной записи")
				continue OuterLoop
			}
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Удаление url из списка хранения
			fmt.Println("Введите имя ссылки, которое нужно удалить")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			_ = text

			// Напишите свой код здесь
			fmt.Println("--=== Список URL ===--")

		default:
			// Если нажата Esc выходим из приложения
			if key == keyboard.KeyEsc {
				fmt.Println("Finished.")
				return
			}
		}
	}
}
