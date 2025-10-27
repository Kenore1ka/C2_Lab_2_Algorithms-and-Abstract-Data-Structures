package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ListNode представляет узел односвязного списка.
type ListNode struct {
	data string
	next *ListNode
}

// LinkedList представляет односвязный список.
type LinkedList struct {
	head *ListNode
	tail *ListNode
}

// Init инициализирует список.
func (ll *LinkedList) Init() {
	ll.head = nil
	ll.tail = nil
}

// AddToHead добавляет элемент в начало списка.
func (ll *LinkedList) AddToHead(value string) {
	newNode := &ListNode{data: value, next: ll.head}
	ll.head = newNode
	if ll.tail == nil {
		ll.tail = ll.head
	}
}

// AddToTail добавляет элемент в конец списка.
func (ll *LinkedList) AddToTail(value string) {
	newNode := &ListNode{data: value}
	if ll.tail != nil {
		ll.tail.next = newNode
	}
	ll.tail = newNode
	if ll.head == nil {
		ll.head = ll.tail
	}
}

// RemoveFromHead удаляет элемент с головы списка.
func (ll *LinkedList) RemoveFromHead() {
	if ll.head == nil {
		return
	}
	ll.head = ll.head.next
	if ll.head == nil {
		ll.tail = nil
	}
}

// RemoveFromTail удаляет элемент с хвоста списка.
func (ll *LinkedList) RemoveFromTail() {
	if ll.tail == nil {
		return
	}
	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
		return
	}
	temp := ll.head
	for temp.next != ll.tail {
		temp = temp.next
	}
	ll.tail = temp
	ll.tail.next = nil
}

// RemoveByValue удаляет узел по значению.
func (ll *LinkedList) RemoveByValue(value string) {
	if ll.head == nil {
		return
	}
	if ll.head.data == value {
		ll.RemoveFromHead()
		return
	}
	temp := ll.head
	for temp.next != nil {
		if temp.next.data == value {
			if temp.next == ll.tail {
				ll.tail = temp
			}
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}
}

// Search ищет элемент по значению.
func (ll *LinkedList) Search(value string) bool {
	temp := ll.head
	for temp != nil {
		if temp.data == value {
			return true
		}
		temp = temp.next
	}
	return false
}

// Print выводит все элементы списка.
func (ll *LinkedList) Print() {
	temp := ll.head
	for temp != nil {
		fmt.Print(temp.data + " ")
		temp = temp.next
	}
	fmt.Println()
}

// Destroy очищает список.
func (ll *LinkedList) Destroy() {
	for ll.head != nil {
		ll.RemoveFromHead()
	}
}

// LoadFromFile загружает элементы из файла.
func (ll *LinkedList) LoadFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ll.AddToTail(scanner.Text())
	}
	return scanner.Err()
}

// SaveToFile сохраняет элементы списка в файл.
func (ll *LinkedList) SaveToFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	temp := ll.head
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
func (ll *LinkedList) AddBefore(target, value string) {
	if ll.head == nil {
		return
	}
	if ll.head.data == target {
		ll.AddToHead(value)
		return
	}
	prev := ll.head
	for prev.next != nil && prev.next.data != target {
		prev = prev.next
	}
	if prev.next == nil {
		return
	}
	newNode := &ListNode{data: value, next: prev.next}
	prev.next = newNode
}

// AddAfter вставляет элемент после указанного.
func (ll *LinkedList) AddAfter(target, value string) {
	node := ll.head
	for node != nil && node.data != target {
		node = node.next
	}
	if node == nil {
		return
	}
	newNode := &ListNode{data: value, next: node.next}
	node.next = newNode
	if node == ll.tail {
		ll.tail = newNode
	}
}

// RemoveBefore удаляет элемент перед указанным.
func (ll *LinkedList) RemoveBefore(target string) {
	if ll.head == nil || ll.head.data == target {
		return
	}
	if ll.head.next != nil && ll.head.next.data == target {
		ll.RemoveFromHead()
		return
	}
	prev := ll.head
	for prev.next != nil && prev.next.next != nil && prev.next.next.data != target {
		prev = prev.next
	}
	if prev.next == nil || prev.next.next == nil {
		return
	}
	nodeToRemove := prev.next
	prev.next = nodeToRemove.next
	if nodeToRemove == ll.tail {
		ll.tail = prev
	}
}

// RemoveAfter удаляет элемент после указанного.
func (ll *LinkedList) RemoveAfter(target string) {
	node := ll.head
	for node != nil && node.data != target {
		node = node.next
	}
	if node == nil || node.next == nil {
		return
	}
	nodeToRemove := node.next
	node.next = nodeToRemove.next
	if nodeToRemove == ll.tail {
		ll.tail = node
	}
}

// runLinkedList выполняет команды над односвязным списком.
func runLinkedList(args []string) {
	list := &LinkedList{}
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
	case "LPUSH":
		if token1 != "" {
			list.AddToHead(token1)
		}
	case "LAPPEND":
		if token1 != "" {
			list.AddToTail(token1)
		}
	case "LREMOVEHEAD":
		list.RemoveFromHead()
	case "LREMOVETAIL":
		list.RemoveFromTail()
	case "LREMOVE":
		if token1 != "" {
			list.RemoveByValue(token1)
		}
	case "LSEARCH":
		if token1 != "" {
			fmt.Println(list.Search(token1))
		}
	case "LPRINT":
		list.Print()
	case "LADDTO":
		if token1 != "" && token2 != "" {
			list.AddBefore(token1, token2)
		}
	case "LADDAFTER":
		if token1 != "" && token2 != "" {
			list.AddAfter(token1, token2)
		}
	case "LREMOVETO":
		if token1 != "" {
			list.RemoveBefore(token1)
		}
	case "LREMOVEAFTER":
		if token1 != "" {
			list.RemoveAfter(token1)
		}
	}

	if fileName != "" {
		list.SaveToFile(fileName)
	}
}