package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Stack - простая реализация стека для строк
type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func (s *Stack) Top() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	index := len(*s) - 1
	element := (*s)[index]
	return element, true
}

func precedence(op rune) int {
	switch op {
	case '*':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

func applyOp(a, b int64, op rune) int64 {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	}
	return 0
}

func evaluate(expression string) (int64, error) {
	var values Stack
	var ops Stack

	for i := 0; i < len(expression); {
		char := rune(expression[i])

		if char == ' ' {
			i++
			continue
		}

		if char == '(' {
			ops.Push("(")
			i++
		} else if unicode.IsDigit(char) {
			var numStr strings.Builder
			for i < len(expression) && unicode.IsDigit(rune(expression[i])) {
				numStr.WriteRune(rune(expression[i]))
				i++
			}
			values.Push(numStr.String())
		} else if char == ')' {
			for {
				opStr, ok := ops.Top()
				if !ok || opStr == "(" {
					break
				}
				ops.Pop()

				val2Str, ok2 := values.Pop()
				val1Str, ok1 := values.Pop()
				if !ok1 || !ok2 {
					return 0, fmt.Errorf("некорректное выражение: нехватка операндов")
				}
				val2, _ := strconv.ParseInt(val2Str, 10, 64)
				val1, _ := strconv.ParseInt(val1Str, 10, 64)

				result := applyOp(val1, val2, []rune(opStr)[0])
				values.Push(strconv.FormatInt(result, 10))
			}
			if _, ok := ops.Pop(); !ok {
				return 0, fmt.Errorf("некорректное выражение: несогласованные скобки")
			}
			i++
		} else {
			for {
				opStr, ok := ops.Top()
				if !ok || opStr == "(" || precedence([]rune(opStr)[0]) < precedence(char) {
					break
				}
				ops.Pop()
				val2Str, ok2 := values.Pop()
				val1Str, ok1 := values.Pop()
				if !ok1 || !ok2 {
					return 0, fmt.Errorf("некорректное выражение: нехватка операндов")
				}
				val2, _ := strconv.ParseInt(val2Str, 10, 64)
				val1, _ := strconv.ParseInt(val1Str, 10, 64)

				result := applyOp(val1, val2, []rune(opStr)[0])
				values.Push(strconv.FormatInt(result, 10))
			}
			ops.Push(string(char))
			i++
		}
	}

	for !ops.IsEmpty() {
		opStr, _ := ops.Pop()
		val2Str, ok2 := values.Pop()
		val1Str, ok1 := values.Pop()
		if !ok1 || !ok2 {
			return 0, fmt.Errorf("некорректное выражение: нехватка операндов")
		}
		val2, _ := strconv.ParseInt(val2Str, 10, 64)
		val1, _ := strconv.ParseInt(val1Str, 10, 64)

		result := applyOp(val1, val2, []rune(opStr)[0])
		values.Push(strconv.FormatInt(result, 10))
	}

	if len(values) != 1 {
		return 0, fmt.Errorf("некорректное выражение")
	}

	finalResult, _ := values.Pop()
	return strconv.ParseInt(finalResult, 10, 64)
}

func runExpressionSolver() {
	fmt.Println("Введите арифметическое выражение для вычисления:")
	reader := bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	result, err := evaluate(expression)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
	} else {
		fmt.Println("Результат:", result)
	}
}