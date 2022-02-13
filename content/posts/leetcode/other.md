
### 约瑟夫环

> 1-N个人，形成一个圈，从编号为1开始报数1，数到M的人x出列，x之后的人再次从1开始喊
> 直到最后一个人，求这个人的编号

#### 数组
遍历数组，数到m的位置标记为-1， 直到数组中仅剩一个非-1的元素，返回元素的位置

#### 环形链表
不做标记，直接从环形链表中删除。
```golang
type Node struct{
    int num,
    next Node,
}
func createCircleLink(int n) Node{
    head := Node{num: 1}
    cur = head
    for i := 2; i < n; i++ {
        tmp := Node{num: i}
        cur.next = tmp
        cur = cur.next
    }
    cur.next = head
    return head
}

func findNum(n, m int) int {
    if (m == 1 || n < 2) return n
    head := createCircleLink(n)
    cur, pre := head, nil
    count := 1
    for head.next != head {
        if count == m {
            pre.next = cur.next
            cur = pre.next
        } else {
            pre = cur
            cur = cur.next
        }
        count++
    }
    return head.num
}
```

### 找到子问题与原问题的规律，递归
```golang

func findNum(n, m int) int {
    if n == 1 {
        return 1
    }
    return (findNum(n - 1, m) + m -1 )%n + 1 
}

```


### 数据流中的中位数

大顶堆维护一半小的数，小顶堆维护一半大的数。中位数为2个堆顶的平均数或者小顶堆的根。

