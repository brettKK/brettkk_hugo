---
title: "数组"
date: 2021-04-05T11:33:56+08:00
draft: true
isCJKLanguage: true
markup: mmark
tags: ["leetcode"]
series: [""]
categories: ["life"]
---

{{< highlight go "linenos=inline" >}}
func search(arr []int, k int) int{

}
{{< / highlight >}}

### 双指针解法

remove duplicates from sorted array

```golang
func Remove(arr []int) int {
    i, j := 0, 0
    for j < len(arr) {
        if i == 0 || arr[j] != arr[j-1] {
            arr[i++] = arr[j]
        } else {
            j++
        }
    }
    return i
}

```
同 11， 42， 283， 80， 1047


### 最大子序和


### 最大的连续子数组和 lc 53

```golang
func get_max_sum(arr []int) int{
    result := min_int
    count := 0
    for i := 0; i < len(arr); i++ {
        count += arr[i]
        if count > result {
            result = count
        }
        if count <= 0 {
            count = 0
        }
    }
    return result
}

func get_max_sum(arr []int) int {
    var sum int
    for i := 0; i< len(arr); i++ {
        if sum + arr[i] < 0 {
            continue
        }
        if arr[i] < 0 {
            sum = max(sum + arr[i], 0)
        }
        sum = sum + arr[i]
    }
    return sum
}


```

```golang
func get_max_sum(arr []int) int {
    dp := make([]int, len(arr)+1)
    dp[0] = arr[0]
    result := dp[0]
    for i := 1; i < len(arr); i++ {
        dp[i] = max(dp[i-1]+arr[i], arr[i])
        if dp[i] > result {
            result = dp[i]
        }
    }
    return result
}
func maxSubArray(nums []int) int{
    maxSum := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i-1] + nums[i] > nums[i] {
            nums[i] += nums[i-1]
        } 
        if nums[i] > maxSum {
            maxSum = nums[i]
        }
    }
    return maxSum 
}

```

### 对一维数组做max pooling

pooling(池化)目的： 压缩张量的规模

```golang
func get_max_pooling(arr []int, k int) []int{
    var deque []int
    var res []int
    for i := 0; i < k; {
        for len(deque) != 0 && deque[len(deque) - 1] < arr[i]{
            // 维护 双端队列，左大 右小 
            // pop 右边小的item
            deque = deque[:len(deque)-1]
        }
        deque = append(deque, i)
    }
    for i := k; i < len(arr); i++ {
        for len(deque) != 0 &&  i - deque[0] >= k {
            // pop 左边超过window区间的item
            deque = deque[1:]
        }
        for len(deque) != 0 && deque[len(deque)-1] < arr[i] {
            // pop 右边小的item
            deque = deque[:len(deque)-1]
        }
        deque = append(deque, i)
        res = append(deque[0])
    }
    return res
}
```

### 求数组中比左边都大，比右边都小的元素，返回这些元素的索引

```golang
// 要求O(n)
func get_index(arr []int) []int {
    var result []int
    leftMax, rightMin := make([]bool, len(arr)), make([]bool, len(arr))
    leftMax, rightMin := MIN_int, MAX_int
    for i := 0; i < len(arr); i++ {
        if i == 0 {
            leftMax[i] = false
            rightMin[len(arr)-i-1] = false
        }
        if arr[i] > leftMax {
            leftMax[i] = true
            leftMax = arr[i]
        }
        if arr[len(arr)-i-1] < rightMin{
            rightMin[i] = true
            rightMin = arr[len(arr)-i-1]
        }
    }
    for i := 0; i < len(arr); i++ {
        if leftMax[i] == rightMin[i] {
            result = append(result, i)
        }
    }
    return result
}

### 旋转矩阵

```golang
func rotation(row, col int) [][]int {
    top, bottom, left, right := 0, row -1, 0, col -1
    i, j := 0
    temp := 1
    result := make([][]int, row)
    for index := range row{
        result[index] = make([]int, col)
    }
    for i >= top && i <= bottom && j >= left && j <= right {
        // left -> right 
        for j <= right {
            result[i][j] = temp
            temp++
            j++
        }
        j--
        top++
        i = top
        // top -> bottom
        for i <= bottom {
            result[i][j] = temp
            temp++
            i++
        }
        i--
        right--
        j = right
        // right -> left
        for j >= left {
            result[i][j] = temp
            temp++
            j--
        }
        j++
        bottom--
        i = bottom
        // bottom -> top
        for i >= top {
            result[i][j] = temp
            temp++
            i--
        }
        i++
        left++
        j = left
    }
    return result
}

```


### 数组中出现次数超过一半的数

数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

分析1: 对数组进行快排，位于数组中间的数字一定是出现次数超过了长度一半的数字。
分析2: 利用map<number, counter> ，遍历数组，发现counter>=len/2时 返回number。
分析3: 最优算法。在遍历数组的时候保存两个值：一个candidate，用来保存数组中遍历到的某个数字；一个nTimes，表示当前数字的出现次数。
有符合条件的数字，则它出现的次数比其他所有数字出现的次数和还要多

```golang

func find_more_half_num(nums []int) int {
    candidate, times := nums[0], 1
    for i := 1; i < len(nums); i++ {
        if times == 0 {
            candidate = nums[i]
            times = 1
        } else {
            if candidate == nums[i] {
                times += 1
            } else {
                times -= 1
            }
        }
    }
    return candidate
}

```



---


https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-858-computer-systems-security-fall-2014/lecture-notes/

https://blog.gopheracademy.com/advent-2018/llvm-ir-and-go/

https://xargin.com/

http://www.nondot.org/sabre/

http://www.aosabook.org/en/index.html

https://gohalo.me/about.html

https://github.com/cloudnativeapp/meetup/tree/master/Kubernetes%20Course

https://www.zhihu.com/column/c_1195294063723929600

https://lihaoquan.me/posts/k8s-crd-develop/

https://jimmysong.io/kubernetes-handbook/concepts/

https://github.com/Junedayday/code_reading

https://www.cnblogs.com/yahuian

tls, quic

架构 功能 演进

数据存储 ，访问方式

搭建业务通路

稳定性 --->减少组建依赖


表设计
+ 元数据表
+ 关系表
+ 数据表