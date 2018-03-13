# istio
***[overview of istio](https://istio.io/docs/concepts/what-is-istio/overview.html)***

istio serves a network that's  load balancing, service-to-service authentication, monitoring, and more.
we need not change any code of microservices. =)

## preparation
1. install kubernetes, >= 1.9 will be better.

in fact, i recommand to use minikube rather than ./hack/local-up-cluster.sh to start a local cluster for testing.

that's becasue you should modify many parameter such as kubectl command .etc. Some accidents will cause unncessary troubles.

start a minikube ref [here](https://github.com/kubernetes/minikube/blob/v0.25.0/README.md#quickstart)

if your host env is linux, that'll be great. run
```shell

minikube start \
	--extra-config=controller-manager.ClusterSigningCertFile="/var/lib/localkube/certs/ca.crt" \
	--extra-config=controller-manager.ClusterSigningKeyFile="/var/lib/localkube/certs/ca.key" \
	--extra-config=apiserver.Admission.PluginNames=NamespaceLifecycle,LimitRanger,ServiceAccount,PersistentVolumeLabel,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota \
	--kubernetes-version=v1.9.0 \
	--vm-driver=none
```

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
Note: If your cluster is running in an environment that does not support an external load balancer (e.g., minikube), the EXTERNAL-IP of istio-ingress says <pending>. You must access the application using the service NodePort, or use port-forwarding instead. [ref here](https://kubernetes.io/docs/concepts/services-networking/service/)
  
 
### check istioctl command
```shell
root@ubuntu:/home/cloud/istio-0.6.0/install/kubernetes# istioctl --help

Istio configuration command line utility.

Create, list, modify, and delete configuration resources in the Istio
system.

Available routing and traffic management configuration types:

        [routerule ingressrule egressrule destinationpolicy]

See https://istio.io/docs/reference/ for an overview of routing rules
and destination policies.

Usage:
  istioctl [command]

Available Commands:
  context-create Create a kubeconfig file suitable for use with istioctl in a non kubernetes environment
  create         Create policies and rules
  delete         Delete policies or rules
  deregister     De-registers a service instance
  gen-deploy     Generates the configuration for Istio's control plane.
  get            Retrieve policies and rules
  help           Help about any command
  kube-inject    Inject Envoy sidecar into Kubernetes pod resources
  proxy-config   Retrieves proxy configuration for the specified pod [kube only]
  register       Registers a service instance (e.g. VM) joining the mesh
  replace        Replace existing policies and rules
  version        Prints out build version information

Flags:
  -h, --help                             help for istioctl
  -i, --istioNamespace string            Istio system namespace (default "istio-system")
  -c, --kubeconfig string                Kubernetes configuration file (default "$KUBECONFIG else $HOME/.kube/config")
      --log_as_json                      Whether to format output as JSON or in plain console-friendly format
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_callers                      Include caller information, useful for debugging
      --log_output_level string          The minimum logging level of messages to output, can be one of "debug", "info", "warn", "error", or "none" (default "info")
      --log_rotate string                The path for the optional rotating log file
      --log_rotate_max_age int           The maximum age in days of a log file beyond which the file is rotated (0 indicates no limit) (default 30)
      --log_rotate_max_backups int       The maximum number of log file backups to keep before older files are deleted (0 indicates no limit) (default 1000)
      --log_rotate_max_size int          The maximum size in megabytes of a log file beyond which the file is rotated (default 104857600)
      --log_stacktrace_level string      The minimum logging level at which stack traces are captured, can be one of "debug", "info", "warn", "error", or "none" (default "none")
      --log_target stringArray           The set of paths where to output the log. This can be any path as well as the special values stdout and stderr (default [stdout])
  -n, --namespace string                 Config namespace
  -p, --platform string                  Istio host platform (default "kube")
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging

Use "istioctl [command] --help" for more information about a command.

```

### how to inject sidercar named `envoy` into your application

##### Installing the Webhook first
```shell
./install/kubernetes/webhook-create-signed-cert.sh \
    --service istio-sidecar-injector \
    --namespace istio-system \
    --secret sidecar-injector-certs
    
kubectl apply -f install/kubernetes/istio-sidecar-injector-configmap-release.yaml


cat install/kubernetes/istio-sidecar-injector.yaml | \
     ./install/kubernetes/webhook-patch-ca-bundle.sh > \
     install/kubernetes/istio-sidecar-injector-with-ca-bundle.yaml
   
kubectl apply -f install/kubernetes/istio-sidecar-injector-with-ca-bundle.yaml

# make sure this pod is running.
kubectl -n istio-system get deployment -listio=sidecar-injector
NAME                     DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
istio-sidecar-injector   1         1         1            1           1d
```

1. inject it into the namespaces
e.g.:
```shell
kubectl create namespace istio-test
namespace "istio-test" created

kubectl label namespace istio-test istio-injection=enabled
namespace "istio-test" labeled

kubectl get ns --show-labels
NAME           STATUS    AGE       LABELS
default        Active    49m       <none>
istio-system   Active    42m       <none>
istio-test     Active    11m       istio-injection=enabled
kube-public    Active    49m       <none>
kube-system    Active    49m       <none>

```

2. inject it into application directly
```shell
kubectl create -f <(istioctl kube-inject -f <your-app-spec>.yaml)
```
Attention for this way[ref here](https://istio.io/docs/reference/commands/istioctl.html#istioctl kube-inject):
kube-inject manually injects envoy sidecar into kubernetes workloads. Unsupported resources are left unmodified so it is safe to run kube-inject over a single file that contains multiple Service, ConfigMap, Deployment, etc. definitions for a complex application. Its best to do this when the resource is initially created.


what happened when injecct a sidecar to our own App？
![beforeAndAfterInjection.png](pic/beforeAndAfterInjection.png)

```shell
# execute it in two namespaces, one is injected with sidecar, the other one is not.
kubectl apply -f samples/sleep/sleep.yaml 


# without injection
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: 2018-03-13T02:46:57Z
  generateName: sleep-776b7bcdcd-
  labels:
    app: sleep
    pod-template-hash: "3326367878"
  name: sleep-776b7bcdcd-bfh7c
  namespace: test
  ownerReferences:
  - apiVersion: extensions/v1beta1
    blockOwnerDeletion: true
    controller: true
    kind: ReplicaSet
    name: sleep-776b7bcdcd
    uid: cc215325-2668-11e8-b230-5254005c78d5
  resourceVersion: "4151"
  selfLink: /api/v1/namespaces/test/pods/sleep-776b7bcdcd-bfh7c
  uid: cc223632-2668-11e8-b230-5254005c78d5
spec:
  containers:
  - command:
    - /bin/sleep
    - infinity
    image: tutum/curl
    imagePullPolicy: IfNotPresent
    name: sleep
    resources: {}
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /var/run/secrets/kubernetes.io/serviceaccount
      name: default-token-2plv6
      readOnly: true
  dnsPolicy: ClusterFirst
  nodeName: ubuntu
  restartPolicy: Always
  schedulerName: default-scheduler
  securityContext: {}
  serviceAccount: default
  serviceAccountName: default
  terminationGracePeriodSeconds: 30
  tolerations:
  - effect: NoExecute
    key: node.kubernetes.io/not-ready
    operator: Exists
    tolerationSeconds: 300
  - effect: NoExecute
    key: node.kubernetes.io/unreachable
    operator: Exists
    tolerationSeconds: 300
  volumes:
  - name: default-token-2plv6
    secret:
      defaultMode: 420
      secretName: default-token-2plv6
status:
  conditions:
  - lastProbeTime: null
    lastTransitionTime: 2018-03-13T02:46:57Z
    status: "True"
    type: Initialized
  - lastProbeTime: null
    lastTransitionTime: 2018-03-13T02:46:59Z
    status: "True"
    type: Ready
  - lastProbeTime: null
    lastTransitionTime: 2018-03-13T02:46:57Z
    status: "True"
    type: PodScheduled
  containerStatuses:
  - containerID: docker://29c73a3c2d5d136d4e32bd5d1961543eeeaeecad726aa811866908a4b3d68a6e
    image: tutum/curl:latest
    imageID: docker-pullable://tutum/curl@sha256:b6f16e88387acd4e6326176b212b3dae63f5b2134e69560d0b0673cfb0fb976f
    lastState: {}
    name: sleep
    ready: true
    restartCount: 0
    state:
      running:
        startedAt: 2018-03-13T02:46:58Z
  hostIP: 192.168.122.40
  phase: Running
  podIP: 172.17.0.10
  qosClass: BestEffort
  startTime: 2018-03-13T02:46:57Z
  
  
# with injection
apiVersion: v1
kind: Pod
metadata:
  annotations:
    sidecar.istio.io/status: '{"version":"1771e5b7d0647c27f709e5194dc91147762ea6dccc6581f791d2de544250cdc1","initContainers":["istio-init"],"containers":["istio-proxy"],"volumes":["istio-envoy","istio-certs"]}'
  creationTimestamp: 2018-03-13T02:24:27Z
  generateName: sleep-776b7bcdcd-
  labels:
    app: sleep
    pod-template-hash: "3326367878"
  name: sleep-776b7bcdcd-cvjx4
  namespace: default
  ownerReferences:
  - apiVersion: extensions/v1beta1
    blockOwnerDeletion: true
    controller: true
    kind: ReplicaSet
    name: sleep-776b7bcdcd
    uid: a71c4e9a-2665-11e8-b230-5254005c78d5
  resourceVersion: "3042"
  selfLink: /api/v1/namespaces/default/pods/sleep-776b7bcdcd-cvjx4
  uid: a724d201-2665-11e8-b230-5254005c78d5
spec:
  containers:
  - command:
    - /bin/sleep
    - infinity
    image: tutum/curl
    imagePullPolicy: IfNotPresent
    name: sleep
    resources: {}
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /var/run/secrets/kubernetes.io/serviceaccount
      name: default-token-cznsq
      readOnly: true
  - args:
    - proxy
    - sidecar
    - --configPath
    - /etc/istio/proxy
    - --binaryPath
    - /usr/local/bin/envoy
    - --serviceCluster
    - sleep
    - --drainDuration
    - 45s
    - --parentShutdownDuration
    - 1m0s
    - --discoveryAddress
    - istio-pilot.istio-system:15003
    - --discoveryRefreshDelay
    - 1s
    - --zipkinAddress
    - zipkin.istio-system:9411
    - --connectTimeout
    - 10s
    - --statsdUdpAddress
    - istio-mixer.istio-system:9125
    - --proxyAdminPort
    - "15000"
    - --controlPlaneAuthPolicy
    - NONE
    env:
    - name: POD_NAME
      valueFrom:
        fieldRef:
          apiVersion: v1
          fieldPath: metadata.name
    - name: POD_NAMESPACE
      valueFrom:
        fieldRef:
          apiVersion: v1
          fieldPath: metadata.namespace
    - name: INSTANCE_IP
      valueFrom:
        fieldRef:
          apiVersion: v1
          fieldPath: status.podIP
    image: docker.io/istio/proxy:0.6.0
    imagePullPolicy: IfNotPresent
    name: istio-proxy
    resources: {}
    securityContext:
      privileged: false
      readOnlyRootFilesystem: true
      runAsUser: 1337
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /etc/istio/proxy
      name: istio-envoy
    - mountPath: /etc/certs/
      name: istio-certs
      readOnly: true
  dnsPolicy: ClusterFirst
  initContainers:
  - args:
    - -p
    - "15001"
    - -u
    - "1337"
    image: docker.io/istio/proxy_init:0.6.0
    imagePullPolicy: IfNotPresent
    name: istio-init
    resources: {}
    securityContext:
      capabilities:
        add:
        - NET_ADMIN
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
  nodeName: ubuntu
  restartPolicy: Always
  schedulerName: default-scheduler
  securityContext: {}
  serviceAccount: default
  serviceAccountName: default
  terminationGracePeriodSeconds: 30
  tolerations:
  - effect: NoExecute
    key: node.kubernetes.io/not-ready
    operator: Exists
    tolerationSeconds: 300
  - effect: NoExecute
    key: node.kubernetes.io/unreachable
    operator: Exists
    tolerationSeconds: 300
  volumes:
  - name: default-token-cznsq
    secret:
      defaultMode: 420
      secretName: default-token-cznsq
  - emptyDir:
      medium: Memory
    name: istio-envoy
  - name: istio-certs
    secret:
      defaultMode: 420
      optional: true
      secretName: istio.default
status:
  conditions:
  - lastProbeTime: null
    lastTransitionTime: 2018-03-13T02:24:43Z
    status: "True"
    type: Initialized
  - lastProbeTime: null
    lastTransitionTime: 2018-03-13T02:25:09Z
    status: "True"
    type: Ready
  - lastProbeTime: null
    lastTransitionTime: 2018-03-13T02:24:27Z
    status: "True"
    type: PodScheduled
  containerStatuses:
  - containerID: docker://b0e28d870f93d71cbdbdb75c16a5e95403aef0003972805c237adb2e8369346d
    image: istio/proxy:0.6.0
    imageID: docker-pullable://istio/proxy@sha256:51ec13f9708238351a8bee3c69cf0cf96483eeb03a9909dea12306bbeb1d1a9d
    lastState: {}
    name: istio-proxy
    ready: true
    restartCount: 0
    state:
      running:
        startedAt: 2018-03-13T02:25:09Z
  - containerID: docker://6cbbe86812f15056a8825cf9ac1fcc3ddfb004b9f279ff08a5d32f26ea19b170
    image: tutum/curl:latest
    imageID: docker-pullable://tutum/curl@sha256:b6f16e88387acd4e6326176b212b3dae63f5b2134e69560d0b0673cfb0fb976f
    lastState: {}
    name: sleep
    ready: true
    restartCount: 0
    state:
      running:
        startedAt: 2018-03-13T02:25:08Z
  hostIP: 192.168.122.40
  initContainerStatuses:
  - containerID: docker://0aeb597589618d75e0dccdcc85771a1871aa119d3cea394f4ce93d4981d18911
    image: istio/proxy_init:0.6.0
    imageID: docker-pullable://istio/proxy_init@sha256:bd1cb7b79e3e3398729d49b2307dbc3335d3182540c740e468340c9490b2880b
    lastState: {}
    name: istio-init
    ready: true
    restartCount: 0
    state:
      terminated:
        containerID: docker://0aeb597589618d75e0dccdcc85771a1871aa119d3cea394f4ce93d4981d18911
        exitCode: 0
        finishedAt: 2018-03-13T02:24:43Z
        reason: Completed
        startedAt: 2018-03-13T02:24:43Z
  phase: Running
  podIP: 172.17.0.9
  qosClass: BestEffort
  startTime: 2018-03-13T02:24:27Z
```

See from the examples above we know that the deplpoyment has been changed or modified by
`install/kubernetes/istio-sidecar-injector-with-ca-bundle.yaml`
which is
```shell
# GENERATED FILE. Use with Kubernetes 1.9+
# TO UPDATE, modify files in install/kubernetes/templates and run install/updateVersion.sh
apiVersion: v1
kind: Service
metadata:
  name: istio-sidecar-injector
  namespace: istio-system
  labels:
    istio: sidecar-injector
spec:
  ports:
  - name: https-webhook # optional
    port: 443
  selector:
    istio: sidecar-injector
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: istio-sidecar-injector-service-account
  namespace: istio-system
---
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: istio-sidecar-injector
  namespace: istio-system
  labels:
    istio: sidecar-injector
spec:
  replicas: 1
  template:
    metadata:
      name: istio-sidecar-injector
      labels:
        istio: sidecar-injector
    spec:
      serviceAccountName: istio-sidecar-injector-service-account
      containers:
        - name: webhook
          image: docker.io/istio/sidecar_injector:0.6.0
          imagePullPolicy: IfNotPresent
          args:
            - --tlsCertFile=/etc/istio/certs/cert.pem
            - --tlsKeyFile=/etc/istio/certs/key.pem
            - --injectConfig=/etc/istio/inject/config
            - --meshConfig=/etc/istio/config/mesh
            - --healthCheckInterval=2s
            - --healthCheckFile=/health
          volumeMounts:
          - name: config-volume
            mountPath: /etc/istio/config
            readOnly: true
          - name: certs
            mountPath: /etc/istio/certs
            readOnly: true
          - name: inject-config
            mountPath: /etc/istio/inject
            readOnly: true
          livenessProbe:
            exec:
              command:
                - /usr/local/bin/sidecar-injector
                - probe
                - --probe-path=/health
                - --interval=2s
            initialDelaySeconds: 4
            periodSeconds: 4
          readinessProbe:
            exec:
              command:
                - /usr/local/bin/sidecar-injector
                - probe
                - --probe-path=/health
                - --interval=2s
            initialDelaySeconds: 4
            periodSeconds: 4
      volumes:
      - name: config-volume
        configMap:
          name: istio
      - name: certs
        secret:
          secretName: sidecar-injector-certs
      - name: inject-config
        configMap:
          name: istio-inject
          items:
          - key: config
            path: config
---
apiVersion: admissionregistration.k8s.io/v1beta1
kind: MutatingWebhookConfiguration
metadata:
  name: istio-sidecar-injector
webhooks:
  - name: sidecar-injector.istio.io
    clientConfig:
      service:
        name: istio-sidecar-injector
        namespace: istio-system
        path: "/inject"
      caBundle: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM1ekNDQWMrZ0F3SUJBZ0lCQVRBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwdGFXNXAKYTNWaVpVTkJNQjRYRFRFNE1ETXhNekF4TkRRMU5Wb1hEVEk0TURNeE1EQXhORFExTlZvd0ZURVRNQkVHQTFVRQpBeE1LYldsdWFXdDFZbVZEUVRDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBSlR0Cm9XMkEzelpHRE9FQ1hQMnlBMWs5Z21KWnVVQU91V1BQREkwa1A0U1hXMjBXZGhDaWlqVmZkMEc3T1lpclI0dmMKRURxOGIvek02TjZqeCtwblhpek1QTnlkak5zV215eUU5QkFRWVkxa1FTcDB1Mm9HeVpON1Q1N0RMMUlRVVVVKwpzdEYyYmNOZkRyV2NaSWJUcFF5b1orYjlDMXBQYXB6RWQ1QzNWbVF0bnk0Z0M2QTJHejY1eEdXZjVTQnJ5a0QvCm5EWHBSendpVEFTcGpVUko1L2RHZWprazlaWEFrSk51UmdPT1F0T3FOVHFmbktHblVEbTVTOURlQVBoQ1lIWFAKQkFHb0txWWtETDI3U0gwMStrVENjcTlmb3BKK0wyYjhhSUtsYlk1dDRNc0tteFhtTjVmb0Z5YjlRd2JYWm9ETgpJSjMrQStFTVVzVWl3NVpGTWQ4Q0F3RUFBYU5DTUVBd0RnWURWUjBQQVFIL0JBUURBZ0trTUIwR0ExVWRKUVFXCk1CUUdDQ3NHQVFVRkJ3TUNCZ2dyQmdFRkJRY0RBVEFQQmdOVkhSTUJBZjhFQlRBREFRSC9NQTBHQ1NxR1NJYjMKRFFFQkN3VUFBNElCQVFBMmN4QjR6RXdSeHNUbXpIRUlUYkJQaFhOTVlIa2R6UStXU21ydXVqNW41aE9pcW1DTwpXTTd5RE9FVGNjSmx2S21Bdm15VlpWaHoxdWtUM2s5dEMxWUlXc2JHeTFwWTFMMStyQ3lLYkVpOStLUm1tcFV0CnVJY0JhQ2VVM05sR0VWdGlZR3lldXpmcUFpRDdvck84YzFrRzRNSURqUHVzT1NNaTlpYkd1aDBmTzh3bUVxb1IKTmtMQVhhYUFMQjZvZC9ybDA1VlNnS2JCYzh5RUNDZUtCRDJwYWkyTDdTYjBaeUxnSEZGaUhnRXF1UTIyM2l6VAppS2RIS1hRZDNTNGpEMjBBcTFHQTA2Mm82dXZRYUUvUStpdW95clVZSk9lRU1LK0ZYZXRCd0xRem5oaXI3RFpSCkVtRkpkMDhBZVNUcnNpQWNKd0xkTVZ3MlM3SzRHSG92SW55TwotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
    rules:
      - operations: [ "CREATE" ]
        apiGroups: [""]
        apiVersions: ["v1"]
        resources: ["pods"]
    namespaceSelector:
      matchLabels:
        istio-injection: enabled
---
```

the original APP script is 
```yaml
apiVersion: v1
kind: Service
metadata:
  name: sleep
  labels:
    app: sleep
spec:
  ports:
  - port: 80
    name: http
  selector:
    app: sleep
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: sleep
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: sleep
    spec:
      containers:
      - name: sleep
        image: tutum/curl
        command: ["/bin/sleep","infinity"]
        imagePullPolicy: IfNotPresent
```

The sleep pod which uses istio has one more containers. That is envoy-proxy.
