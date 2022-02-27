---
title: "linux-base"
date: 2021-05-05T11:33:56+08:00
draft: true
isCJKLanguage: true
markup: mmark
tags: ["linux"]
series: [""]
categories: ["技术"]
---

### rust 安装

https://rustup.rs

rustup update 
rustc --version
rustup doc （好用）

cheats.rs
rust-gdb
rust-lldb

docs.rs/std 可以定向到doc.rust-lang.org/stable/std/


### rust 开发环境

IDE： vscode 
插件： rust-analyzer（用）, Rust(别用, 与rust-analyzer有冲突)。
      crates。
       TOML Language Support .
       REST Client
      Cargo.

      Error Lens, One Dark Pro, Code LLDB.

playground

rustc -V, cargo -V 查看环境版本。

rustup 管理rust编程环境。 
rustup which rustc 查看rustc的位置
rustup component add rust-src 下载rust源码。

cargo check快速检查语法问题。


### 资料

rust权威指南： https://kaisery.github.io/trpl-zh-cn/

rustprimer:  https://rustcc.gitbooks.io/rustprimer/content/

rust 编程之道 张汉东

welcome to rust 101: https://www.ralfj.de/projects/rust-101/main.html

rust by example

https://tourofrust.com/

youtube channels: jon gjengset

### blog and news

https://www.rust-lang.org/community

https://this-week-in-rust.org/

https://blog.rust-lang.org/

https://blog.rust-lang.org/inside-rust/index.html


### 软件

serde json

tokio, Ethereum, Redox, Tikv, clippy

strato virt, tch-rs, https://gamedev.rs

ThinLTO, PGO

---

text -> tokens -> ast -> hir ->mir -> llvm ir -> llvm -> 1110


Rust程序设计



+ 安装

https://rustup.rs/

+ 第一个程式


cargo -h
cargo new --bin hello
cargo run 

rustup doc
cargo doc --open

std::collections
Vec, VecDeque, HashMap, String


geektime-rust , codes , source_code , cargo doc --open
阅读顺序： 类型自身，实现的方法，实现的trait

## 工程，包crate，模块mod

```
project
│   README.md
│   file001.txt
└───src
│   │   lib.rs
│   └───package1
│       │   mod.rs
│       │   a.rs
│       │   b.rs
│   
└───folder2
    │   file021.txt
    │   file022.txt
```

工程目录结构参考 pingcap/tikv

## 引入依赖
工程的toml文件引入第三方库
+ 引入rust官方仓库 
  + 中央仓库crates.io，docs.rs
+ 引入git依赖
+ 引入本地依赖

工程的代码中加载依赖 extern crate xxx as yyy;

## 数组array

数据的类型： 元素类型+数组长度（长度也是类型的一部分）

[Type; N] , [1, 2, 3], [val; n]

## enum

包含各种子类型的集合。

## 编写测试

rust中的测试是带有标注test属性的函数。 #[test]

cargo test.

## 基本语法概念

12个。

rust, open source. MIT and Apache License 2.0. 
v1.0 -> 2015. play.rust-lang.org

println!("{:?}", ..), mut, shadow vaviable, constants

types. String  UTF-8.

tuple(different type and fix length), Array( same type and fix length)

Expressions and statements.  Functions, loop, while, for。

+ 介绍思路
      + data type. Option, Result, Interator
      + control flow. match, for, if let, while let
      + error handler
      + ownership
      + trait. Iterator, Deref, Index, Send/Sync
      + project. Chalk, tokio, Servo.


## 变量的生命周期，所有权，借用，move，ref

+ 变量let()与常量const, static
      + 变量let声明
            + 变量shadow
      + 常量const声明，且一定要声明数据的类型， 编译时copy到使用的地方
      + static变量，可变的静态变量不是线程安全的，unsafe scope起来使用
+ 所有权
      + 对象的所有权
      + 对象所在内存的所有权（主人）



哪些操作会导ownership move
+ let binding
+ 传参数到function
+ match 表达式
+ methods。 impl block里任何method第一个参数是self
+ closure中

```rust
let s = Some(String::from("hallo"));
if let Some(_s) = s {

}
// _s 会发生绑定， 所有权移动
if let Some(_) = s {

}
// _ 不会绑定 所有权没有移动
```

### 借用（引用 同意）
rust中所用的参数传递都是传值。 
引用实现了copy trait。 按照copy语义，引用会复制一份交给调用的函数。

late bound, early bound. Quiz


### 操作符号重载
rust中的操作符几乎都与一个trait相对应。
解引用*x, 重载std::ops::Deref

不可重载的运算符： &, &mut, 

### 标准库-- 集合

标准库的集合在std::collections模块中。
sequence: Vec, VecDeque, LinkedList
Map: HashMap, BTreeMap
Set: HashSet, BTreeSet
Misc: BinaryHeap

std::vec::Vec 不在collections模块中。

标准库--Prelude模块。 会被自动引入到项目中。具体有哪些？


## 智能指针

Box<T>

Rc<T> ,Arc<T> 

---

Cell<T>, RefCell<T> 扩展所有权机制的灵活性。

通常修改一个值，2个方式
+ 成为值的所有者，并声明mut
+ &mut 获取可变借用

Cell， RefCell则不受编译静态规则的限制。
Cell<T> 只能用于T实现了Copy trait的情况。
RefCell<T> 

---

Deref, Drop trait

Weak<T>

## 元编程

反射， 宏

## unsafe

## 异步编程

Future, Tokio

## web 框架