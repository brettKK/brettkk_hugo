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
### 利用栈对树进行前中后遍历

#### 前
```golang
func pre_order(root *TreeNode) []*TreeNode {
    var result []*TreeNode
    var st Stack
    if root != nil {
        st.push(root)
    }
    for len(st) != 0 {
        node := st.pop()
        result = append(result, node)
        if node.Right != nil {
            st.push(node.Right)
        }
        if node.Left != nil {
            st.push(node.Left)
        }
    } 
    return result
}
```

#### 中
```golang
func in_order(root *TreeNode) []*TreeNode{
    var result []*TreeNode
    var st Stack
    if root != nil {
        st.push(root)
    }
    cur := root
    for cur != nil || len(st) != nil {
        if cur != nil {
            st.push(cur.left)
            cur = cur.left //左
        } else {
            cur := st.pop()
            result = append(result, cur) //中
            cur = cur.right // 右
        }
    }
    return result
}
```

#### 后
```golang
// pre 中左右 --》 中右左-- 》 反转 左右中
func post_order(root *TreeNode) []*TreeNode{
    var st Stack
    var result []*TreeNode
    if root == nil {
        return result
    }
    st.push(root)
    for len(st) != 0 {
        node := st.pop()
        result = append(result, node)
        if node.Left != nil {
            st.push(node.Left)
        }
        if node.right != nil {
            st.push(node.right)
        }
    }
    reverse(result)
    return result
}
```

### 利用栈对树进行前中后遍历 （统一）

要处理的节点放入栈之后，紧接着放入一个空指针作为标记。 统一三种遍历的写法

#### 前
```golang
func pre_order(root *TreeNode) []*TreeNode{
    var result []*TreeNode
    var st Stack
    if root != nil {
        st.push(root)
    }
    for len(st) != 0 {
        node := st.top()
        if node != nil {
            node := st.pop() // 将该节点弹出，避免重复操作，下面再将右中左节点添加到栈中
            if node.Right != nil {
                st.push(node.Right) //右
            }
            if node.Left != nil {
                st.push(node.Left) // 左
            }
            st.push(node)  //中
            st.push(nil)
        } else {
            st.pop() // pop nil
            node := st.pop()
            result = append(result, node)
        }
    }
    return result
}

```


#### 中

```golang

func in_order(root *TreeNode) []*TreeNode{
    var result []*TreeNode
    var st Stack
    if root != nil {
        st.push(root)
    }
    for len(st) != 0 {
        node := st.top()
        if node != nil {
            node := st.pop() // 将该节点弹出，避免重复操作，下面再将右中左节点添加到栈中
            if node.Right != nil {
                st.push(node.Right)
            }
            st.push(node)
            st.push(nil)
            if node.Left != nil {
                st.push(node.Left)
            }
        } else {
            st.pop() // pop nil
            node := st.pop()
            result = append(result, node)
        }
    }
    return result
}

```

#### 后
```golang
func post_order(root *TreeNode) []*TreeNode{
    var result []*TreeNode
    var st Stack
    if root != nil {
        st.push(root)
    }
    for len(st) != 0 {
        node := st.top()
        if node != nil {
            node := st.pop()
            st.push(node)  //中
            st.push(nil)
            if node.Right != nil {
                st.push(node.Right) 
            }
            if node.Left != nil {
                st.push(node.Left)
            }
        } else {
            st.pop() // pop nil
            node := st.pop()
            result = append(result, node)
        }
    }
    return result
}

```

### 递归函数的return

+ no return  (need search total tree, 113)
+ have return  lc 112， lc 113 path_sum , lc 136.


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


### 二叉树的深度

#### 二叉树的最小深度

```golang

func getMinDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := getMinDepth(root.left)
    right := getMinDepth(root.right)
    // 后序遍历的原因： 需要比较递归之后的结果
    if root.left == nil && root.right != nil {
        return 1 + right   
    }
    if root.right == nil && root.left != nil{
        return 1 + left
    }

    return 1 + min(left, right)
}

```

