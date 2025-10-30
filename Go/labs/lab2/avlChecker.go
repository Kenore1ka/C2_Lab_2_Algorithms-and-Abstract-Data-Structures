package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Value string
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func (t *BinaryTree) insertRec(node *TreeNode, value string) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value}
	}
	if value < node.Value {
		node.Left = t.insertRec(node.Left, value)
	} else if value > node.Value {
		node.Right = t.insertRec(node.Right, value)
	}
	return node
}

func (t *BinaryTree) Insert(value string) {
	t.Root = t.insertRec(t.Root, value)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + int(math.Max(float64(height(node.Left)), float64(height(node.Right))))
}

func (t *BinaryTree) isAVLBalancedRec(node *TreeNode) bool {
	if node == nil {
		return true
	}
	leftHeight := height(node.Left)
	rightHeight := height(node.Right)

	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return false
	}
	return t.isAVLBalancedRec(node.Left) && t.isAVLBalancedRec(node.Right)
}

func (t *BinaryTree) IsAVLBalanced() bool {
	return t.isAVLBalancedRec(t.Root)
}

func runAvlChecker() {
	tree := &BinaryTree{}
	fmt.Println("Введите последовательность целых чисел, оканчивающуюся нулем:")
	for {
		var number int
		_, err := fmt.Scan(&number)
		if err != nil || number == 0 {
			break
		}
		tree.Insert(fmt.Sprint(number))
	}

	if tree.IsAVLBalanced() {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}