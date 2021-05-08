---
title: "go timer"
date: 2021-05-05T18:34:48+08:00
draft: false
isCJKLanguage: true
tags: ["golang"]
series: [""]
categories: ["技术"]
---

Go提供了两种定时器，分为一次性定时器、周期性定时器。 两种定时器内部实现机制完全相同
src/runtime/time.go:addtimerLocked()函数负责添加timer

+ 一次性定时器：定时器只计时一次，结束便停止  Timer只执行一次就结束
+ 周期性定时器：定时器周期性进行计时  周期性定时器Ticker

```golang
type Timer struct { // Timer代表一次定时，时间到来后仅发生一个事件。
    C <-chan Time
    r runtimeTimer
}
```

```golang
//src/time/tick.go:Ticker
type Ticker struct {
    C <-chan Time // The channel on which the ticks are delivered.
    r runtimeTimer
}
```


系统协程把runtimeTimer存放在数组中，并按照when字段对所有的runtimeTimer进行堆排序，定时器触发时执行runtimeTimer中的预定义函数f
即完成了一次定时任务

```golang
//src/time/sleep.go:runtimeTimer定义了其数据结构
type runtimeTimer struct {
    tb uintptr                          // 存储当前定时器的数组地址
    i  int                              // 存储当前定时器的数组下标

    when   int64                        // 当前定时器触发时间
    period int64                        // 当前定时器周期触发间隔
    f      func(interface{}, uintptr)   // 定时器触发时执行的函数
    arg    interface{}                  // 定时器触发时执行函数传递的参数一
    seq    uintptr                      // 定时器触发时执行函数传递的参数二(该参数只在网络收发场景下使用)
}
```

```golang
func NewTimer(d Duration) *Timer {
    c := make(chan Time, 1)  // 创建一个管道
    t := &Timer{ // 构造Timer数据结构
        C: c,               // 新创建的管道
        r: runtimeTimer{
            when: when(d),  // 触发时间
            f:    sendTime, // 触发后执行函数sendTime
            arg:  c,        // 触发后执行函数sendTime时附带的参数
        },
    }
    startTimer(&t.r) // 此处启动定时器，只是把runtimeTimer放到系统协程的堆中，由系统协程维护
    return t
}
```


```golang
// 设定超时时间
func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(1 * time.Second)
	select {
	case <-conn:
		timer.Stop()
		return true
	case <-timer.C: // 超时
		println("WaitChannel timeout!")
		return false
	}
}
```


```golang
//After()
func AfterDemo() {
	log.Println(time.Now())
	<-time.After(1 * time.Second)
	log.Println(time.Now())
}

```

```golang
//AfterFunc()
func AfterFuncDemo() {
	log.Println("AfterFuncDemo start: ", time.Now())
	time.AfterFunc(1*time.Second, func() {
		log.Println("AfterFuncDemo end: ", time.Now())
	})

	time.Sleep(2 * time.Second) // 等待协程退出
}

```


周期性ticker
```golang
// TickerDemo 用于演示ticker基础用法
func TickerDemo() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Ticker tick.")
	}
}

```
