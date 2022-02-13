---
title: "栈"
date: 2021-04-05T11:33:56+08:00
draft: true
isCJKLanguage: true
markup: mmark
tags: ["leetcode"]
series: [""]
categories: ["life"]
---

```golang

type Stack struct {
    items []int
}

func NewStack() *Stack{
    return &Stack{
        items: []int{}
    }
}

func (s *Stack)Push(value int) {
    s.items = append(s.items, value)
}

func (s *Stack)Pop() int {
    if len(s.items) == 0 {
        return -1
    }
    item := s.items[len(s.items) - 1]
    s.items = s.items[0: len(s.items) - 1]
    return item
}

func (s *Stack)Len() int {
    return len(s.items)
}

```

### 双栈排序

```golang
func sort(sk Stack, tmp Stack) Stack {
    for s.Len() != 0 {
        val := sk.Pop()
        if val >= tmp.Peek() {
            // sk.Pop()
            tmp.Push(val)
        } else {
            for tmp.Len() != 0 {
                sk.Push(tmp.Pop())
                if val <= tmp.Peek() {
                    break
                }
            }
            tmp.Push(val)
        }
    }
    return tmp
}
```