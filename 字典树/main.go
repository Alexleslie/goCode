/*
 * @Author: Alexleslie
 * @Date: 2022-03-07 17:23:22
 * @LastEditors: Alexleslie
 * @LastEditTime: 2022-03-08 15:21:45
 * @FilePath: \goCode\src\字典树\main.go
 * @Description: 实现一个字典树
 */
package main

import (
	"fmt"
)

type TreeNode struct {
	val    byte
	childs map[byte]*TreeNode
}

func main() {
	collect := []string{"hello", "hi", "hiking", "helloChicken"}
	root := buildTree(collect)
	fmt.Printf("%v, %v \n", string(root.val), root.childs)

	fmt.Println(root.findWord("hik"))

	travese(*root)

}

func buildTree(text []string) *TreeNode {
	root := &TreeNode{val: '/', childs: nil}
	root.childs = make(map[byte]*TreeNode)

	for i := range text {
		root.addWord(text[i])
	}
	return root
}

func (t *TreeNode) addWord(words string) {
	node := t

	for i := range words {
		word := words[i]
		if node.childs[word] == nil {
			temp := make(map[byte]*TreeNode)
			newnode := &TreeNode{val: word, childs: temp}
			node.childs[word] = newnode
		}
		node = node.childs[word]
	}
}

func (t *TreeNode) findWord(words string) int {
	node := t

	for i := range words {
		word := words[i]
		if node.childs[word] != nil {
			node = node.childs[word]
		} else {
			return -1
		}
	}
	return 0
}

func travese(root TreeNode) {
	ch := make(chan TreeNode, 10000)
	ch <- root

	result := ""
	for len(ch) > 0 {
		node := <-ch
		for i := range node.childs {
			result += string(i) + "-"
			ch <- *node.childs[i]
		}
	}
	fmt.Print(result)
}
