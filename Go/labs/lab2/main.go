package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printHelp() {
	fmt.Println("Программа для выполнения заданий лабораторной работы №2.")
	fmt.Println("Использование: ./lab2_program <номер_задания> [аргументы...]")
	fmt.Println()
	fmt.Println("Доступные задания:")
	fmt.Println("  1: Вычислитель выражений (интерактивный режим).")
	fmt.Println("     Пример: ./lab2_program 1")
	fmt.Println()
	fmt.Println("  2: Сопоставление с паттерном.")
	fmt.Println("     Пример: ./lab2_program 2 --query \"MATCH emails.txt @.ru\"")
	fmt.Println()
	fmt.Println("  3: Проверка АВЛ-сбалансированности (интерактивный режим).")
	fmt.Println("     Пример: ./lab2_program 3")
	fmt.Println()
	fmt.Println("  4: Хеш-функция (метод свертки).")
	fmt.Println("     Пример: ./lab2_program 4 --query \"FOLD 523456795\"")
	fmt.Println()
	fmt.Println("  5: LRU Кэш.")
	fmt.Println("     Пример: ./lab2_program 5 \"cap = 2, Q = 2 Queries = SET 1 2 GET 1\"")
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	taskNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка: первый аргумент должен быть числом - номером задания.")
		printHelp()
		return
	}

	args := os.Args[2:]

	switch taskNumber {
	case 1:
		runExpressionSolver()
	case 2:
		runPatternMatcher(args)
	case 3:
		runAvlChecker()
	case 4:
		runFoldingHasher(args)
	case 5:
		if len(args) > 0 {
			// Объединяем все аргументы в одну строку для парсинга
			runLruCache(strings.Join(args, " "))
		} else {
			fmt.Fprintln(os.Stderr, "Ошибка: для задания 5 требуется строка с параметрами кэша.")
		}
	default:
		fmt.Fprintf(os.Stderr, "Ошибка: неизвестный номер задания '%d'.\n", taskNumber)
		printHelp()
		return
	}
}