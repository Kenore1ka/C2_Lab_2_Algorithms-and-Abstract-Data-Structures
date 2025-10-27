package main

import (
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 4 {
		return
	}

	var fileName, query string
	for i := 0; i < len(args); i++ {
		if args[i] == "--file" && i+1 < len(args) {
			fileName = args[i+1]
			i++
		} else if args[i] == "--query" && i+1 < len(args) {
			query = args[i+1]
			i++
		}
	}

	if fileName == "" || query == "" {
		return
	}

	// Выбор функции для запуска на основе первой буквы команды.
	switch query[0] {
	case 'M':
		runDynamicArray(args)
	case 'L':
		runLinkedList(args)
	case 'D':
		runDLinkedList(args)
	case 'Q':
		runQueue(args)
	case 'S':
		runStack(args)
	// В предоставленном C++ коде не было реализации для 'H' (HashTable) и 'T' (BinaryTree),
	// поэтому они здесь опущены.
	}
}