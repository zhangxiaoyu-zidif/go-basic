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


## 创建pod方式
Kubelet负责pod的创建，pod的来源kubelet当前支持三种类型的podSource 
- FileSource: 通过kubelet的启动参数–pod-manifest-path来指定pod manifest文件所在的路径或者文件都可以．Kubelet会读取文件里面定义的pod进行创建．常常我们使用来定义kubelet管理的static pod 
- HTTPSource: 通过kubelet的启动参数–manifest-url –manifest-url-header来定义manifest url. 通过http Get该manifest url获取到pod的定义 
- ApiserverSource: 通过定义跟kube-apiserver进行通过的kubeclient, 从kube-apiserver中获取需要本节点创建的pod的信息．
