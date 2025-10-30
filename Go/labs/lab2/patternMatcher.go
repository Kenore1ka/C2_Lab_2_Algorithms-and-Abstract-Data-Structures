package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func match(text, pattern string) bool {
	n, m := len(text), len(pattern)
	i, j := 0, 0
	starIdx, matchIdx := -1, 0

	for i < n {
		if j < m && (pattern[j] == '?' || pattern[j] == text[i]) {
			i++
			j++
		} else if j < m && pattern[j] == '*' {
			starIdx = j
			matchIdx = i
			j++
		} else if starIdx != -1 {
			j = starIdx + 1
			matchIdx++
			i = matchIdx
		} else {
			return false
		}
	}

	for j < m && pattern[j] == '*' {
		j++
	}

	return j == m
}

func runPatternMatcher(args []string) {
	var query string
	for i, arg := range args {
		if arg == "--query" && i+1 < len(args) {
			query = args[i+1]
			break
		}
	}

	if query == "" {
		fmt.Fprintln(os.Stderr, "Ошибка: не найден аргумент --query.")
		fmt.Fprintln(os.Stderr, "Пример: --query \"MATCH <файл_с_данными> <шаблон>\"")
		return
	}

	parts := strings.Fields(query)
	if len(parts) != 3 || parts[0] != "MATCH" {
		fmt.Fprintln(os.Stderr, "Ошибка: некорректный формат запроса.")
		fmt.Fprintln(os.Stderr, "Пример: --query \"MATCH data.txt h*o\"")
		return
	}
	dataFileName, pattern := parts[1], parts[2]

	file, err := os.Open(dataFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Файл \"%s\" не найден.\n", dataFileName)
		return
	}
	defer file.Close()

	var stringList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringList = append(stringList, scanner.Text())
	}

	if len(stringList) == 0 {
		fmt.Printf("Файл \"%s\" пуст.\n", dataFileName)
		return
	}

	fmt.Printf("Строки из файла \"%s\", соответствующие шаблону \"%s\":\n", dataFileName, pattern)
	foundMatches := false
	for _, currentString := range stringList {
		if match(currentString, pattern) {
			fmt.Println(currentString)
			foundMatches = true
		}
	}

	if !foundMatches {
		fmt.Println("Совпадений не найдено.")
	}
}