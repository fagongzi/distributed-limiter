# distributed-limiter
Distributed-limiter是一个分布式流控实现，分为服务端和客户端，客户端使用API和服务端交互来实施流控。可以使用Distributed-limiter来对后端资源的访问做流量控制（例如数据库，Restful接口等等）。

## Feature
* 高可用
* Auto-Scale
* Auto-Rebalance
* 高性能
* 支持多种流控策略(例如：单位时间并发控制)

## 概念
### Resource
Distributed-limiter把所有需要被流控的后端统一称为Resource，每一个Resource都有一个ID和单位时间的并发数。

## 服务端 
Distributed-limiter为了提升并发性能，把管理的所有的Resource进行分组，在整个服务器集群中，每个服务端负责一部分Resource，系统保证每个服务端上的负载均衡（如果不均衡会进行调度，大致保证每个服务器上的并发均衡）。可以动态的增加和减少服务节点。可以通过简单的增加机器来获得线性的性能增长。
