# cgroup相关测试研究

## case 1: 修改container的memory limit值

测试步骤：
1. 下载一份kubernetes代码
2. 配置go 1.10.3的环境变量
3. 在kubernetes代码的文件夹下执行，./hack/install-etcd.sh，并配置etcd的环境变量
4. 相同文件夹下，执行./hack/local-up-cluster.sh，并配置相关环境参数
5. kubectl create -f 以下文件：
```yaml
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: busybox-apps-cp
  labels:
    app: busybox-apps-cp
spec:
  replicas: 1
  selector:
    matchLabels:
      app: busybox-apps-v1beta1
  template:
    metadata:
      labels:
        app: busybox-apps-v1beta1
    spec:
      containers:
      - name: busybox-3
        image: busybox:1.25
        imagePullPolicy: Never
        resources:
          limits:
            memory: "200Mi"
          requests:
            cpu: 0.1
            memory: "100Mi"
        command: ["/bin/sh"]
        args: ["-c", "while true; do echo hello; sleep 10;done"]
      - name: busybox-4
        image: busybox:1.25
        imagePullPolicy: Never
        resources:
          limits:
            memory: "100Mi"
          requests:
            cpu: 0.1
            memory: "100Mi"
        command: ["/bin/sh"]
        args: ["-c", "while true; do echo hello; sleep 10;done"]
```
6. 通过docker ps命令确认需要修改memory limit的container的container id。这个还需要配合kubectl get pod xxx -o json的container status内容一起看。

**memory-swap和memory设置成相同就是为了让container无法使用多余swap的值**
* Memory and MemorySwap are set to the same value, this prevents containers from using any swap.

```shell
docker update --memory-swap 300M --memory 300M dc5a14edf475
```
7. 通过docker stats命令
```shell
CONTAINER           CPU %               MEM USAGE / LIMIT     MEM %               NET I/O             BLOCK I/O           PIDS
0f3012abd99f        0.08%               6.465 MiB / 170 MiB   3.80%               972 kB / 2.42 MB    0 B / 0 B           8
316b9240603c        0.00%               80 KiB / 300 MiB      0.03%               648 B / 648 B       0 B / 0 B           2
dc5a14edf475        0.00%               76 KiB / 300 MiB      0.02%               648 B / 648 B       0 B / 0 B           2
bdbf99f6fd36        0.00%               44 KiB / 3.701 GiB    0.00%               648 B / 648 B       0 B / 0 B           1
2bf99e0874c4        0.10%               312 KiB / 200 MiB     0.15%               1.3 kB / 648 B      0 B / 0 B           2
01254140df02        0.13%               172 KiB / 200 MiB     0.08%               1.3 kB / 648 B      0 B / 0 B           2
6de9c3a6216f        0.00%               44 KiB / 3.701 GiB    0.00%               1.3 kB / 648 B      0 B / 0 B           1
69ff14c1aa29        0.00%               40 KiB / 3.701 GiB    0.00%               972 kB / 2.42 MB    1.46 MB / 0 B       1
```

8. 查看pod下的cgroup信息
```shell
ls /sys/fs/cgroup/memory/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-podcdd85f83_7871_11e8_9fa0_080027f4a74d.slice
```
这个文件夹下的memory.limit_in_bytes没有发生变化。

9. 查看container的cgroup信息
```shell
ls /sys/fs/cgroup/memory/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-podcdd85f83_7871_11e8_9fa0_080027f4a74d.slice/docker-dc5a14edf4758b21542537a97d46d100ac478a8f7db317cf3e0d90e3608b46ec.scope
```

可以发现memory的limit值是被修改的
```shell
[root@localhost docker-dc5a14edf4758b21542537a97d46d100ac478a8f7db317cf3e0d90e3608b46ec.scope]# cat memory.limit_in_bytes 
314572800
[root@localhost docker-dc5a14edf4758b21542537a97d46d100ac478a8f7db317cf3e0d90e3608b46ec.scope]# cat memory.memsw.limit_in_bytes 
314572800
```
