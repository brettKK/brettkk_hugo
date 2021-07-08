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