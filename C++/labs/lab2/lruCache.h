#pragma once

#include <string>

// --- Новые структуры данных, специфичные для LRU кэша ---

// Узел, который будет храниться в двусвязном списке.
// Он содержит и ключ, и значение.
struct CacheNode {
    int key;
    int value;
    CacheNode* prev;
    CacheNode* next;
};

// Узел для нашей собственной хеш-таблицы.
// Хранит ключ и указатель на узел в двусвязном списке.
struct MapNode {
    int key;
    CacheNode* listNode; // Указатель на узел в списке
    MapNode* next;       // Для разрешения коллизий методом цепочек
};

// --- Основная структура LRU кэша ---
struct LRUCache {
    int capacity;
    int currentSize;

    // Двусвязный список для хранения данных и порядка использования
    CacheNode* head;
    CacheNode* tail;

    // Хеш-таблица для быстрого доступа к узлам списка
    MapNode** map;
    int mapSize;

    // Функции для работы с кэшем
    void init(int cap);
    void destroy();
    int get(int key);
    void set(int key, int value);

    // Вспомогательные функции для работы со списком
    void moveNodeToHead(CacheNode* node);
    void removeNodeFromTail();
};

// Функция для запуска задания из командной строки
void runLruCache(int argc, char* argv[]);