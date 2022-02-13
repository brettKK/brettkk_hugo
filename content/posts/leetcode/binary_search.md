---
title: "二分查找"
date: 2021-04-05T11:33:56+08:00
draft: true
isCJKLanguage: true
markup: mmark
tags: ["leetcode"]
series: [""]
categories: ["life"]
---

> Donald Ervin Knuth（高德纳）
> Although the basic idea of binary search is comparatively straightforward, the details can be surprisingly tricky





```
index  0 1 2 3 4 5 6 7
data   1 2 3 5 5 5 8 9
```

+ 找第一个大于等于5的元素index （3）
+ 找第一个大于5的元素index （6）
+ 找最后一个小于等于5的元素index （5）
+ 找最后一个小于5的元素index （2）

---

```
0 1 2 3 .. k-1 ]   [ k, k+1, ...N
    white                 black
```

``` golang
left = -1
right = N
for (left + 1) != right {
    m = (left + right )/2
    if isWhite(m) {
        left = m
    } else {
        right = m
    }
}
return left or right
```
---

基本写法
```golang
func binSearch(arr []int, k int) int {
    left, right := 0, len(arr) - 1
    for left <= right {
        mid := (left+right) / 2
        if arr[mid] == k{
            return mid
        } else if arr[mid] < k {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}

```

```golang
func binSearchRecursive(arr []int, target int, left, right int) int {
    if left > right {
        return -1
    }
    mid := (left+right)/2
    if target == arr[mid] {
        return mid
    }
    if target > arr[mid] {
        return binSearchRecursive(arr, target, mid+1, right)
    }
    if target < arr[mid] {
        return binSearchRecursive(arr, target, left, mid-1)
    }
    return -1
}  
```

```golang
func Search(n int, f func(int) bool) int {
    i, j := 0, n
    for i < j {
        h := int(uint(i+j)>>1)
        if !f(h) {
            i = h + 1
        } else {
            j = h
        }
    }
    return i
}

func main() {
    nums := []int64{2, 3,5,6,7,1,2,3}
    sort.Search(len(nums), func(i int) bool {
        return nums[i] == 1
    })
}

```


### 找一个元素第一次出现的位置和最后一次出现的位置

```golang
// [1,3,3,3,5,7] target=3
func leftFirst(arr []int, int target) int {
    left, right := 0, len(arr) - 1
    for left < right {
        mid := (right+left) / 2
        if nums[mid] == target {
            // 当找到时，收缩右边界
            right = mid
        } else if nums[mid] > target {
            right = mid 
        } else {
            left = mid + 1
        }
    }
    return left
}
func rightFirst(arr []int, int target) int {
    left, right := 0, len(arr) - 1
    for left < right {
        mid := (right + left) / 2
        if nums[mid] = target {
            // 当找到时，收紧左边界
            left = mid + 1
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left - 1
}

```


```golang
func bsFirst(arr []int, k int) {
    left, right := 0, len(arr) - 1
    for left < right {
        mid := (left+right)/2
        if arr[mid] < k {
            left = mid + 1 
        }else {
            right = mid
        }
    }
    return left
}
func bsLast(arr []int, k int) {
    left, right := 0, len(arr) - 1
    for left < right {
        mid := (left+right)/2
        if arr[mid] > k {
            right = mid - 1
        }else {
            left = mid
        }
    }
    return left
}

```

### 旋转数组中找一个元素的位置

``` golang
func search(arr []int, k int) int{

}


```

{{< youtube yTXCPGAD3SQ >}}