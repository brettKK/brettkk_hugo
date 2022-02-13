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

cargo check快速检查语法问题。


### 资料

rust权威指南： https://kaisery.github.io/trpl-zh-cn/

rustprimer:  https://rustcc.gitbooks.io/rustprimer/content/

rust 编程之道 张汉东

welcome to rust 101: https://www.ralfj.de/projects/rust-101/main.html

rust by example

https://tourofrust.com/

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

## array

[Type; N] , [1, 2, 3], [val; n]

## 编写测试

rust中的测试是带有标注test属性的函数。 #[test]

cargo test.


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


## 智能指针

Box<T>

Rc<T>

RefCell<T>

Deref, Drop trait

Weak<T>

## 元编程

反射， 宏

## unsafe

## 异步编程

Future, Tokio

## web 框架