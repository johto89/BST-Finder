package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Tree struct {
	root *Node
}
type Node struct {
	key   string
	left  *Node
	right *Node
}

// Tree
func (t *Tree) insert(data string) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

// Node
func (n *Node) insert(data string) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

// Node - Left - Right
func printPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%v \n", n.key)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

// Left - Right - Node
func printPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%v \n", n.key)
	}
}

// Left - Node - Right
func printInOrder(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		fmt.Printf("%v \n", n.key)
		printPostOrder(n.right)
	}
}

// Searc content in tree
func search(n *Node, key string) bool {
	if n == nil {
		return false
	}
	if key < n.key {
		return search(n.left, key)
	}
	if key > n.key {
		return search(n.right, key)
	}
	return true
}

func main() {
	var t Tree

	//// Insert Tree
	// t.insert("aa")
	// t.insert("ab")
	// t.insert("ae")
	// t.insert("cc")
	// t.insert("cd")
	// t.insert("a5")
	// t.insert("a7")
	// t.insert("0a")
	// t.insert("2")
	// t.insert("1")
	// t.insert("ah")
	// t.insert("aj")

	file, err := os.Open("test.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	for _, each_ln := range text {
		split := strings.Split(each_ln, ":")
		t.insert(split[0])
	}

	//fondr := "f8d92b4001cc11b535e099bb795059d4" //==> false
	fondr := "d8d3240e40939a72e153cd9069ec54e3" //==>true

	//print tree
	fmt.Printf("Pre Order: ")
	printPreOrder(t.root)
	// fmt.Println()
	// fmt.Printf("Post Order: ")
	// printPostOrder(t.root)
	// fmt.Println()
	// fmt.Printf("In Order: ")
	// printInOrder(t.root)
	// fmt.Println()

	result := search(t.root, fondr)
	fmt.Println(result)

	fmt.Println("Done")

}
