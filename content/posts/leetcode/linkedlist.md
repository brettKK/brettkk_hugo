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

同lc 237， 141， 92， 25