# what is cgroup?

[cgroup](https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/Documentation/cgroup-v1/cgroups.txt)

control group，控制一组进程对于资源的使用。cgroup实际是上绑定一个进程集合到一个或者多个子系统（subsystem）。

参考：https://blog.csdn.net/huang987246510/article/details/80765628


向用户提供操作接口：虚拟文件系统类型（cgroupfs）。

所有对于cgroup的修改都是在cgroupfs中完成的。

# How could we do resource management?
