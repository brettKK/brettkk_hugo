




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


###