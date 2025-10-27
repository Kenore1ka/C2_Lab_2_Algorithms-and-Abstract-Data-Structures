#include "hashTable.h"

#include <cstring>
#include <fstream>
#include <iostream>
#include <string>

using namespace std;

// Определение структуры узла хеш-таблицы
struct HashNode {
    string key;       // Ключ для хеширования
    string value;     // Значение, связанное с ключом
    HashNode* next;   // Указатель на следующий узел в случае коллизии
};

const int TABLE_SIZE = 10; // Размер хеш-таблицы

// Инициализация хеш-таблицы, представляющей собой массив указателей на HashNode
HashNode* hashTable[TABLE_SIZE];

// Функция инициализации хеш-таблицы
void initTable() {
    for (int i = 0; i < TABLE_SIZE; i++) {
        hashTable[i] = nullptr;
    }
}

// Простая хеш-функция, возвращающая индекс на основе суммы ASCII кодов символов ключа
int hashFunction(const string& key) {
    int hash = 0;
    for (char ch : key) {
        hash += ch;
    }
    return hash % TABLE_SIZE;
}

// Вставка нового ключа и значения в хеш-таблицу
void insert(const string& key, const string& value) {
    int index = hashFunction(key);
    HashNode* newNode = new HashNode{key, value, nullptr};

    if (hashTable[index] == nullptr) { // Если ячейка пуста, вставляем новый узел
        hashTable[index] = newNode;
    } else { // Если есть коллизия, проходим по цепочке и вставляем узел в конец
        HashNode* current = hashTable[index];
        while (current->next != nullptr) {
            if (current->key == key) { // Если ключ уже существует, обновляем значение
                current->value = value;
                delete newNode;
                return;
            }
            current = current->next;
        }
        if (current->key == key) { // Обновляем значение, если ключ совпадает
            current->value = value;
            delete newNode;
            return;
        }
        current->next = newNode; // Вставляем новый узел в конец цепочки
    }
}

// Получение значения по ключу из хеш-таблицы
string get(const string& key) {
    int index = hashFunction(key);
    HashNode* current = hashTable[index];

    while (current != nullptr) { // Проходим по цепочке до совпадения ключа
        if (current->key == key) {
            return current->value;
        }
        current = current->next;
    }
    return "Ключ не найден"; // Возвращаем сообщение, если ключ не найден
}

// Удаление узла по ключу из хеш-таблицы
void remove(const string& key) {
    int index = hashFunction(key);
    HashNode* current = hashTable[index];
    HashNode* prev = nullptr;

    while (current != nullptr) {
        if (current->key == key) {
            if (prev == nullptr) { // Если удаляемый узел первый в цепочке
                hashTable[index] = current->next;
            } else { // Если узел в середине или конце цепочки
                prev->next = current->next;
            }
            delete current;
            return;
        }
        prev = current;
        current = current->next;
    }
}

// Вывод всех элементов хеш-таблицы
void printTable() {
    for (int i = 0; i < TABLE_SIZE; i++) {
        cout << "Индекс " << i << ": ";
        HashNode* current = hashTable[i];
        while (current != nullptr) { // Выводим всю цепочку для данного индекса
            cout << "[" << current->key << ": " << current->value << "] ";
            current = current->next;
        }
        cout << endl;
    }
}

// Сохранение хеш-таблицы в файл
void saveToFile(const string& fileName) {
    ofstream file(fileName);
    for (int i = 0; i < TABLE_SIZE; i++) {
        HashNode* current = hashTable[i];
        while (current != nullptr) { // Записываем каждый узел в файл
            file << current->key << " " << current->value << endl;
            current = current->next;
        }
    }
    file.close();
}

// Загрузка хеш-таблицы из файла
void loadFromFile(const string& fileName) {
    ifstream file(fileName);
    string key, value;
    while (file >> key >> value) { // Читаем файл и вставляем ключи и значения в таблицу
        insert(key, value);
    }
    file.close();
}

// Освобождение памяти, занятой хеш-таблицей
void freeTable() {
    for (int i = 0; i < TABLE_SIZE; i++) {
        HashNode* current = hashTable[i];
        while (current != nullptr) { // Освобождаем каждый узел цепочки
            HashNode* temp = current;
            current = current->next;
            delete temp;
        }
    }
}

// Функция для запуска хеш-таблицы с параметрами командной строки
void runHashTable(int argc, char* argv[]) {
    initTable();

    string fileName;
    string query;

    for (int i = 1; i < argc; i++) {
        if (strcmp(argv[i], "--file") == 0 && i + 1 < argc) { // Обработка аргумента для имени файла
            fileName = argv[i + 1];
            i++;
        } else if (strcmp(argv[i], "--query") == 0 && i + 1 < argc) { // Обработка запроса
            query = argv[i + 1];
            i++;
        }
    }

    loadFromFile(fileName);

    string command;
    size_t pos = query.find(' ');
    if (pos != string::npos) {
        command = query.substr(0, pos);
        query = query.substr(pos + 1);
    } else {
        command = query;
    }

    if (command == "HSET") { // Вставка или обновление значения по ключу
        size_t pos = query.find(' ');
        string key = query.substr(0, pos);
        string value = query.substr(pos + 1);
        insert(key, value);
        saveToFile(fileName);
    } else if (command == "HGET") { // Получение значения по ключу
        cout << get(query) << endl;
    } else if (command == "HDEL") { // Удаление узла по ключу
        remove(query);
        saveToFile(fileName);
    } else if (command == "HPRINT") { // Вывод всех значений хеш-таблицы
        printTable();
    }

    freeTable(); // Освобождение памяти
}
