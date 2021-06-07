---
title: "数组"
date: 2021-04-05T11:33:56+08:00
draft: false
isCJKLanguage: true
markup: mmark
tags: ["leetcode"]
series: [""]
categories: ["life"]
---

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

