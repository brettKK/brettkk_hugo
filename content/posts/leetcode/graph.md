---
title: "图"
date: 2022-05-18T11:33:56+08:00
draft: true
isCJKLanguage: true
markup: mmark
tags: ["leetcode"]
series: [""]
categories: ["life"]
---

所有的数据结构的存储方式只有两种：内存上的顺序存储（数组存储），离散存储（链表存储）。
图也不例外。图的两种表示方法，邻接表就是链表，邻接矩阵就是二维数组。
+ 邻接表
    + 优点：节省空间
    + 缺点：操作的效率上肯定比不过邻接矩阵
+ 二维数组
    + 优点：判断连通性迅速
    + 缺点：图比较稀疏的话很耗费空间

概念字典：无向图，有向图， 加权图， prim， kruskal， 最小生成树，Dijkstra，网络流， 拓扑排序。

```golang



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