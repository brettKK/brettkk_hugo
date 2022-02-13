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


### 归并

稳定排序的算法。

```golang
package main

import "fmt"

func main () {
	inputArr := []int{80, 38, 97, 65, 76, 13, 27}
	fmt.Println("before sort", inputArr)
	outputArr := MergeSort(inputArr, 0, len(inputArr) -1)
	fmt.Println("end sort", outputArr)
}

func MergeSort(arr []int, left, right int) []int {
	if left > right {
		return nil
	}
	if left == right {
		return []int{arr[left]}
	}
	mid := (left + right) / 2
	arr1 := MergeSort(arr, left, mid)
	arr2 := MergeSort(arr, mid+1, right)
	return mergeSortArray(arr1, arr2)
}

func mergeSortArray(arr1, arr2 []int) []int {
	resArr := make([]int, len(arr1) + len(arr2))
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			resArr[k] = arr1[i]
			k++
			i++
		} else {
			resArr[k] = arr2[j]
			k++
			j++
		}
	}
	for i < len(arr1) {
		resArr[k] = arr1[i]
		k++
		i++
	}
	for j < len(arr2) {
		resArr[k] = arr2[j]
		k++
		j++
	}
	return resArr
}

```

### 堆排序

思路：
> 将待排序的数组构造为大顶堆

> 堆顶元素与数组的末尾数交换，剩余带排序的数据为n-1

> 将n-1的数组构造为大顶堆，循环执行


```golang
func heapUp(heap []int, int k) {
	var son, parent, x int
	x = heap[k]
	son = k
	parent = (son - 1) / 2
	for son > 0 {
		if heap[parent] >= x {
			break
		}
		heap[son] = heap[parent]
		son = parent
		parent = (son - 1) / 2
	}
	heap[son] = x
}

func heapDown(heap []int, k int, n int) {
	x := heap[k]
	parent := k
	son := 2 * k + 1
	for son >=0 {
		if son + 1 <= n && heap[son + 1] > heap[son] {
			son++
		}
		if x >= heap[son] { 
			break
		} 	
		heap[parent] = heap[son]
		parent = son
		son = parent * 2 + 1
	}
	heap[parent] = x
}

func buildHeap(heap []int, int n) {
	for i := (n - 1) / 2; i>= 0 ;i-- {
		heapDown(heap, i, n)
	}
}

func heapsort(heap []int, n int) {
	for i := n; i > 0; i-- {
		swap(heap, 0, i)
		heapDown(heap, 0, i - 1)
	}
}

```