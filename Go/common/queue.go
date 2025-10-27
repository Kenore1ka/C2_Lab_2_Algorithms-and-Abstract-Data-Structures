package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// QueueNode представляет узел очереди.
type QueueNode struct {
	data string
	next *QueueNode
}

// Queue представляет очередь.
type Queue struct {
	front *QueueNode
	rear  *QueueNode
}

// Init инициализирует очередь.
func (q *Queue) Init() {
	q.front = nil
	q.rear = nil
}

// Enqueue добавляет элемент в конец очереди.
func (q *Queue) Enqueue(value string) {
	newNode := &QueueNode{data: value}
	if q.rear != nil {
		q.rear.next = newNode
	}
	q.rear = newNode
	if q.front == nil {
		q.front = q.rear
	}
}

// Dequeue удаляет элемент из начала очереди.
func (q *Queue) Dequeue() {
	if q.front == nil {
		return
	}
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
}

// Print выводит содержимое очереди.
func (q *Queue) Print() {
	temp := q.front
	for temp != nil {
		fmt.Print(temp.data + " ")
		temp = temp.next
	}
	fmt.Println()
}

// Destroy очищает очередь.
func (q *Queue) Destroy() {
	for q.front != nil {
		q.Dequeue()
	}
}

// LoadFromFile загружает очередь из файла.
func (q *Queue) LoadFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		q.Enqueue(scanner.Text())
	}
	return scanner.Err()
}

// SaveToFile сохраняет очередь в файл.
func (q *Queue) SaveToFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	temp := q.front
	for temp != nil {
		_, err := writer.WriteString(temp.data + "\n")
		if err != nil {
			return err
		}
		temp = temp.next
	}
	return writer.Flush()
}

// runQueue выполняет команды над очередью.
func runQueue(args []string) {
	queue := &Queue{}
	queue.Init()
	defer queue.Destroy()

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
		queue.LoadFromFile(fileName)
	}

	parts := strings.SplitN(query, " ", 2)
	command := parts[0]
	var value string
	if len(parts) > 1 {
		value = parts[1]
	}

	switch command {
	case "QPUSH":
		queue.Enqueue(value)
		if fileName != "" {
			queue.SaveToFile(fileName)
		}
	case "QPOP":
		queue.Dequeue()
		if fileName != "" {
			queue.SaveToFile(fileName)
		}
	case "QPRINT":
		queue.Print()
	}
}