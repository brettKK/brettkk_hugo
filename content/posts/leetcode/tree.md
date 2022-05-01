---
title: "tree"
date: 2021-04-05T11:33:56+08:00
draft: true
isCJKLanguage: true
markup: mmark
tags: ["leetcode"]
series: [""]
categories: ["life"]
---

### 种类

+ 树
    + 多叉树
        + B树
        + 字典树
    + 二叉树
        + 二叉排序树
        + 二叉平衡树 AVL
        + 伸展树 
        + 红黑树


```golang

type TreeNode struct {
    value int
    left *TreeNode
    right *TreeNode
}

// bin search tree
type BinTree struct {
    root *TreeNode
}

func (t *BinTree) Insert(value int) {
    node := &TreeNode{value, nil, nil}
    if t.root == nil {
        t.root = node
    } else {
        insertNode(t.root, node)
    }
}
func insertNode(root, tmp *TreeNode) {
    if tmp.value < root.value {
        if root.left == nil {
            root.left = tmp
            return
        } else {
            insertNode(root.left, tmp)
        }
    } else {
        if root.right == nil {
            root.right = tmp
            return 
        } else {
            insertNode(root.right, tmp)
        }
    }
}


func (t *BinTree) Inorder(f func(int)) {
    inorder(t.root, f)
}
func inorder(root *TreeNode, f func(int)) {
    if root != nil {
        inorder(root.left, f)
        f(root.value)
        inorder(root.right, f)
    }
}

func (t *BinTree) getMin() int {
    if t == nil {
        return -1
    }
    for t.left != nil {
        t = t.left
    }
    return t.value
}

func (t *BinTree) getMax() int {
    if t== nil {
        return -1
    }
    for t.right != nil {
        t = t.right
    }
    return t.value
}

```

### 获取树的高度

```golang
func getHeight(p *TreeNode) int {
    if p == nil {
        return 0
    }
    left := getHeight(p.Left)
    right := getHeight(p.Right)
    if right >= left {
        return right + 1
    }
    return left + 1
}

```

层次遍历bfs求树的高度

```golang
func getHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }
    res := 0
    que := []*TreeNode{}
    que = append(que, root)
    for len(que) != 0 {
        size := len(que)
        for i := 0; i < size; i++ {
            cur := que[0] //get first
            que = que[1:] //pop first
            if cur.Left != nil {
                que = append(que, cur.Left)
            }
            if cur.Right != nil {
                que = append(que, cur.Right)
            }
        }
        res++
    }
    return res
}

```


### 合并两个二叉树

leetcode 617。

```golang
func mergeTree(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil {
        return t2
    }
    if t2 == nil {
        return t1
    }
    t1.val += t2.val
    t1.left = mergeTree(t1.left, t2.left)
    t1.rirght = mergeTree(t1.right, t2.right)
    return t1
}

func mergeTree(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1== nil {
        return t2
    }
    if t2 == nil {
        return t1
    }
    root := &TreeNode{}
    root.val = t1.val + t2.val
    root.left = mergeTree(t1.left, t2.left)
    root.right = mergeTree(t1.right, t2.right)
    return root
}

// use queue, level travel
func mergeTree(t1 *TreeNode, t2 *TreeNode) *TreeNode{
    if t1 == nil {
        return t2
    }
    if t2 == nil {
        return t1
    }
    que := make([]*TreeNode, 0)
    que = append(que, t1)
    que = append(que, t2)
    for len(que) != 0 {
        node1 := que[len(que)-1]
        que = que[:len(que)-1]
        node2 := que[len(que)-1]
        que = que[:len(que)-1]
        node1.val += node2.val
        if node1.left != nil && node2.left != nil {
            que = append(que, node1.left)
            que = append(que, node2.left)
        }
        if node1.right != nil && node2.right != nil {
            que = append(que, node1.right)
            que = append(que, node2.right)
        }
        if node1.left == nil && node2.left != nil {
            node1.left = node2.left
        }
        if node1.right == nil && node2.right != nil {
            node1.right = node2.right
        }
    }
    return t1
}
```

### 判断2棵树是不是子树关系

