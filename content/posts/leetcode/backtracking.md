---
title: "dfs & backtracking"
date: 2022-01-13T11:33:56+08:00
draft: true
isCJKLanguage: true
tags: ["leetcode"]
series: [""]
categories: ["life"]
---

## 全排列，组合，子集

### 求数组的全部子集

#### 数组没有重复元素的情况下，输出所有子集合

```golang
var res []string
func getAllSet(arr []byte) []string {
    dfsYesNo(arr, 0, []byte{})
    // print res

    dfs(arr, 0, []byte{})
    // print res
    return res
}

func dfsYesNo(arr []byte, position int, tempArr []byte) {
    if start == len(arr) -1 {
        res = append(res, tempArr)
        return
    }
    for i := position; i < len(arr); i++ {
        //选 i位置的元素
        tempArr = append(tempArr, arr[i])
        dfsYesNo(arr, i + 1, tempArr)
        // 不选i位置的元素
        tempArr = tempArr[:len(tempArr)-1]
        dfsYesNo(arr, i + 1, tempArr)
    }
}

func dfs(arr []byte, start int, tempArr []byte) {
    if start < len(arr) {
        res = append(res, temp)
    }
    for i := start; i < len(arr); i++ {
        temp = append(temp, arr[i])
        dfs(arr, i+1, temp)
        temp = temp[:len(temp)-1]
    }
}
```

#### 数组里有重复元素时，输出所有子集合

res 换成map去重



TODO 