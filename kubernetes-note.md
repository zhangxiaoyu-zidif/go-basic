#### 打印制定的pod name和uid
```shell
kubectl get pods -o=jsonpath='{range .items[*]}{.metadata.name}{"\t\t\t"}{.metadata.uid}{"\n"}'
```

1.关闭 swap
具体操作：
vi /etc/fstab
注掉/删除  /home/swap swap swap default 0 0
或者在KUBELET_ARGS参数中添加： --fail-swap-on=false

2.删除原来的KUBELET_API_SERVER ， 通过/etc/kubernetes/kubelet.kubeconfig向apiserver中的server发起注册

3.修改启动参数
原1.6:
KUBELET_ARGS="--kubeconfig=/etc/kubernetes/kubelet.kubeconfig --pod-manifest-path=/etc/kubernetes/manifests --log-dir=/etc/kubernetes/logs --cluster-dns=10.254.0.10 --cluster-domain=cluster.local --logtostderr=true --allow-privileged=true"


通过curl命令创建pod的demo：
```shell
curl --header "Content-Type: application/json" \
--request POST \
--data '{"apiVersion": "v1", "kind": "Pod", "metadata": { "name": "nginx-aaaa", "namespace": "default"}, "spec": { "containers": [{"name": "aaaa","image": "gcr.io/google_containers/busybox:latest","imagePullPolicy": "Never"}]}}' \
http://103.103.103.183:8080/api/v1/namespaces/default/pods
```

1.8:
KUBELET_ARGS="--cgroup-driver=cgroupfs --port=10250 --log-dir=/etc/kubernetes/logs --cluster-dns=10.254.0.10 --enable-controller-attach-detach=true --logtostderr=false --address=0.0.0.0 --hostname-override=3.3.3.198 --allow-privileged=true --kubeconfig=/etc/kubernetes/kubelet.kubeconfig --pod-manifest-path=/etc/kubernetes/manifests"





```shell
root      1066     1  1 09:03 ?        00:01:00 /usr/bin/kube-proxy --logtostderr=true --v=8 --master=http://127.0.0.1:8080
root      1375     1  0 09:03 ?        00:00:06 /usr/bin/flanneld -etcd-endpoints=http://all-in-one:2379 -etcd-prefix=/kube-centos/network
kube      7099     1  1 09:40 ?        00:00:15 /usr/bin/kube-scheduler --logtostderr=true --v=8 --master=http://127.0.0.1:8080
root     10193  7793 16 09:52 pts/0    00:00:29 /root/gopath/src/k8s.io/kubernetes/_output/local/bin/linux/amd64/hyperkube apiserver --authorization-mode=Node,RBAC --runtime-config=admissionregistration.k8s.io/v1alpha1,settings.k8s.io/v1alpha1 --cloud-provider= --cloud-config= --v=3 --vmodule= --cert-dir=/var/run/kubernetes --client-ca-file=/var/run/kubernetes/client-ca.crt --service-account-key-file=/tmp/kube-serviceaccount.key --service-account-lookup=true --enable-admission-plugins=Initializers,LimitRanger,ServiceAccount,SecurityContextDeny,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota,PodPreset --disable-admission-plugins= --admission-control-config-file= --bind-address=0.0.0.0 --secure-port=6443 --tls-cert-file=/var/run/kubernetes/serving-kube-apiserver.crt --tls-private-key-file=/var/run/kubernetes/serving-kube-apiserver.key --insecure-bind-address=127.0.0.1 --insecure-port=8080 --storage-backend=etcd3 --etcd-servers=http://127.0.0.1:2379 --service-cluster-ip-range=10.0.0.0/24 --feature-gates=AllAlpha=false --external-hostname=localhost --requestheader-username-headers=X-Remote-User --requestheader-group-headers=X-Remote-Group --requestheader-extra-headers-prefix=X-Remote-Extra- --requestheader-client-ca-file=/var/run/kubernetes/request-header-ca.crt --requestheader-allowed-names=system:auth-proxy --proxy-client-cert-file=/var/run/kubernetes/client-auth-proxy.crt --proxy-client-key-file=/var/run/kubernetes/client-auth-proxy.key --cors-allowed-origins=/127.0.0.1(:[0-9]+)?$,/localhost(:[0-9]+)?$
root     10489  7793  0 09:52 pts/0    00:00:00 sudo /root/gopath/src/k8s.io/kubernetes/_output/local/bin/linux/amd64/hyperkube proxy --config=/tmp/kube-proxy.yaml --master=https://localhost:6443 --v=3
root     10490  7793  4 09:52 pts/0    00:00:06 /root/gopath/src/k8s.io/kubernetes/_output/local/bin/linux/amd64/hyperkube scheduler --v=3 - kubeconfig /var/run/kubernetes/scheduler.kubeconfig --feature-gates=AllAlpha=false --master=https://localhost:6443
root     10504 10489  0 09:52 pts/0    00:00:01 /root/gopath/src/k8s.io/kubernetes/_output/local/bin/linux/amd64/hyperkube proxy --config=/tmp/kube-proxy.yaml --master=https://localhost:6443 --v=3
root     10664  7793  0 09:52 pts/0    00:00:00 sudo -E /root/gopath/src/k8s.io/kubernetes/_output/local/bin/linux/amd64/hyperkube kubelet --v=3 --vmodule= --chaos-chance=0.0 --container-runtime=docker --rkt-path= --rkt-stage1-image= --hostname-override=127.0.0.1 --cloud-provider= --cloud-config= --address=127.0.0.1 --kubeconfig /var/run/kubernetes/kubelet.kubeconfig --feature-gates=AllAlpha=false --cpu-cfs-quota=true --enable-controller-attach-detach=true --cgroups-per-qos=true --cgroup-driver=systemd --keep-terminated-pod-volumes=true --eviction-hard=memory.available<100Mi,nodefs.available<10%,nodefs.inodesFree<5% --eviction-soft= --eviction-pressure-transition-period=1m --pod-manifest-path=/var/run/kubernetes/static-pods --fail-swap-on=false --cluster-dns=10.0.0.10 --cluster-domain=cluster.local --port=10250
root     10666 10664  5 09:52 pts/0    00:00:08 /root/gopath/src/k8s.io/kubernetes/_output/local/bin/linux/amd64/hyperkube kubelet --v=3 --vmodule= --chaos-chance=0.0 --container-runtime=docker --rkt-path= --rkt-stage1-image= --hostname-override=127.0.0.1 --cloud-provider= --cloud-config= --address=127.0.0.1 --kubeconfig /var/run/kubernetes/kubelet.kubeconfig --feature-gates=AllAlpha=false --cpu-cfs-quota=true --enable-controller-attach-detach=true --cgroups-per-qos=true --cgroup-driver=systemd --keep-terminated-pod-volumes=true --eviction-hard=memory.available<100Mi,nodefs.available<10%,nodefs.inodesFree<5% --eviction-soft= --eviction-pressure-transition-period=1m --pod-manifest-path=/var/run/kubernetes/static-pods --fail-swap-on=false --cluster-dns=10.0.0.10 --cluster-domain=cluster.local --port=10250
```
