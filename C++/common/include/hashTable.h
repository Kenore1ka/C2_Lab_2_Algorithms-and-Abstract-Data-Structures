#pragma once

#include <string>

// Объявления всех функций, которые могут понадобиться другим модулям
void initTable();
void insert(const std::string& key, const std::string& value);
std::string get(const std::string& key);
void remove(const std::string& key);
void printTable();
void saveToFile(const std::string& fileName);
void loadFromFile(const std::string& fileName);
void freeTable();
void saveSetToFile(const std::string& fileName);
void loadSetFromFile(const std::string& fileName);

// Основная функция для запуска операций с хеш-таблицей из командной строки
void runHashTable(int argc, char* argv[]);