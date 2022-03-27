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

### 单调栈

#### 输出数组每个元素与右边比自己大的元素的距离

```golang
// result 每个位置存index of next greater item  <lc 739 每日温度>
func nextGreater(arr []int) []int{
    var stack []int
    result := make([]int, len(arr))
    for i := range len(arr) {
        result[i] = -1
    }
    for i := range len(arr) {
        for len(stack) != 0 && arr[stack[len(stack)-1]] < arr[i] {
            // 维持递减栈， 遇到比栈顶元素时出栈，计算栈顶元素遇到比自己大的元素时的距离
            index := stack[len(stack) -1]
            stack = stack[:len(stack)]
            result[index] = i - index 
        }
        stack = append(stack, i)
    }
    return result
}

// result 每个位置存next greater item  <lc496 下一个更大元素>
func nextGreater(arr []int) []int{
    var stack []int
    result := make([]int, len(arr))
    for i := len(arr) - 1; i >= 0; i-- {
        for len(stack) != 0 && arr[i] >= stack[len(stack)-1] {
            // 维护递减栈
            stack = stack[:len(stack)-1]
        }
        //
        result[i] = len(stack) == 0 ? -1 : stack[len(stack)-1]
        stack = append(stack, arr[i])
    }
}

```

#### 输出最大矩形面积 <lc 84>

input = [2,1,5,6,2,3]; output = 5*2 = 10

```golang
func largestRectangle(arr []int) int {
    result := 0
    var stack []int
    // 保证栈为空， 所有元素参与了计算
    arr = append(arr, 0) 
    for i := range len(arr) {
        for len(stack) != 0 && arr[stack[len(stack)-1]] > arr[i] {
            // 维护递增栈，遇到比栈顶元素小时出栈计算
            index := stack[len(stack) -1] 
            stack = stack[:len(stack)-1]
            result = max(result, arr[index] * (i - index))
        }
        stack = append(stack, i)
    }
    return result
}

```
#### 输出二维数组中包含1的最大矩形面积

```golang

func maxRectange(matrix [][]int) int {
    rows := len(matrix)
    cols := len(matrix[0])
    arr := make([]int, cols)
    result := 0
    for i := range rows {
        for j := range cols {
            if matrix[i][j] = 0 {
                arr[j] = 0
            } else {
                arr[j] = arr[j] + 1
            }
        }
        tmp := largestRectangle(arr)
        result = max(result, tmp)
    }
    return result
}

```

#### lc 402 移掉K位数字

#### lc 581 最短无序连续子数组

#### lc 42 接雨水
从左往右维护单调递减栈，栈里存index。
当cur 大于栈顶元素，出现了凹地，可以接雨水。

#### lc 316 去除重复字符