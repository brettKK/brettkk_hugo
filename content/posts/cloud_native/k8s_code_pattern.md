---
title: "k8s 编码模式学习"
date: 2021-04-05T11:33:56+08:00
draft: false
isCJKLanguage: true

tags: ["cloud native"]

---

+ k8s 组件
  + kubectl
  + client-go
  + kube-apiserver
  + kube-controller
  + kube-scheduler
  + kubelet
  + kube-proxy

### 周围项目

k8s.io/gengo
k8s.io/client-go


### 项目组织结构


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

《kubernets 源码分析》 郑东旭 
```golang
type Visitor interface {
	Visit(VisitorFunc) error
}
type VisitorFunc func() error

type VisitorList []Vistor

func (l VisitorLList)Visit(fn VisitorFunc) error {
	for i := range l {
		if err := l[i].Visit(func() error {
			fmt.Println("in visitorlist before fn")
			fn()
			fmt.Println("in visitorlist after fn")
		}); err != nil {
			return err
		}
	}
	return nil
}

type Visitor1 struct {
}

type (v Visitor1) Visit(fn VisitorFunc) error {
	fmt.Println("in visitor1 before fn")
	fn()
	fmt.Println("in visitor1 after fn")
}

type Visitor2 struct {
	visitor Visitor
}
func (v Visitor2) Visit(fn VisitorFunc) error {
	v.visitor.Visit(func() error{
		fmt.Println("in visitor2 before fn")
		fn()
		fmt.Println("in visitor2 after fn")
		return nil
	})
	return nil
}

type Visitor3 struct {
	visitor Visitor
}

func (v Visitor3) Visit(fn VisitorFunc) error {
	v.visitor.Visit(func() error{
		fmt.Println("in visitor3 before fn")
		fn()
		fmt.Println("in visitor3 after fn")
		return nil
	})
	return nil
}

func maiin() {
	var visitor Visitor
	var visitors []Visitor

	visitor = Visitor1{}
	visitors = append(visitors, visitor)
	visitor = Visitor2{VisitorList(visitors)}
	visitor = Visitor3{visitor}

	visitor.Visit(func() error{
		fmt.Println("In visitfunc")
		return nil
	})
}

/*
in visitor1 before fn
in visitorlist before fn
in visitor2 before fn
in visitor3 before fn
in visistfunc
in  visitor3 after fn
in visitor2 after fn
in visitorlist after fn
in visitor1 after fn
*/

```

### client-go 编程式交互


### 观察者模式

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