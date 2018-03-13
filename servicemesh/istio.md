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


![beforeAndAfterInjection.png](pic/beforeAndAfterInjection.png)
```shell

```

