package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	ll       *list.List
}

type cacheEntry struct {
	key   int
	value int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		ll:       list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		return elem.Value.(*cacheEntry).value
	}
	return -1
}

func (c *LRUCache) Set(key, value int) {
	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		elem.Value.(*cacheEntry).value = value
		return
	}

	if c.ll.Len() >= c.capacity {
		// Удаляем самый старый элемент
		oldest := c.ll.Back()
		if oldest != nil {
			c.ll.Remove(oldest)
			delete(c.cache, oldest.Value.(*cacheEntry).key)
		}
	}

	entry := &cacheEntry{key, value}
	elem := c.ll.PushFront(entry)
	c.cache[key] = elem
}

func runLruCache(fullQuery string) {
	var capacity, numQueries int
	// Простой парсинг строки запроса
	parts := strings.Fields(fullQuery)
	for i, p := range parts {
		if p == "=" && i > 0 {
			switch parts[i-1] {
			case "cap":
				capacity, _ = strconv.Atoi(parts[i+1])
			case "Q":
				numQueries, _ = strconv.Atoi(parts[i+1])
			}
		}
	}

	cache := NewLRUCache(capacity)
	var results []string

	queryIndex := -1
	for i, p := range parts {
		if p == "Queries" && parts[i+1] == "=" {
			queryIndex = i + 2
			break
		}
	}
	if queryIndex == -1 {
		fmt.Println("Ошибка: не найдена часть 'Queries =' в запросе")
		return
	}

	queryParams := parts[queryIndex:]

	i := 0
	for q := 0; q < numQueries && i < len(queryParams); q++ {
		op := queryParams[i]
		i++
		if op == "SET" && i+1 < len(queryParams) {
			key, _ := strconv.Atoi(queryParams[i])
			value, _ := strconv.Atoi(queryParams[i+1])
			cache.Set(key, value)
			i += 2
		} else if op == "GET" && i < len(queryParams) {
			key, _ := strconv.Atoi(queryParams[i])
			resultVal := cache.Get(key)
			results = append(results, strconv.Itoa(resultVal))
			i++
		}
	}

	fmt.Println(strings.Join(results, " "))
}