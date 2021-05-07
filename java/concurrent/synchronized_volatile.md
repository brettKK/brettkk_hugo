+ synchronized锁的优化 -> 锁级别从低到高：无锁，偏向锁 轻量锁(CAS) 重量锁(依赖OS的mutex lock)

+ volatile
  + 保证共享变量在线程间的可见性
  + 保证happens-before
  + 什么时候使用volatile变量
    + 一个线程读写，其他线程只读
    + 所有线程对变量的写不依赖与之前的状态；例如boolean的写入，温度等

1 可见性

线程中读取volatile变量不从cpu的cache中取， 而是从主存中取指，写volatile变量不仅写cpu的cache，并且写到主存中

![image](image/java-volatile-1.png)

2 防止指令重排

---
+ java锁
  + synchronized
  + Lock接口 AQS
    + 独占锁 ReetrantLock 只能有一个线程拿到锁
    + 共享锁 CountDownLatch 可以多个线程拿到锁

---
+ java 内存模型 语义

---
+ 多线程死锁
+ 数据库死锁（事务之间）

+ 避免死锁
  + 加锁顺序
  + 加锁时间
  + 死锁检测
