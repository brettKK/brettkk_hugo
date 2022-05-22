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

## Dynamic programming

题型：优化问题（min, max）, 计数问题（有多少种方式）。

种类：矩阵坐标， 序列型（LIS， LCS）， Knapsack背包（coin， sub set sum）, 区间型。

寻找递推公式（大问题依赖于之前的小问题）

+ 套路
    + 第I类基本类型。单序列I (时间序列型，今天的状态只依赖昨天的状态 )
        + lc_198, lc_213, lc_123, lc_309, lc_376, lc_276
        + 权利， lc_487,  lc_1186, 
    + 第II类基本类型。单序列II  （时间序列型，今天的状态依赖于之前的某一个状态）
        + lc_300, lc_368, lc_1105
    + 双序列型
        + lc_1143, longest common subsequences
        + lc_1092, shortest common suppersequences
        + lc_72, edit distance
        + lc_97, lc_115, lc_727, lc_583, lc_712, lc_1035, lc_1216, lc_1312
    + 第I类区间型 （分K个连续区间，计算这些区间的最优性质）
        + 状态定义： dp[i][k]表示为s[1:i]f分成K个区间得到的最优解
        + 搜索最后一个区间的起始位置j, 将dp[i][k] 分小为dp[j-1][k-1] 和 s[j:i]两部分
        + 最优的结果 dp[N][K]
        + lc_1278, lc_813, lc_410, lc_1335, 
    + 第II类区间型 
        + lc 516, lc_312, lc_375, lc_1246, lc_1000
    + 背包问题 (背包九讲)
        + lc_494, lc 518, lc_1049, lc_474, lc_879, lc_956
    + 状态压缩 
        + lc_691, lc 1349, lc_943
    + others
        + lc_887, lc_920

### knapsack without repetition (0-1背包)

input:   v[i] i=1..n, ith item 的value；
 w[i] i=1..n, ith item 的weight；
 C 背包的最大承受的weight。
 note: only one copy of each item at most。

 output： max value

分析： naive 把所有的组合都试一次。 时间复杂度为O（2^n）。

两个维度：f(i,j), i item {1,2,..i}给你选，背包的重量为j。


```golang

func knapsack01(v []int, w []int, C int) int {
    dp := make([][]int, len(v)+1)
    for i := 0; i <= len(v); i++ {
        dp[i] = make([]int, C+1)
        dp[i][0] = 0
    }
    for i := 0; i <= C; i++ {
        dp[0][i] = 0
    }
    for i := 1; i <= len(v); i++ {
        for j := 1; j <= C; j++ {
            dp[i][j] = dp[i-1][j]  
            if w[i-1] <= j {
                dp[i][j] = max(dp[i][j], dp[i][j-w[i-1]] + v[i-1])
            }
        }
    }
    return dp[len(v)][C]
}

```

### knapsack with repetition

note: infinite copies of each item. output: max value.


item的重量是小数时怎么办？ 乘于factor变为整数。

```golang
func knapsack(v []int, w []int, C int) int {
    dp := make([]int, C + 1)
    for i := 1; i <= C; i++ {
        for j := 0; j < n; j++ {
            if i >= w[j] {
                dp[i] = max(dp[i], dp[i-w[j]] + v[j])
            }
        }
    }
    return dp[C]
}

LC 322 换零钱。

2, 5, 7 三种硬币，能否拼出X元。如果能拼出，最少用多少个枚硬币拼出X

f(x) 最少用多少枚硬币拼出X。 f(x) = min(f(x-2) + 1, f(x-5) + 1, f(x-7) + 1)

```golang

func change_coins(coins []int, X int) int {
    dp := make([]int, X + 1)
    dp[0] = 0
    for i := 1; i <= X; i++ {
        dp[i] = MAX_INT
        for _, j := range coins {
            if i >= j && dp[i-j] != MAX_INT {
                dp[i] = min(dp[i-j]+1, dp[i])
            }
        }
    }
    return dp[X]
}

