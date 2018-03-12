# istio
***[overview of istio](https://istio.io/docs/concepts/what-is-istio/overview.html)***

istio serves a network that's  load balancing, service-to-service authentication, monitoring, and more.
we need not change any code of microservices. =)

## preparation
1. install kubernetes, >= 1.9 will be better.

in fact, i recommand to use ./hack/local-up-cluster.sh to start a local cluster for testing.

2. get binaries and configs from git.io

```shell
curl -L https://git.io/getLatestIstio | sh -
```

add path into ~/.profile and remember to source it to make it work at once!:

export PATH="$PATH:/home/cloud/istio-0.6.0/bin"

## installation

```shell
kubectl apply -f install/kubernetes/istio.yaml
...

root@ubuntu:/home/cloud/istio-0.6.0/install/kubernetes# kb get pods --all-namespaces
NAMESPACE      NAME                             READY     STATUS    RESTARTS   AGE
istio-system   istio-ca-7754766889-9cg8p        1/1       Running   0          1m
istio-system   istio-ingress-544c6657bd-fk8rr   1/1       Running   0          1m
istio-system   istio-mixer-59c44f5fb7-7ttwd     3/3       Running   0          1m
istio-system   istio-pilot-d8ff96dc8-8j6m6      2/2       Running   0          1m
kube-system    kube-dns-5c6c5b55b-8pd2q         3/3       Running   0          7m

```

## Verifying the installation

```
root@ubuntu:/home/cloud/istio-0.6.0/install/kubernetes# kubectl get svc -n istio-system
NAME            TYPE           CLUSTER-IP   EXTERNAL-IP   PORT(S)                                                            AGE
istio-ingress   LoadBalancer   10.0.0.101   <pending>     80:31661/TCP,443:31485/TCP                                         13m
istio-mixer     ClusterIP      10.0.0.158   <none>        9091/TCP,15004/TCP,9093/TCP,9094/TCP,9102/TCP,9125/UDP,42422/TCP   13m
istio-pilot     ClusterIP      10.0.0.194   <none>        15003/TCP,8080/TCP,9093/TCP,443/TCP
```
Note: If your cluster is running in an environment that does not support an external load balancer (e.g., minikube), the EXTERNAL-IP of istio-ingress says <pending>. You must access the application using the service NodePort, or use port-forwarding instead.
  

