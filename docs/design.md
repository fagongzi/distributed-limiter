# 设计
描述Distributed-limiter的设计方案

## 使用场景
Distributed-limiter被设计为在一个集群环境中，针对后端的Resource（例如某一个后端服务）进行流控。可以被用作统一接入层的一个基础组件。例如: 一个统一接入层（比如API网关），这个接入层本身是集群的，有Node1,Node2,Node3三个节点，流量从外部进入统一接入层，通过接入层流入到后端Resource。如果我们需要在统一接入层针对后端的Resource统一做流量控制，那么Distributed-limiter就可以派上用场了。

## 目标
* Auto-Scale
* High-Available
* Auto-Reblance

## Auto-Scale
要做到高性能，就必须利用整个集群的机器，所以传统的Master-Slave模式是不适用的，Master-Salve结构下，只有Master节点才能提供服务，性能不能通过增加Salve机器来提升。对于分布式流控的使用场景看，系统的并发是非常大的，所以单节点的性能不能满足。

针对这个场景，我们设计为Distributed-limiter集群中所有的机器都是Master，都可以提供服务。我们把Distributed-limiter保护的所有后端称为Resource，每一个Resource分配一个ID，这个ID是连续的，针对这个连续的ID分片，一个机器负责一组分片。这样每个机器都可以是这一组分片的Master。

## High-Available
当Distributed-limiter集群中的某一台机器故障后，这台机器上所有Master的分片，会被其他机器上的Slave取代，成为Master，继续对外提供服务。在Master-Slave切换的时候，存在一个服务不可用的时间，针对流控的场景，转化为在客户端根据客户端（统一接入层的节点）的数量，对流量做均分的流控，比如某一个Resource的SLA为1000，现在有10个客户端节点，那么每个客户端就针对这个Resource做SLA为100的流控，等到Master-Salve切换完成后，恢复为的分布式流控。

## Auto-Reblance
当Distributed-limiter集群中有机器退出或者加入，就会造成新加入的机器和老机器之间的负载不均衡，这时候需要调度，让每个机器上的复杂尽可能的均衡。具体设计如下：

* Distributed-limiter集群在启动的时候，自动选择一个节点作为调度节点
* Distributed-limiter集群中的所有节点通过心跳汇报当前节点的数据（Resource数目，每个Resource单位时间的请求数量）到调度节点
* Distributed-limiter的调度节点根据上报的信息发出调度指令(增加，删除，变更Master)，用来达到负载均衡，目标是每台机器上的每秒请求数量大致相同
* 当调度节点故障时，其他任意节点都可以变成新的调度节点
