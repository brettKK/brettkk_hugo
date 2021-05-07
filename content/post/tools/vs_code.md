---
title: "vs code 小结"
date: 2021-05-05T11:33:56+08:00
draft: false
isCJKLanguage: true
markup: mmark
tags: ["soft_tools"]
series: [""]
categories: ["技术"]
---

+ 快捷键内置查询表 菜单栏-》帮助-》快捷键参考-》very good
+ 自定义快捷键： 文件—》首选项-》键盘快捷方式

command+shift+P (F1) 命令面板， 输入：format， 格式化代码;

ctrl + shift + N: 打开新的编译器窗口

command+J 打开关闭控制台

command+B 打开关闭侧边栏

command+P 搜索文件
+ 在代码里穿梭
 + 输入跳转到函数定义的位置, F12（ctrl+鼠标左键）
 + 原路返回 ctrl + "-"， 撤销返回 ctrl+shift+"-"
 + 跳到指定行 ctrl + G, 跳到指定标签页 ctrl+数字
 + 跳到文件顶部、底部， cmd+up,down
 + shift+F12 显示所有引用
 + @, 查找属性或者函数
 + #， 根据名字查找symbol
 + cmd + "." 修复语法错误

+ vscode 常见功能键 good tools
    + Mac下 vs code设置文件：$HOME/Library/Application Support/Code/User/settings.json
        + 最开始使用run test时无输出， 在settings.json里添加"go.testFlags":["-V"],
    +   预览md格式（⇧⌘V）
+ Configure Visual Studio Code for Go
    + extensions: go
    + view -> command palette... -> goinstall