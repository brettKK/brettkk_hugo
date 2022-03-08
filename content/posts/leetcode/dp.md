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