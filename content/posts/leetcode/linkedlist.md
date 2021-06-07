---
title: "链表"
date: 2021-04-05T11:33:56+08:00
draft: false
isCJKLanguage: true
markup: mmark
tags: ["leetcode"]
series: [""]
categories: ["life"]
---

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