```



### 机器人

机器人从左上角，每次可下或者可右一步，有多少种不同的方式走到右下角

```golang
func unique_paths(m, n int) int {
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    for i := 1; i < m; i++ {     //row: top to bottom
        for j := 1; j < n; j++ { //column: left to right
            if i == 0 || j == 0 {
                // 在边上，只有一条path
                dp[i][j] = 1
            }
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[m-1][n-1]
}

```
### 青蛙跳

```goalng
//dp[x] = or (0-x 所有可能性)

func can_jump(arr []int) bool {
    dp := make([]bool, len(arr))
    dp[0] = true
    for i := 1; i < len(arr); i++ {
        // 枚举i之前的所有可能的石头。
        for j := 0; j < i; j++ {
            if j + arr[j] >= i && dp[j] == true {
                dp[i] = true
                break
            }
        }
    }
    return dp[len(arr)-1]
}
```


斐波那契， 连续子数组最大和lc 53，连续子数组最大乘积lc 152， 最长递增子序列lc 300， 最长公共子序列lc 1143，最长公共子串，不同子序列lc 115

### 斐波那契

```golang

func fib(int n) {
    dp := make([]int, n+1)
    dp[0] =0
    dp[1] = 1
    for i := 2; i < n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}

```

### 连续子数组最大和 lc 53

```golang
// dp[i] 以第i个结尾的最大子数组和
func max_sub_array_sum(nums []int) int{
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    var max int
    for i := 1; i < len(nums); i++ {
        // 实际上是在判断dp[i-1] 是否小于0，并没有直接使用dp之前的数据，所以也可以不用dp
        dp[i] = max(dp[i-1]+nums[i], nums[i])
        if dp[i] > max{
            max = dp[i]
        }
    }
    return max
}
```

### 连续子数组最大乘积 lc 152

数组中出现负数， 最大子数组 可能由负负得正。

开两个dp数组，dp——max 记录乘积的最大值，dp-min记录乘积的最小值。

```golang
dpmax[i] = max(dpmax[i-1]*nums[i], dpmin[i-1]*nums[i], nums[i])
dpmin[i] = min(dpmax[i-1]*nums[i], dpmin[i-1]*nums[i], nums[i])

```

### 最长递增子序列 (lc 300 longest increasing subsequence)

当前的dp[i]与dp[0 ～ i-1]的每一个状态均有关系。
当前的dp[i]只与dp[i-1]的状态有关系。


```golang
// dp[i] 为以nums[i]结尾的最长递增子序列的长度。
func max_lis(nums []int) int {
    dp := make([]int, len(nums))
    max_len := 1
    dp[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        temp_max := 0
        for j := 0; j < i; j++ { // o(n2) 这里可以优化
            // 统计i之前，结尾数字小于当前i数字，且长度为记录的最长
            if nums[j] < nums[i] && dp[j] > temp_max {
                temp_max = dp[j]
            }
        }
        dp[i] = temp_max + 1
        if dp[i] > max_len{
            max_len = dp[i]
        }
    }
    return max_len
}

```

similar: lc 673. number of longest increasing subsequence


### largest divisible subset lc 368

给定一个数组S。求最大的子集，使得里面的所有元素之间都可以相互整除。

状态定义： dp[i]以S[i]结尾，满足限制的最大子集的元素数量。

状态转移： 在0-i之间寻找最优的前驱状态j，将dp[i] 与dp[j] 产生联系。

```golang

func get_largest(nums []int) int {
    sort(nums)
    for i := 1; i < len(nums); i++ {
        for j := 1; j < i; j++ {
            if nums[i] % nums[j] == 0 {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
    }
    ans := max(dp[i]) // i = ...
    return ans
}

```

### filling bookcase shelves lc 1105

给定N本书（宽度高度各异）的序列，要求按照所给出的顺序摆放。相邻的若干本书可以放一层，但同一层的宽度不能超过W。
问书架的最矮可以有多高？

将数组分成若干个sub array。 min（每个subarray的最大值之和）。

第i本书所在的这一层有多高？取决于上一层的最后一本书在哪儿。


```golang

// width[i],height[i] 表示第i本书的宽度和高度。
// dp[i] 表示0-i本书放到书架上的最矮高度
func min_height(width []int, height []int, limit_width int) int {
    dp := make([]int, len(width))
    for i := 0; i < len(width); i++ {
        // 寻找上一层最后一本书的位置。
        for j := 0; j < i; j++ {
            if get_width(width, j+1, i) <= limit_width {
                dp[i] = min(dp[i], dp[j] + max_height(height, j+1 , i))
            }
        }
    }
    return dp[len(width) - 1]
}

```


### 最长公共子序列lc 1143 longest common subsequences

给定2个字符串，返回2个字符串的最长公共子序列的长度lcs。



```golang
// dp[i][j] 表示 以text1以第i个结尾，text2以第j个结尾的最长公共子序列长度。
// 状态转移方程：
// text1[i] == text2[j]  dp[i][j] = dp[i-1][j-1] + 1
// text1[i] != text2[j]  dp[i][j] = max(dp[i-1][j], dp[i][j-1])

func max_lcs(s1 []byte, s2 []byte) int {
    dp := make([][]int, len(s1) + 1)
    for i := 0; i < len(s1); i++ {
        // 2行 也行
        dp[i] = make([]int, len(s2)+1)
    }
    for i := 0; i < len(s1); i++ {
        for j := 0; j < len(s2); j++ {
            if s1[i] == s2[j] {
                dp[i+1][j+1] =dp[i][j] + 1
            } else {
                dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
            }
        }
    }
    return dp[len(s1)][len(s2)]
}

```

### 最长公共子串

之前是给一个数组，可以给2个数组。

子串： 连续。
给定2个字符串，返回2个字符串的最长公共子串的长度。

// text1[i] == text2[j] dp[i][j] = dp[i-1][j-1] + 1
// text1[i] != text2[j] dp[i][j] = 0

### 不同子序列lc 115

给定一个字符串s和一个字符串t，计算在s的子序列中t出现的次数。

s的子序列是指通过删除一些字符得到的新字符串。

```golang



```

### 最小编辑距离 lc 72 edit distance

编辑： 插入，删除，替换
// text1[i] == text2[j] dp[i][j] = dp[i-1][j-1]
// text1[i] != text2[j] dp[i][j] = 1 + min(dp[i-1][j] + dp[i][j-1] + dp[i-1][j-1])

```golang

func get_min_change(s1 []byte, s2 []byte) int {
    dp := make([][]int, len(s1) + 1)
    for i := range len(s1) + 1{
        dp[i] = make([]int, len(s2) + 1)
        for j := range len(s2) + 1 {
            if i == 0 {
                dp[0][j] = j
            }
            if j == 0 {
                dp[i][0] = i
            }
        }
    }
    for i := range (1,len(s1)) {
        for j := range (1,len(s2)) {
            if s1[i-1] == s2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            }else {
                dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
            }
        }
    }
    return dp[len(s1)][len(s2)]
} 

```


### largest sum of averages lc 813


## 第II类区间型

### longest palindromic subsequence lc 516

给定一个string， 求是回文串的最长subsequence的长度

状态定义： dp[i][j]表示为字符串S[i:j]里是回文串的最长subsequence的长度。

状态转移：第一层循环是区间大小；第二层循环是起始点；想办法把大区间的dp[i][j]往小区间上转移。

```golang
func longest_palin(s []byte) int {
    for i := 0; i < len(s); i++ {
        for j := 0; i + j < len(s); j++ {
            k := j + i
        }
    }
}
```

### burst balloons lc 312

给一排气球及其价值S。每戳爆一个气球的得分： 气球本身的分值*（左边存留的气球分值）* （右边存留的气球分值）。
如何戳破气球，得到max score。

dp[i][j] = max(dp[i][j], dp[i][k-1] + s[i-1] * s[k] * s[j+1] + dp[k+1][j])


### guess number higher or lower II (lc 375)

### palindrome removal (lc 1246)

### minimum cost to merge stones (lc 1000)

结合第I类和第II区间型DP算法的hard题。


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


### 打家劫舍 lc 198

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
#### 打家劫舍 （首尾相连）lc 213

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

### lc 123 

### wiggle subsequence lc 376

fun longest_wiggle_subsequence(nums []int) int{
    // dp rows = len(nums) + 1, cols = 2
    // dp[i][0] 以当前元素结尾且上升
    // dp[i][1] 以当前元素结尾且下降
    dp := make([][]int, len(nums))
    if i = 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            dp[i][0] = dp[i-1][1] + 1
        }else {
            dp[i][1] = dp[i-1][0] + 1
        }
    }
    return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

### paint fence lc 276

给出cost[i][j]表示第i个房子喷第j种颜色的价格。限制相邻的房子不能涂相同的颜色。求喷涂完所有房子的最小价格？

```golang

// dp[i][j] 喷涂从0到i的房子，且第i个房子为颜色j的价格。
// dp[i][j] = min(dp[i-1][!j] + cost[i][j])
// result = min(dp[N][j] j=...)

func min_cost(cost [][]int) int {
    house_nums, color_nums:= len(cost), len(cost[0])
    dp = make([][]int, len(house_nums))
    for i := 0; i < house_nums; i++ {
        dp[i] = make([]int, len(color_nums))
        dp[i][0] = cost[i][0]
    }
    for i := 0; i < color_nums; i++ {
        dp[0][i] = cost[0][i]
    }
    for i := 1; i < house_nums; i++ {
        for j := 1; j < color_nums; j++ {
            var temp_min int
            for k := 1; k < color_nums && k != j; k++ {
                if dp[i-1][k] + cost[i][j] < temp_min {
                    temp_min = dp[i-1][k] +cost[i][j]
                }
            }
            dp[i][j] = temp_min
        }
    }
    var min int
    // min dp[house_nums][...]
    for j := 0; j < color_nums; j++ {
        if dp[house_nums-1][i] < min {
            min = dp[house_nums-1][i]
        }
    }
    return min
}

```
similar: mininum falling path sum II.

### max consecutive ones II  lc 487

给定一个0/1的数组，最多有一次从0翻转为1的权利。问最多可以有多少连续的1？


dp[i][0]: 以当前元素为结尾，未使用翻转权利的最长连续1的个数
dp[i][1]: 以当前元素为结尾，  使用翻转权利的最长连续1的个数

分析：

第i-1轮： 以当前元素结尾且没有使用权力的最长连续1 （dp[i-1][0]） ，  以当前元素结尾且使用权利的最长连续1 (dp[i-1][1])
之间的关系，
第i轮：   以当前元素结尾且没有使用权力的最长连续1 dp[i][0]，  以当前元素结尾且使用权利的最长连续1(dp[i][1])


```golang

func get_ones_nums(arr []int) int {
    dp := make([][2]int, len(arr))
    if arr[0] == 0 {
        dp[0][0] = 0
        dp[0][1] = 1
    } else {
        dp[0][0] = 1
        dp[0][1] = 1
    }
    for i := 1; i < len(arr); i++ {
        if arr[i] == 0 {
            // 需要使用翻转权利
            dp[i][0] = 0
            dp[i][1] = dp[i-1][0] + 1

        } else {
            // 不需要使用翻转权利
            dp[i][0] = dp[i-1][0]+ 1
            dp[i][1] = dp[i-1][1] + 1
        }
    }
    // for all dp[i][j] to get max value
    // not min(dp[len(arr)-1][0], dp[len(arr)-1][1])
    return get_min_value(dp)
}

```

### maximum subarray sum with one deletion lc 1186

给定一个数组，最多有一次删除一个数的权利。问sum最大的subarray？

涉及到权利状态时放到dp的维度里去。

dp[i][0]: 以当前元素为结尾且未使用权利的最大sum
dp[i][1]: 以当前元素为结尾且使用权利的最大sum

dp[i][0] = max(dp[i-1][0] + arr[i], arr[i])
dp[i][1] = max(dp[i-1][0], dp[i-1][1] + arr[i])




### 熊出没

给定一个装满苹果的山洞，只能从洞口和洞尾拿苹果。 每轮两只熊交替拿，最后谁拿到的重量大谁赢。

```golang
// input: 1kg, 3kg, 14kg, 7kg

// 若每轮大熊贪心拿最大的： 7kg + 3kg = 10kg; 熊二 14kg + 1 kg = 15kg

// 没想明白..

// dfs 

func CanFirstBearWin(nums []int) bool {
    return picker(0, len(nums)-1, nums) >= 0
}
func picker(i, j int, nums []int) int {
    if i == j {
        return nums[i]
    }
    head := nums[i] - picker(i + 1, j, nums)
    end := nums[j] - picker(i, j-1, nums)
    return max(head, end)
}

// dp
func CanFirstBearWin(nums []int) bool {
    length := len(nums)
    dp := make([][]int, length)
    for i:=0; i<length; i++ {
        dp[i] = make([]int, length)
        dp[i][i] = nums[i]
    }
    for j := 1; j <length; j++ {
        for i := j - 1; i >= 0; i-- {
            head := nums[i] - dp[i+1][j]
            end := nums[j] - dp[i][j-1]
            dp[i][j] = max(head, end)
        }
    }
    return dp[0][length-1] >= 0
}

```

### 