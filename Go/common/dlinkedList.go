package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// DListNode представляет узел двусвязного списка.
type DListNode struct {
	data string
	next *DListNode
	prev *DListNode
}

// DLinkedList представляет двусвязный список.
type DLinkedList struct {
	head *DListNode
	tail *DListNode
}

// Init инициализирует список.
func (dll *DLinkedList) Init() {
	dll.head = nil
	dll.tail = nil
}

// AddToHead добавляет элемент в начало списка.
func (dll *DLinkedList) AddToHead(value string) {
	newNode := &DListNode{data: value, next: dll.head}
	if dll.head != nil {
		dll.head.prev = newNode
	}
	dll.head = newNode
	if dll.tail == nil {
		dll.tail = dll.head
	}
}

// AddToTail добавляет элемент в конец списка.
func (dll *DLinkedList) AddToTail(value string) {
	newNode := &DListNode{data: value, prev: dll.tail}
	if dll.tail != nil {
		dll.tail.next = newNode
	}
	dll.tail = newNode
	if dll.head == nil {
		dll.head = dll.tail
	}
}

// RemoveFromHead удаляет элемент с головы списка.
func (dll *DLinkedList) RemoveFromHead() {
	if dll.head == nil {
		return
	}
	dll.head = dll.head.next
	if dll.head != nil {
		dll.head.prev = nil
	} else {
		dll.tail = nil
	}
}

// RemoveFromTail удаляет элемент с хвоста списка.
func (dll *DLinkedList) RemoveFromTail() {
	if dll.tail == nil {
		return
	}
	dll.tail = dll.tail.prev
	if dll.tail != nil {
		dll.tail.next = nil
	} else {
		dll.head = nil
	}
}

// RemoveByValue удаляет узел по значению.
func (dll *DLinkedList) RemoveByValue(value string) {
	temp := dll.head
	for temp != nil {
		if temp.data == value {
			if temp.prev != nil {
				temp.prev.next = temp.next
			} else {
				dll.head = temp.next
			}
			if temp.next != nil {
				temp.next.prev = temp.prev
			} else {
				dll.tail = temp.prev
			}
			return
		}
		temp = temp.next
	}
}

// Search ищет элемент по значению.
func (dll *DLinkedList) Search(value string) bool {
	temp := dll.head
	for temp != nil {
		if temp.data == value {
			return true
		}
		temp = temp.next
	}
	return false
}

// Print выводит все элементы списка.
func (dll *DLinkedList) Print() {
	temp := dll.head
	for temp != nil {
		fmt.Print(temp.data + " ")
		temp = temp.next
	}
	fmt.Println()
}

// Destroy очищает список.
func (dll *DLinkedList) Destroy() {
	for dll.head != nil {
		dll.RemoveFromHead()
	}
}

// LoadFromFile загружает элементы из файла.
func (dll *DLinkedList) LoadFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dll.AddToTail(scanner.Text())
	}
	return scanner.Err()
}

// SaveToFile сохраняет элементы списка в файл.
func (dll *DLinkedList) SaveToFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	temp := dll.head
	for temp != nil {
		_, err := writer.WriteString(temp.data + "\n")
		if err != nil {
			return err
		}
		temp = temp.next
	}
	return writer.Flush()
}

// AddBefore вставляет элемент перед указанным.
func (dll *DLinkedList) AddBefore(target, value string) {
	if dll.head == nil {
		return
	}
	if dll.head.data == target {
		dll.AddToHead(value)
		return
	}
	node := dll.head.next
	for node != nil && node.data != target {
		node = node.next
	}
	if node == nil {
		return
	}
	newNode := &DListNode{data: value, next: node, prev: node.prev}
	if node.prev != nil {
		node.prev.next = newNode
	}
	node.prev = newNode
}

// AddAfter вставляет элемент после указанного.
func (dll *DLinkedList) AddAfter(target, value string) {
	node := dll.head
	for node != nil && node.data != target {
		node = node.next
	}
	if node == nil {
		return
	}
	newNode := &DListNode{data: value, next: node.next, prev: node}
	if node.next != nil {
		node.next.prev = newNode
	} else {
		dll.tail = newNode
	}
	node.next = newNode
}

// RemoveBefore удаляет элемент перед указанным.
func (dll *DLinkedList) RemoveBefore(target string) {
	if dll.head == nil || dll.head.data == target {
		return
	}
	if dll.head.next != nil && dll.head.next.data == target {
		dll.RemoveFromHead()
		return
	}
	node := dll.head.next
	for node != nil && node.data != target {
		node = node.next
	}
	if node == nil || node.prev == nil {
		return
	}
	toRemove := node.prev
	before := toRemove.prev
	node.prev = before
	if before != nil {
		before.next = node
	} else {
		dll.head = node
	}
}

// RemoveAfter удаляет элемент после указанного.
func (dll *DLinkedList) RemoveAfter(target string) {
	node := dll.head
	for node != nil && node.data != target {
		node = node.next
	}
	if node == nil || node.next == nil {
		return
	}
	toRemove := node.next
	after := toRemove.next
	node.next = after
	if after != nil {
		after.prev = node
	} else {
		dll.tail = node
	}
}

// runDLinkedList выполняет команды над двусвязным списком.
func runDLinkedList(args []string) {
	list := &DLinkedList{}
	list.Init()
	defer list.Destroy()

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
		list.LoadFromFile(fileName)
	}

	parts := strings.Fields(query)
	command := parts[0]
	var token1, token2 string
	if len(parts) > 1 {
		token1 = parts[1]
	}
	if len(parts) > 2 {
		token2 = parts[2]
	}

	switch command {
	case "DPUSH":
		if token1 != "" {
			list.AddToHead(token1)
		}
	case "DAPPEND":
		if token1 != "" {
			list.AddToTail(token1)
		}
	case "DREMOVEHEAD":
		list.RemoveFromHead()
	case "DREMOVETAIL":
		list.RemoveFromTail()
	case "DREMOVE":
		if token1 != "" {
			list.RemoveByValue(token1)
		}
	case "DSEARCH":
		if token1 != "" {
			fmt.Println(list.Search(token1))
		}
	case "DPRINT":
		list.Print()
	case "DADDTO":
		if token1 != "" && token2 != "" {
			list.AddBefore(token1, token2)
		}
	case "DADDAFTER":
		if token1 != "" && token2 != "" {
			list.AddAfter(token1, token2)
		}
	case "DREMOVETO":
		if token1 != "" {
			list.RemoveBefore(token1)
		}
	case "DREMOVEAFTER":
		if token1 != "" {
			list.RemoveAfter(token1)
		}
	}

	if fileName != "" {
		list.SaveToFile(fileName)
	}
}