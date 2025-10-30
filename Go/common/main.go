package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Получаем все аргументы, переданные программе, кроме имени самой программы.
	args := os.Args[1:]

	// Если аргументов нет, выводить нечего.
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Ошибка: не указаны аргументы.")
		fmt.Fprintln(os.Stderr, "Пример: --file data.txt --query \"HSET mykey myvalue\"")
		return
	}

	// Ищем ключевой аргумент --query, так как он определяет, что нужно делать.
	var query string
	for i := 0; i < len(args); i++ {
		// Проверяем, что текущий аргумент это "--query" и что после него есть еще один (значение).
		if args[i] == "--query" && i+1 < len(args) {
			query = args[i+1]
			break // Нашли, выходим из цикла.
		}
	}

	// Если --query не был найден, программа не может продолжить.
	if query == "" {
		fmt.Fprintln(os.Stderr, "Ошибка: обязательный аргумент --query не найден.")
		return
	}

	// Разбиваем строку запроса на слова, чтобы получить команду.
	parts := strings.Fields(query)
	if len(parts) == 0 {
		fmt.Fprintln(os.Stderr, "Ошибка: запрос в --query не может быть пустым.")
		return
	}
	command := parts[0]

	switch {
	case strings.HasPrefix(command, "H"):
		runHashTable(args)
	case strings.HasPrefix(command, "SET"):
		runSet(args)

	// Ваши оригинальные обработчики
	case strings.HasPrefix(command, "M"):
		runDynamicArray(args)
	case strings.HasPrefix(command, "L"):
		runLinkedList(args)
	case strings.HasPrefix(command, "D"):
		runDLinkedList(args)
	case strings.HasPrefix(command, "Q"):
		runQueue(args)
	case strings.HasPrefix(command, "S"):
		runStack(args)
	default:
		fmt.Fprintf(os.Stderr, "Ошибка: неизвестная команда или группа команд в запросе: '%s'\n", command)
	}
}