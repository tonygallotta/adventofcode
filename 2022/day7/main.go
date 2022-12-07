package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileSystemObject struct {
	name        string
	isDirectory bool
	size        int
}

type TreeNode struct {
	value    FileSystemObject
	parent   *TreeNode
	children []*TreeNode
}

type FileSystemTree struct {
	root *TreeNode
}

func NewDirectory(name string) FileSystemObject {
	return FileSystemObject{name, true, 0}
}

func NewFile(size, name string) FileSystemObject {
	var sizeIntValue, _ = strconv.Atoi(size)
	return FileSystemObject{name, false, sizeIntValue}
}

func main() {
	data, _ := os.ReadFile("input.txt")
	var fileAsString = string(data)
	var lines = strings.Split(fileAsString, "\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var tree FileSystemTree = buildTree(lines)
	var matchingDirectories = subDirectoriesBelowSize(tree.root, 100000)
	var result = 0
	for _, dir := range matchingDirectories {
		// A bit wasteful to recompute this, we'll see if it ends up mattering
		result += objectSize(dir)
	}
	fmt.Println("Part 1: ", result)
}

func part2(lines []string) {
	var tree FileSystemTree = buildTree(lines)
	var usedSpace = objectSize(tree.root)
	var spaceNeeded = 30_000_000 - (70_000_000 - usedSpace)
	var matchingDirectories = subDirectoriesAboveSize(tree.root, spaceNeeded)
	var result = 70_000_000
	fmt.Println("Used space is ", usedSpace)
	fmt.Println("Looking for ", spaceNeeded, " across ", len(matchingDirectories), " directories")
	for _, dir := range matchingDirectories {
		var currentSize = objectSize(dir)
		if currentSize < result {
			result = currentSize
		}
	}
	fmt.Println("Part 2: ", result)
}

func subDirectoriesBelowSize(node *TreeNode, threshold int) []*TreeNode {
	var result = []*TreeNode{}
	for _, child := range node.children {
		if child.value.isDirectory {
			var size = objectSize(child)
			if size < threshold {
				result = append(result, child)
			}
			result = append(result, subDirectoriesBelowSize(child, threshold)...)
		}
	}
	return result
}

func subDirectoriesAboveSize(node *TreeNode, threshold int) []*TreeNode {
	var result = []*TreeNode{}
	for _, child := range node.children {
		if child.value.isDirectory {
			var size = objectSize(child)
			if size > threshold {
				result = append(result, child)
			}
			result = append(result, subDirectoriesAboveSize(child, threshold)...)
		}
	}
	return result
}

func objectSize(node *TreeNode) int {
	var result = 0
	if !node.value.isDirectory {
		return node.value.size
	}
	for _, child := range node.children {
		result += objectSize(child)
	}
	return result
}

func buildTree(lines []string) FileSystemTree {
	var root = TreeNode{NewDirectory("/"), nil, []*TreeNode{}}
	var tree = FileSystemTree{&root}
	var currentNode = &root
	for _, line := range lines[1:] {
		var parts = strings.Split(line, " ")
		if parts[0] == "$" {
			var cmd = parts[1]
			switch cmd {
			case "cd":
				var directoryName = parts[2]
				if directoryName == ".." {
					currentNode = currentNode.parent
				} else {
					var newNode = TreeNode{NewDirectory(directoryName), currentNode, []*TreeNode{}}
					currentNode.children = append(currentNode.children, &newNode)
					currentNode = &newNode
				}
			}
		} else {
			// We'll discover directories as we cd to them, so only add files here
			if parts[0] != "dir" {
				var newNode = TreeNode{NewFile(parts[0], parts[1]), currentNode, []*TreeNode{}}
				currentNode.children = append(currentNode.children, &newNode)
			}
		}
	}
	return tree
}
