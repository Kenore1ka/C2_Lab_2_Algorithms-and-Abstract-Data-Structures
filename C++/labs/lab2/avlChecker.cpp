// Задание 4 | Вариант 3

#include "avlChecker.h"
#include "binaryTree.h" 

#include <iostream>
#include <string>
#include <vector>

void runAvlChecker(int argc, char* argv[]) {
    // 1. Создаем экземпляр нашего бинарного дерева.
    BinaryTree tree;

    std::cout << "Введите последовательность целых чисел, оканчивающуюся нулем:" << std::endl;

    int number;
    // 2. Читаем числа из стандартного ввода, пока не встретим 0.
    while (std::cin >> number && number != 0) {
        // 3. Вставляем каждое число в дерево, предварительно сконвертировав в строку.
        tree.insert(std::to_string(number));
    }

    // 4. Вызываем наш новый метод для проверки сбалансированности.
    if (tree.isAVLBalanced()) {
        std::cout << "YES" << std::endl;
    } else {
        std::cout << "NO" << std::endl;
    }
}