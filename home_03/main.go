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
	"encoding/json"
	"fmt"
	"io"
	"log"
	"maps"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type Tags []string

func (t Tags) String() string {
	var s string
	for _, v := range t {
		s += fmt.Sprintf("#%s ", v)
	}
	return strings.TrimSpace(s)
}

type Record struct {
	Date time.Time
	Url  string
	Desc string
	Tags
}

func (r Record) GetUrl() string {
	return r.Url
}

func (d Record) GetDesc() string {
	return d.Desc
}

func (d Record) GetTime() string {
	return d.Date.Format("01/02/2006 15:04")
}

func (d Record) String() string {
	var s strings.Builder
	s.WriteString("Имя: ")
	s.WriteString(d.GetDesc())
	s.WriteString(";\n")
	s.WriteString("URL: ")
	s.WriteString(d.GetUrl())
	s.WriteString(";\n")
	s.WriteString("Теги: ")
	s.WriteString(d.String())
	s.WriteString(";\n")
	s.WriteString("Дата: ")
	s.WriteString(d.GetTime())
	s.WriteString(";\n")
	return s.String()
}

// пока подразумеваем что не может быть двух ссылок с полностью одинаковыми названиями. Иначе результат нужно собрать в слайс
func searchRec(m map[string]Record, s string) string {
	for url, rec := range m {
		if strings.EqualFold(rec.Desc, s) {
			return url
		}
	}
	return ""
}

// опять-таки как самый простой вариант. Кажется, что в действительности лучше собирать отдельную мапу слайсов с ключём-тэгом
func searchTag(m map[string]Record, s string) []string {
	var r []string
	for url, rec := range m {
		for _, t := range rec.Tags {
			if strings.EqualFold(t, s) {
				r = append(r, url)
			}
		}
	}
	return r
}

const fn = "urllist.json"

func writeFile(m map[string]Record, f string) {
	j, _ := json.Marshal(m)

	err := os.WriteFile(f, j, 0644)
	if err != nil {
		fmt.Printf("Ошибка сохранения файла: %s\n", f)
		fmt.Println(err)
	} else {
		fmt.Printf("Список сохранен в файл %s\n\n", fn)
	}

}

func readFile(f string) map[string]Record {
	var m map[string]Record
	fmt.Printf("Чтение списка из файла: %s\n", f)
	j, err := os.Open(f)
	if err != nil {
		fmt.Printf("Невозможно открыть файл: %s\n", f)
		fmt.Println(err)
	} else {
		fmt.Println("...")
		if byteValue, err := io.ReadAll(j); err != nil {
			fmt.Println("Невозможно прочитать список")
			fmt.Println(err)
		} else {
			fmt.Println("Список успешно прочитан")
			json.Unmarshal(byteValue, &m)
		}
	}
	defer j.Close()

	return m
}

