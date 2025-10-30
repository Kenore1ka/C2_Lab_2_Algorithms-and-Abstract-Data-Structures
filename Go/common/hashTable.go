package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const tableSize = 10

// HashNode представляет узел в цепочке (связном списке) для разрешения коллизий.
type HashNode struct {
	Key   string
	Value string
	Next  *HashNode
}

// HashTable - это основная структура хеш-таблицы.
type HashTable struct {
	Table [tableSize]*HashNode
}

// NewHashTable создает и инициализирует новую хеш-таблицу.
func NewHashTable() *HashTable {
	// В Go массив указателей по умолчанию инициализируется значениями nil,
	// так что дополнительная инициализация не требуется.
	return &HashTable{}
}

// hashFunction вычисляет хеш для ключа.
func hashFunction(key string) int {
	var hash int
	for _, char := range key {
		hash += int(char)
	}
	return hash % tableSize
}

// Insert вставляет или обновляет пару ключ-значение.
func (ht *HashTable) Insert(key, value string) {
	index := hashFunction(key)
	newNode := &HashNode{Key: key, Value: value, Next: nil}

	if ht.Table[index] == nil {
		ht.Table[index] = newNode
	} else {
		current := ht.Table[index]
		for {
			// Если ключ уже существует, обновляем значение и выходим.
			if current.Key == key {
				current.Value = value
				return
			}
			// Если дошли до конца цепочки, добавляем новый узел.
			if current.Next == nil {
				current.Next = newNode
				return
			}
			current = current.Next
		}
	}
}

// Get находит значение по ключу. Возвращает значение и true, если найдено,
// иначе пустую строку и false.
func (ht *HashTable) Get(key string) (string, bool) {
	index := hashFunction(key)
	current := ht.Table[index]

	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}
	return "", false
}

// Remove удаляет пару ключ-значение по ключу.
func (ht *HashTable) Remove(key string) {
	index := hashFunction(key)
	current := ht.Table[index]
	var prev *HashNode

	for current != nil {
		if current.Key == key {
			if prev == nil { // Узел является головой списка
				ht.Table[index] = current.Next
			} else { // Узел в середине или в конце
				prev.Next = current.Next
			}
			return
		}
		prev = current
		current = current.Next
	}
}

// Print выводит содержимое всей хеш-таблицы.
func (ht *HashTable) Print() {
	for i := 0; i < tableSize; i++ {
		fmt.Printf("Индекс %d: ", i)
		current := ht.Table[i]
		for current != nil {
			fmt.Printf("[%s: %s] ", current.Key, current.Value)
			current = current.Next
		}
		fmt.Println()
	}
}

// SaveToFile сохраняет хеш-таблицу в файл.
func (ht *HashTable) SaveToFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < tableSize; i++ {
		current := ht.Table[i]
		for current != nil {
			fmt.Fprintf(writer, "%s %s\n", current.Key, current.Value)
			current = current.Next
		}
	}
	return writer.Flush()
}

// LoadFromFile загружает данные из файла в хеш-таблицу.
func (ht *HashTable) LoadFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Если файла нет, это не ошибка, просто таблица будет пустой.
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) == 2 {
			ht.Insert(parts[0], parts[1])
		}
	}
	return scanner.Err()
}

// runHashTable - основная функция для запуска операций с хеш-таблицей.
func runHashTable(args []string) {
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

	ht := NewHashTable()
	if err := ht.LoadFromFile(fileName); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка загрузки файла:", err)
		return
	}

	parts := strings.Fields(query)
	command := parts[0]
	
	switch command {
	case "HSET":
		if len(parts) == 3 {
			ht.Insert(parts[1], parts[2])
			if err := ht.SaveToFile(fileName); err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка сохранения файла:", err)
			}
		}
	case "HGET":
		if len(parts) == 2 {
			if value, found := ht.Get(parts[1]); found {
				fmt.Println(value)
			} else {
				fmt.Println("Ключ не найден")
			}
		}
	case "HDEL":
		if len(parts) == 2 {
			ht.Remove(parts[1])
			if err := ht.SaveToFile(fileName); err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка сохранения файла:", err)
			}
		}
	case "HPRINT":
		ht.Print()
	}
}