#include <cstring>
#include <iostream>
#include <string>

// Подключаем заголовочные файлы всех наших заданий
#include "avlChecker.h"
#include "expressionSolver.h"
#include "foldingHasher.h"
#include "lruCache.h"
#include "patternMatcher.h"

// Объявляем новую функцию-обертку для expressionSolver
void runExpressionSolver(int argc, char* argv[]);

void print_help() {
    std::cout << "Программа для выполнения заданий лабораторной работы №2." << std::endl;
    std::cout << "Использование: ./lab2_program <номер_задания> [аргументы...]" << std::endl << std::endl;
    std::cout << "Доступные задания:" << std::endl;
    std::cout << "  1: Вычислитель выражений (интерактивный режим)." << std::endl;
    std::cout << "     Пример: ./lab2_program 1" << std::endl << std::endl;
    std::cout << "  2: Сопоставление с паттерном." << std::endl;
    std::cout << "     Пример: ./lab2_program 2 --query \"MATCH emails.txt *@*.ru\"" << std::endl
              << std::endl;
    std::cout << "  3: Проверка АВЛ-сбалансированности (интерактивный режим)." << std::endl;
    std::cout << "     Пример: ./lab2_program 3" << std::endl << std::endl;
    std::cout << "  4: Хеш-функция (метод свертки)." << std::endl;
    std::cout << "     Пример: ./lab2_program 4 --query \"FOLD 523456795\"" << std::endl << std::endl;
    std::cout << "  5: LRU Кэш." << std::endl;
    std::cout << "     Пример: ./lab2_program 5 \"cap = 2, Q = 2 Queries = SET 1 2 GET 1\"" << std::endl;
}

int main(int argc, char* argv[]) {
    if (argc < 2) {
        print_help();
        return 1;
    }

    int task_number = 0;
    try {
        task_number = std::stoi(argv[1]);
    } catch (const std::exception& e) {
        std::cerr << "Ошибка: первый аргумент должен быть числом - номером задания." << std::endl;
        print_help();
        return 1;
    }

    // Мы передаем в функции заданий аргументы, НАЧИНАЯ со второго (argv + 1),
    // чтобы они не видели номер задания, а работали как раньше.
    switch (task_number) {
        case 1:
            runExpressionSolver(argc - 1, argv + 1);
            break;
        case 2:
            runPatternMatcher(argc - 1, argv + 1);
            break;
        case 3:
            runAvlChecker(argc - 1, argv + 1);
            break;
        case 4:
            runFoldingHasher(argc - 1, argv + 1);
            break;
        case 5:
            // Для LRU кэша все аргументы сливаются в один
            if (argc > 2) {
                runLruCache(argc - 1, argv + 1);
            } else {
                std::cerr << "Ошибка: для задания 5 требуется строка с параметрами кэша." << std::endl;
            }
            break;
        default:
            std::cerr << "Ошибка: неизвестный номер задания '" << task_number << "'." << std::endl;
            print_help();
            return 1;
    }

    return 0;
}