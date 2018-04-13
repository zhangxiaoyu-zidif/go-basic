# CKA

Certified Kubernetes Administrator (CKA) 

## syllabus

### first of all
we should know how to install a cluster and create/distribute the rsa .etc

ref here: https://github.com/markthink/cka-kubernetes

### 5% - Scheduling

#### Use label selectors to schedule Pods.
[ref here](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)

one usage scenario：

```shell
kubectl label node 127.0.0.1 accelerator=nvidia-tesla-p100
```
some label limits:


```yaml
apiVersion: v1
kind: Pod
metadata:
  name: cuda-test
spec:
  containers:
    - name: cuda-test
      image: "k8s.gcr.io/cuda-vector-add:v0.1"
      resources:
        limits:
          nvidia.com/gpu: 1
  nodeSelector:
    accelerator: nvidia-tesla-p100
```

#### Understand the role of DaemonSets.
A DaemonSet ensures that all (or some) Nodes run a copy of a Pod.
[ref here](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/)

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: busybox-apps-v1
  labels:
    app: busybox-apps-v1  # important
spec:
  selector:
    matchLabels:
      name: busybox-apps-v1  # important
  template:
    metadata:
      labels:
        name: busybox-apps-v1
    spec:
      containers:
      - name: busybox-apps-v1
        image: busybox
        imagePullPolicy: IfNotPresent
```


#### Understand how resource limits can affect Pod scheduling.
[ref here](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/)

When you create a Pod, the Kubernetes scheduler selects a node for the Pod to run on. Each node has a maximum capacity for each of the resource types: the amount of CPU and memory it can provide for Pods. The scheduler ensures that, for each resource type, the sum of the resource requests of the scheduled Containers is less than the capacity of the node. Note that although actual memory or CPU resource usage on nodes is very low, the scheduler still refuses to place a Pod on a node if the capacity check fails. This protects against a resource shortage on a node when resource usage later increases, for example, during a daily peak in request rate.

actual memory or CPU resource = maximum allocable capacity - sum of all requested resources.
It doest not care the really actual usage.


#### Understand how to run multiple schedulers and how to configure Pods to use them.
```shell
# create docker file
cat <<EOF > DOCKERFILE
FROM busybox
ADD ./_output/dockerized/bin/linux/amd64/kube-scheduler /usr/local/bin/kube-scheduler
EOF
docker build -t my-kube-scheduler:1.0 .
docker load -i my-kube-scheduler:1.0

```

if we set RBAC to enable, we should do， if not, skip it.
```shell
$ kubectl edit clusterrole system:kube-scheduler
- apiVersion: rbac.authorization.k8s.io/v1
  kind: ClusterRole
  metadata:
    annotations:
      rbac.authorization.kubernetes.io/autoupdate: "true"
    labels:
      kubernetes.io/bootstrapping: rbac-defaults
    name: system:kube-scheduler
  rules:
  - apiGroups:
    - ""
    resourceNames:
    - kube-scheduler
    - my-scheduler
    resources:
    - endpoints
    verbs:
    - delete
    - get
    - patch
    - update
```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: annotation-second-scheduler
  labels:
    name: multischeduler-example
spec:
  schedulerName: my-scheduler   # specify the scheduler name
  containers:
  - name: pod-with-second-annotation-container
    image: k8s.gcr.io/pause:2.0
```

#### Manually schedule a pod without a scheduler.
stop the kube-scheduler, and use `nodeName` to assign one pod to a specified node.
```yaml
kind: Pod
apiVersion: v1
metadata:
  name: test-pod
spec:
  nodeName: 192.165.1.72    # use nodeName to specify the node without kube-scheduler
  containers:
  - name: test-pod
    image: busybox:1.25
    imagePullPolicy: Never
  restartPolicy: "Never"

```

daemonset's replicas could deploy without scheduler.
```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: busybox-ds
  labels:
    k8s-app: busybox-ds
spec:
  selector:
    matchLabels:
      name: busybox-ds
  template:
    metadata:
      labels:
        name: busybox-ds
    spec:
      containers:
      - name: busybox-ds
        image: busybox:1.25
        resources:
          limits:
            memory: 200Mi
          requests:
            cpu: 100m
            memory: 200Mi
```

#### Display scheduler events.

```shell
kubectl get events
```

#### Know how to configure the Kubernetes scheduler.
[ref here](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/scheduling/scheduler_extender.md)