```golang
// level traval

func getMinDepth(root *TreeNode) int {
    if root = nil {
        return 0
    }
    depth := 1
    var que []*TreeNode 
    que.append(root)
    for len(que) != 0 {
        depth++
        level_size := len(que)
        for i := 0; i < level_size; i++ {
            node := que[l0]
            que = que[1:]
            if node.left != nil {
                que = append(que, node.left)
            }
            if node.right != nil {
                que = append(que, node.right)
            }
            if node.left == nil && node.right == nil {
                return depth
            }
        }
    }
    return depth
}

```

#### 判断是否为平衡二叉树  lc 110

```golang
var is_balance bool
func judge_balance(root *TreeNode) bool {
    if root == nil {
        return true
    }
    is_balance = true
}

func helper(root *TreeNode) int {
    if is_balance == false {
        return 0
    }
    if root == nil {
        retun 0
    }
    if is_balance == true {
        left := helper(root.left)
        right := helper(root.right)
        if math.abs(left, right) > 1 {
            is_balance = false
        }
    }
    return 1 + max(left, right)
}

```

#### 二叉树的最大深度 lc 104

根节点的高度。

```golang
func get_depth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.left ==nil && root.right == nil {
        return 1
    }
    left_depth := get_depth(root.left)
    right_depth := get_depth(root.right)
    return 1 + max(left_depth, right_depth) 
}

```

### 反转二叉树 invert-binary-tree

遍历时把每个节点的左右孩子交换一下。 

前序，后序，层次遍历都可行。

中序不行。因为先左孩子交换孩子，再根交换孩子（做完后，右孩子已经变成了原来的左孩子），再右孩子交换孩子（此时其实是对原来的左孩子做交换）。


```golang

func inverse(root *TreeNode) *TreeNode{
    if root == nil || (root.left == nil && root.right == nil) {
        return root
    }
    temp := root.left
    root.left = root.right
    root.right = tem
    inverse(root.left)
    inverse(root.right)
    return root
}

```

### 二叉树的所有路径 lc 257

路径时从根到叶子，所以前序比较合适。

```golang

// 前序遍历，stack
func getAllPath(root *TreeNode) [][]*TreeNode {
    var result [][]*TreeNode

    var stack Stack
    if root == nil {
        return result
    }
    stack.push(root)

    path := []*TreeNode{}
    for len(stack) != 0 {
        node = stack.pop()
        path = append(path, node)
        if node.left == nil && node.right == nil {
            result = append(result, path)
            continue
        }
        if node.right != nil {
            stack.push(node.right)
        }
        if node.left != nil {
            stack.push(node.left)
        }
    }
    return result
}


```


```golang
func getAllPath(root *TreeNode) [][]*TreeNode{
    var result [][]*TreeNode
    if root == nil {
        return result
    }
    dfs(root, []*TreeNode{})
    return result
}

func dfs(root *TreeNode, temp []*TreeNode) {
    temp = append(temp, root)
    if root.left == nil && root.right == nil {
        // 找叶子节点
        result = append(result, temp)
        return
    }
    if root.left != nil {
        dfs(root.left, temp)
        temp = temp[:len(temp)-1]
    }
    if root.right != nil {
        dfs(root.right, temp)
        temp = temp[:len(temp)-1]
    }
}

func binaryTreePaths(root *TreeNode) []string {
    res := make([]string, 0)
    var travel func(node *TreeNode, s string)
    travel = func(node *TreeNode, s string) {
        if node.Left == nil && node.Right == nil {
            v := s + strconv.Itoa(node.Val)
            res = append(res, v)
            return
        }
        s = s + strconv.Itoa(node.Val) + "->"
        if node.Left != nil {
            travel(node.Left, s)
        }
        if node.Right != nil {
            travel(node.Right, s)
        }
    }
    travel(root, "")
    return res
}

```

