#### ELk 特点

1. 运维成本高，每增加一个日志收集，需要手动修改配置

1. 监控缺失，无法准备获取logstash的状态

1. 无法定制化开发以及维护

App -> log Agent -> Kafka -> ElasticResearch
                       |---> Hadoop
                       |---> Storm
                          
Kafka 用于解耦日志管理，如果有新的需求后，仅需要在kafka增加一个topic即可。
