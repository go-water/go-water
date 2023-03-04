二叉树排序在很多场景下用得着，也常常是很多面试官爱提问的知识点，这里使用go语言实现一个二叉树数据结构，同时实现排序，它核心就一个结构体，每个节点就是一个结构体，它同时又包含两个同类型指针，一个左枝，一个右枝，小的值存入左枝，大值存入右枝。

<div class="text-center">![二叉树](/images/2023/2023022608282963.jpg)
</div>代码
```
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func main() {
	tree := create(10)
	fmt.Println("The value of the root of the tree is", tree.Value)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root of the tree is", tree.Value)
}
```
输出结果
```
F:\github.com\iissy\goexample>go run main.go
The value of the root of the tree is 3
0 2 3 4 9 11 12 13 16 17 18 19        
-10 -2 0 2 3 4 9 11 12 13 16 17 18 19 
The value of the root of the tree is 3
```