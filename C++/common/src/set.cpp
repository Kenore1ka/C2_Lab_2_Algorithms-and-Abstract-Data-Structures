#include "set.h"
#include "hashTable.h" // Подключаем хеш-таблицу
#include <cstring>
#include <fstream>
#include <iostream>
#include <sstream>
#include <string>

// Функция-обработчик для ВСЕХ команд, связанных с множеством
void runSet(int argc, char* argv[]) {
    std::string fileName;
    std::string query;

    // Парсим аргументы командной строки
    for (int i = 1; i < argc; i++) {
        if (strcmp(argv[i], "--file") == 0 && i + 1 < argc) {
            fileName = argv[i + 1];
            i++;
        } else if (strcmp(argv[i], "--query") == 0 && i + 1 < argc) {
            query = argv[i + 1];
            i++;
        }
    }

    if (query.empty()) {
        std::cerr << "Ошибка: аргумент --query не найден." << std::endl;
        return;
    }

    std::stringstream ss(query);
    std::string command, token1, token2;
    ss >> command >> token1 >> token2;

    initTable(); // Инициализируем хеш-таблицу

    if (command == "SETADD") {
        if (fileName.empty() || token1.empty()) {
            std::cerr << "Ошибка: для SETADD требуются --file и значение." << std::endl;
        } else {
            loadSetFromFile(fileName); // ИСПОЛЬЗУЕМ НОВУЮ ФУНКЦИЮ
            insert(token1, token1);
            saveSetToFile(fileName);   // ИСПОЛЬЗУЕМ НОВУЮ ФУНКЦИЮ
        }
    } else if (command == "SETDEL") {
        if (fileName.empty() || token1.empty()) {
            std::cerr << "Ошибка: для SETDEL требуются --file и значение." << std::endl;
        } else {
            loadSetFromFile(fileName); // ИСПОЛЬЗУЕМ НОВУЮ ФУНКЦИЮ
            remove(token1);
            saveSetToFile(fileName);   // ИСПОЛЬЗУЕМ НОВУЮ ФУНКЦИЮ
        }
    } else if (command == "SET_AT") {
        if (fileName.empty() || token1.empty()) {
            std::cerr << "Ошибка: для SET_AT требуются --file и значение." << std::endl;
        } else {
            loadSetFromFile(fileName); // ИСПОЛЬЗУЕМ НОВУЮ ФУНКЦИЮ
            if (get(token1) != "Ключ не найден") {
                std::cout << "true" << std::endl;
            } else {
                std::cout << "false" << std::endl;
            }
        }
    } else if (command == "SET_UNION") {
        if (token1.empty() || token2.empty()) {
            std::cerr << "Ошибка: для SET_UNION требуются два имени файла в запросе." << std::endl;
        } else {
            loadSetFromFile(token1); // Загружаем первое множество
            loadSetFromFile(token2); // Загружаем второе, дубликаты не добавятся
            std::cout << "Объединение множеств (" << token1 << " U " << token2 << "):" << std::endl;
            printTable(); // printTable можно оставить, она для отладки выводит и ключ и значение
        }
    } else if (command == "SET_INTERSECTION") {
        if (token1.empty() || token2.empty()) {
            std::cerr << "Ошибка: для SET_INTERSECTION требуются два имени файла в запросе." << std::endl;
        } else {
            loadSetFromFile(token1); // Загружаем первое множество (A)

            std::ifstream fileB(token2);
            if (!fileB.is_open()) {
                std::cerr << "Ошибка: не удалось открыть файл " << token2 << std::endl;
            } else {
                std::cout << "Пересечение множеств (" << token1 << " n " << token2 << "):" << std::endl;
                std::string key;
                // Читаем файл второго множества
                while (fileB >> key) { 
                    if (get(key) != "Ключ не найден") { // Если элемент из B есть в A
                        std::cout << key << std::endl;
                    }
                }
                fileB.close();
            }
        }
    } else if (command == "SET_DIFFERENCE") {
        if (token1.empty() || token2.empty()) {
            std::cerr << "Ошибка: для SET_DIFFERENCE требуются два имени файла в запросе." << std::endl;
        } else {
            loadSetFromFile(token1); // Загружаем первое множество (A)

            std::ifstream fileB(token2);
            if (!fileB.is_open()) {
                std::cerr << "Ошибка: не удалось открыть файл " << token2 << std::endl;
            } else {
                std::string key;
                // Читаем файл второго множества
                while (fileB >> key) { 
                    remove(key); // Удаляем из A все элементы, которые есть в B
                }
                fileB.close();

                std::cout << "Разность множеств (" << token1 << " - " << token2 << "):" << std::endl;
                printTable();
            }
        }
    } else {
        std::cerr << "Неизвестная команда для множества: " << command << std::endl;
    }

    freeTable(); // Освобождаем память
}