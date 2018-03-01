1.关闭 swap
具体操作：
vi /etc/fstab
注掉/删除  /home/swap swap swap default 0 0
或者在KUBELET_ARGS参数中添加： --fail-swap-on=false

2.删除原来的KUBELET_API_SERVER ， 通过/etc/kubernetes/kubelet.kubeconfig向apiserver中的server发起注册

3.修改启动参数
原1.6:
KUBELET_ARGS="--kubeconfig=/etc/kubernetes/kubelet.kubeconfig --pod-manifest-path=/etc/kubernetes/manifests --log-dir=/etc/kubernetes/logs --cluster-dns=10.254.0.10 --cluster-domain=cluster.local --logtostderr=true --allow-privileged=true"

1.8:
KUBELET_ARGS="--cgroup-driver=cgroupfs --port=10250 --log-dir=/etc/kubernetes/logs --cluster-dns=10.254.0.10 --enable-controller-attach-detach=true --logtostderr=false --address=0.0.0.0 --hostname-override=3.3.3.198 --allow-privileged=true --kubeconfig=/etc/kubernetes/kubelet.kubeconfig --pod-manifest-path=/etc/kubernetes/manifests"
