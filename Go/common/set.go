package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// runSet - функция-обработчик для всех команд, связанных с множеством.
func runSet(args []string) {
	var fileName, query string
	for i := 0; i < len(args); i++ {
		if args[i] == "--file" && i+1 < len(args) {
			fileName = args[i+1] // Файл по умолчанию для операций с одним множеством
			i++
		} else if args[i] == "--query" && i+1 < len(args) {
			query = args[i+1]
			i++
		}
	}

	if query == "" {
		fmt.Fprintln(os.Stderr, "Ошибка: аргумент --query не найден.")
		return
	}

	parts := strings.Fields(query)
	command := parts[0]
	
	ht := NewHashTable()

	switch command {
	case "SETADD":
		if fileName != "" && len(parts) == 2 {
			ht.LoadFromFile(fileName)
			// Для множества ключ и значение одинаковы
			ht.Insert(parts[1], parts[1])
			ht.SaveToFile(fileName)
		} else {
			fmt.Fprintln(os.Stderr, "Ошибка: для SETADD требуются --file и значение.")
		}
	case "SETDEL":
		if fileName != "" && len(parts) == 2 {
			ht.LoadFromFile(fileName)
			ht.Remove(parts[1])
			ht.SaveToFile(fileName)
		} else {
			fmt.Fprintln(os.Stderr, "Ошибка: для SETDEL требуются --file и значение.")
		}
	case "SET_AT":
		if fileName != "" && len(parts) == 2 {
			ht.LoadFromFile(fileName)
			if _, found := ht.Get(parts[1]); found {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		} else {
			fmt.Fprintln(os.Stderr, "Ошибка: для SET_AT требуются --file и значение.")
		}
	case "SET_UNION":
		if len(parts) == 3 {
			fileA, fileB := parts[1], parts[2]
			ht.LoadFromFile(fileA)
			ht.LoadFromFile(fileB) // Дубликаты просто перезапишутся, что для множества нормально
			fmt.Printf("Объединение множеств (%s U %s):\n", fileA, fileB)
			ht.Print()
		} else {
			fmt.Fprintln(os.Stderr, "Ошибка: для SET_UNION требуются два имени файла в запросе.")
		}
	case "SET_INTERSECTION":
		if len(parts) == 3 {
			fileA, fileB := parts[1], parts[2]
			ht.LoadFromFile(fileA) // Загружаем множество A
			
			file, err := os.Open(fileB)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Ошибка: не удалось открыть файл %s\n", fileB)
				return
			}
			defer file.Close()

			fmt.Printf("Пересечение множеств (%s n %s):\n", fileA, fileB)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				// Берем ключ из файла B
				key := strings.Fields(scanner.Text())[0]
				// Если элемент из B есть в A, печатаем его
				if _, found := ht.Get(key); found {
					fmt.Println(key)
				}
			}
		} else {
			fmt.Fprintln(os.Stderr, "Ошибка: для SET_INTERSECTION требуются два имени файла в запросе.")
		}
	case "SET_DIFFERENCE":
		if len(parts) == 3 {
			fileA, fileB := parts[1], parts[2]
			ht.LoadFromFile(fileA) // Загружаем множество A

			file, err := os.Open(fileB)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Ошибка: не удалось открыть файл %s\n", fileB)
				return
			}
			defer file.Close()
			
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				// Удаляем из A все элементы, которые есть в B
				key := strings.Fields(scanner.Text())[0]
				ht.Remove(key)
			}
			
			fmt.Printf("Разность множеств (%s - %s):\n", fileA, fileB)
			ht.Print()
		} else {
			fmt.Fprintln(os.Stderr, "Ошибка: для SET_DIFFERENCE требуются два имени файла в запросе.")
		}
	default:
		fmt.Fprintln(os.Stderr, "Неизвестная команда для множества:", command)
	}
}