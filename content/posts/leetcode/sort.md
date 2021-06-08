---
title: "排序"
date: 2021-04-05T11:33:56+08:00
draft: true
isCJKLanguage: true
markup: mmark
tags: ["leetcode"]
series: [""]
categories: ["life"]
---

### 第k大的元素

> [3, 2, 1, 5, 6, 4] k=2, output= 5


```golang
// k大小的最小堆

func KthLargest(arr []int, k int) {
    heap := NewHeap()
    for i := 0; i < len(arr); i++ {
        if heap.Size() < k || arr[i] >= heap.Peek() {
            heap.Add(arr[i])
        }
    }
    return heap.Peek()
}
```

```golang
// 快速排序

```

同 347， 253， 295， 767， 703