1 排序后根据index取
public int findKthLargest(int[] nums, int k) {
        int len = nums.length;
        Arrays.sort(nums);
        return nums[len - k];
}

2 小顶堆 装数组，poll (len-k-1)后， peek为第k大的元素

3 大顶堆 装数组， poll (k-1)后， peek为第k大的元素

4 容量为k的小顶堆， peek为第k大的元素

+ 常见漏洞
+ 敏感信息
+ 敏感逻辑 支付 下单 发红包
+ 签名方式 
+ 传输过程 https 劫持风险

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


3Blue1Brown的系列视频 线性代数的本质
