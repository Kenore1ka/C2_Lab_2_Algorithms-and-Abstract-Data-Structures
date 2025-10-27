#include "set.h"
#include "hashTable.h" // Включаем, чтобы использовать функции insert, remove, get и др.

#include <iostream>
#include <fstream>
#include <string>
#include <cstring>

// Функция-обработчик для команд, связанных с множеством
void runSet(int argc, char* argv[]) {
    // 1. Инициализируем глобальную хеш-таблицу
    initTable();

    std::string fileName;
    std::string query;

    // 2. Парсим аргументы командной строки
    for (int i = 1; i < argc; i++) {
        if (strcmp(argv[i], "--file") == 0 && i + 1 < argc) {
            fileName = argv[i + 1];
            i++;
        } else if (strcmp(argv[i], "--query") == 0 && i + 1 < argc) {
            query = argv[i + 1];
            i++;
        }
    }

    // 3. Загружаем данные в хеш-таблицу (которая представляет наше множество)
    if (!fileName.empty()) {
        loadFromFile(fileName);
    }

    // 4. Разбираем команду и значение
    std::string command;
    std::string value;
    size_t pos = query.find(' ');
    if (pos != std::string::npos) {
        command = query.substr(0, pos);
        value = query.substr(pos + 1);
    } else {
        std::cerr << "Ошибка: Некорректный формат запроса." << std::endl;
        freeTable();
        return;
    }

    // 5. Выполняем команду
    if (command == "SETADD") {
        // Для множества ключ и значение одинаковы
        insert(value, value);
        if (!fileName.empty()) {
            saveToFile(fileName);
        }
    } else if (command == "SETDEL") {
        remove(value);
        if (!fileName.empty()) {
            saveToFile(fileName);
        }
    } else if (command == "SET_AT") {
        // Используем функцию get, чтобы проверить наличие
        if (get(value) != "Ключ не найден") {
            std::cout << "true" << std::endl;
        } else {
            std::cout << "false" << std::endl;
        }
    } else {
        std::cout << "Неизвестная команда для множества: " << command << std::endl;
    }

    // 6. Освобождаем память, занятую хеш-таблицей
    freeTable();
}