```golang
func getAllPath_error(root *TreeNode) {
    var result [][]*TreeNode
    if root == nil {
        return result
    }
    dfs(root, []*TreeNode{})
    return result
}

/*
            1
        2        3
    4                    
        5
*/
// 走到4时 1， 2， 4 会被放到path里
func dfs_error(root *TreeNode, temp []*TreeNode) {
    if root == nil {
        result <- temp
        return
    }
    temp = append(temp, root)
    dfs_error(root.left, temp)
    dfs_error(root.right, temp)
    temp = temp[:len(temp)-1]
}

```


### 判断二叉树 是否为对称的
/*
                1
            2       2
          3   4   4   3 
*/


```golang
func is_sys(root *TreeNode) bool {
    if root == nil return true
    return helper(root.left, root.right)
}

func helper(left, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    if left == nil {
        return false
    }
    if right == nil {
        return false
    }
    if left.val != right.val {
        return false
    }
    return helper(left.left, right.right) && helper(left.right, right.left)
}

```

### 求二叉树中最大路径和 lc 124

```golang
var result int
func get_max_path(root *TreeNode) int {
    helper(root)
    return result
}

func helper(root *TreeNode) int {
    if root == nil {
        return nil
    }
    left := helper(root.left)
    right := helper(root.right)
    result = max(result, root.val + left + right)
    return root.val + max(left, right)
}

```


### 监控二叉树 lc 968

给定一个二叉树，我们在树的节点上安装摄像头。
节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。
计算监控树的所有节点所需的最小摄像头数量。

分析： 摄像头可以覆盖上中下三层，如果把摄像头放在叶子节点上，就浪费的一层的覆盖。
先给叶子节点父节点放个摄像头，然后隔两个节点放一个摄像头，直至到二叉树头结点。

采用哪种方式遍历？ 

从下往上推到，使用后序遍历的方式。

```golang
func traversal(root *TreeNode) {
    if terminate return
    
    left := traversal(root.left)  //左
    right := traversal(root.right) //右

    // do  中

    return 
}

```

如何隔两个节点放摄像头？

给节点增加状态： 0节点无覆盖，1节点摄像头，2节点有覆盖。

```golang
var result int
func traversal(root *TreeNode) int {
    if root == nil {  //空节点，该节点有覆盖
        return 2
    }
    left := traversal(root.left)
    right := traversal(root.right)

     // 左右节点都有覆盖
    if (left == 2 && right == 2) return 0
    
    // left == 0 && right == 0 左右节点无覆盖
    // left == 1 && right == 0 左节点有摄像头，右节点无覆盖
    // left == 0 && right == 1 左节点有无覆盖，右节点摄像头
    // left == 0 && right == 2 左节点无覆盖，右节点覆盖
    // left == 2 && right == 0 左节点覆盖，右节点无覆盖
    if (left == 0 || right == 0) {
        result++;
        return 1;
    }

    // left == 1 && right == 2 左节点有摄像头，右节点有覆盖
    // left == 2 && right == 1 左节点有覆盖，右节点有摄像头
    // left == 1 && right == 1 左右节点都有摄像头
    // 其他情况前段代码均已覆盖
    if (left == 1 || right == 1) return 2;

    //逻辑不会走到这里
    return -1
}

func getMinCarmera(root *TreeNode) int {
    if traversal(root) == 0 { // 根节点为无覆盖，+1
        result++
    }
    return result
}

```

### 层次遍历

#### lc 513 找二叉树的左下角的值

层次遍历：每次记录每层的第一个元素，最后一层的第一个元素即是左下角的值。

递归的方式， 怎么找在树的最后一行找到最左边的值？

使用递归法，如何判断是最后一行? 用max_dep记录最大深度。
如果找最左边的？可以使用前序遍历，先优先左边搜索。

```golang
var value int
var max_dep int
func find_bottom_left_value(root *TreeNode) int{
    find_left_value(root, 0)
    return value
}

func find_left_value(root *TreeNode, dep int) {
    if root == nil {
        return
    }
    if root.left == nil && root.right == nil {
        if dep > max_dep {
            max_dep = dep
            value = root.val
        }
        return
    }
    if root.left != nil {
        find_left_value(root.left, dep+1)
    }
    if root.right != nil {
        find_left_value(root.right, dep+1)
    }
}

```