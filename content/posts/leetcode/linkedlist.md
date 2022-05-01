---
title: "链表"
date: 2021-04-05T11:33:56+08:00
draft: true
isCJKLanguage: true
markup: mmark
tags: ["leetcode"]
series: [""]
categories: ["life"]
---

### 链表

```golang

type ListNode struct {
    value int
    next *ListNode
}
type LinkedList struct{
    head *ListNode
    size int
}

func (l *LinkedList) Append(value int) {
    listNode := &ListNode(value, nil)
    if l.head == nil {
        ll.head = listNode
    } else {
        cur := l.head
        for {
            if cur.next == nil {
                break
            }
            cur = cur.next
        }
        cur.next = listNode
    }
    l.size++
}

func (l *LinkedList) Insert(index int, value int) error {
    if index < 0 || index > l.size {
        return fmt.Errorf("")
    }
    addNode := &Node(value, nil)
    if i == 0 {
        addNode.next = l.head
        l.head= addNode
        return nil
    }
    cur := l.head
    i := 0
    for i < index - 1 {
        i++
        cur = cur.next
    }
    addNode.next = cur.next
    cur.next = addNode
    l.size++
    return nil
}

func (l *LinkedList) RemoveAt(index int) (*ListNode, error) {
    if index < 0 || l.size > index {
        return nil, fmt.Errorf("")
    }
    cur := l.head
    i := 0 
    for i < index - 1 {
        i++
        cur = cur.next
    }
    removeNode := cur.next
    cur.next = removeNode.next
    l.size--
    return removeNode, nil
}

func deleteNode(head *ListNode, val int) *ListNode {
    if head.val == val {
        return head.next
    }
    pre, cur := head, head.next
    for cur != nil && cur.val != val {
        pre = cur
        cur = cur.next
    }
    pre.next = cur.next
    return head
}

func (l *LinkedList) IndexOf(value int) int {
    cur := l.head
    i := 0
    for {
        if cur.value == value {
            return i
        }
        if cur.next == nil {
            return -1
        }
        cur = cur.next
        i++
    }
}

```

### 反转链表

递归

```golang
func reverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := reverse(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}
```

双指针
```golang
func reverse(head *ListNode) *ListNode {
    var pre *LListNode
    for head != nil {
        next := head.Next
        head.Next = pre
        pre = head
        head = next
    }
    return pre
}
```

### 反转指定区间内的链表

```golang
// lc 92
func reverseLR(left, right int, head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    dummy := &ListNode{}
    dummy.next = head
    pre, cur := dummy, head
    for i := 0; i < left; i++ {
        cur = cur.next
        pre = pre.next
    }
    if cur == nil || cur.next == nil {
        return head
    }
    rangeLeft := pre
    // reverse
    for i := left; i < right; i++ {
        tmp := cur.next
        cur.next.next =  cur
        cur.next = nil
        cur = tmp
    }
    rangeLeft.next.next = cur.next
    rangeLeft.next = cur
    return dummy.next
} 
```

```golang
// simple 
func reverseRange(left, right int, head *ListNode) *ListNode{
    dummy := &ListNode{}
    dummy.next = head
    p = dummy
    // p 指向left的前一个节点
    for i := 0; i < left - 1; i++ {
        p = p.next
    }
    m, n := p.next, p.next.next
    for i := 0; i < right - left; i++ {
        tmp := cur.next
        cur.next = pre
        pre = cur
        cur = tmp
    }
    // 连接
    p.next.next = cur
    p.next = pre
}
```

同lc 237， 141， 92， 25

### 判断链表是否有环

```golang

func has_circle(head *ListNode) bool {
    if head == nil || head.next == nil {
        return false
    }
    fast, slow := head, head
    for fast != nil && fast.next != nil {
        fast = fast.next.next
        slow = slow.next
        if fast == slow {
            return true
        }
    }
    return false
}
```

### 找链表中环的入口

+ 思路1: map 存下遍历过的节点，第一次重复的节点为环的入口。
+ 思路2: 快慢指针，


```golang 

```

### 合并2个有序链表

```golang
func merge2SortList(l1, l2 *ListNode) *ListNode {
    // dummy 停留在链表头，用于函数的返回
    dummy := &ListNode{}
    // cur 用于合并时的移动
    cur := dummy
    for l1 != nil && l2 != nil {
        if l1.val < l2.val {
            cur.next = l1
            l1 = l1.next
        } else {
            cur.next = l2
            l2 = l2.next
        }
        cur = cur.next
    }
    if l1 != nil {
        cur.next = l1
    }
    if l2 != nil {
        cur.next = l2
    }
    return dummy.next
}
```

### 合并K个升序链表

### 链表的快速排序

