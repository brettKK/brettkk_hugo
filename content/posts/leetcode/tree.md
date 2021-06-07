


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


###