// Задание 4 | Вариант 3

#include "patternMatcher.h"
#include "array.h" 

#include <iostream>
#include <string>
#include <cstring>
#include <sstream>

bool match(const std::string& text, const std::string& pattern) {
    int i = 0, j = 0;
    int star_idx = -1, match_idx = 0;
    int n = text.length();
    int m = pattern.length();

    while (i < n) {
        if (j < m && (pattern[j] == '?' || pattern[j] == text[i])) {
            i++; j++;
        } else if (j < m && pattern[j] == '*') {
            star_idx = j;
            match_idx = i;
            j++;
        } else if (star_idx != -1) {
            j = star_idx + 1;
            match_idx++;
            i = match_idx;
        } else {
            return false;
        }
    }

    while (j < m && pattern[j] == '*') {
        j++;
    }

    return j == m;
}

/**
 * @brief Основная функция. Загружает строки из файла в DynamicArray и 
 *        проверяет каждую на соответствие шаблону.
 */
void runPatternMatcher(int argc, char* argv[]) {
    std::string query;

    for (int i = 1; i < argc; i++) {
        if (strcmp(argv[i], "--query") == 0 && i + 1 < argc) {
            query = argv[i + 1];
            i++;
            break;
        }
    }

    if (query.empty()) {
        std::cerr << "Ошибка: не найден аргумент --query." << std::endl;
        std::cerr << "Пример: --query \"MATCH <файл_с_данными> <шаблон>\"" << std::endl;
        return;
    }
    
    std::stringstream ss(query);
    std::string command, dataFileName, pattern;
    ss >> command >> dataFileName >> pattern;

    if (command != "MATCH" || dataFileName.empty() || pattern.empty()) {
        std::cerr << "Ошибка: некорректный формат запроса." << std::endl;
        std::cerr << "Пример: --query \"MATCH data.txt h*o\"" << std::endl;
        return;
    }

    // --- ИСПОЛЬЗОВАНИЕ DYNAMICARRAY ---

    // 1. Создаем экземпляр нашего динамического массива
    DynamicArray stringList;

    // 2. Загружаем в него все строки из указанного файла
    stringList.loadFromFile(dataFileName);

    if (stringList.length() == 0) {
        std::cout << "Файл \"" << dataFileName << "\" пуст или не найден." << std::endl;
        return;
    }

    std::cout << "Строки из файла \"" << dataFileName << "\", соответствующие шаблону \"" << pattern << "\":" << std::endl;
    
    bool foundMatches = false;
    // 3. Проходим по всем элементам массива
    for (int i = 0; i < stringList.length(); ++i) {
        std::string currentString = stringList.get(i);
        
        // 4. Вызываем нашу вспомогательную функцию match для каждой строки
        if (match(currentString, pattern)) {
            std::cout << currentString << std::endl;
            foundMatches = true;
        }
    }

    if (!foundMatches) {
        std::cout << "Совпадений не найдено." << std::endl;
    }
}