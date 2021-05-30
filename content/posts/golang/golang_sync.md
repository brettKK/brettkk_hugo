---
title: "go同步编码方式"
date: 2021-05-06T18:34:48+08:00
draft: false
isCJKLanguage: true
tags: ["golang"]
series: [""]
categories: ["技术"]
---
golang提供了比较便捷的并发编程方式。golang的并发单元是goroutine。目前实现多个goroutine同步的主要方式有： sync包， runtime包下的channel，context包。

### sync.WaitGroup

标准库sync里的Waitgroup，用来阻塞主协程去等待所有协程执行完。WaitGroup主要三个方法：Add方法添加等待的协程数量，Done方法等于Add(-1)减少等待的协程数量，Wait方法阻塞主协程。原理示意图如下：

![image](/golang/golang_sync.png))


#### WaitGroup使用说明
WaitGroup在工作中经常会使用到。例如：需要并行处理一些子任务，需要起多个协程对一些复杂结构体的各个字段进行初始化等。简单的使用说明如下：
```golang
func fun() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        //process
       wg.Done() 
    }()
    wg.Wait()
}

```

### channel

通道是golang中协程通信的一种方式，可以灵活、高效支持协程之间的通信。通道与协程之间的交互流程图因是否存在缓冲区、缓冲区是否填满、是否存在发送等待队列等，会在过程上会存在不同。下面给出的示意图是针对利用无缓冲通道进行通信的其中一种情况，主要是为了简化流程和突出主要功能点。图中hchan是通道的底层数据结构，其中主要包含： buf为数据存储区， recvq和sendq为在此通道上等待接受和发送的协程队列，lock为保护通道数据的锁等。P为运行时的协程调度器。原理示意图如下：

![image](/golang/hchan.png)

#### channel 使用说明

通过无缓冲的通道，阻塞一个协程，等待另一个协程唤醒。例如：main协程在通道的接受队列中等待， 等worker协程写入数据后，main协程才会唤醒并继续执行。通信的数据是直接写一个协程的stack。这样可以省去去操作通道的buf，减少内存拷贝。简单的使用说明如下：

```golang
func worker(done chan bool) {
    //process
    done <- true
}
func main(){
    done := make(chan bool)
    go worker(done)
    <- done
}
```

channel读写超时的处理。Go 语言的 channel 本身是不支持 timeout 的，一般实现 channel 的读写超时采用 select + time的方式。实例说明如下：

```golang
func main() {
    var ch chan int
    go func(){
        //process   
        ch<-1
    }()
    select{
    case <-time.After(time.Second):
        return
    case <-ch 
        //process
    }
}
```


#### channel 错误使用示例

```

// 死锁
func main() {
    ch := make(chan int)
    <- ch //block
    go func() {
        ch <- 1
    }
}

// 死锁
func main() {
    ch := make(chan int)
    ch <- 1 // block
    go func() {
        <-ch
    }
}

// 死锁
func main() {
        var ch chan int  //nil channel

        go func(c chan int) {
                for v := range c { // <-ch  channel read
                        fmt.Println(v)
                }
        }(ch)

        ch <- 1  // channel write
}

```

```
// goroutine泄漏
func main() {
    ch := make(chan int)
    go func() {
        // process more than 100ms
        ch <- 1
    }
    select {
    case <-ch:
        // process done
        return    
    case <- time.After(100*time.Millisecond) 
        // process timeout
        return
    }
}

```

错误操作引起Panic 

```
// send on close channel 往已关闭的 channel 写数据会 panic
// closing on close channel  写入端无法知道 channel 是否已经关闭
// send or read on nil channel channel初始化使用错误

1. 不要尝试在读取端关闭 channel 
2. 一个写入端，在这个写入端可以放心关闭 channel
3. 多个写入端时，不要在写入端关闭 channel 
```

### context

context是在context包里定义的一种接口。在context包中可以通过WithCancel、WithDeadline、WithTimeout、WithValue这四个方法生成新 Context。WithValue方法生成的context可以做到了跨API的传值功能。WithTimeout方法是对WithDeadline的封装。WithDeadline方法生成的context实现了定时cancel的功能。

下面以WithCancel为入口，以源码（go1.12.7）为内容作介绍。当 parent 的 Done() 关闭的时候，孩子 ctx 的 Done() 也会关闭， 这种情况是如何实现的？具体过程分两种情况： 1 当传入的context是库里支持的cancelCtx、timeCtx时，child加入parent的衍生map中； 2 当传入的context为自定义实现类时，标准库是新起goroutine监听信号并执行cancel子context的过程。

```golang
// 在 parent 和 child 之间同步取消的信号，保证在 parent 被取消时，child 也会收到对应的信号，不会发生状态不一致的问题。
func propagateCancel(parent Context, child canceler) {
    if parent.Done() == nil { //parent不会触发cancel信号，则返回
        return // parent is never canceled
    }
    if p, ok := parentCancelCtx(parent); ok { //当前context为库中ctx时
        p.mu.Lock()
        if p.err != nil {
            // parent has already been canceled
            child.cancel(false, p.err)
        } else {
            if p.children == nil {
                p.children = make(map[canceler]struct{})
            }
            p.children[child] = struct{}{} //加入p的children map，等待p释放信号
        }
        p.mu.Unlock()
    } else { //非库中cancelCtx,timeCtx时，新起协程监听父ctx的信号
        go func() {
            select {
            case <-parent.Done(): // parent关闭时，关闭child
                child.cancel(false, parent.Err())
            case <-child.Done():
            }
        }()
    }
}
```

cancel方法的逻辑是关闭cancelCtx的Done通道，并递归地关闭context的所有子context。

```golang
func (c *cancelCtx) cancel(removeFromParent bool, err error) {
    if err == nil { // 外层传errors.New("context canceled") 这里不会进入
        panic("context: internal error: missing cancel error")
    }
    c.mu.Lock() //保护
    if c.err != nil { // 对已经cancel的context进行cancel 则返回
        c.mu.Unlock()
        return // already canceled
    }
    c.err = err  //context的err为 context cancel error
    if c.done == nil { 
        c.done = closedchan
    } else {
        close(c.done)
    }
    for child := range c.children { //遍历子context对其执行cancel操作
        // NOTE: acquiring the child's lock while holding parent's lock.
        child.cancel(false, err) 
    }
    c.children = nil 
    c.mu.Unlock()
    if removeFromParent {
        removeChild(c.Context, c)
    }
}
```

#### context使用说明

etcd中有一些select , context, channel 的例子

参考： https://blog.golang.org/context