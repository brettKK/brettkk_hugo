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

+ UnsupportedOperationException 运行时异常的子类
  + 对不可修改的集合做修改时会出现这个异常 
