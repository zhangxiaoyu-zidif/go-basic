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

#### Know various ways to configure applications.Know how to scale applications.

#### Understand the primitives necessary to create a self-healing application.

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

#### Understand persistent volumes and know how
to create them.

#### Understand access modes for volumes.

#### Understand persistent volume claims primitive.

#### Understand Kubernetes storage objects.

#### Know how to configure applications with persistent storage.

### 10% - Troubleshooting

#### Troubleshoot application failure.

#### Troubleshoot control plane failure.

#### Troubleshoot worker node failure.

#### Troubleshoot networking.

### 19% - Core Concepts

#### Understand the Kubernetes API primitives.

#### Understand the Kubernetes cluster architecture.

#### Understand Services and other network primitives.

### 11% - Networking

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

#### Know where to get the Kubernetes release binaries.

#### Provision underlying infrastructure to deploy a Kubernetes cluster.

#### Choose a network solution.

#### Choose your Kubernetes infrastructure configuration.

#### Run end-to-end tests on your cluster.

#### Analyse end-to-end tests results.

#### Run Node end-to-end tests.