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


### 快排

```golang
func Test_arr(t *testing.T) {
	arr := []int{5,2,10, 2,3,8,1, 3,7, 4}
	left := partition1(arr, 0, len(arr) - 1)
	fmt.Println("left=", left)
	for i, v := range arr {
		fmt.Println("i",i, "v", v)
	}
}

func partition1(arr[] int, left int, right int) int{
	pivot := arr[left]
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}

		swap(arr, left, right)
		for left < right && arr[left] <= pivot {
			left++
		}
		swap(arr, left, right)

	}
	return left
}

func partition(arr[]int, left int, right int) int {
	pivot := arr[left]
	for left < right {
		for left < right && arr[right] >= pivot{
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= pivot {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	return left
}
func swap(arr[]int, a, b int){
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

```

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