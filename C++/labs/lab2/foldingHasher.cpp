// Задание 6 | Вариант 1
#include "foldingHasher.h"

#include <algorithm> 
#include <cstring>
#include <iostream>
#include <sstream> 
#include <string>

long long foldingHash(const std::string& key, int chunkSize) {
    long long totalSum = 0;

    std::cout << "Вычисление: ";
    bool isFirstPart = true;

    // Идем по строке с шагом, равным размеру части
    for (size_t i = 0; i < key.length(); i += chunkSize) {
        // Вырезаем очередную часть из строки
        std::string partStr = key.substr(i, chunkSize);

        // Конвертируем строку в число. Используем stoll для long long.
        long long partValue = std::stoll(partStr);

        // Добавляем к общей сумме
        totalSum += partValue;

        // Выводим процесс вычисления, как в примере
        if (!isFirstPart) {
            std::cout << "+";
        }
        std::cout << partStr;
        isFirstPart = false;
    }

    // Выводим итоговый результат
    std::cout << " = " << totalSum << std::endl;

    return totalSum;
}

void runFoldingHasher(int argc, char* argv[]) {
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
        std::cerr << "Пример использования: --query \"FOLD <число>\"" << std::endl;
        return;
    }

    std::stringstream ss(query);
    std::string command, numberKey;
    ss >> command >> numberKey;

    if (command != "FOLD" || numberKey.empty()) {
        std::cerr << "Ошибка: некорректный формат запроса." << std::endl;
        std::cerr << "Пример использования: --query \"FOLD 523456795\"" << std::endl;
        return;
    }

    // Проверка, что ключ состоит только из цифр
    for (char const& c : numberKey) {
        if (std::isdigit(c) == 0) {
            std::cerr << "Ошибка: ключ должен состоять только из цифр." << std::endl;
            return;
        }
    }

    // В примере используется разбиение на части по 3 символа.
    const int chunkSize = 3;

    std::cout << "Ввод: " << numberKey << std::endl;
    std::cout << "Вывод: ";
    foldingHash(numberKey, chunkSize);
}