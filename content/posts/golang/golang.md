---
title: "golang"
date: 2021-05-07T18:54:41+08:00
draft: false
isCJKLanguage: true
tags: ["golang"]
series: [""]
categories: ["技术"]
---

## 资料
 
 1. 目前golang的主要maintainer: Russ Cox
 2. 个人主页：swtch.com/~rsc/

 3. awesome-go系列，https://github.com/gocn/knowledge

练习操场playground: go.dev/play。  后有rust playground。

---
## 多态，继承， 范型

golang中面向对象的实现方式
### 继承

组合结构体，一个结构体嵌入另一个结构体，能够实现对嵌入结构体的字段以及其实现的方法的继承。

{{< highlight go "linenos=table,hl_lines=8 15-17,linenostart=1" >}}
type person struct {
    name string
    age  int
}

type student struct {
    person //   匿名字段，通过组合的方式
    school string 
}

func (p * person) talk() {
    fmt.Println(p.name, p.age)
}

func (s *student) getSchool() {
    return s.school
}

func main() {
    s := student{person{"brettkk", 18}, "yidu"}
    fmt.Println(s.name) // 继承person的name字段
    s.talk() // 继承person实现的方法
}
{{< / highlight >}}

### 多态

go没有 implements, extends 关键字， golang是去鸭子类型：看起来像鸭子， 那么它就是鸭子。

多态： 父类指针或引用去调用方法，在执行的时候，能够根据子类的类型去执行子类当中的方法。

为什么是在执行的时候才能根据子类的类型执行子类的方法？

```golang

type Person interface {
    func gender() string
}

type Men struct{
    Gender string
    Age int
    Name string
}

type Woman struct {
    Gender string
    Age int
    Name strings
}

func (m Men) gender() {
    return m.Gender
}

func (w Woman) gender() {
    return w.Gender
}

func JudgeFunc() bool {
    // may has many code
    // can't not get boolean result until running this JudgeFunc
}

func main() {
    var a Person
    // so compiler can't not known clearly about the type info of a
    a = JudgeFunc() ? Men{} : Woman{}

    var b Person
    // compiler can judge the type info of a 
    b = Men{} // or b = Woman{}
}

```

### 范型

空interface{} 类型，没有方法集的接口，只需要存类型和类型对应的值即可。

空interface{} = (type, value) ==> (nil, nil) == nil 


```golang
type eface struct {
	_type *_type
	data  unsafe.Pointer
}

```


### 接口的使用实践

golang源码里的较好实现方式是： 使用小的接口，尤其是只包含一个方法的接口；通过小接口的组合来定义大接口。

好处是：使用者可以只依赖于必要功能的最小接口。

```golang
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}

func store(reader Reader) error {
    // 入参可以只依赖于必要功能的最小接口， 让函数的返回更广
}
func store(readWriter ReadWriter) error {
    // 入参依赖了不必要功能的接口，让限制了函数的使用范围为更小
}

```

## 模糊点

### rune类型

字符串有字符组成， 字符分2种， 一种是占1byte的字符，如英文字符；一种是占1-4 byte的字符用rune表示，如中文。


在builtin/builtin.go中定义
```golang

type rune = int32
fmt.Println(len("hallo中文"))// 5 + 3 + 3
fmt.Println(len([]rune("hallo中文"))) //5 + 1 + 1

```

### 数组与切片

数组的初始化 （元素类型， 数组大小）

2种初始化方式，显式的指定数组的大小 
1. arr := [3]int{1, 2, 3}
2. [...]T 声明数组 // arr := [...]int{1, 2, 3}


<br/>

### 函数的接受者

函数接受者为指针类型:
1. 函数可以修改接受者的内容
2. 避免方法调用时变量的拷贝，而是指针的拷贝

<br/>

### golang值拷贝

go严格上只有复制拷贝传递
  
总是创建副本按值传递，只不过这个副本可以是变量的副本，也可以是指针的副本

<br/>


### 闭包

闭包=函数+引用环境。

