# what is cgroup?

[cgroup](https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/Documentation/cgroup-v1/cgroups.txt)

control group，控制一组进程对于资源的使用。cgroup实际是上绑定一个进程集合到一个或者多个子系统（subsystem）。

参考：https://blog.csdn.net/huang987246510/article/details/80765628


向用户提供操作接口：虚拟文件系统类型（cgroupfs）。

所有对于cgroup的修改都是在cgroupfs中完成的。


## CPU

### prequist

狂跑CPU的脚本：
```shell
x=0
while [ True ];do
    x=$x+1
done;
```

### cfs_quota_us


### cfs_period_us


### cpu.rt_period_us


### cpu.rt_runtime_us


### nr_periods


### nr_throttled


### cpuset

#### cpuset.cpus 和 cpuset.mems 

用来限制进程可以使用的 cpu 核心和内存节点的。
这两个参数中 cpu 核心、内存节点都用 id 表示，之间用 “,” 分隔。比如 0,1,2 。也可以用 “-” 表示范围，如 0-3 。两者可以结合起来用。如“0-2,6,7”。在添加进程前，cpuset.cpus、cpuset.mems 必须同时设置，而且必须是兼容的，否则会出错

### cpuacct.stat

统计了该控制组中进程用户态和内核态的 cpu 使用量，单位是 USER_HZ，也就是 jiffies、cpu 滴答数。每秒的滴答数可以用 getconf CLK_TCK 来获取，通常是 100。将看到的值除以这个值就可以换算成秒。


## Memory

### prequist

```shell
狂耗费Memory的脚本：
x="a"
while [ True ];do
    x=$x$x
done;
```

## Disk

```shell
dd if=/dev/sda of=/dev/null 
通过iotop看io占用情况，磁盘速度到了284M/s

30252 be/4 root      284.71 M/s    0.00 B/s  0.00 %  0.00 % dd if=/dev/sda of=/dev/null
```
验证io读写


## Net


# How could we do resource management?
