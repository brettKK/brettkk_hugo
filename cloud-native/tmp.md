+ is power of 2
```
func isPowerOf2(n int) bool {
    return n != 0 && n &(n-1) == 0
}

```
Go语言的sync/atomic包

```
type Once struct {
    m    Mutex
    done uint32
}
func (o *Once) Do(f func()) {
    if atomic.LoadUint32(&o.done) == 1 {
        return
    }
    // Slow-path.
    o.m.Lock()
    defer o.m.Unlock()
    if o.done == 0 {
        defer atomic.StoreUint32(&o.done, 1)
        f()
    }
}
```

用原子操作来替换mutex锁
其主要原因是，原子操作由底层硬件支持，而锁则由操作系统提供的API实现。若实现相同的功能，前者通常会更有效率



1 排序后根据index取
public int findKthLargest(int[] nums, int k) {
        int len = nums.length;
        Arrays.sort(nums);
        return nums[len - k];
}

2 小顶堆 装数组，poll (len-k-1)后， peek为第k大的元素

3 大顶堆 装数组， poll (k-1)后， peek为第k大的元素

4 容量为k的小顶堆， peek为第k大的元素

//gorename
// CMPQ
// JLS Jump on lower or the same


+ 常见漏洞
+ 敏感信息
+ 敏感逻辑 支付 下单 发红包
+ 签名方式 
+ 传输过程 https 劫持风险
+ redis 内存回收机制
	+ 过期回收
		+ 定时过期回收
		+ 懒性过期回收
		+ 定期过期回收
	+ 内存上限 主动回收
		+ LRU
+ 

+ 一致性 
    + 强一致性 （复制是同步）
    + 若一致性  （复制是异步）
    + 使一个从库是同步的，而其他的则是异步的 
    + 对于某些特定的内容，都从主库读
    + 数据一致性 （复制是导致出现数据一致性问题的唯一原因）
        + 最终一致性
            + 因果一致性
            + 读已之所写
            + 会话一致性
            + 单调读一致性
            + 单调写一致性
    + 事务一致性（ACID）
