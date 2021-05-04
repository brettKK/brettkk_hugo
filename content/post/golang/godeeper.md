---
title: "go todo"
date: 2021-05-03T18:34:48+08:00
draft: false
isCJKLanguage: true
tags: ["golang"]
series: [""]
categories: ["技术"]
---


+ 深入解析go
    + 从词法语法解析、类型检查、中间代码生成以及机器码生成的编译相关全链路；
        + cmd/compile/internal 生成不同cpu下的mr
        + src/cmd/compile/internal/gc
    + 数组、哈希表和字符串等数据结构的表示以及基本操作的实现原理；
    + 理解 Go 语言中的函数、方法、闭包和上下文等语言特性；
    + 常见并发编程使用的 WaitGroup、Once 以及互斥锁等结构的实现；
    + 语言中的万能类型 interface 的实现原理；
    + 各种关键字 make、new、defer、select 以及 for 循环的处理过程；
    + Goroutine 和 Channel 相关的结构、调度策略以及原理

    + 源码目录文件夹(src, test, 
    + 源码中的汇编（crypto, math包，为了性能和特殊的cpu操作）
    + 基本数据结构，slice， map， channel
        + 创建数组（创建时类型和长度固定， 数组是否在heap上 编译时已经确定）
        + slice=动态数组
        + map
        + channel
    + gdb 跟踪源码
        + go如何做具体类型转为interface类型的（runtime.convT2E）
    + delve 跟踪
    + 打印调用栈信息 import "runtime/debug"   debug.Stack() 

+ 优点
  + 没有异常，没有继承（也就没有多态）
  + 允许多返回值, 有闭包
  + 自带GC(...)
  + 并发编程简单，just go, channel(自带线程安全), WaitGroup, RWLock
+ 缺点
  + error 没有类型
    + 2008.9.10 errorno int64 系统调用返回的错误类型； 2:19pm type Error struct {s string}
    + 2009.4.17 type Error interface { String() string}
    + github.com/pkg/errors
    + gopkg.in/errgo.v2
    + github.com/hashicorp/errwrap
    + upspin.io/errors
    + github.com/spacemonkeygo/errors
  + init() 函数滥用
  + 包版本管理器 vendor，govendor，godep不成熟
  + slice溢出时自动拷贝

+ 类型断言 空接口类型 运行时 转为 普通类型
+ golang中参数传递时 只有 参数值拷贝 然后 传入函数体，参数类型分为值类型和引用类型


+ gin服务器
    + sync.Pool 复用context对象
    + context.go gin.go routergroup.go tree.go
    + 维护url路径的前缀树


----

### an overview of go's tooling

+ table of content 
+ installing tooling
    + $ GO111MODULE=on go get golang.org/x/tools/cmd/stresss
    + go get -u, 强制使用网络去更新包和它的依赖包, 自动完成编译和安装 (git clone + go install)
+ viewing environment information
    + go env
    + go env GOPATH GOOS GOARCH
    + go help environment (go help get)
+ development
    + GOOS=windows go build
        + go build -gcflags="-m -l" // 看变量分配在stack or heap
    + go run ./cmd/foo
    + go run, go test, go build 可以静默添加外面依赖
    + go get github.com/foo/bar@v1.2.3 也可以指定特别的版本号
        + go get -v 显示操作流程的日志及信息，方便检查错误
        + go get -u 重新下载依赖包
        + go get -d 只下载，不安装
        + go get -insecure  允许使用不安全的 HTTP 方式进行下载操作
    + go list -m all
    + go mod why -m golang.org/x/sys
    + go mod graph
    + go clean -modcache (cache: GOPATH/pkg/mod)
    + gofmt -d -w -r 'Add -> ADD' .
    + go doc -src strings.Replace (view go documentation) , go doc fmt Printf
+ testing
    + go test ./foo/bar
    + go clean -testcache
    + go test -v -run=^TestFooBar$ .
    + go test -cover ./...
    + go test -run=^TestFooBar$ -count=500 . (stress testing)
    + go test -covermode=count -coverprofile=/tmp/profile.out ./...
    + go tool cover -html=/tmp/profile.out
+ pre_commit checks
    + gofmt -w -s -d .
    + go vet foo.go
    + golint foo.go
    + go mod tidy, go mod verify
+ build and deployment
    + go clean -cache 
    + go install -> $GOPATH/bin/demo
    + go list -f '{{ .Name}}'
    + go tool dist list (os/architecture)
    + GOOS=linux GOARCH=amd64 go build -o=/tmp/linux_amd64/foo .
    + https://github.com/golang/go/wiki/WebAssembly
+ diagnosing problems and making optimizations
    + go tool pprof --nodefraction=0.1  -http=:9999 /tmp/cpuprofile.out
    + go tool trace
        + goroutine的时间阶段： Execution Time， Network Wait Time， Sync Block Time， Blocking Syscall Time， Scheduler Wait Time， GC Sweeping， GC Pause。

+ managing dependencies
    + go list -m -u all
+ upgrading to a new go release
    + go fix
+ reporting bugs
    + go bug
+ chetsheet

+ goimports

+ godef
+ gopls
    + golang.org/wiki/gopls
    + go get golang.org/x/tools/gopls@latest

---

+ 反射 Reflection (interface对象=（Type, Value）反射类型 reflect.Type, reflect.Value)  三点
    + 反射可以将interface类型的变量 转换为 反射对象（reflect.Type reflect.Value）
    + 反射可以将反射对象 还原为interface对象
    + 反射对象可以修改，提前是value值是变量的指针  var x float= 3.1 ; v:= reflect.ValueOf(&x); v.Elem().SetFloat(3.4);
+ 动态调用方法， reflect.Value(interface) -> reflect.Value->value.MethodByName->Call

---
+ 函数接受者为指针类型
  + 函数可以修改接受者的内容
  + 避免方法调用时变量的拷贝，而是指针的拷贝
+ go严格上只有复制拷贝传递
  + 总是创建副本按值传递，只不过这个副本可以是变量的副本，也可以是指针的副本

+ io操作
    + io基本接口，io/ioutil 封装实用函数
    + fmt 格式化io
    + bufio 带缓冲io

```
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

+ Readers
  + os.File
  + bufio.Reader
  + net.Conn
  + compress/gzip, crypto/tls,...
  + bytes.Buffer

+ Goroutines
+ Channels (buffer, un buffer)
+ Range and Close
+ Select

+ golang 并发控制
    + channel
    + waitgroup
    + context
    + sync 
        + sync.Mutex 
        + sync.Once
        + sync.RWMutex
        + sync.WaitGroup
        + sync.Cond
+ golang 内存管理
    + 内存分配原理
        + src/runtime/mheap.go:mspan
    + 内存回收原理
        + 三色标记法 （未使用，已使用，带处理）
    + 逃逸分析（对象是否被函数外面引用 例如闭包； 对象过大也可能在堆中）
        + 逃逸分析的目的是决定变量的内存分配地址是在栈还是堆
        + 逃逸分析在编译阶段完成的
        + 传递指针真的比传值高效吗？不一定 由于指针传递会产生逃逸，可能会使用堆，增加GC负担。
        + go build -gcflags=-m // escapes to heap


+ interface{} 底层实现原理
    + 接口的作用是 通过一组方法 定义 一个对象 的行为；本质是
    + 思维层次
        + 基本原理(底层数据结构)
        + 类型断言的过程
        + 动态派发机制

    + eface(不含method的interface结构), runtime.convT2E（x := 1; var iInter interface{} = x, 汇编里调用该方法） 
        + 类型和指向数据的2个指针， 所以任意类型的对象 均可以转换为 interface{}类型的对象
    + iface, runtime.convT2I
        + itab
        + todo 
```
type eface struct { //16 bytes
    _type   *_type
    data    unsafe.Pointer
}

type iface struct { //16 bytes
    tab     *itab 
    data    unsafe.Pointer
}

type itab struct { //32 bytes
    inter *interfacetype //对 _type 类型的简单封装
    _type *_type    //_type 字段是 Go 语言类型在运行时的内部结构，每一个 _type 结构体中都包含了类型的大小、对齐以及哈希等信息
    hash uint32     // copy of _type.hash 从 interface 到具体类型的切换时用于快速判断目标类型和接口中类型是否一致
    _ [4]byte
    fun [1]uintptr //当前数组中内容为空就表示 _type 没有实现 inter 接口
}

type _type struct {
    size       uintptr
    ptrdata    uintptr // size of memory prefix holding all pointers
    hash       uint32
    tflag      tflag
    align      uint8
    fieldalign uint8
    kind       uint8
    alg        *typeAlg
    // gcdata stores the GC type data for the garbage collector.
    // If the KindGCProg bit is set in kind, gcdata is a GC program.
    // Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
    gcdata    *byte
    str       nameOff
    ptrToThis typeOff
}

```

----        

+ channel, goroutine make concurrent programming in go more convenient
+ channel 结构 (src/runtime/chan.go中实现)
    + 带缓存的，不带缓存的
    + 发送数据 3种情况
        + hchan.buf为空时，直接将数据发送给goroutine
        + hchan.buf不为空时，将数据放入buffer里
        + hchan.buf满时，阻塞当前goroutine
    + closechan 关闭channel的函数
        + closed置为1
        + 唤醒recvq队列里的阻塞goroutine
        + 唤醒sendq队列里的阻塞goroutine
    + gracefully close channels
```
ch := make(chan int, 3)
内存channel的初始状态
qcount = 0
dataqsiz = 3
buf = [3]int{0,0,0}
sendx = 0
recvx = 0

```

```
// channel 底层结构
type hchan struct {
    qcount uint,            // 队列中数据个数
    dataqsiz uint;          // channel 大小
    buf unsafe.Pointer      // 存放数据的环型数组
    elemsize uint16;        // channel中数据类型的大小
    closed uint32;          // 表示channel是否关闭
    elemtype *_type;        // 元数据的类型
    uintgo sendx;           // send index
    uintgo recvx;            // rece index
    WaitQ recvq;             //由 recv 行为（也就是 <-ch）阻塞在 channel 上的 goroutine 队列
    WaitQ sendq;             // 由 send 行为 (也就是 ch<-) 阻塞在 channel 上的 goroutine 队列
    mutex Lock;              // protect field
}

stuct WaitQ {
    first *sudog;
    last *sudog;
}

type sudog struct {

}

```


----

+ GPM结构  sched调度器相关的数据结构
    + M-os线程，P-go代码执行时需要的资源， 在P上实现工作流窃取的调度器
    + go关键字 -> 均有runtime.newproc（size, f, args）创建groutinue 

```

struct G
{
    uintptr stackguard; //分段栈的可用空间下界 防止栈溢出 base+64k 的为止
    uintptr stackbase;  //分段栈的栈基址
    Gobuf sched;        //切换时，sched存上下文信息
    uintptr stack0;     //
    FuncVal* fnstart;   //goroutine运行时的函数
    void* param;        //用于传递参数
    int16 status;       //goroutine状态，如Gidle, Grunnable, Grunning, Gsyscall, Gwaiting, Gdead等
    int64 goid;         //goroutine id号
    G* schedlink;       //
    M* m;               // for debug
    M* lockedm;         //G被锁定只能在m上运行
    uintptr gopc;       //创建这个goroutine的go表达式的pc
    ...
}

struct GoBuf
{
    unintptr sp;    // 当前栈指针   
    byte* pc;       // 程序计数器    
    G* g;           // goroutine本身
    ...
}

// Machine 对物理线程的抽象
struct M
{
    G* g0;                      // 带有调度栈的goroutine, M对应的线程的栈
    G* gsignal;                 // 处理信号的goroutine
    void (*mstartfn)(void);     // 
    G* curg;                    // M中当前运行的goroutine
    P* p;                       // 关联p去执行go代码，（没有执行go代码时，p为nil）
    P* nextp;
}

struct P
{
    Lock;
    uint32 status;  //Pidle, Prunning, Psyscall, Pgcstop, Pdead
    P* link;
    M* m // 链接到关联的M
    G* runq[256]
    int32 runqhead;
    int32 runqtail;
}

// 多数信息已在M G P 中定义， Sched是一个空客的封装
struct Sched 
{
    Lock;
    uint64 goidgen;

    M* midle;
}

```
----
+ nil
    + what is nil
        + nil is a kind of zero
    + what is nil in go
        + 各种类型的零值
        + zero value of all type value
            + bool false, number 0, string ""
            + pointer, slice, map, channel, function, interface{} ==> nil
            + structure{}
        + kind of nil 
            + interface{} = (type, value) ==> (nil, nil) == nil 
    + what does nil mean
    + is nil useful ? yes

+ 普通类型如何转为interface类型， convT2E

+ unsafe
+ unsafe.Pointer()
  + T1指针与 T2指针类型之间的转换
  + 修改内存数据 uintptr 常用于与 unsafe.Pointer配合，用于做指针运算

+ go ast,llvm
+ gdb单步调试go程序的执行过程    Delve 更好

+ 工欲善其事，必先利其器

+ vscode 常见功能键 good tools
    + Mac下 vs code设置文件：$HOME/Library/Application Support/Code/User/settings.json
        + 最开始使用run test时无输出， 在settings.json里添加"go.testFlags":["-V"],
    +   预览md格式（⇧⌘V）
+ vim go: https://octetz.com/docs/2019/2019-04-24-vim-as-a-go-ide/

+ Guru 导航go 代码的编辑器集成工具
    + 变量，函数的声明地点
    + 变量，函数的所有引用地点
    + 实现此接口的所有具体类型
    + 
    + golang.org/x/tools/cmd/guru
    + http://golang.org/s/using-guru

----

+ golang fuzz testing

--- 
+ avo
    + github.com/mmcloughlin/avo
+ go assembly
    + 优势： 避免编译器的错误优化，使用特殊的硬件指令，在数学计算，密码，系统调用，底层运行细节
+ plan 9
    + show known better
    + https://lrita.github.io/2017/12/12/golang-asm/#go%E5%87%BD%E6%95%B0%E8%B0%83%E7%94%A8
    + https://lrita.github.io/images/posts/go/GoFunctionsInAssembly.pdf
    + https://github.com/buptbill220/gotls/tree/master/interesting
    + 查看当前cpu架构的指令集合 cat /proc/cpuinfo | grep flags | head -1
    + 指令查询： http://68k.hax.com/
    + 命令查询：https://9p.io/magic/man2html/1/8a
    + plan9 文档： https://9p.io/sys/doc/
+ https://defuse.ca/online-x86-assembler.htm#disassembly2
  +  x86 or x64 assembly instructions encoding<==> decoding binary code
  +  It uses GCC and objdump behind the scenes