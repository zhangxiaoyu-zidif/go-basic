# pod priority and preemption

why we need it?
-- Priority indicates the importance of a Pod relative to other Pods. When a Pod cannot be scheduled, the scheduler tries to preempt (evict) lower priority Pods to make scheduling of the pending Pod possible.

## pod priority 

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    env: test
spec:
  containers:
  - name: nginx
    image: nginx
    imagePullPolicy: IfNotPresent
  priorityClassName: high-priority
```

```yaml
apiVersion: scheduling.k8s.io/v1alpha1
kind: PriorityClass
metadata:
  name: high-priority
value: 1000000
globalDefault: false
description: "This priority class should be used for XYZ service pods only."
```

## preemption

1. enable the feature gates:
add '--feature-gates=PodPriority=true' in apiserver, scheduler, and kubelet startup parameters.
add  enable scheduling.k8s.io/v1alpha1 API and Priority admission controller in API server`--runtime-config=scheduling.k8s.io/v1alpha1=true --enable-admission-plugins=Controller-Foo,Controller-Bar,...,Priority`
enable scheduling.k8s.io/v1alpha1 API and Priority admission controller in API server

1. create PriorityClasses

1. add PriorityClasses to Pod template of deployment as Pod spec.
