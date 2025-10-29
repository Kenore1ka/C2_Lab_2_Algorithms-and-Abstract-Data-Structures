#include "lruCache.h"

#include <cstring>
#include <iostream>
#include <sstream>

#include "array.h"

// Вспомогательные функции для внутренней хеш-таблицы
int lruHash(int key, int mapSize) { return key % mapSize; }

// Реализация методов LRUCache
void LRUCache::init(int cap) {
    capacity = cap;
    currentSize = 0;
    head = nullptr;
    tail = nullptr;
    mapSize = (capacity == 0) ? 1 : capacity * 2;
    map = new MapNode*[mapSize];
    for (int i = 0; i < mapSize; ++i) {
        map[i] = nullptr;
    }
}

void LRUCache::destroy() {
    CacheNode* current = head;
    while (current != nullptr) {
        CacheNode* next = current->next;
        delete current;
        current = next;
    }
    head = tail = nullptr;

    for (int i = 0; i < mapSize; ++i) {
        MapNode* currentMapNode = map[i];
        while (currentMapNode != nullptr) {
            MapNode* next = currentMapNode->next;
            delete currentMapNode;
            currentMapNode = next;
        }
    }
    delete[] map;
}

void LRUCache::moveNodeToHead(CacheNode* node) {
    if (node == head) return;
    if (node->prev) node->prev->next = node->next;
    if (node->next) node->next->prev = node->prev;
    if (node == tail) tail = node->prev;
    node->next = head;
    node->prev = nullptr;
    if (head) head->prev = node;
    head = node;
    if (tail == nullptr) tail = head;
}

void LRUCache::removeNodeFromTail() {
    if (tail == nullptr) return;

    int key_to_remove = tail->key;
    int index = lruHash(key_to_remove, mapSize);
    MapNode* currentMapNode = map[index];
    MapNode* prevMapNode = nullptr;

    while (currentMapNode != nullptr) {
        if (currentMapNode->key == key_to_remove) {
            if (prevMapNode == nullptr)
                map[index] = currentMapNode->next;
            else
                prevMapNode->next = currentMapNode->next;
            delete currentMapNode;
            break;
        }
        prevMapNode = currentMapNode;
        currentMapNode = currentMapNode->next;
    }

    CacheNode* temp = tail;
    if (tail->prev) {
        tail = tail->prev;
        tail->next = nullptr;
    } else {
        head = tail = nullptr;
    }
    delete temp;
    currentSize--;
}

int LRUCache::get(int key) {
    int index = lruHash(key, mapSize);
    MapNode* currentMapNode = map[index];
    while (currentMapNode != nullptr) {
        if (currentMapNode->key == key) {
            moveNodeToHead(currentMapNode->listNode);
            return currentMapNode->listNode->value;
        }
        currentMapNode = currentMapNode->next;
    }
    return -1;
}

void LRUCache::set(int key, int value) {
    int index = lruHash(key, mapSize);
    MapNode* currentMapNode = map[index];
    while (currentMapNode != nullptr) {
        if (currentMapNode->key == key) {
            currentMapNode->listNode->value = value;
            moveNodeToHead(currentMapNode->listNode);
            return;
        }
        currentMapNode = currentMapNode->next;
    }
    if (currentSize >= capacity && capacity > 0) {
        removeNodeFromTail();
    }
    if (capacity > 0) {
        CacheNode* newListNode = new CacheNode{key, value, nullptr, head};
        if (head) head->prev = newListNode;
        head = newListNode;
        if (tail == nullptr) tail = head;
        currentSize++;
        MapNode* newMapNode = new MapNode{key, newListNode, map[index]};
        map[index] = newMapNode;
    }
}

void runLruCache(int argc, char* argv[]) {
    if (argc < 2) {
        std::cout << "Пример использования: ./program 5 \"cap = 2, Q = 2 ...\"" << std::endl;
        return;
    }

    std::string full_query = argv[1];
    std::stringstream ss(full_query);

    std::string token;
    int capacity = 0;
    int num_queries = 0;

    ss >> token;
    ss >> token;
    ss >> capacity;
    ss >> token;
    ss >> token;
    ss >> token;
    ss >> num_queries;
    ss >> token;
    ss >> token;

    LRUCache cache;
    cache.init(capacity);

    DynamicArray results;

    for (int i = 0; i < num_queries; ++i) {
        ss >> token;
        if (token == "SET") {
            int key, value;
            ss >> key >> value;
            cache.set(key, value);
        } else if (token == "GET") {
            int key;
            ss >> key;
            int result_val = cache.get(key);
            results.add(std::to_string(result_val));
        }
    }

    // Вывод результатов из DynamicArray
    for (int i = 0; i < results.length(); ++i) {
        std::cout << results.get(i) << (i == results.length() - 1 ? "" : " ");
    }
    std::cout << std::endl;

    cache.destroy();
}