package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// DynamicArray представляет динамический массив строк.
type DynamicArray struct {
	data     []string
	size     int
	capacity int
}

// NewDynamicArray создает и возвращает новый DynamicArray.
func NewDynamicArray(initialCapacity int) *DynamicArray {
	if initialCapacity <= 0 {
		initialCapacity = 10
	}
	return &DynamicArray{
		data:     make([]string, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
}

// resize изменяет емкость массива.
func (da *DynamicArray) resize(newCapacity int) {
	newData := make([]string, newCapacity)
	copy(newData, da.data[:da.size])
	da.data = newData
	da.capacity = newCapacity
}

// Add добавляет новый элемент в конец массива.
func (da *DynamicArray) Add(value string) {
	if da.size == da.capacity {
		da.resize(da.capacity * 2)
	}
	da.data[da.size] = value
	da.size++
}

// Insert вставляет элемент в заданную позицию.
func (da *DynamicArray) Insert(index int, value string) {
	if index < 0 || index > da.size {
		return
	}
	if da.size == da.capacity {
		da.resize(da.capacity * 2)
	}
	copy(da.data[index+1:], da.data[index:da.size])
	da.data[index] = value
	da.size++
}

// Remove удаляет элемент из массива по индексу.
func (da *DynamicArray) Remove(index int) {
	if index < 0 || index >= da.size {
		return
	}
	copy(da.data[index:], da.data[index+1:da.size])
	da.size--
}

// Get возвращает элемент по индексу.
func (da *DynamicArray) Get(index int) string {
	if index < 0 || index >= da.size {
		return ""
	}
	return da.data[index]
}

// Set устанавливает значение элемента по индексу.
func (da *DynamicArray) Set(index int, value string) {
	if index < 0 || index >= da.size {
		return
	}
	da.data[index] = value
}

// Length возвращает текущий размер массива.
func (da *DynamicArray) Length() int {
	return da.size
}

// Print выводит все элементы массива.
func (da *DynamicArray) Print() {
	for i := 0; i < da.size; i++ {
		fmt.Print(da.data[i] + " ")
	}
	fmt.Println()
}

// Clear очищает массив.
func (da *DynamicArray) Clear() {
	da.capacity = 10
	da.size = 0
	da.data = make([]string, da.capacity)
}

// LoadFromFile загружает данные из файла в массив.
func (da *DynamicArray) LoadFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		da.Add(scanner.Text())
	}
	return scanner.Err()
}

// SaveToFile сохраняет содержимое массива в файл.
func (da *DynamicArray) SaveToFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < da.size; i++ {
		_, err := writer.WriteString(da.data[i] + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// runDynamicArray выполняет команды над динамическим массивом.
func runDynamicArray(args []string) {
	arr := NewDynamicArray(10)

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
		arr.LoadFromFile(fileName)
	}

	parts := strings.SplitN(query, " ", 2)
	command := parts[0]
	var params string
	if len(parts) > 1 {
		params = parts[1]
	}

	switch command {
	case "MPUSH":
		arr.Add(params)
		arr.SaveToFile(fileName)
	case "MINSERT":
		parts := strings.SplitN(params, " ", 2)
		index, _ := strconv.Atoi(parts[0])
		value := parts[1]
		arr.Insert(index, value)
		arr.SaveToFile(fileName)
	case "MDEL":
		index, _ := strconv.Atoi(params)
		arr.Remove(index)
		arr.SaveToFile(fileName)
	case "MSET":
		parts := strings.SplitN(params, " ", 2)
		index, _ := strconv.Atoi(parts[0])
		value := parts[1]
		arr.Set(index, value)
		arr.SaveToFile(fileName)
	case "MLEN":
		fmt.Println(arr.Length())
	case "MPRINT":
		arr.Print()
	case "MGET":
		index, _ := strconv.Atoi(params)
		fmt.Println(arr.Get(index))
	}
}