// Задание 4 | Вариант 3

#include "avlChecker.h"

#include <iostream>
#include <string>
#include <vector>

#include "binaryTree.h"

void runAvlChecker(int argc, char* argv[]) {
    // Создаем экземпляр нашего бинарного дерева.
    BinaryTree tree;

    std::cout << "Введите последовательность целых чисел, оканчивающуюся нулем:" << std::endl;

    int number;
    // Читаем числа из стандартного ввода, пока не встретим 0.
    while (std::cin >> number && number != 0) {
        // Вставляем каждое число в дерево, предварительно сконвертировав в строку.
        tree.insert(std::to_string(number));
    }

    // Вызываем наш новый метод для проверки сбалансированности.
    if (tree.isAVLBalanced()) {
        std::cout << "YES" << std::endl;
    } else {
        std::cout << "NO" << std::endl;
    }
}