
---
title: "compiler"
date: 2021-05-05T11:33:56+08:00
draft: true
isCJKLanguage: true
markup: mmark
tags: ["soft_tools"]
series: [""]
categories: ["技术"]
---

词法分析： code -> token

lex, flex 

递归下降

LALR(1)

词法分析+语法分析： promql

compiler explorer, tensor RT


---

语法分析： token -> ASR

yacc


EBNF grammar


--

## goyacc

安装： go get -u github.com/golang/tools/tree/master/cmd/goyacc

```
type yyLexer interface {
	Lex(lval *yySymType) int
	Error(e string)
}

type yyParser interface {
	Parse(yyLex) int
	Lookahead() int
}
```

csc 135,  finite automate, regular expressions, grammer

csc 151, compiler construction, Ghassan Shobaki

《engineering a compiler》 cooper and torczon

---

scanning(lexing), parsing

finite automata, regular expressions.


---

+ 
+ structure and major components

register allocation

when register is not enough for allocation, variable will spill into l1 cache or main memory.

instruction scheduling 

## 语言

c/c++, rust， llvm

java, c#, golang

python, matlab

lamada 演算。



## 想法

变量初始化都没有读去过。 rust warning， golang can 