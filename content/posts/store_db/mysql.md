---
title: "mysql 小结"
date: 2021-05-05T11:33:56+08:00
draft: false
isCJKLanguage: true
tags: ["存储"]
series: [""]
categories: ["技术"]
---
## 存储

存储引擎-是创建表时设置，存储引擎是表级别的使用
+ innodb: 只有数据文件即时索引文件，聚集索引，一个表对应一个文件
+ mysiam: 既有数据文件，又有索引文件

+ InnoDB表中数据按行存储 类似网络协议栈，由页头，区头，段头，表头
扇区：磁盘的最小存储单位 512b；

磁盘块：文件系统读写数据的最小单位 4KB；

页：内存的最小存储单位 4KB； 16k，包含多行

一个区由64个页组成，为1M

一个段由多个区组成，构成表空间文件ibd （数据段、索引段、回滚段）

+ mysql innodb存储引擎的文件结构
	+ 共享表空间文件（ibdata1）
	+ 独占表空间文件（ibd）
	+ 表结构文件（user.frm， 表名 + .frm组成整个名称），与具体存储引擎无关
	+ 日志文件（redo文件等）

机械磁盘

+ RAID 0 
+ RAID 1 磁盘镜像，可靠性
+ RAID 5 
+ RAID 10  RAID 1 后 RAID 0  

固态硬盘 Sata SSD， PCI-E SSD



---
## 索引

+ 红黑树 二叉树，深度可能很深
+ B树 mongodb，允许一个树节点中存多个数据，中间节点存了数据
+ B+树 innodb
  + 只有叶子节点存数据，叶子节点左从到右，从小到达，叶子之间有指针，方便区间查询

## sql优化

+ 搜索字段上建索引，不在索引字段上使用函数
+ 避免select * ,  这样杜绝去使用覆盖索引
	+ 在innodb引擎下，select a,b,c 可以走辅助索引找到所有的数据，避免走到聚集索引上，减少一次聚集索引查询过程
	+ 减少数据传输量，limit 好习惯
+ 不使用like ‘%xx’, 多用区间操作, 能用 between 就不要用 in 	
+ 复合索引最左前缀 (不是指SQL语句的where顺序要和复合索引一致)
	+ 复合索引底层存储也是一颗B+树，只不过排序key=（a,b,c）, 这样是复合索引使用时需要满足最左前缀的原理
+ 避免大事务（避免夹杂rpc）
+ 避免使用子查询， 子查询产生的临时表再扫描无索引可走，会全表扫描
+ explain 显示select语句的执行计划
	+ select_type: simple等
	+ type： 判断查询是否高效，const, ref, range, index,all
	+ extra： using index, using filesort(不能通过索引达到排序), using temporary(使用了临时表)


## 锁

+ mysql 锁
	+ mysql锁类型，锁机制 （InnoDB 检索数据走索引是行锁，不走索引是表锁）
		+ 行级锁
			+ Record Lockx 对索引项加锁，锁定一条记录
			+ Gap Lockx 锁定一个范围，不包含记录本身
			+ Next-Key Lock = gap + record lock, 锁定一个范围 并锁定记录本身
				+ 查询的索引包含唯一属性时，next-key lock 降级为record lock
		+ 表级锁
+ 死锁的必要条件
	+ 互斥
	+ 请求并保留
	+ 不剥夺
	+ 循环等待
+ 死锁的注意事项：
	+ InnoDB的索引的修改会锁定更多的行
	+ InnoDB加锁是边查询边加锁，逐行获得锁
	+ InnoDB表的并发过多，会达到锁的最大深度，或者锁的最大
	+ 锁的升级也会导致死锁，建议事务开始就获得所有资源



## 事务原理


+ mysql transaction （ACID 锁保证隔离性， relog 保证原子性和持久性， undo保证一致性）
	+ 隔离级别 可在终端进行模拟
		+ read uncommited （tranx will read dity data）
			+ 读脏数据
		+ read committed (no repeatable read)
			+ 不可重复读（事务中两次读的结果不一致）
			+ 事务1等读到事务2的commit， 从数据库对事务ACID角度上讲，违背了隔离性
		+ repeatable read (default mode)
			+ 事务在运行时看到相同的数据，可重复读
			+ 存在幻读（innodb的间隙锁解决）
		+ seriablizable
	+ mvcc
    	+ Read View生成时机的不同，从而造成RC,RR级别下快照读的结果的不同
    	+ 对某条记录的第一次快照读会创建一个快照及Read View，之后的快照读获取的都是同一个Read View， RR
    	+ 每次快照读都会新生成一个快照和Read View, 在RC级别下的事务中可以看到别的事务提交的更新
+ innodb重要的日志，undo logs, redo logs
	+ undo logs: 记录某数据被修改前的值，可以用来在事务失败时进行rollback，保证一致性
		+ 逻辑日志，随机写，每行记录修改， 采用段的方式
		+ 帮助事务回滚以及MVCC功能 
	+ redo logs: 保证事务的原子性和持久性
		+ 物理日志，顺序写，512字节，记录的是页的物理修改操作
		+ 确保redo log写入磁盘，事务提交前需要fsync操作。 因此磁盘的性能决定了事务提交的性能。
+ undo+redo log
	+ 为保证事务的持久性，事务提交前先写redo log持久化（WAL）
	+ 数据不需要在事务提交前写入磁盘，而是在内存的缓存中
	+ redo log保证事务的持久性， undo log保证事务的原子性

---
## 数据同步

+ Master-Slave 原理 （Binlog, Relay log）
	+ 异步复制
		+ 主库
			+ 当从库连接式，binlog dump thread发送binlog event,并持有event的lock
			+ IO thread 读完后释放lock 	 
		+ 从库
			+ IO thread读binlog event，生成relay log
			+ SQL thread 读relay log，执行
	+ Binlog, Relay log
		+ SQL(增删改查，DDL)语句按 提交顺序 生成的日志，用于数据回滚，异步复制
		+ Relay log通过异步复制Master中的binlog，生成slave的执行日志
		+ Binlog的3种存储格式：binlog_format(row, statement, mixed)
			+ statement格式： 每一条会修改数据的sql都会记录在binlog中
			+ row： 不记录sql语句上下文相关信息，仅保存哪条记录被修改
			+ mixed： 混合	（Binlog记录数据库发生的变化,用于replication）
---
## 使用

数据库使用自增主键作为唯一key时可能的问题： 分库分表时，会出现主键重复。

+	字符类 (char,varchar,text,blob)
  + varchar(50)表示最大存放50个字符, char
  + int(10)表示最大显示宽度为10,存储还是占4字节
+ 时间日期类
  + date, '1000-01-01' to '9999-12-31'
  + datetime, '1000-01-01 00:00:00' t0 '9999-12-31 00:00:00',范围广
  + timestamp, '1970-01-01 00:00:01' UTC to '2038-01-19 03:14:07' UTC
