# Special nouns

#### IOPS

IOPS是Input/Output Operations per Second，即每秒能处理的I/O个数，用于表示块存储处理读写（输出/输入）的能力。如果要部署事务密集型应用，需要关注IOPS性能


#### Linux 限制

ulimit -a

#### hash 一致性

最关键的区别就是，对节点和数据，都做一次哈希运算，然后比较节点和数据的哈希值，数据取和节点最相近的节点做为存放节点。这样就保证当节点增加或者减少的时候，影响的数据最少。
