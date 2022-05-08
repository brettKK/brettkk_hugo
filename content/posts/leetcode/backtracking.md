---
title: "dfs & backtracking"
date: 2022-01-13T11:33:56+08:00
draft: true
isCJKLanguage: true
tags: ["leetcode"]
series: [""]
categories: ["life"]
---

## 全排列，组合，子集

### 求数组的全部子集

#### 数组没有重复元素的情况下，输出所有子集合

```golang
var res []string
func getAllSet(arr []byte) []string {
    dfsYesNo(arr, 0, []byte{})
    // print res

    dfs(arr, 0, []byte{})
    // print res
    return res
}

func dfsYesNo(arr []byte, position int, tempArr []byte) {
    if start == len(arr) -1 {
        res = append(res, tempArr)
        return
    }
    for i := position; i < len(arr); i++ {
        //选 i位置的元素
        tempArr = append(tempArr, arr[i])
        dfsYesNo(arr, i + 1, tempArr)
        // 不选i位置的元素
        tempArr = tempArr[:len(tempArr)-1]
        dfsYesNo(arr, i + 1, tempArr)
    }
}

func dfs(arr []byte, start int, tempArr []byte) {
    if start < len(arr) {
        res = append(res, temp)
    }
    for i := start; i < len(arr); i++ {
        temp = append(temp, arr[i])
        dfs(arr, i+1, temp)
        temp = temp[:len(temp)-1]
    }
}
```

#### 数组里有重复元素时，输出所有子集合

res 换成map去重



TODO 


### 二叉树

```golang
// stack 实现的前序 非递归遍历
func dfs(root *TreeNode) {
    if root == nil return 
    stack.Push(root)
    for len(stack) != 0 {
        node = stack.Pop()
        println(node)
        if node.right != nil {
            stack.Push(node.right)
        }
        if node.left != nil {
            stack.Push(node.left)
        }
    }
}


```

### 岛屿数量 lc 200

给定一个由`1`（陆地） 和 `0` （水）组成的二维网格，计算网络中岛屿（被边界或者水包围的区域）的数量 

```golang
// DFS
/* 思路： 扫描二维网格， 如果节点是1 则以节点为根进行dfs。 在dfs的过程中，每个被访问过的节点标记为0。统计启动dfs的根节点的数量，即为岛屿数量 */

func get_island_num(area [][]int) int {
    var island_nums int
    rows, cols := len(area), len(area[0])
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if area[i][j] == 1 {
                island_nums++
                dfs(area, i, j, rows, cols)
            }
        }
    }
    return island_nums
}

func dfs(area [][]int, i, j, rows, cols int) {
    area[i][j] = 0
    if i + 1 < rows && area[i+1][j] == 1 {
        dfs(area, i + 1, j, rows, cols, temp_try_pixel)
    }
    if i - 1 >= 0 && area[i-1][j] == 1 {
        dfs(area, i - 1, j, rows, cols, temp_try_pixel)
    }
    if j + 1 < cols && area[i][j+1] == 1{
        dfs(area, i, j + 1, rows, cols, temp_try_pixel)
    }
    if j - 1 >= 0 && area[i][j-1] == 1{
        dfs(area, i, j - 1, rows, cols, temp_try_pixel)
    }
}

func is_isolate_pixel(area [][]int, i, j, rows, cols int) bool {
    if i >= rows || j >= cols {
        return false
    }
    if area[i][j] == 0 {
        return false
    }
    if i - 1 >=0 && area[i-1][j] == 0 {
        //左 为 0
        if i + 1 < rows && area[i+1][j] == 0 {
            //右为0
            //...
            return true
        }
    }
    return false
}

```

```golang
// BFS
func get_island_nums(area [][]int) int {
    var island_nums int
    rows, cols := len(area), len(area[0])
    var que []tuple(a, b)
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if area[i][j] == 1 {
                island_nums++
                area[i][j] = 0
                que = append(que, (i, j))
                for len(que) != 0 {
                    (x, y) = que[0]
                    // pop
                    que = que[1:]
                    if x + 1 < rows && area[x + 1][y] == 1 {
                        area[x+1][y] = 0
                        que = append(que, (x+1, y))
                    }
                    if x - 1 >= 0 && area[x - 1][y] == 1 {
                        area[x-1][y] = 0
                        que = append(que, (x+1, y))
                    }
                    if y + 1 < cols && area[x][y + 1] == 1 {
                        area[x][y+1] = 0
                        que = append(que, (x+1, y))
                    }
                    if y - 1 >= 0 && area[x][y - 1] == 1 {
                        area[x][y-1] = 0
                        que = append(que, (x+1, y))
                    }
                }
            }
        }
    }
    return island_nums
}

```