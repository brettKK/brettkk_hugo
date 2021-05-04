control cgroup

/sys/fs/cgroup

cgroup把每种资源定义为子系统 典型的有以下几种
+ cpu
+ cpuacct
+ cpuset
+ memory
+ blkio
+ devices
+ net_cls
+ freezer
+ ns 

+ who use cgroup
  + docker ( )
  + rkt
  + LXC
  + YARN (Hadoop)
  + systemd
  + OpenVZ
  + Tupperware(FB) 
  + runc (libcontainer -> runc)