func main() {

	fmt.Println("** Программа для добавления url в список **")
	fmt.Println("")

	urlMap := make(map[string]Record)

	if _, err := os.Stat(fn); err == nil {
		oldList := readFile(fn)
		if len(oldList) > 0 {
			urlMap = oldList
		}
	}

	fmt.Println("=====================================")
	fmt.Println("a - добавить запись")
	fmt.Println("l - показать список")
	fmt.Println("r - удалить запись")
	fmt.Println("s - найти запись по описанию")
	fmt.Println("t - найти запись по тэгу")
	fmt.Println("w - сохранить список в файл")
	fmt.Println("g - прочитать список из файла")
	fmt.Println("-------------------------------------")
	fmt.Println("Для выхода из приложения нажмите Esc")
	fmt.Println("=====================================")

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

		fmt.Print("Выберите режим (a/r/l/s/t/w/g/ESC): ")
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n\n", string(char))

		switch char {
		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			// Добавление нового url в список хранения
			fmt.Println("\nВведите новую запись в формате <url описание_ссылки теги,через,запятую>:")
			fmt.Print("> ")

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

			// проверка, записан ли уже этот url
			if urlMap[url].Url != "" {
				fmt.Println("Ошибка. Этот URL уже записан")
				fmt.Println("")
				continue OuterLoop
			}

			urlMap[url] = Record{time.Now(), args[0], desc, tags}

			fmt.Println("Запись добавлена")
			fmt.Println("")

		case 'l':
			// Вывод списка добавленных url. Выведите количество добавленных url и список с данными url
			// Вывод в формате
			// Имя: <Описание>
			// URL: <url>
			// Теги: <Теги>
			// Дата: <дата>

			// Напишите свой код здесь
			// проверка на наличие записей
			if len(urlMap) < 1 {
				fmt.Println("Ошибка. Нет ни одной записи")
				fmt.Println("")
				continue OuterLoop
			}
			// вывод результата
			fmt.Println("\n--=== Список URL ===--")
			for _, r := range urlMap {
				fmt.Print(r.String())
				fmt.Println("--=== ===--")
			}

		case 'r':
			if len(urlMap) < 1 {
				fmt.Println("Ошибка. Нет ни одной записи")
				fmt.Println("")
				continue OuterLoop
			}
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Удаление url из списка хранения
			fmt.Println("\nВведите имя ссылки, которое нужно удалить:")
			fmt.Print("> ")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			//			_ = text
			text = strings.TrimRight(text, "\r\n")
			text = strings.TrimRight(text, "\r")
			// Напишите свой код здесь
			if url := searchRec(urlMap, text); url != "" {
				delUrl := urlMap[url].Url
				delete(urlMap, url)
				fmt.Printf("Запись %s (%s) удалена\n\n", text, delUrl)
			} else {
				fmt.Printf("\nОшибка. Нет записи %s", text)
				continue OuterLoop
			}

		case 's':
			if len(urlMap) < 1 {
				fmt.Println("Поиск невозможен. Нет ни одной записи")
				fmt.Println("")
				continue OuterLoop
			}
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Поиск url в списке хранения
			fmt.Println("Введите имя ссылки для поиска:")
			fmt.Print("> ")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')

			text = strings.TrimRight(text, "\r\n")
			text = strings.TrimRight(text, "\r")

			if url := searchRec(urlMap, text); url != "" {
				fmt.Println("\n--=== Результат поиска ===--")
				fmt.Print(urlMap[url].String())
				fmt.Println("--=== ===--")
			} else {
				fmt.Printf("\nЗапись %s не найдена\n\n", text)
				continue OuterLoop
			}

		case 't':
			if len(urlMap) < 1 {
				fmt.Println("Поиск невозможен. Нет ни одной записи")
				fmt.Println("")
				continue OuterLoop
			}
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Поиск url в списке хранения
			fmt.Println("Введите тэг для поиска:")
			fmt.Print("> ")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')

			text = strings.TrimRight(text, "\r\n")
			text = strings.TrimRight(text, "\r")

			if urls := searchTag(urlMap, text); len(urls) > 0 {
				fmt.Println("\n--=== Результат поиска ===--")
				for _, u := range urls {
					fmt.Print(urlMap[u].String())
					fmt.Println("--=== ===--")
				}
			} else {
				fmt.Printf("\nЗапись по тэгу %s не найдена\n\n", text)
				continue OuterLoop
			}
		case 'w':
			if len(urlMap) < 1 {
				fmt.Println("Ошибка. Нет ни одной записи")
				fmt.Println("")
				continue OuterLoop
			}

			writeFile(urlMap, fn)

		case 'g':
			if len(urlMap) < 1 {
				urlMap = readFile(fn)
			} else {
				maps.Copy(urlMap, readFile(fn))
			}

		default:
			// Если нажата Esc выходим из приложения
			if key == keyboard.KeyEsc {
				fmt.Println("Finished.")
				return
			}
		}
	}
}
