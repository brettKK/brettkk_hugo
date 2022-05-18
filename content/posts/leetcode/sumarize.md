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

### 所有可能路径 lc797

```golang
// input graph 邻接表， 是有向无环图
var result List<[]int>
func get_all_path(graph [][]int) List<[]int>{
    traversal(graph,0, []int{})
    return result
}
func traversal(graph [][]int, start int, path []int){
    path = append(path, start)
    if len(path) == len(graph) {
        result = append(result, path)
        return
    }
    for _, i := range graph[start]{
        traversal(graph, i, path)
    }
    path = path[:len(path)-1]
}

```

### 判断有向图中是否有环 （拓扑排序 or dfs）

lc 207 课程表

课程有依赖 -》转换为有向图。 有向图中有环则表示有循环依赖。

拓扑排序的对象是有向无环图

```golang
var graph []*ListNode // 邻接表的方式存图 graph[i] 为节点i指向的所有其他节点。
var hasCycle bool
var on_path []bool
var visited []bool

func build_graph(num_courses int, prerequisites[][]int) []*ListNode{
    graph = make([]*ListNode, nums_courses)
    for _, edge := range prerequisites{
        from, to := edge[1], edge[0]
        graph[from].Add(to)
    }
    return graph
}

// dfs
func has_cycle(num_courses int, prerequisites[][]int)
    graph = build_graph(num_courses, prerequisites)
    for i := 0; i < num_courses; i++ {
        // 遍历图中的所有节点
        traversal(graph, i)
    }
    return hasCycle
}

func traversal(graph []*ListNode, start int) {
    if on_path[start] {
        hasCycle = true
    }
    if visited[start] || hasCycle {
        return
    }
    visited[start] = true
    on_path[start] = true
    for _, i := range graph[start] { // list range...
        traversal(graph, i)
    }
    on_path[start] = true
}

```


lc 210

后序遍历的结果进行反转，就是拓扑排序的结果

var visited []boolean
// 记录后序遍历结果
var postorder []int

func findOrder(numCourses int, prerequisites [][]int) []int{
    // 先保证图中无环
    if (!has_cycle(numCourses, prerequisites)) {
        return new int[]{}
    }
    // 建图
    graph := build_graph(numCourses, prerequisites)
    // 进行 DFS 遍历
    visited = make([]bool, numCourses)
    for (int i = 0; i < numCourses; i++) {
        traverse(graph, i);
    }
    // 将后序遍历结果反转
    reverse(postorder)
    return postorder
}

void traverse(graph []*ListNode, int s) {
    if (visited[s]) {
        return
    }

    visited[s] = true;
    for (int t : graph[s]) {
        traverse(graph, t);
    }
    // 后序遍历位置
    postorder.add(s);
}

### 