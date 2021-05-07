+ Object通知线程的方法
  + wait, 需要先获得对象锁，所以方法需放到同步代码中。一旦调用wait方法，线程会释放对象锁
    + wait方法有3种重载方式， 不同的参数的wait方法
    + 从jdk文档上看， 4种可能唤醒线程的方式
      + notify/notifyAll
      + Interrupt
      + wait的时间timeout
    + 可能跑出3种异常
      + IllegalArgumentException, 当wait的时间参数为负数时
      + IllegalMonitorStateException, 当前线程没有拥有obj的监视器（obj.wait()）
      + InterrupedException, 有其他线程中断当前线程
  + notify， notifyAll， 需要先获得对象锁，需放到同步代码中
  + 若没有获取对象锁，调用wait和notify，抛IllegalMonitorStateException(属于runtime exception)
  + 不要在String或者Global object上调用wait方法, 不同线程都去wait同一个对象上的signal
+ 拥有对象的监视器的三种方式
  + 执行对象的一个synchronized instance方法
  + syhchronized(obj)
  + 执行对象对应的Class的一个static synchronized方法

+ 抛InterruptedException异常
  + 说明该方法耗时， 但是可以被中断
  + 代表方法
    + Object类 wait方法
    + Thread类 sleep
    + Thread类 join
    + 这些方法内容会不断去检查中断状态，从而跑出中断异常

+ 中断线程
  + 通过Thread.interrupt(),修改线程的中断状态来告诉线程
  + 对于可中断的线程，比如Thread.sleep(), Object.wait(), Thread.join(),收到中断信号后，抛出InterruptedException, 同时将中断状态改为true
+ 线程进入阻塞状态
  + sleep
  + wait
  + 等IO
  + 进入同步块
+ 想终止处于阻塞的线程
  + 

+ 锁池
  + 线程A，B，C需要获取D对象的锁，则A，B，C进入D对象的锁池
+ 等待池
  + 线程A调用B对象的wait方法，线程A就会释放该对象的锁，然后进入B对象的等待池中