and [here](https://github.com/kubernetes/community/blob/master/contributors/devel/scheduler.md)

for example:
add startup parameter in kube-scheduler, with --policy-config-file="xxx.json"

xxx.json refer below:
```json
{
  "predicates": [
    {
      "name": "HostName"
    },
    {
      "name": "MatchNodeSelector"
    },
    {
      "name": "PodFitsResources"
    }
  ],
  "priorities": [
    {
      "name": "LeastRequestedPriority",
      "weight": 1
    }
  ],
  "extenders": [
    {
      "urlPrefix": "http://127.0.0.1:12345/api/scheduler",
      "filterVerb": "filter",
      "enableHttps": false
    }
  ]
}
```

### 5% - Logging/Monitoring
```shell
kubectl logs <pod name>
```
ref here: https://kubernetes.io/docs/concepts/cluster-administration/logging/


#### Understand how to monitor all cluster components.
ref: https://kubernetes.io/docs/tasks/debug-application-cluster/debug-cluster/


#### Understand how to monitor applications.

```shell
$ kubectl create -f deploy/kube-config/influxdb/
$ kubectl create -f deploy/kube-config/rbac/heapster-rbac.yaml

#and run:
kubectl top node/pod <nodename>/<podname>
```

#### Manage cluster component logs.
see: /var/log/<component name>.log

#### Manage application logs.
```shell
kubectl logs <pod name>
```

### 8% - Application Lifecycle Management

#### Understand Deployments and how to perform rolling updates and rollbacks.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: busybox-deployment
  labels:
    app: busybox
spec:
  replicas: 2
  selector:
    matchLabels:
      app: busybox
  template:
    metadata:
      labels:
        app: busybox
    spec:
      containers:
      - name: busybox
        image: busybox:1.25
```

##### updates
```yaml
kubectl set image deployment/nginx-deployment nginx=nginx:1.91
...
kubectl edit deployment/nginx-deployment
```

##### rollbacks
if you want to review the rollout history, remember add `--record=true` when create the deployment
```shell
[root@k8s:/home/ubuntu/yaml]$  kubectl create -f deployment.yaml --record
...
[root@k8s:/home/ubuntu/yaml]$ kubectl rollout history deployment/busybox-deployment
deployments "busybox-deployment"
REVISION  CHANGE-CAUSE
1         kubectl create --filename=deployment.yaml --record=true
2         kubectl edit deployment/busybox-deployment
```

if you do not add `--record`, you can get CHANGE-CAUSE which always is <none>.

```shell
kubectl rollout undo deployment/busybox-deployment --to-revision=1
```

#### Know various ways to configure applications.Know how to scale applications.
1. edit the deployment directly
```shell
replicas: 2 # change it to other value
```

2. execute kubectl
```shell
kubectl scale deployment busybox-deployment --replicas=4
```

3. kubectl patch
```shell
kubectl patch deployment busybox-deployment --patch '{"spec": {"replicas": 5}}'
```

#### Understand the primitives necessary to create a self-healing application.

livenessProbe
Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. 


readinessProbe
Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. 


restartPolicy
set it to be `Always`? I am not sure...

### 11% - Cluster Maintenance

#### Understand Kubernetes cluster upgrade process.

#### Facilitate operating system upgrades.

#### Implement backup and restore methodologies.

### 12% - Security

#### Know how to configure authentication
and authorization.

#### Understand Kubernetes security primitives.
Know to configure network policies.

#### Create and manage TLS certificates for cluster components.

#### Work with images securely.

#### Define security contexts.

#### Secure persistent key value store.

#### Work with role-based access control.

### 7% - Storage

#### Understand persistent volumes and know how to create them.

#### Understand access modes for volumes.

#### Understand persistent volume claims primitive.

#### Understand Kubernetes storage objects.

#### Know how to configure applications with persistent storage.

### 10% - Troubleshooting

#### Troubleshoot application failure.
kbuectl logs <pod name>

kubectl get event
scheduling / mount / oom / image pulled failed 


#### Troubleshoot control plane failure.

component(sheduler, controller-manager, apiserver) disabled.

#### Troubleshoot worker node failure.

node notready, resource(memory, cpu, disk pressure)

#### Troubleshoot networking.

flanneld/calico disabled.

IP pool is empty



### 19% - Core Concepts


#### Understand the Kubernetes API primitives.

curl / kubectl command.


#### Understand the Kubernetes cluster architecture.


#### Understand Services and other network primitives.

service / endpoint / ingress

### 11% - Networking

##### Flannel
[flannel](https://github.com/coreos/flannel#flannel)

```shell
[Unit]
Description=Flanneld overlay address etcd agent
After=network.target
After=network-online.target
Wants=network-online.target
After=etcd.service
Before=docker.service

[Service]
Type=notify
ExecStart=/usr/local/bin/flanneld \
  -etcd-cafile=/etc/kubernetes/ssl/ca.pem \
  -etcd-certfile=/etc/kubernetes/ssl/flanneld.pem \
  -etcd-keyfile=/etc/kubernetes/ssl/flanneld-key.pem \
  -etcd-endpoints=${ETCD_ENDPOINTS} \
  -etcd-prefix=${FLANNEL_ETCD_PREFIX} \
  --iface=${NODE_IP} # 虚拟机需要这里填写节点的IP地址
ExecStartPost=/usr/local/bin/mk-docker-opts.sh -k DOCKER_NETWORK_OPTIONS -d /run/flannel/docker
Restart=on-failure

[Install]
WantedBy=multi-user.target
RequiredBy=docker.service
```

##### Calico
[Calico](https://docs.projectcalico.org/v3.1/introduction/)

install: https://docs.projectcalico.org/v3.1/getting-started/kubernetes/

#### Understand the networking configuration on the cluster nodes.



#### Understand Pod networking concepts.



#### Understand service networking.


#### Deploy and configure network load balancer.



#### Know how to use Ingress rules.

#### Know how to configure and use the cluster DNS.

#### Understand CNI.

### 12% - Installation, Configuration & Validation

#### Design a Kubernetes cluster.

#### Install Kubernetes masters and nodes, including the use of TLS bootstrapping.

#### Configure secure cluster communications.



#### Configure a Highly-Available Kubernetes cluster.

3 masters + 3 etcds + N work node.

see https://kubernetes.io/docs/tasks/administer-cluster/highly-available-master/


#### Know where to get the Kubernetes release binaries.

https://github.com/kubernetes/kubernetes/releases

#### Provision underlying infrastructure to deploy a Kubernetes cluster.

openstack VM

VM

bare metal with Linux/Windows host OS

#### Choose a network solution.

flannel: networking

calico: network policy

#### Choose your Kubernetes infrastructure configuration.



#### Run end-to-end tests on your cluster.

go run 


#### Analyse end-to-end tests results.



#### Run Node end-to-end tests.