```golang
func isSubStructure(A, B TreeNode) bool {
    if A == nil || B == nil {
        return false
    }
    return isSubTree(A, B) || isSubStructure(A.left, B) || isSubStructure(A.right, B)
}

func isSubTree(A, B TreeNode) bool {
    if A == nil && B == nil {
        return true
    }
    if B == nil {
        return true
    }
    if A == nil {
        return false
    }
    if A.val != B.val {
        return false
    }
    return isSubTree(A.left, B.left) && isSubTree(A.right, B.right)
}

```

### 前序和中序的数组重建二叉树

```golang

func buildTree(preOrder, inOrder []int) *TreeNode{
    preLen, inLen := len(preOrder), len(inOrder)
    inMap := make(map[int]int, inLen)
    for i := 0; i < inLen; i++ {
        inMap[inOrder[i]] = i
    }
    return build(preOrder, inOrder, 0, preLen - 1, 0, inLen - 1, inMap)
}

func build(preOrder, inOrder []int, preLeft, preRight, inLeft, inRight int, inMap map[int]int) *TreeNode {
    if preLeft > preRight || inLeft > inRight {
        return nil
    }
    root := &TreeNode{val: preOrder[preLeft]}
    index := map[root.val]
    root.left := build(preOrder, inOrder, preLeft + 1, preLeft + index - inLeft, inLeft, index - 1, inMap)
    root.right := build(preOrder, inOrder, preLeft + index - inLeft + 1, preRight, index + 1, inRight, inMap)
    return root
}

```

### 中序和后序的数组重建二叉树
前序和后序不能唯一确定树。
```golang
func build(inOrder, postOrder []int) *TreeNode{
    // 边界条件
    if len(postOrder) == 0 {
        return nil
    }
    // 后序数组的最后一个item，为root
    rootValue := postOrder[len(postOrder) - 1]
    root := &TreeNode{val: rootValue}
    //叶子结点
    if len(postOrder) == 1 {
        return root
    }
    // 找到中序数组中的切割点
    var delimiterIndex int
    for delimiterIndex = 0; delimiterIndex < len(inOrder); delimiterIndex++ {
        if inOrder[delimiterIndex] == rootValue {
            break
        }
    }
    // 切割中序数组
    // 切割后序数组

    root.left = build(中序left, 后序left)
    root.right = build(中序right, 后序right)
    return root
}

```


### 字典树

```golang
type TrieNode struct {
    isEnd bool
    son []TrieNode
}

func insert(word string, root TrieNode) {
    cur := root
    for ch := range word {
        index := int(ch - 'a')
        if cur.son[index] == nil {
            cur.son[index] = TrieNode{}
        }
        cur = cur.son[index]
    }
    cur.isEnd = true
}

func search(word string, root TrieNode) bool {
    cur := root
    for ch := range word {
        index := int(ch - 'a')
        if cur.son[index] == nil {
            return false
        }
        cur = cur.son[index]
    }
    return cur.isEnd
}

func searchPrefix(word string) bool {
    cur := root
    for ch := range word {
        index := int(ch - 'a')
        if cur.son[index] == nil {
            return false
        }        
        cur = cur.son[index]
    }
    return true
}

```

```golang

type TrieNode struct{
    isEnd bool
    sonMap Map<char, TrieNode>
}

```


### 二叉搜索树

二叉搜索树BST的中序遍历是从小到大的有序集合。

#### 二叉搜索树中的搜索 leetcode 700。

```golang
func searchBST(root *TreeNode, target int) *TreeNode{
    if root == nil || root.val == target {
        return root
    }
    if root.val > target {
        return searchBST(root.left, target)
    }
    if root.val < target {
        return searchBST(root.right, target)
    }
    return nil
}

```

```golang
//迭代法
func serachBST(root *TreeNode, target int) *TreeNode{
    for root != nil {
        if root.val ==target {
            return root
        } else if root.val < target {
            root = root.right
        } else {
            root = root.left
        }
    }
    return nil
}

```

#### 验证二叉搜索树

lc 98。

思路1: 中序遍历二叉树，看是否为有序集合。


