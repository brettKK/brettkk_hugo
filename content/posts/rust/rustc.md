---
title: "rustc"
date: 2021-05-05T11:33:56+08:00
draft: true
isCJKLanguage: true
markup: mmark
tags: ["rustc", "compiler"]
series: [""]
categories: ["技术"]
---

### rustc task list

+ 1 command line args, user environment, toolchain, load compilation cache
+ 2 lexing, parsing
+ 3 macro expansion, feature gating, various compiler magic

1,2,3 基于ast， 能够对程序进行语法操作。

+ 4 type inference
+ 5 type checking
+ 6 trait solving/checking

4,5,6 基于HIR

+ 7 pattern exhaustiveness checking

7 基于THIR

+  Borrow checking （限制 local， scope是一个方法内部）
+ Constant evaluation
+ Rust-level Optimizations
+ Monomorphization
+ Saving compilation cache
+ LLVM IR generation

基于 MIR， control flow graph。 方便borrow checking, dataflow for correctness checks and opts， monomorphize, optimize, generate code。

+ LLVM (more optimization, binary generation, linking, linking-time optimization...)