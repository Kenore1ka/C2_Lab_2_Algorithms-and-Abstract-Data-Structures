package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// StackNode представляет узел стека.
type StackNode struct {
	data string
	next *StackNode
}

// Stack представляет стек.
type Stack struct {
	top *StackNode
}

// Init инициализирует стек.
func (s *Stack) Init() {
	s.top = nil
}

// Push добавляет элемент на вершину стека.
func (s *Stack) Push(value string) {
	newNode := &StackNode{data: value, next: s.top}
	s.top = newNode
}

// Pop удаляет элемент с вершины стека.
func (s *Stack) Pop() {
	if s.top == nil {
		return
	}
	s.top = s.top.next
}

// Print выводит содержимое стека.
func (s *Stack) Print() {
	temp := s.top
	for temp != nil {
		fmt.Print(temp.data + " ")
		temp = temp.next
	}
	fmt.Println()
}

// Destroy очищает стек.
func (s *Stack) Destroy() {
	for s.top != nil {
		s.Pop()
	}
}

// LoadFromFile загружает стек из файла.
func (s *Stack) LoadFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Используем временный срез для хранения строк, чтобы потом добавить их в стек в правильном порядке.
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Добавляем элементы в стек в обратном порядке, чтобы вершина стека соответствовала последней строке файла.
	for i := len(lines) - 1; i >= 0; i-- {
		s.Push(lines[i])
	}
	return nil
}

// SaveToFile сохраняет стек в файл.
func (s *Stack) SaveToFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	temp := s.top
	for temp != nil {
		_, err := writer.WriteString(temp.data + "\n")
		if err != nil {
			return err
		}
		temp = temp.next
	}
	return writer.Flush()
}

// runStack выполняет команды над стеком.
func runStack(args []string) {
	stack := &Stack{}
	stack.Init()
	defer stack.Destroy()

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

	if fileName != "" {
		stack.LoadFromFile(fileName)
	}

	parts := strings.SplitN(query, " ", 2)
	command := parts[0]
	var value string
	if len(parts) > 1 {
		value = parts[1]
	}

	switch command {
	case "SPUSH":
		stack.Push(value)
		if fileName != "" {
			stack.SaveToFile(fileName)
		}
	case "SPOP":
		stack.Pop()
		if fileName != "" {
			stack.SaveToFile(fileName)
		}
	case "SPRINT":
		stack.Print()
	}
}