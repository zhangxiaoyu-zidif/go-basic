# kubernetes common questions


1. 怎样使pod 和 pod之间进行通讯？ 一个请求进来，pod之间有时也需要通讯。
kubernetes是要安装网络插件才能正常工作。

2. 一般研究k8s技术，有多少工程师一起研究？

3. k8s的监控系统 能够报警吗？比如微信公众号，或有这样的接口，让我们自己开发？

4. 可以怎样倒出每个pod内的日志？ kubectl logs 命令如果要看最后200行，或全部的日志用什么命令？

5. 有哪些帮助trouble shooting的常用命令？

6. 部署的脚本用ansible，如您所说，可以参看openshift 的部署脚本，您这儿有没有部署openshift的脚本呢？

7. spring cloud和k8s的比较和优劣点？

8. RBAC role based authorization control

9. Istio 技术跟踪：有什么好的论坛，qq群等？目前的进展

10. 能否介绍一下 calico, cni？

11. 写yaml文件时，发现kind有很多种，pod，service，serviceAccount， deployment，ingress，那么在哪里看到到底有多少种，每种是代表什么意思。pod，service，deployment我懂，service account，entrypoint等不懂。

12. 如果需要一个测试环境，开发环境，发布环境，用 k8s 怎么实现？网上说，放在不同的namespace就可以啊？怎么放？

13. configmap 和 env 的区别，他们都是key value，区别在哪里？

14. kubectl create -f xx.yaml  和  kubectl apply -f xx.yaml 有什么区别？

15. 怎样设计使得有状态的组件获得高可靠性？