闭包里引用局部变量时，在堆上创建该变量的一个拷贝，并把该变量地址和函数闭包组成一个结构体，并把该结构体传出来作为返回值


<br/>

### unsafe包

uintptr与unsafe.Pointer（类似void*）

1. T1指针与 T2指针类型之间的转换
2. 修改内存数据， uintptr 常用于与 unsafe.Pointer配合，用于做指针运算

```golang
//结构体的成员变量在内存存储上是一段连续的内存
type Num struct{
	i string
	j int64
}

func main(){
	n := Num{i: "EDDYCJY", j: 1}
	nPointer := unsafe.Pointer(&n)
	// 结构体的初始地址就是第一个成员变量的内存地址
	niPointer := (*string)(unsafe.Pointer(nPointer))
	*niPointer = "wuhan"
	// 基于结构体的成员地址去计算偏移量。就能够得出其他成员变量的内存地址
	//uintptr 是 Go 的内置类型。返回无符号整数，可存储一个完整的地址。后续常用于指针运算
	//unsafe.Offsetof：返回成员变量 x 在结构体当中的偏移量。更具体的讲，就是返回结构体初始位置到 x 之间的字节数
	//uintptr 类型是不能存储在临时变量中的 临时变量是可能被垃圾回收掉的 之后的内存操作有迷茫了
	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
	*njPointer = 91

	fmt.Printf("n.i: %s, n.j: %d", n.i, n.j)
}
```


指针是对内存区域的地址， 与指针相配合的类型
说明区域有哪些属性，如何去解析。

<br/>

### nil

+ what is nil
    + nil is a kind of zero
+ what is nil in go
    + 各种类型的零值
    + 基本类型的零值不为nil
        + bool为false
        + number为0
        + string为""
    + 非基本类型的零值为nil
        + pointer, slice, map, channel, function, interface{} ==> nil
        + interface{}类型 需要type和value均为nil时，才会为nil
        + structure{}，因为有结构type存在，type不为nil，所以结构体变量不会是nil

<br/>

### 反射

反射 Reflection ，interface对象=（Type, Value）

反射类型 reflect.Type, reflect.Value) 

三种反射的基本操作：
1. 反射可以将interface类型的变量 转换为 反射对象（reflect.Type reflect.Value）
2. 反射可以将反射对象 还原为interface对象
3. 反射对象可以修改，提前是value值是变量的指针  

```golang
var x float= 3.1 
v:= reflect.ValueOf(&x) //需要传入x的指针
v.Elem().SetFloat(3.4);
```

作用：
1. 运行时动态调用方法
2. 运行时构造函数进行mock插桩

<br/>

### 内存管理

1. 基本想法：局部P上（per cpu）和全局队列想结合的均衡思想。局部分配无需加锁，局部分配无法满足时加锁全局分配

+ src/runtime/mheap.go:mspan
+ 内存回收原理
  + 三色标记法 （未使用，已使用，带处理）
  + root对象开始，BFS遍历标记，
+ 逃逸分析（对象是否被函数外面引用 例如闭包； 对象过大也可能在堆中）
  + 逃逸分析的目的是决定变量的内存分配地址是在栈还是堆
  + 逃逸分析在编译阶段完成的
  + 传递指针真的比传值高效吗？不一定 由于指针传递会产生逃逸，可能会使用堆，增加GC负担。
  + go build -gcflags=-m // escapes to heap

<br/>

## go 汇编

ssa里的优化pass：https://github.com/golang/go/blob/acbe242f8a2cae8ef4749806291a37d23089b572/src/cmd/compile/internal/ssa/compile.go#L418

Generate x86 Assembly with Go

https://github.com/mmcloughlin/avo

+ 汇编风格分为2类
1. intel风格
2.  AT&T风格(左 --> 右) go是plan9， 属于此类

+ 助记符    call, add
+ register  ax, sb, pc
+ 立即数 $1 , $0x100
+ 内存寻址 （R1）

+ plan 9
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

+ goroot/src/syscall --》 建议使用 golang.org/x/sys