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