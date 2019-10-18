#### 什么是Mesos 

一个大规模集群管理平台，可以资源共享，支持多种framework共享一个集群，比如把 Hadoop、Spark放到一个共同管理的环境。

#### Mesos想解决的问题

在同一个大规模集群系统中，要想运行多个framework，该如何设计调度算法来适应不同framework的需求和资源的分布。

#### 特性

- 两级调度器，让开发者能够创建自定义的frameworks’ schedulers 来运行task。
- 超过10,000 node的扩展性
- 对task的资源隔离
- CPU和memory高利用率的资源调度算法
- ZooKeeper保证了master的高可用
- 可用来监控集群状态的Web UI

#### 架构

简单描述一下Mesos系统的话，它有多个master和多个slave node,但只有一个master是leader。

Mesos master主要功能是把可用的slave资源告知Frameworks’ schedulers。

slave向Master汇报自己的空闲资源和任务的状态。

然后Frameworks’ schedulers 选择要用的资源运行需要的task，因此framework必须实现mesos提供的API,从而对master传送给过来的resource offer作出反应。

 根据论文中的运行WordCount MapReduce job on Hadoop 的测试，集群的对slots的利用率可以达到100%。

###### 设计思想

Mesos的设计理念是仅设定最低限度的接口，来达到资源的合理分配，而把剩下对resource offer的选择权交给framework。

这种设计思想的好处是

1. framework具有自由度
2. Mesos免于繁复功能，达到轻量级与可扩展。

##### 资源分配

###### Dominant Resource Fairness policy

DRF的主要思想是保持各个framework对于资源（CPU,内存，网络带宽）的消耗的平衡。

对于 n 个 frameworks DRF有着 O(logn) 的时间复杂度。

###### Revocation

Mesos 可以杀死有bug的或者是过度占用资源的task。

当一个framework的allocation低于它的guaranteed allocation，Mesos不会杀死它的task。

但如果framework的allocation高于它的guaranteed allocation，那么mesos会先向 executor 询问，如果 executor 没有反应就能够杀死task。

###### Making Resource Offers Scalable and Robust

- 在master上使用ﬁlter，避免向那些总是拒绝特定资源的framework发送offer。
- 激励framework对offer有更短的反应时间。
- 如果一个framework太久没有对offer作出反应，mesos会解除这个offer，然后重新发一个offer。

##### 隔离性

mesos通过Linux containers来使得运行在同一个slave上面的framework executor互相是隔离的。

##### 可扩展性

根据论文中的实验，mesos在50000个slave结点的情况下，也能达到对平均时长为30秒的task小于一秒的overhead。

需要注意的是，当结点数量过多时， mesos希望能够分摊allocation cost, 于是mesos的设计是在batch intervals的时候进行allocation。

##### framework应当如何操作

- 避免Short tasks
- 避免等待更小的资源
- 不要接受不属于自身的资源

##### 容错性

master是整个架构的核心，那么要使得系统的容错性较高，mesos的设计是把master设计成是有状态的，但是这个状态属于soft state，master仅仅是把状态信息推送给slave和scheduler。

mesos使用了冗余来实现高可用。只有一个master是leader,其余的是hot standby，这些hot standby在leader故障时接管leader的工作。一旦原有的master出现故障，就用ZooKeeper来在master里面选举leader。一旦新的leader选举出， slave和Schedulers会将信息发给新的leader,未被处理的消息也会再次传给新的leader。

根据论文中的测试，slave连接到新的leader的平均时间是7-8秒，对不同大小的task都有95%的置信度。 

##### metrics

如何给大规模集群系统顶下性能的指标，mesos使用了以下的标准。

- 一个新的framework达到它的allocation的时间
- 完成一个job的时间
- 集群的利用率

###### 名词解释

**Framework**: al so known as a Mesos application, which registers with the master to receive resource *offers*, and one or more *executors*, which launches *tasks* on slaves. Examples of Mesos frameworks include Marathon, Chronos, and Hadoop

**Task**: a unit of work that is scheduled by a framework, and is executed on a slave node. A task can be anything from a bash command or script, to an SQL query, to a Hadoop job

**soft state** is state which is useful for efficiency, but not essential, as it can be regenerated or replaced if needed.  

**Hot standby** is a redundant method in which one system runs simultaneously with an identical primary system. Upon failure of the primary system, the hot standby system immediately takes over, replacing the primary system. However, data is still mirrored in real time. Thus, both systems have identical data.

**slots** are identical slices of machines