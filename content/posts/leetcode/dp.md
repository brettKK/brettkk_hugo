---
title: "动态规划"
date: 2021-04-05T11:33:56+08:00
draft: true
isCJKLanguage: true
markup: mmark
tags: ["leetcode"]
series: [""]
categories: ["life"]
---

### 0-1背包问题

```golang
// dp[i][j] : 背包重量为j时，从0～i种物品选入装入背包的价格最大值
func bag01(weight, value []int, bagweight int) int {
   dp := make([][]int, len(weight))
   for i, _ := range dp {
       dp[i] = make([]int, bagweight + 1)
   } 
   for i := bagweight; i >= weight[0]; i--  {
       dp[0][i] = dp[0][i-weight[0]] + value[i]
   }
   for i := 1; i < len(weight); i++ {
       for j := 0; j <= bagweight; j++ {
           if weight[i] <= j {
               dp[i][j] = max(dp[i-1][j - weight[i]] + value[i], dp[i-1][j]
           }else {
               dp[i][j] = dp[i-1][j]
           }
       }
   }
   return dp[len(weight)-1][bagweigth]
}

```

### next 


### 打家劫舍

相邻的房间不能偷，能获取的最大金额

```golang
// input [2, 7, 9, 3, 1] 
// output 2 + 9 + 1 = 12

// dp[i] 前 i 间房屋能偷窃到的最高总金额

dp[0] = nums[0]
dp[1] = nums[1]
dp[k] = max(dp[k-1], dp[k-2] + nums[i])
// 也可以不申请dp数组， 因为dp[i] 只与前两项有关
```
#### 打家劫舍 （首尾相连）

```golang
// 不选择第一间， [1 ~ n]
// 不选择最后一间，[0 ~ n-1]

func getMax(nums []int) int {
    return max(robRange(nums, 0, len(nums-2)), robRange(nums, 1, len(nums-1)))
}

func robRange(nums []int, left int, right int) {
    first, second := nums[left], max(nums[left], nums[left+1])
    for i := left + 2; i <= right; i++ {
        temp := max(first + nums[i], second)
        first = second
        second = temp
    }
    return second
}

```

