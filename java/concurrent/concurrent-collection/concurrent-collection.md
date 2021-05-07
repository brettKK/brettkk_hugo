BlockingQueue： 用于实现生产者-消费者模式，当队列满时添加会阻塞；当队列空时删除会阻塞

---
PriorityQueue 优先队列(java.util包下，非线程安全)，底层使用堆结构的数组实现

+ 插入时，插入的数组尾部，从下到上调整堆
+ 删除时，删除数组的头部，从上到下调整堆

---
DelayQueue 延迟队列
+ 队列里的元素为实现了Delayed接口的元素，是一个具有过期时间元素
+ 根据队列里元素的某些属性排序的顺序队列PriorityQueue

内部结构
+ 可重入锁
+ 根据元素delay时间排序的优先队列PriorityQueue
+ 优化阻塞通知的leader线程
+ 实现阻塞和通知的Condition对象

---
ThreadLocal是一个包装类，通过ThreadLocal访问线程对象自身的线程局部变量

ThreadLocal的get方法（入参为当前线程），获取当前线程对应的线程局部变量
get方法-》从当前线程的threadlocalmap里 -》取出线程局部变量
set方法-》第一次时，为当前线程的threadlocal变量初始化设置值

+ 每个线程的变量副本是存储在哪里的？
  + 存在线程对象自身的threadlocal变量中
+ 源码中的threadlocal的初始值是什么时机设置的？
  + ThreadLocal的set方法，当线程第一次调用set方法时，给当前线程的threadlocal变量创建ThreadLocalMap并设置相应的值

---

CopyOnWriteArrayList， CopyOnWriteArraySet
  + 写时复制容器，添加完成后，使得原容器的引用执行新的容器
  + 读写分离的思想，读和写在不同的容器上

ConcurrentHashMap

----

+ CAS (compare and swap)
  + ABA问题，变量记录版本号，AtomicStampedReference类
  + 可能在CPU的自旋时间很长
  + 只能保证一个共享变量的原子操作
  + AtomicInteger利用CAS

---

+ volatile
  + 可见性
  + 禁止java编译器指令重排

+ synchronized
  + 锁代码块， 插入monitorenter，monitorexist
  + 锁方法， 设置方法区中方法的标志位ACC_SYNCHRONIZED

---

+ ReentrantLock

+ AQS

---
