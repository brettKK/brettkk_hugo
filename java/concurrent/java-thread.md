+ Runnable, Callable, Future, FutureTask 区别

```
public interface Runnable{
  public void fun();
  //  返回值为void，所以执行任务后无法返回任何结果
}
```

----

```
publiic interface Callable<V>{
  V call() throws Exception;
  // 返回值为V类型，可以跑出异常，配合ExecutorService使用
}
```

```
public interface Future<V> {
  // 取消任务， 参数
  boolean cancel(boolean mayInterruptIfRunning);

  // 任务是否已取消成功
  boolean isCancelled();

  // 任务是否已完成
  boolean isDone();

  // get方法会阻塞，直到任务返回结果
  V get() throws InterruptedException, ExecutionException;

  // 在指定的时间内获取结果，否则返回null
  V get(long timeout, TimeUnit unit) throws InterruptedException, ExecutionException, TimeoutException;
}
```

Future是对具体的Runnable, Callable任务的执行结果进行取消，查询是否完成，获取结果，设置结果操作
Future有3个功能：判断任务是否完成，能够中断任务，能够获取任务的执行结果

RunnableFuture继承了Runnable接口和Future接口
```
public interface RunnableFutrue<V> extends Runnable, Future<V> {
  void run();
}
```

FutureTask实现了RunnableFuture接口，则既可以作为Runnable被执行，又可以作为Future来得到Callable的返回
```
public class FutureTask<V> implements RunnableFuture<V> {
...
}
```

Executor是Runnable和Callable的调度容器

Executors framework主要接口和类：

+ Executor (interface)
  + void execute(Runnable task); 基本接口
+ ExecutorService  (interface)
    + extends Executor ， 它扩展了Executor接口，提供更多的接口方法
    + 一个Executor的生命周期是：运行，关闭，终止
    + Future<?> submit(Runnable task); 2种提交的方法
    + <T> Future<T> submit(Callable<T task);
    + void shutdown(); 平滑关闭ExecutorService
    + ....
+ Executors类
  + 工厂类
  + 提供一系列工厂方法创建线程池，返回的线程池均实现了ExecutorService接口
  + ExecutorService new FixedThreadPool(int threads);
  + some common Executor implementations
    + ThreadPoolExecutor
      + worker threads in pool
      + inner class Worker
    + ForkJoinPool
    + ExecutionCompletionService
      + BlockingQueue<E>
+ Future<V>  (interface)
  + ...

+ ThreadPoolExecutor的execute方法
  + worker数量小于corePoolSize时，新建worker对处理提交的任务
  + 如果第一步失败，将提交的任务放入阻塞队列
  + 第二步失败，worker数量小于maxmumPoolSize, 新建worker处理提交的任务
  + 第三步失败，调用拒绝策略处理该任务
+ addWorker方法
  + 新建一个Woker线程
  + 加入到woker集合中
  + 启动（start）worker线程

----
![image](executor_framewok.png
----
+ 中断正在执行的线程
+ 如何中断正在io阻塞的线程
  + 通过Future对象调用cancel失效，线程无法响应Interrupt
+ 如何中断正在执行nio的线程
  + cancel生效, 线程可以响应Interrupt

+ ABC三个线程如何保证顺序执行
  + 设置int value， lock， value=1, value=2, value=3
  + join方法
+ AB线程交替打印奇数和偶数
