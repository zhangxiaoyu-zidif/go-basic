#### ELk 特点

1. 运维成本高，每增加一个日志收集，需要手动修改配置

1. 监控缺失，无法准备获取logstash的状态

1. 无法定制化开发以及维护

App -> log Agent -> Kafka -> ElasticSearch

                       |---> Hadoop
                       
                       |---> Storm
                          
Kafka 用于解耦日志管理，如果有新的需求后，仅需要在kafka增加一个topic即可。

logAent 可以读取etcd的配置，进而获取日志。日志收集客户端，用于收集服务器上的日志。

Kafka，高吞吐量的分布式队列，linkin开发，apache的开源项目。主要和应用解耦，将内容放入消息队列，而不关心谁去消费它。
1.  异步处理，把非关键流程异步化，提高响应时间和健壮性

1. 应用解耦，通过消息队列

1. 流量削峰

Zookeeper

1. 服务注册和服务发现（服务增加，扩缩容）

1. 配置中心（作为应用的）

1. 分布式锁

a. 安装JDK

b. 安装zookeeper from http://apache.fayea.com/zookeeper/

mv conf/zoo_sample.cfg conf/zoo.cfg
编辑conf/zoo.cfg，修改dataDir=D:\zookeeper-xxx\data\
运行bin/zkServer.cmd

c. 安装kafk from http://kafka.apache.org/downloads
打开config目录下的server.properties，修改log.dirs为D:\kafka_logs,修改advertised.host.name=服务器p
启动kafka ./bin/windows/kafka-server-start.bat ./configserver.preperties

kafka向Zookeeper进行注册。

log agent: Kafka, config, tailf, log.

ES， elasticSearch，开源搜索引擎，提供基于http restful的web接口。

Hadoop，分布式计算框架，能够对大量数据进行分布式处理的平台。

Storm 实时计算框架。

