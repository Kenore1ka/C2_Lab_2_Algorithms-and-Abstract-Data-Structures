package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func foldingHash(key string, chunkSize int) int64 {
	var totalSum int64
	fmt.Print("Вычисление: ")
	isFirstPart := true

	for i := 0; i < len(key); i += chunkSize {
		end := i + chunkSize
		if end > len(key) {
			end = len(key)
		}
		partStr := key[i:end]

		partValue, _ := strconv.ParseInt(partStr, 10, 64)
		totalSum += partValue

		if !isFirstPart {
			fmt.Print("+")
		}
		fmt.Print(partStr)
		isFirstPart = false
	}
	fmt.Printf(" = %d\n", totalSum)
	return totalSum
}

func runFoldingHasher(args []string) {
	var query string
	for i, arg := range args {
		if arg == "--query" && i+1 < len(args) {
			query = args[i+1]
			break
		}
	}

	if query == "" {
		fmt.Fprintln(os.Stderr, "Ошибка: не найден аргумент --query.")
		fmt.Fprintln(os.Stderr, "Пример использования: --query \"FOLD <число>\"")
		return
	}

	parts := strings.Fields(query)
	if len(parts) != 2 || parts[0] != "FOLD" {
		fmt.Fprintln(os.Stderr, "Ошибка: некорректный формат запроса.")
		fmt.Fprintln(os.Stderr, "Пример использования: --query \"FOLD 523456795\"")
		return
	}
	numberKey := parts[1]

	for _, c := range numberKey {
		if !unicode.IsDigit(c) {
			fmt.Fprintln(os.Stderr, "Ошибка: ключ должен состоять только из цифр.")
			return
		}
	}

	const chunkSize = 3
	fmt.Println("Ввод:", numberKey)
	fmt.Print("Вывод: ")
	foldingHash(numberKey, chunkSize)
}