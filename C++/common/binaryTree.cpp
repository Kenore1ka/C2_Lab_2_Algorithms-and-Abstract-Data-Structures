#include "binaryTree.h"
#include <cstring>
#include <fstream>
#include <iostream>
#include <queue>
#include <string> // для std::stoi

using namespace std;

/*
 Конструктор: инициализирует пустое дерево (root = nullptr).
*/
BinaryTree::BinaryTree() : root(nullptr) {}

/*
 Деструктор: освобождает память, удаляя все узлы.
*/
BinaryTree::~BinaryTree() {
    deleteTree(root);
}

/*
 deleteTree: рекурсивно удаляет поддерево, начиная с node.
*/
void BinaryTree::deleteTree(Node* node) {
    if (!node) return;
    deleteTree(node->left);
    deleteTree(node->right);
    delete node;
}

/*
 insert (public): публичный метод для вставки ключа.
 Запускает рекурсивную вставку от корня.
*/
void BinaryTree::insert(const std::string& key) {
    root = insert(root, key);
}

/*
 insert (private): рекурсивный метод для вставки в бинарное дерево ПОИСКА.
*/
BinaryTree::Node* BinaryTree::insert(Node* node, const std::string& key) {
    // Если дошли до пустого места - создаем новый узел
    if (node == nullptr) {
        return new Node(key);
    }

    // Преобразуем строки в числа для сравнения
    int key_val = std::stoi(key);
    int node_val = std::stoi(node->key);

    // Идем в левое или правое поддерево
    if (key_val < node_val) {
        node->left = insert(node->left, key);
    } else if (key_val > node_val) {
        node->right = insert(node->right, key);
    }

    // Если ключи равны, ничего не делаем, дубликаты не вставляем
    return node;
}


/*
 search: поиск ключа в бинарном дереве ПОИСКА.
 Возвращает true если найден, иначе false. Это эффективнее, чем BFS.
*/
bool BinaryTree::search(const std::string& key) {
    Node* current = root;
    int key_val = std::stoi(key);

    while (current != nullptr) {
        int current_val = std::stoi(current->key);
        if (key_val == current_val) {
            return true; // Ключ найден
        } else if (key_val < current_val) {
            current = current->left; // Идём налево
        } else {
            current = current->right; // Идём направо
        }
    }
    return false; // Ключ не найден
}

/*
 saveNode: сохраняет узлы в файл в порядке PRE-ORDER (корень-левый-правый),
 чтобы правильно восстановить структуру дерева при загрузке.
*/
void BinaryTree::saveNode(Node* node, std::ofstream& file) {
    if (!node) return;
    file << node->key << std::endl; 
    saveNode(node->left, file);      
    saveNode(node->right, file);     
}

/*
 saveToFile: сохраняет дерево в файл.
*/
void BinaryTree::saveToFile(const string& fileName) {
    ofstream file(fileName);
    if (!file.is_open()) {
        cout << "Не удалось открыть файл для записи." << endl;
        return;
    }
    saveNode(root, file);
    file.close();
}

/*
 loadNode: читает ключи из потока и вставляет их в дерево.
*/
void BinaryTree::loadNode(ifstream& file) {
    string key;
    while (file >> key) {
        insert(key);
    }
}

/*
 loadFromFile: загружает дерево из файла.
*/
void BinaryTree::loadFromFile(const string& fileName) {
    ifstream file(fileName);
    if (!file.is_open()) {
        cout << "Не удалось открыть файл для чтения." << endl;
        return;
    }
    
    deleteTree(root);
    root = nullptr;
    
    loadNode(file);
    file.close();
}

/*
 isFullNode: рекурсивная проверка свойства "full" для поддерева.
*/
bool BinaryTree::isFullNode(Node* node) const {
    if (!node) return true;
    
    bool hasLeft = (node->left != nullptr);
    bool hasRight = (node->right != nullptr);
    
    if (hasLeft && hasRight) {
        return isFullNode(node->left) && isFullNode(node->right);
    } else if (!hasLeft && !hasRight) {
        return true;
    } else {
        return false;
    }
}

/*
 isFull: публичная оболочка для проверки, является ли дерево full.
*/
bool BinaryTree::isFull() const {
    return isFullNode(root);
}

/*
 inorder: рекурсивный обход дерева в порядке in-order (левый-корень-правый).
*/
void BinaryTree::inorder(Node* node) {
    if (!node) return;
    inorder(node->left);
    cout << node->key << " ";
    inorder(node->right);
}

/*
 preorder: рекурсивный обход дерева в порядке pre-order (корень-левый-правый).
*/
void BinaryTree::preorder(Node* node) {
    if (!node) return;
    cout << node->key << " ";
    preorder(node->left);
    preorder(node->right);
}

/*
 postorder: рекурсивный обход дерева в порядке post-order (левый-правый-корень).
*/
void BinaryTree::postorder(Node* node) {
    if (!node) return;
    postorder(node->left);
    postorder(node->right);
    cout << node->key << " ";
}

/*
 printInorder: публичный метод для in-order обхода.
*/
void BinaryTree::printInorder() {
    cout << "Inorder обход: ";
    inorder(root);
    cout << endl;
}

/*
 printPreorder: публичный метод для pre-order обхода.
*/
void BinaryTree::printPreorder() {
    cout << "Preorder обход: ";
    preorder(root);
    cout << endl;
}

/*
 printPostorder: публичный метод для post-order обхода.
*/
void BinaryTree::printPostorder() {
    cout << "Postorder обход: ";
    postorder(root);
    cout << endl;
}

/*
 printBFS: публичный метод для BFS обхода.
*/
void BinaryTree::printBFS() {
    if (!root) {
        cout << endl;
        return;
    }
    
    queue<Node*> q;
    q.push(root);
    bool first = true;
    
    while (!q.empty()) {
        Node* cur = q.front(); 
        q.pop();
        
        if (!first) cout << " ";
        cout << cur->key;
        first = false;
        
        if (cur->left) q.push(cur->left);
        if (cur->right) q.push(cur->right);
    }
    cout << endl;
}

/*
 runBinaryTree: функция интерфейса командной строки.
*/
void runBinaryTree(int argc, char* argv[]) {
    BinaryTree tree;

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

    if (!fileName.empty()) tree.loadFromFile(fileName);

    string command;
    string arg;
    size_t pos = query.find(' ');
    if (pos != string::npos) {
        command = query.substr(0, pos);
        arg = query.substr(pos + 1);
    } else {
        command = query;
    }

    if (command == "TINSERT") {
        tree.insert(arg);
        if (!fileName.empty()) tree.saveToFile(fileName);
    } else if (command == "TGET") {
        if (tree.search(arg)) cout << arg << endl;
        else cout << "Ключ не найден" << endl;
    } else if (command == "TFULL") {
        cout << (tree.isFull() ? "true" : "false") << endl;
    } else if (command == "TSEARCH") {
        cout << (tree.search(arg) ? "true" : "false") << endl;
    } else if (command == "TINORDER") {
        tree.printInorder();
    } else if (command == "TPREORDER") {
        tree.printPreorder();
    } else if (command == "TPOSTORDER") {
        tree.printPostorder();
    } else if (command == "TBFS") {
        cout << "BFS обход: ";
        tree.printBFS();
    }
}