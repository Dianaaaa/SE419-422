# Storage

## 1. Storage Madia

- Floppy disk （软盘）

  容量:  1.2MB-1.44MB

  成本：5,000 RMB per GB

  转速：300-360 rpm

  在当前时代暂无优点。

  缺点：容量小，读写速度慢。

- Tape （磁带）

  容量：容量不等

  成本：0.5 RMB per GB

  密度：200-6250 bpi (bytes per inch)

  速度：1-30 ips (inches per second)

  优点：单位容量价格低廉，保存年限长，低能耗

  缺点：同容量下磁带比磁盘和ssd体积大，读写速度慢；出错率高

- CD/DVD

  容量：1.5-17GB

  成本：0.7 RMB per GBs

  读写速度：1.4-33.2MB/s

  优点：比软盘容量大，可存放各种多媒体内容。

  缺点：相较于现在的便携式存储介质携带不方便；使用需要光驱；容量较小。

- HDD

  容量：500GB-2TB

  成本：0.2-0.4 RMB per GB

  读写速度：80-160 MB/s

  优点：可廉价地存储大量数据

  缺点：读写速度不如SSD，不能满足高性能的需求

- SSD (flash）

  容量：128GB-100TB

  成本：1.9-2.6 RMB per GB

  读写速度：500MB-3000MB/s

  优点：速度快，高性能

  缺点：价格贵，制造技术要求高，需要考虑寿命问题（磨损均衡算法）

- NVM (non-volatile memory)

  优点：数据持久，按字节存取，存储密度高，低能耗，读写性能接近DRAM

  缺点：昂贵，读写速度相差有些大（读远快于写），寿命有限（磨损均衡算法），还未有成熟的产品。

## 2.SSD (solid-state drive)

- NAND & NOR flash

  ![](./images/1.png)

  如上图所示是NAND和NOR的优势雷达图，可以看到它们各有优势。NAND便宜、容量大、写速度快，NOR读速度快、能耗低。总体NAND比NOR更好，但NOR也有NAND取代不了的优势。

  ![operations of NAND & NOR](./images/2.png)

  可以看到，NOR flash的read speed是相当的快，但其写速度与读速度相差甚远，也慢于NAND flash，Erase time甚至是NAND的450倍，因其自身特性，可以Random access地访问数据，传输数据的效率极高，可执行程序可以在芯片内执行（eXecuteInPlace），常用于计算机的BIOS载体。NAND的各项指标比较均衡，数据密度是NOR flash的十几倍，经济又实惠，在当下适合作为HDD的升级。

  SSD使用的是NAND flash。

- SSD & HDD

  在第一部分我们可以看到，SSD比HDD读写速度快6-20倍，比HDD贵7-9倍，因其更加出色的读写性能，目前driver市场SSD与HDD并存，SSD有逐渐取代HDD的趋势。

  SSD有寿命限制，每个flash单元可擦除的次数有限。通过磨损均衡算法，可以最大延长SSD的使用寿命，甚至达到跟HDD一样的时长，使得寿命不再是SSD的短板。

- SSD nonvolatile memory type

  - SLC (Single-Level Cell): 一个cell内存储一字节数据
  - MLC (Moutiple-Level Cell): 存储容量提升两倍
  - TLC (Triple-Level Cell)
  - QLC ......

- Type of SSD driver

  - mSATA
  -  M.2 Module 
  - PCI-E

  根据主板接口选择SSD类型。

- HDD的改进：SMR (Shingled Magnetic Recording)

  ![](./images/3.png)

  SMR又称叠瓦记录技术，顾名思义，是将tracks叠放， tracks之间有重叠，因此较传统的HDD能够存放更多的数据。现在的SMR的容量已经可以达到单碟2GB。但SMR如果对已写过的数据直接进行修改操作，将会覆盖与其相邻的磁轨上的数据，所以在覆写之前，需要将可能被影响的数据转移，无疑降低了写的性能。因其低廉的价格和大容量，用作一次性写入的仓库式存储是很好的选择。

## 3. NVM (non-volatile memory)

![](./images/4.png)

NVM的中文是非易失性内存。具有非易失、bytes addressable等特点，读写性能接近DRAM但是还是有差距，下图是NVM与DRAM、NAND的对比：

![](./images/5.png)

NVM的延迟与DRAM在同一个量级，寿命却与NAND flash差不多，所以也需要考虑磨损均衡。使用NVM替代DRAM是一个设想，但在当下还无法实现，NVM的读写速度与DRAM还有差距。

目前出现了DRAM与NAND flash结合的NVM：NVDIMM

![](./images/6.png)

平时跟一块DRAM内存差不多，到断电的时候，将由内部的超级电容器供电，把DRAM中的数据存储到NAND中，实现数据的非易失。

目前NVM的存储材质有很多种，更多的技术还在研发中，因其出色的读写性能和非易失性，在将来可能取代内存和硬盘。

- 3D-XPOINT

  3D-XPOINT是英特尔研发的，据说比NAND快1000倍，寿命长1000倍。使用非易失性材料，目前已有产品。

# Key Indicators

- Density

  数据存放密度，密度越高，相同容量的存储介质体积越小。

- Rotation speed

  用于磁盘等需要转动的存储介质。Rotation speed不等于read/write speed，还应考虑到density等因素

- Read/write speed

  读取和写入的速度，不同的存储介质存在着读写速度的差异，同一个存储器的读速度和写速度也存在差异，一般读速度大于写速度。

- Durability

  存储介质的寿命，影响到数据存放的可靠性，寿命越长的存储介质更适合档案式存储，但寿命再长的存储介质终究是有寿命期限的。

- Price per GB

  每GB数据的成本。一般性能越好价格越高，价格低廉的在性能上必定有所妥协（如SSD与HDD）。有时需要根据实际应用场景在性能与价格间做trade-off。

# Make A Choice

存储器种类众多，排除掉已经被淘汰的产品（如软盘），如何在这些存储器中选择最合适的？

首先必须了解到这些存储介质各自的应用场景。对于data center而言，需要的自然是以SSD和HDD为主的非易失性大容量存储器。SSD价格昂贵，HDD性能不如SSD，对于档案式存放的数据，可使用价格低的HDD，对于常读常写的数据使用SSD。以这样的方式提高性价比，降低成本。

以上只是个小例子，每种存储介质实际上应用场景重叠不大，一般根据情况，也只会剩下2-3种备选方案，再充分考虑到当前预算和预期成本，自然会浮现出合适的方案。

- RAID

  RAID中文为独立冗余磁盘阵列。利用多个磁盘的组合以提供优于单个硬盘的读写性能和可用性。虽然RAID的初衷是节省成本，但后来更多的还是为了提高性能与可用性。

  RAID分为0~10个不同级别，从0-10性能逐渐下降、可用性逐渐提高，根据自己的实际情况选择RAID的级别，目前最常用的是RAID5.

