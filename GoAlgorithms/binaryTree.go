package main

import "fmt"

type Tree struct {
	value string
	left  *Tree
	right *Tree
}

//前序根左右
func front(tree *Tree) {
	if tree == nil {
		return
	}
	fmt.Println(tree.value)
	front(tree.left)
	front(tree.right)
}

//中序左根右
func middle(tree *Tree) {
	if tree == nil {
		return
	}
	fmt.Println(tree.left.value)
	front(tree)
	front(tree.right)
}

//后序左右根
func backend() {

}

func main() {

	var leftTree Tree
	leftTree.value = "b"
	leftTree.left = nil
	leftTree.right = nil

	var rightTree Tree
	rightTree.value = "c"
	rightTree.left = nil
	rightTree.right = nil

	var tree Tree
	tree.value = "a"
	tree.left = &leftTree
	tree.right = &rightTree

	middle(&tree)
}
