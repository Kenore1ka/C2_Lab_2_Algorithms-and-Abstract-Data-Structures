// Задание 7 | Вариант 1
#include "lruCache.h"
#include <iostream>
#include <sstream>
#include <cstring>
#include <vector> // Используем только для парсинга ввода, как требует задание

// --- Вспомогательные функции для внутренней хеш-таблицы ---

int lruHash(int key, int mapSize) {
    return key % mapSize;
}

// --- Реализация методов LRUCache ---

void LRUCache::init(int cap) {
    capacity = cap;
    currentSize = 0;
    head = nullptr;
    tail = nullptr;

    // Инициализируем нашу хеш-таблицу.
    // Размер таблицы можно сделать больше, чем capacity, для уменьшения коллизий.
    mapSize = capacity * 2; 
    map = new MapNode*[mapSize];
    for (int i = 0; i < mapSize; ++i) {
        map[i] = nullptr;
    }
}

void LRUCache::destroy() {
    // Очищаем двусвязный список
    CacheNode* current = head;
    while (current != nullptr) {
        CacheNode* next = current->next;
        delete current;
        current = next;
    }
    head = tail = nullptr;

    // Очищаем хеш-таблицу
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

// Перемещает существующий узел в голову списка
void LRUCache::moveNodeToHead(CacheNode* node) {
    if (node == head) return; // Уже в голове

    // 1. Вырезаем узел из текущей позиции
    if (node->prev) node->prev->next = node->next;
    if (node->next) node->next->prev = node->prev;

    // Если узел был хвостом, обновляем хвост
    if (node == tail) {
        tail = node->prev;
    }

    // 2. Вставляем узел в голову
    node->next = head;
    node->prev = nullptr;
    if (head) {
        head->prev = node;
    }
    head = node;

    // Если список был пуст, хвост - это тоже наш узел
    if (tail == nullptr) {
        tail = head;
    }
}

// Удаляет узел из хвоста (самый старый)
void LRUCache::removeNodeFromTail() {
    if (tail == nullptr) return;

    // --- Удаление из хеш-таблицы ---
    int key_to_remove = tail->key;
    int index = lruHash(key_to_remove, mapSize);
    MapNode* currentMapNode = map[index];
    MapNode* prevMapNode = nullptr;

    while (currentMapNode != nullptr) {
        if (currentMapNode->key == key_to_remove) {
            if (prevMapNode == nullptr) {
                map[index] = currentMapNode->next;
            } else {
                prevMapNode->next = currentMapNode->next;
            }
            delete currentMapNode;
            break; 
        }
        prevMapNode = currentMapNode;
        currentMapNode = currentMapNode->next;
    }

    // --- Удаление из списка ---
    CacheNode* temp = tail;
    if (tail->prev) {
        tail = tail->prev;
        tail->next = nullptr;
    } else {
        // Это был единственный элемент
        head = tail = nullptr;
    }
    delete temp;
    currentSize--;
}


int LRUCache::get(int key) {
    int index = lruHash(key, mapSize);
    MapNode* currentMapNode = map[index];

    // 1. Ищем ключ в хеш-таблице
    while (currentMapNode != nullptr) {
        if (currentMapNode->key == key) {
            // 2. Если нашли, перемещаем узел в начало списка
            moveNodeToHead(currentMapNode->listNode);
            // 3. Возвращаем значение
            return currentMapNode->listNode->value;
        }
        currentMapNode = currentMapNode->next;
    }
    
    // 4. Если не нашли
    return -1;
}

void LRUCache::set(int key, int value) {
    int index = lruHash(key, mapSize);
    MapNode* currentMapNode = map[index];

    // 1. Ищем, есть ли уже такой ключ
    while (currentMapNode != nullptr) {
        if (currentMapNode->key == key) {
            // Если нашли, обновляем значение и перемещаем в голову
            currentMapNode->listNode->value = value;
            moveNodeToHead(currentMapNode->listNode);
            return;
        }
        currentMapNode = currentMapNode->next;
    }

    // 2. Если ключа нет - это новый элемент
    // Проверяем, не переполнен ли кэш
    if (currentSize >= capacity) {
        removeNodeFromTail();
    }

    // 3. Создаем новый узел для списка
    CacheNode* newListNode = new CacheNode{key, value, nullptr, head};
    if (head) {
        head->prev = newListNode;
    }
    head = newListNode;
    if (tail == nullptr) {
        tail = head;
    }
    currentSize++;

    // 4. Создаем новый узел для хеш-таблицы и добавляем его
    MapNode* newMapNode = new MapNode{key, newListNode, map[index]};
    map[index] = newMapNode;
}


// --- Функция-запускальщик ---

void runLruCache(int argc, char* argv[]) {
    if (argc < 2) {
        std::cout << "Пример использования: ./program <входная строка>" << std::endl;
        return;
    }

    std::string full_query = argv[1];
    std::stringstream ss(full_query);
    
    std::string token;
    int capacity = 0;
    int num_queries = 0;
    
    // Парсим cap и Q
    ss >> token; // "cap"
    ss >> token; // "="
    ss >> capacity;
    ss >> token; // ","
    ss >> token; // "Q"
    ss >> token; // "="
    ss >> num_queries;
    ss >> token; // "Queries"
    ss >> token; // "="

    LRUCache cache;
    cache.init(capacity);
    
    std::vector<int> results;

    for(int i = 0; i < num_queries; ++i) {
        ss >> token;
        if (token == "SET") {
            int key, value;
            ss >> key >> value;
            cache.set(key, value);
        } else if (token == "GET") {
            int key;
            ss >> key;
            results.push_back(cache.get(key));
        }
    }
    
    // Вывод результатов
    for(size_t i = 0; i < results.size(); ++i) {
        std::cout << results[i] << (i == results.size() - 1 ? "" : " ");
    }
    std::cout << std::endl;
    
    cache.destroy();
}