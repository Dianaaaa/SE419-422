# Google Borg Investigation

## 1. What is Borg

> Google's Borg system is a cluster manager that runs hundreds of thousands of jobs,  from many thousands of different applications, across a number of clusters each with up to tens of thousands of machines. 

Borg系统是一个集群管理者，通过rule和policy管理多个集群，使成千上万的应用在上面正常运行。

用户通过购买quote (CPU, RAM, disk等资源)，可以在Borg集群中运行自己的applications。Borg通过特有的设计保证集群的高可用性、高可靠性和安全性。

## 2. The Structure of Borg

Borg的结构如下图：

![](./images/1.png)

一个Borg cell是Borg系统的集群单位，里面包含BorgMaster和worker (论文中并未取名) 两种机器。另外还有Google的Chubby server作为保持分布式一致性的辅助。

- BorgMaster

  BorgMaster只有1台，另有4台replicas，每台机器都存放有Borg cell的状态（内存和persistent store中）。Borg cell的状态又称checkpoint，可用于回退、恢复。

  BorgMaster往往从5台机器中通过选举产生，使用Chubby server的lock机制来保持选举结果的一致性。当master宕机后，在剩下的replicas中会自动地产生新的master，接替原来master的工作。

  Master由两部分构成：

  - BorgMaster主进程：负责处理client的RPC (Remote Procedure Call)；管理系统中对象的状态机；与Borglets交互； 提供web UI作为Sigma (让用户查看jobs状态和cell状态的service) 的后备。
  - 调度进程：负责对任务进行分配。

- Borglet

  每台worker上都运行了一个名为Borglet的进程，它负责运行分配到的tasks，管理本地资源，定期向Master和其他monitor报告本机状态 (Master没隔几秒就会对Borglet进行轮询)。

以下是一些重要概念：

- job: job是用户提交到集群上的需要运行的程序，一个job包含多个tasks。

- task: task是调度的最小单位，多个tasks构成一个job。

- priority: 每个job都有一个priority，用以标记此job的优先级别。

## 3. Policy

### BNS (Borg Name Service)

为了让每台机器都都能找到task，Borg会创建一个BNS用以task的定位。如名为cc的cell中ubar用户的名为jfoo的job的第55个task，可以通过Borg name “50.jfoo.ubar.cc.borg.gooble.com"定位。

### Scheduling

BorgMaster的sheduler进程负责对job进行调度。

当一个job被提交后，Borgmaster会记录进Paxos store，并将其加入pending queue。

之后sheduler对pending queue 按照优先级别进行扫描，当有足够的资源时，将task分配给该机器。采用round-robin算法来保证公平公正。

此调度算法分为两个部分：

- feasibility checking： 找到足够运行task的所有机器。
- scoring：对上述机器根据制定好的评分标准进行打分，择优将task分配过去。

当无法找到适合当前task的机器时，scheduler将跳过这个task，转向队列中的下一个task，重复以上操作。

由于scoring过程开销过大，将scoring结果cache起来，只有当机器和task状态发生变化时才重新计算score。

为了减少task的start up时间，优先考虑已具备必要package的机器。

当Cell机器数量过大时，无论是feasibility checking还是scoring都是很大的开销，于是scheduler尝试寻找局部最优解（Relaxed randomization）减少这两步骤的开销。

### Communications

Master每隔几秒地对workers进行轮询，获取workers的状态。若有机器无法回应，超过一定次数后将被标记为down，其上面在跑的所有tasks将被重新调度到其他可用机器上。

Master负责准备发送给Borglets的信息，并根据response修改cell的状态。而每个replica通过link shard与Borglets保持联系，分摊了Master的工作。

## 4. Availability & Utilization

Borg采取了很多措施来增强可用性。

- 自动重新调度失败tasks。
- 减少错误的依赖传递
- 在进行维护活动时，限制允许的task中断率，和同时宕机的task数。
- 对各operations采用幂等的设计。

研究发现，CPI与单机上的CPU使用率和单机上的task数成正比。并且在shared cells中提升更为显著。

保证Large Cells，小的cells实现相同规模的cluster需要更多的机器。

Borg用户资源时，CPU以核为单位，memory和disk以byte为单位。

当task用尽其固定资源时，将会被中止。

## 5. Isolation

Borg使用chroot jail作为主要的安全隔离机制。保证同一个机器上task间的安全运行。在保证安全的同时，Borg也需要保证performance的隔离，即一个task的运行不会影响到同机其他task的performance。

