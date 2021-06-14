---
title: "k8s 编码模式学习"
date: 2021-04-05T11:33:56+08:00
draft: false
isCJKLanguage: true

tags: ["cloud native"]

---


### cobra命令行的使用

cmd下的main函数为各个k8s组建的入口。    

```golang

```


### 函数选项模式

```golang

```

### 访问者模式

kubectl里有使用。

```golang

type Visitor interface {
	Visit(VisitorFunc) error
}

type VisitorFunc func(*Info, error) error

type VisitorList []Visitor

// Visit implements Visitor
func (l VisitorList) Visit(fn VisitorFunc) error {
	for i := range l {
		if err := l[i].Visit(fn); err != nil {
			return err
		}
	}
	return nil
}

```

### channel的使用方式

#### 内部具体channel的使用示例

stopCh <-chan struct{}，  channel的使用方式

#### 抽象channel的使用模式



### 工厂模式

#### 调度器算法的插件

```golang

```

### 

```golang

```