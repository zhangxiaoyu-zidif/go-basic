# kubelet


## 主要功能
Kubelet组件运行在Node节点上，维持运行中的Pods以及提供kuberntes运行时环境，主要完成以下使命： 
１．监视分配给该Node节点的pods 
２．挂载pod所需要的volumes 
３．下载pod的secret 
４．通过docker/rkt来运行pod中的容器 
５．周期的执行pod中为容器定义的liveness探针 
６．上报pod的状态给系统的其他组件 
７．上报Node的状态 
