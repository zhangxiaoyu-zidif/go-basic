#

add `--feature-gates=PersistentLocalVolumes=true,LocalStorageCapacityIsolation=true` to apiserver and kubelet

```yaml
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: busybox-apps-v1beta1
  labels:
    app: busybox-apps-v1beta1
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
      - name: busybox
        image: busybox
        command:
        - sleep
        - "3600"
        imagePullPolicy: Never
        resources:
          limits:
            ephemeral-storage: "50Mi"
```
execute:
```shell
kubectl cp some-big-file busybox:/
```
and you will see:
```shell
busybox-apps-v1beta1-7f8dd8d89-kh6xc   1/1       Running   0          19m
busybox-apps-v1beta1-7f8dd8d89-mg7ls   0/1       Evicted   0          21m
[root@k8s-master-controller:/]$ kubectl describe pod busybox-apps-v1beta1-7f8dd8d89-mg7ls
Name:           busybox-apps-v1beta1-7f8dd8d89-mg7ls
Namespace:      default
Node:           172.160.134.17/
Start Time:     Mon, 23 Apr 2018 09:27:02 +0800
Labels:         app=busybox-apps-v1beta1
                pod-template-hash=394884845
Annotations:    kubernetes.io/created-by={"kind":"SerializedReference","apiVersion":"v1","reference":{"kind":"ReplicaSet","namespace":"default","name":"busybox-apps-v1beta1-7f8dd8d89","uid":"6c817aea-4695-11e8-9103-f...
Status:         Failed
Reason:         Evicted
Message:        The node was low on resource: ephemeral-storage.
IP:
Created By:     ReplicaSet/busybox-apps-v1beta1-7f8dd8d89
Controlled By:  ReplicaSet/busybox-apps-v1beta1-7f8dd8d89
Containers:
  busybox:
    Image:  busybox
    Port:   <none>
    Command:
      sleep
      3600
    Limits:
      ephemeral-storage:  50Mi
    Requests:
      ephemeral-storage:  50Mi
    Environment:          <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-7tchh (ro)
Volumes:
  default-token-7tchh:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-7tchh
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     <none>
Events:
  Type     Reason                 Age   From                     Message
  ----     ------                 ----  ----                     -------
  Normal   Scheduled              22m   default-scheduler        Successfully assigned busybox-apps-v1beta1-7f8dd8d89-mg7ls to 172.160.134.17
  Normal   SuccessfulMountVolume  22m   kubelet, 172.160.134.17  MountVolume.SetUp succeeded for volume "default-token-7tchh"
  Normal   Pulled                 22m   kubelet, 172.160.134.17  Container image "busybox" already present on machine
  Normal   Created                22m   kubelet, 172.160.134.17  Created container
  Normal   Started                22m   kubelet, 172.160.134.17  Started container
  Warning  Evicted                19m   kubelet, 172.160.134.17  pod ephemeral local storage usage exceeds the total limit of containers {{52428800 0} {<nil>} 50Mi BinarySI}
  Normal   Killing                19m   kubelet, 172.160.134.17  Killing container with id docker://busybox:Need to kill Pod
[root@k8s-master-controller:/home]$ kubectl describe rs busybox-apps-v1beta1-7f8dd8d89
Name:           busybox-apps-v1beta1-7f8dd8d89
Namespace:      default
Selector:       app=busybox-apps-v1beta1,pod-template-hash=394884845
Labels:         app=busybox-apps-v1beta1
                pod-template-hash=394884845
Annotations:    deployment.kubernetes.io/desired-replicas=1
                deployment.kubernetes.io/max-replicas=2
                deployment.kubernetes.io/revision=1
Controlled By:  Deployment/busybox-apps-v1beta1
Replicas:       1 current / 1 desired
Pods Status:    1 Running / 0 Waiting / 0 Succeeded / 1 Failed
Pod Template:
  Labels:  app=busybox-apps-v1beta1
           pod-template-hash=394884845
  Containers:
   busybox:
    Image:  busybox
    Port:   <none>
    Command:
      sleep
      3600
    Limits:
      ephemeral-storage:  50Mi
    Environment:          <none>
    Mounts:               <none>
  Volumes:                <none>
Events:
  Type    Reason            Age   From                   Message
  ----    ------            ----  ----                   -------
  Normal  SuccessfulCreate  53m   replicaset-controller  Created pod: busybox-apps-v1beta1-7f8dd8d89-mg7ls
  Normal  SuccessfulCreate  51m   replicaset-controller  Created pod: busybox-apps-v1beta1-7f8dd8d89-kh6xc
```
