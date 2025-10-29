// Задание 1 | Вариант 1

#include "expressionSolver.h"

#include <cctype>
#include <iostream>
#include <stdexcept>
#include "stack.h"

int precedence(char op) {
    if (op == '*') return 2;
    if (op == '+' || op == '-') return 1;
    return 0;
}

long long applyOp(long long a, long long b, char op) {
    switch (op) {
        case '+':
            return a + b;
        case '-':
            return a - b;
        case '*':
            return a * b;
    }
    return 0;
}

long long evaluate(const std::string& expression) {
    Stack values;
    values.init();
    Stack ops;
    ops.init();

    for (size_t i = 0; i < expression.length(); ++i) {
        if (expression[i] == ' ') {
            continue;
        }

        if (expression[i] == '(') {
            ops.push("(");
        } else if (isdigit(expression[i])) {
            std::string num_str;
            while (i < expression.length() && isdigit(expression[i])) {
                num_str += expression[i];
                i++;
            }
            values.push(num_str);
            i--;
        } else if (expression[i] == ')') {
            while (ops.top != nullptr && ops.top->data != "(") {
                if (values.top == nullptr || values.top->next == nullptr)
                    throw std::runtime_error("Некорректное выражение: нехватка операндов.");
                long long val2 = std::stoll(values.top->data);
                values.pop();
                long long val1 = std::stoll(values.top->data);
                values.pop();

                if (ops.top == nullptr)
                    throw std::runtime_error("Некорректное выражение: несогласованные скобки.");
                char op = ops.top->data[0];
                ops.pop();

                values.push(std::to_string(applyOp(val1, val2, op)));
            }
            if (ops.top == nullptr)
                throw std::runtime_error("Некорректное выражение: несогласованные скобки.");
            ops.pop();  // Удаляем (
        } else {
            while (ops.top != nullptr && ops.top->data != "(" &&
                   precedence(ops.top->data[0]) >= precedence(expression[i])) {
                if (values.top == nullptr || values.top->next == nullptr)
                    throw std::runtime_error("Некорректное выражение: нехватка операндов.");
                long long val2 = std::stoll(values.top->data);
                values.pop();
                long long val1 = std::stoll(values.top->data);
                values.pop();

                char op = ops.top->data[0];
                ops.pop();

                values.push(std::to_string(applyOp(val1, val2, op)));
            }
            ops.push(std::string(1, expression[i]));
        }
    }

    while (ops.top != nullptr) {
        if (values.top == nullptr || values.top->next == nullptr)
            throw std::runtime_error("Некорректное выражение: нехватка операндов.");
        long long val2 = std::stoll(values.top->data);
        values.pop();
        long long val1 = std::stoll(values.top->data);
        values.pop();

        char op = ops.top->data[0];
        ops.pop();

        values.push(std::to_string(applyOp(val1, val2, op)));
    }

    if (values.top == nullptr || values.top->next != nullptr) {
        throw std::runtime_error("Некорректное выражение.");
    }

    long long result = std::stoll(values.top->data);

    values.destroy();
    ops.destroy();

    return result;
}

void runExpressionSolver(int argc, char* argv[]) {
    (void)argc;  // Подавляем предупреждение о неиспользуемых переменных
    (void)argv;  // Подавляем предупреждение о неиспользуемых переменных

    std::string expression;
    std::cout << "Введите арифметическое выражение для вычисления:" << std::endl;
    std::getline(std::cin, expression);

    try {
        long long result = evaluate(expression);
        std::cout << "Результат: " << result << std::endl;
    } catch (const std::exception& e) {
        std::cerr << "Ошибка: " << e.what() << std::endl;
    }
}