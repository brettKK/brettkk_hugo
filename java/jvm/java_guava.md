### java and guava

+ 基本数据类型

+ BigDecimal(小数和精度)
+ StringBuilder, StringBuffer

----
Strings

+ Ints
	+ readed it

---

+ Throwable
	+ Error
		+ 描述底层资源不足或系统错误
		+ OutOfMemoryError
		+ StackOverflowError
		+ 一般不需要catch
	+ Exception
		+ RuntimeException	 	

guava

+ Throwables
	+ getCausalChain
	+ getRootCause
	+ propagate
	+ getStackTraceAsString

---
日志与标准输出／错误流 对比

+ System.out.printXX方法有synchronized关键词，多线程下会有锁的竞争，影响性能
+ 无法定义日志级别
+ 无法按天将日志转为文件，异常发邮件出来等需求
+ 无法直接输出到文件

+ slf4j , logback（实现了slf4j）,源码
	+ 日志级别:trace<debug<info<warn<error
+ logback
	+ Logger 产生日志事件
	+ Appender 记录到数据库，等存储
	+ Layout 格式化日志， html 	 	 

+ Tomcat
	+ 登陆跳板机（公司不允许直连服务器，可以从跳板机上ssh到server）
	+ 进入applicatioin目录，／home／q／www／application
	+ grep  -C
	+ zgrep 利用zgrep可以在压缩文件中查找关键词
	+ wc
	+ awk 	
	+ 高QPS时，采样打印日志

+ 日志规范
	+ Logger必须是org.slf4j.Logger
	+ 使用占位符代替字符串拼接
	+ 打印堆栈时，不需要占位符，且Throwable必须是最后一个参数

---
+ 代码规范
+ 设计规范 	
+ Javadoc格式注释
+ 《编写可读性代码的艺术》
+ 接口
	+ 即通信协议
	+ 接口职责单一
	+ RPC接口返回java原生容器，避免返回第三方库的容器 	
---

JDK Collections

+ List
	+ ArrayList
		+ resizable array
		+ 扩容原来的1.5倍
	+ LinkedList
		+ doubly linked list
+ Set based on map
	+ HashSet
	+ LinkedHashSet
	+ TreeSet
+ Map  
	+ HashMap
		+ array + LinkedList
		+ Capacity default 16
		+ Load factor default 0.75
		+ 扩容为原来的2倍
	+ LinkedHashMap
	+ TreeMap
		+ red-black tree 	 

Collections工具类 Ordering类

TreeMap的key是排序的，TreeSet，与Collections.sort区别

---
LinkedHashMap -> LRU map

EnumMap

ConcurrentHashMap

---
+ pool
+ cache
+ buffer

Java IO Guava IO

装饰器模式
