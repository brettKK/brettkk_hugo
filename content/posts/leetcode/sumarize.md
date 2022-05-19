---
title: "code"
date: 2021-04-05T11:33:56+08:00
draft: true
isCJKLanguage: true
markup: mmark
tags: ["leetcode"]
series: [""]
categories: ["life"]
---

+ why data structure
    + 在不同的应用场景，尽可能高效地增删查改
+ what data structure
    + 所有的数据结构的存储方式只有两种：内存上的顺序存储（数组存储），离散存储（链表存储）。
    + 所有的数据结构的操作方式：（增，删，查，改）。
+ how to make data structure 
    + Array
    + String
    + LinkedList
    + Binary Search
    + Math, Matrix, random
    + Tree
    + Back Tracking
    + Dynamic Programming
    + DFS & BFS
    + stack
    + priority queue
    + bit manipulation
    + topological sort
    + graph
    + Trie



github.com/soulmachine/leetcode


### 如何查，遍历

遍历的方式： 数组的遍历， 链表的遍历（迭代 or 递归）

数组的线性遍历方式
``` golang
func traversal(arr []int) {
    for i := 0; i < len(arr); i++ {
        //visit index i
    }
}
```

链表的遍历方式： 迭代 or 递归

```golang
type ListNode struct {
    val int
    next *ListNode
}
// 迭代遍历
func traversal(head *ListNode) {
    for tmp := head; tmp != nil; tmp = tmp.next{
        //visit
    }
}
// 递归遍历
func traversal(head *ListNode) {
    if head == nil {
        return nil
    }
    //visit 
    traversal(head.next)
}
```

二叉树的遍历：迭代遍历（利用栈） or 递归遍历

```golang
type TreeNode struct{
    val int
    left, right *TreeNode
}

func traversal(root *TreeNode) {
    if root == nil {
        return
    }
    traversal(root.left)
    traversal(root.right)
}
```

多叉树，N叉树的遍历

```golang
type TreeNode struct {
    val int
    childs []*TreeNode
}

func traversal(root *TreeNode) {
    if root == nil {
        return
    }
    for _, child := range root.childs{
        traversal(child)
    }
}
```

图的bfs遍历

```golang
var graph Graph
var visited []bool // 图是可能包含环的, 记录访问过的节点 可以 避免死循环。

func traversal(graph Graph, i int) {
    if visited[i] return 
    visited[i] = true
    for _, neighbor := range graph.neighbors {
        // 为什么不像backtrack一样将visisted放到循环内？
        // 循环内会漏掉起始点的遍历
        traversal(neighbor) 
    }
    visited[i] = false
}

```