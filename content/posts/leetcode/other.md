
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

#### 找到子问题与原问题的规律，递归
```golang

func findNum(n, m int) int {
    if n == 1 {
        return 1
    }
    return (findNum(n - 1, m) + m -1 )%n + 1 
}

```

---

### 数据流中的中位数

大顶堆维护一半小的数，小顶堆维护一半大的数。中位数为2个堆顶的平均数或者小顶堆的根。

---

### 接雨水

leetcode: trapping rain water.

思路： 
位置x能承多少水？ x的左边最大高度left_max, x的右边最大高度right_max。
位置x的盛水量等于 min(left_max, right_max) - h(x)。

```golang

func getWater(arr []int) int {
    leftMaxArr := make([]int, len(arr))
    leftMaxArr[0] = 0
    rightMaxArr := make([]int, len(arr))
    rightMaxArr[len(arr) - 1] = 0
    for i := 1; i < len(arr); i++ {
        leftMaxArr[i] = max(leftMaxArr[i - 1], arr[i])
    }
    for i := len(arr) - 2; i >=0; i-- {
        rightMaxArr[i] = max(rightMaxArr[i - 1], arr[i])
    }
    sum := 0 
    for i := 0; i < len(arr); i++ {
        if i == 0 || i == len(arr) - 1 {
            continue
        }
        cur := min(leftMaxArr[i], rightMaxArr[i]) - arr[i]
        if cur <= 0 {
            continue
        }
        sum += cur
    }
    return sum
}

```

```golang

// 单调栈的思路

维护 递减栈

```

---


### 计算任意两个年月日之间相隔的天数

写程序判断 别人的实现是否正确？ Fuzz 


### 取石子的游戏

巴什博弈。

N颗石头，一次可取1-M颗。

小a，小b轮番取石头，谁最后取完石头则胜利。

例如： N=8， M =3。

N mod(M + 1) == 0, 首先取石头的小a必败。
N mod(M + 1) != 0, 首先取石头的小a必胜。