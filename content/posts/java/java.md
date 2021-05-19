---
title: "java 小结"
date: 2021-03-05T11:33:56+08:00
draft: false
isCJKLanguage: true
markup: mmark
tags: ["java"]
series: [""]
categories: ["技术"]
---

Class文件格式图
![image](/os/java-class-file-internal-structure.png)

1  class文件由十个部分组成 javap查看字节码文件
    + my very cute animal turns savage in full moon areas
    +  magic,  0xCAFEBABE ，pdf png文件内都有魔数
    +  version, 生成class文件时的java版本
    +  constant pool,  javap -v Helloworld
    +  access flag, 2字节 标识一个类为final, abstract, public 之类的
    +  this class,  确定类的继承关系 指向常量池的索引
    +  super class, 确定类的继承关系
    +  interface,   确定类的继承关系
    +  field, 字段表
       +  变长
       +  field_info结构 
          +  access_flags 字段访问标识
          +  name_index 字段名索引 指向常量池的字符串常量
          +  descriptor_index  字段类型描述符的索引  引用类型“L;” -> “Ljava/lang/String;” 指向常量池的字符串常量
          +  attributes_count 属性个数
          +  attribute_info  属性集合
             +  ConstantValue
             +  Synthetic
             +  Signature
             +  Deprecated
             +  Runtime-Visible Annotationns
             +  Runtime-Invisible Annotations
    +  method, 类中定义的方法存储在这里
       +  变长
       +  method_info
          +  access_flags
          +  name_index
          +  decriptor_index 方法描述符 指向常量池类型为constant_uft8_innfo的字符串常量项，表述(参1的类型，参2的类型...) 返回值类型
          +  attributes_count
          +  attribute_info
    +  attribute 属性表
       +  attribute_info 
          +  attribute_name_index 属性名索引 指向常量池的字符串常量
          +  attribute_length info 数组长度
          +  info
             +  ConstantValue属性
             +  Code属性 重要的部分 包含了方法的字节码

2 字节码指令
+ switch-case的两种字节码指令的实现
   + tableswitch , case间隔紧密时使用 o(1)
   + lookupswitch , case间隔稀疏时使用  o(lgn)
+ try-catch 字节码分析 JVM的异常处理是通过编译期间确定下来的异常表，在运行时利用异常表来实现的
   + code属性里有异常表 表项4元组：from, to, target, exception type
   + 在from,to 范围内的字节码发生指定的exception type时跳转到target字节码的位置
   + finally 字节码分析
      + 为什么finally一定执行： javac生成字节码时在try和catch中所有退出之前加入invokevirtual 调用finally里代码块
      + finally语句中有return语句时，javac生成字节码时会将try 和catch里的return语句的值临时放入局部变量表里，只返回finally里的return语句
+ 对象相关的字节码
   + <init> 方法 ，类的构造方法，非静态变量的初始化，对象初始化代码块 
      + 创建对象需要三条指令： new, dup, invokespecial 
   + <clinit> 类的静态初始化方法 ， 类静态初始化块 静态变量初始化
      + static {};
      + new, getstatic, putstatic, invokestatic 触发

3 字节码进阶 （方法调用 反射 invokedynamic 线程同步）

+ 5条方法调用指令的联系和区别 （以invoke开头的指令）
   + invokestatic 调用static静态方法
      + 调用的方法在编译期间确定，静态绑定
      + 不需要将对象加载到操作数栈上，仅参数入栈，执行invokestatic就行
   + invokespecial  调用特殊的实例方法，构造器方法
      + 调用的方法在编译期间确定，效率比invokevirtual高
      + 实例构造器方法<init>, private修饰的私有方法， super关键词调用的父类方法
   + invokevirtual 运行时根据对象的实际类型，执行具体子类的实现方法。 vtable
      + 调用的方法在运行时才能根据对象的实际类型确定
      + 需要将对象应用，参数入栈
   + invokeinterface 调用接口方法  itable
      + 调用的方法在运行时根据对象的类型确定目标方法
      + invokeinterface从itbale表找对应的方法
   + invokedynamic 调用动态方法
      + 前4条指令的分派规则固化在jvm中
      + invokedynamic 是把如何查找目标方法的决定权给了具体的用户代码中

+ JVM方法分配机制与vtable,itable原理， 方法分派
   + vtable, itable是java实现多态的基础
   + 对象的继承与多态
   + 子类继承父类的vtable。 一个没有方法的类的vtable也有5条来自Object类
   + final, static, private 方法不能被继承和重写，所以不会出现在vtable中
   + itable = offset table + method table 去支持invokeinterface指令 
+ invokedynamic指令，lambda表达式中的作用
   + java8中出现的lambda表达式才第一次使用上这条指令
   + 匿名内部类是在编译期间生成新的class文件来实现的
   + asm在运行时生成class对象
+ 从字节码角度理解范型擦除
   + 范型类型会被擦除为Object类型
   + ArrayList<int> 不允许，因为原始类型不能存储到Object类型上
+ synchronized 关键字的字节码原理
   + 插入monitorenter (获取栈顶对象的锁)， 插入monitorexit(释放锁，一定插入至少2个，原因是正常退出和异常退出均会插入）
   + 一定会在code属性里生成异常处理的描述表，即在所有退出的地方插入monitorexit 保证锁在任何退出的节点均被释放（try-catch-finaly monitorexit）


类加载主要做三件事：
+ 通过类名从class文件中获取字节流
+ 将字节流存储在方法区上的运行时数据结构
+ 在堆上生成Class对象，作为这个类在方法区上的访问入口


+ 生成字节码的框架
  + javassit，cglib，asm
+ 读和写字节码
  + 一个CtClass对象 - 对应Class对象
  + ClassPool是由CtClass对象组成的hashtable
    + obtain from ClassPool
---

一个对象在内存中的结构

![image](/os/object-memory-structure.png)

对象头由2部分组成 - mark word（32位或者64位）： 25bit 哈希码，2bit 锁标志位 ； 二  类型指针（指向Class对象）

在对象上加锁的过程
+ 复制对象头到执行栈中的锁记录中
+ 修改对象头中的mark word， 2点修改：修改mark word的前30bit为存放锁记录的地址，修改锁标志位


+ 如何自实现一个Immutable类
  + class定义为final，避免继承
  + 所有成员熟悉定义为final
  + 构造函数不要引用外部对象，构造时深度拷贝
  + 其实在类加载是可以做任何字节码修改
+ 判断一个类是否是Immutable的

+ Concurrent Modification（并发修改集合对象）
  + 遍历集合类的同时，修改集合类的结构，产生ConcurrentModificationException
+ java里有2种iterator的实现方式： fail-safe , fail-fast(快速失败)
  + fail fast
    + fail fast iterator throw ConcurrentModifacationException
    + iterator use original collection to travel over
  + fail safe
    + iterator required extra memory to clone the original collections
    + iterator use the copy collection to traverse over
    + allow modification a collection while iterating over it

+ 扩容过程
  + 经过rehash之后，元素的位置要么是在原位置，要么是在原位置再移动2次幂的位置
+ HashMap 扩容时，size增大一倍， newsize=oldsize*2
+ ArrayList（动态数组）扩容时，扩大0.5倍，newsize=oldsize*1.5
+ ArrayQueue， 当size小于64时，+2；size大于64时，newsize=oldsize*1.5

