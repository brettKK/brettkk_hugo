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
rustup doc

docs.rs/std 可以定向到doc.rust-lang.org/stable/std/

### rust 开发环境

IDE： vscode 
插件： rust-analyzer（用）, Rust(别用, 与rust-analyzer有冲突)。
      crates。
       TOML Language Support .
       REST Client
      Cargo.

playground

### 软件

serde json

tokio, Ethereum, Redox, Tikv, clippy

strato virt, tch-rs, https://gamedev.rs

ThinLTO, PGO

---




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


## array

[Type; N] , [1, 2, 3], [val; n]

## 编写测试

rust中的测试是带有标注test属性的函数。 #[test]

cargo test.


## 变量的生命周期，所有权，借用，move，ref

哪些操作会导ownership move
+ let binding
+ 传参数到function
+ match 表达式
+ methods。 impl block里任何method第一个参数是self
+ closure中

### 借用（引用 同意）
rust中所用的参数传递都是传值。 
引用实现了copy trait。 按照copy语义，引用会复制一份交给调用的函数。


## 智能指针

Box<T>

Rc<T>

RefCell<T>

Deref, Drop trait

Weak<T>
