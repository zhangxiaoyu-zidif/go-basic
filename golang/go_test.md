# 使用pprof进行性能调优

1. 首先执行
```shell
go test -bench . -cpuprofile cpu.out

ls | cpu.out

go tool pprof cpu.out

web

```

将生成一个图，将大框的内容优化掉，就会降低执行时间。

安装Graphviz(www.graphviz.org/download)

map[] 要比 []int效率高的多。
