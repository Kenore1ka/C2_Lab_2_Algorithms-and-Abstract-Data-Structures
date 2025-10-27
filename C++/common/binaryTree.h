#pragma once

#include <string>
#include <fstream> 

class BinaryTree {
public:
    // Конструктор и деструктор
    BinaryTree();
    ~BinaryTree();

    // Основные операции
    void insert(const std::string& key);
    bool search(const std::string& key);

    // Работа с файлами
    void saveToFile(const std::string& fileName);
    void loadFromFile(const std::string& fileName);

    // Проверка, является ли дерево полным (full)
    // "Полное" (full) дерево - это дерево, в котором каждый узел имеет 0 или 2 потомка.
    bool isFull() const;

    // Обходы дерева
    void printInorder();   // Левый — Корень — Правый
    void printPreorder();  // Корень — Левый — Правый
    void printPostorder(); // Левый — Правый — Корень
    void printBFS();       // Уровень за уровнем (обход в ширину)

private:
    // Внутренняя структура для узла дерева
    struct Node {
        std::string key;
        Node* left;
        Node* right;

        // Конструктор для удобного создания узлов
        Node(const std::string& k) : key(k), left(nullptr), right(nullptr) {}
    };

    Node* root; // Указатель на корневой узел

    // Вспомогательные рекурсивные функции
    Node* insert(Node* node, const std::string& key); // Изменено для BST
    void deleteTree(Node* node);
    void saveNode(Node* node, std::ofstream& file);
    void loadNode(std::ifstream& file);
    bool isFullNode(Node* node) const;
    void inorder(Node* node);
    void preorder(Node* node);
    void postorder(Node* node);
};

// Интерфейс командной строки
void runBinaryTree(int argc, char* argv[]);