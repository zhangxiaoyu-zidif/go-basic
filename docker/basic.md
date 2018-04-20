### 1.12.6 docker由 docker-client ,dockerd,containerd,docker-shim,runc组成

dockerd本身实属是对容器相关操作的api的最上层封装，直接面向操作用户。

dockerd实际真实调用的还是containerd的api接口（rpc方式实现），containerd是dockerd和runc之间的一个中间交流组件。

docker-shim是一个真实运行的容器的真实垫片载体，每启动一个容器都会起一个新的docker-shim的一个进程， 
他直接通过指定的三个参数：容器id，boundle目录（containerd的对应某个容器生成的目录，一般位于：/var/run/docker/libcontainerd/containerID）

runc是一个命令行工具端，他根据oci（开放容器组织）的标准来创建和运行容器.
