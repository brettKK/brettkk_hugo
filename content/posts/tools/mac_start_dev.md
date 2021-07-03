---
title: "mac安装开发环境"
date: 2021-05-05T11:33:56+08:00
draft: false
isCJKLanguage: true
markup: mmark
tags: ["soft_tools"]
series: [""]
categories: ["技术"]
---


新电脑构建开发环境，安装指南

+ iterm2
+ brew
+ zsh, ohmyzsh
+ golang
+ vscode
+ hugo
+ github

---


+ 安装终端窗口工具

下载地址：https://iterm2.com/


+ 安装 brew

/bin/bash -c "$(curl -fsSL https://cdn.jsdelivr.net/gh/ineo6/homebrew-install/install.sh)"

+ 安装，切换为zsh，并安装oh my zsh -- shell辅助工具

可能遇到的问题

```
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

curl: (7) Failed to connect to raw.githubusercontent.com port 443: Connection refused

```
解决方式： https://github.com/ohmyzsh/ohmyzsh/issues/9465


+ 安装 golang

https://golang.org/doc/install

https://sourabhbajaj.com/mac-setup/Go/README.html

vim .bashrc

export GOPATH=$HOME/go-workspace # don't forget to change your path correctly!
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin


+ 安装 vscode

https://code.visualstudio.com/docs/setup/mac

插件 

+ 安装blog

Brew install Hugo

git pull origin master --allow-unrelated-histories

git submodule update --init --recursive



+ 本地配置GitHub访问权限

在～下创建.ssh文件夹 ssh-keygen -t rsa -C “xxx@xx.com"
将公钥添加到github

go env -w GOPROXY=https://goproxy.cn



