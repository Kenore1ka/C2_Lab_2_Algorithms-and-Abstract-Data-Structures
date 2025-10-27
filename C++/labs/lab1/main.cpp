#include <cstdlib>
#include <cstring>
#include <iostream>
#include <string>

#include "array.h"
#include "binaryTree.h"
#include "hashTable.h"
#include "linkedList.h"
#include "dlinkedList.h"
#include "queue.h"
#include "stack.h"
#include "set.h"

using namespace std;

int main(int argc, char* argv[]) {
    if (argc < 5) {
        return 1;
    }

    string fileName;
    string query;

    for (int i = 1; i < argc; i++) {
        if (strcmp(argv[i], "--file") == 0 && i + 1 < argc) {
            fileName = argv[i + 1];
            i++;
        } else if (strcmp(argv[i], "--query") == 0 && i + 1 < argc) {
            query = argv[i + 1];
            i++;
        }
    }

    if (fileName.empty() || query.empty()) {
        return 1;
    }

    // Получаем саму команду (слово до первого пробела)
    string command = query.substr(0, query.find(' '));

    // Проверяем, с чего начинается команда
    if (command.rfind("M", 0) == 0) {
        runDynamicArray(argc, argv);
    } else if (command.rfind("L", 0) == 0) {
        runLinkedList(argc, argv);
    } else if (command.rfind("D", 0) == 0) {
        runLLinkedList(argc, argv);
    } else if (command.rfind("Q", 0) == 0) {
        runQueue(argc, argv);
    } else if (command.rfind("S", 0) == 0 && command.rfind("SET", 0) != 0) {
        runStack(argc, argv);
    } else if (command.rfind("SET", 0) == 0) {
        runSet(argc, argv);
    } else if (command.rfind("H", 0) == 0) {
        runHashTable(argc, argv);
    } else if (command.rfind("T", 0) == 0) {
        runBinaryTree(argc, argv);
    } else {
        cerr << "Неизвестный тип команды." << endl;
    }

    return 0;
}