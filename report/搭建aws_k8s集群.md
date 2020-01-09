#### 

#### 安装Docker

```
sudo snap install docker
```

#### 安装microk8s

```
sudo snap install microk8s --classic
```

启动dns dashboard registry

```
sudo microk8s.enable dns dashboard registry
```



### 配置集群

master结点添加node

```
microk8s.add-node
```

slave结点

```
microk8s.join ip-172-31-20-243:25000/DDOkUupkmaBezNnMheTBqFYHLWINGDbf
```

查看node

```
microk8s.kubectl get no
```

#### 启动dashboard管理后台

dashboard是kubernetes提供的容器服务管理后台，可视化界面，用来进行机器负载，集群管理，镜像扩容，配置数据等相关操作

启动dashboard

```
microk8s.enable dashboard
```

获取token

```
token=$(microk8s.kubectl -n kube-system get secret | grep default-token | cut -d " " -f1)
microk8s.kubectl -n kube-system describe secret $token
```

暴露端口给外部

```
microk8s.kubectl port-forward -n kube-system service/kubernetes-dashboard 10443:443 --address 0.0.0.0
```

![avatar](https://github.com/Dianaaaa/SE419-422/blob/project/report/images/1.png)

## Backend Deployment

#### 启动deployment和service

```

microk8s.kubectl apply -f back-dm.yaml
microk8s.kubectl apply -f back-svc.yaml

```

通过nodeport暴露32200端口，可直接通过 IP 访问后端

![avatar](https://github.com/Dianaaaa/SE419-422/blob/project/report/images/3.png) 

POST请求

```
curl -i -X POST -H 'Content-type':'application/json' -d {\"path\":\"www.bilibili.com\"} http://172.31.1.7:32200/generate
```

GET请求

```
 curl 172.31.1.7:32200/get-url?shortlink=5HNjJo8U
```

![avatar](https://github.com/Dianaaaa/SE419-422/blob/project/report/images/post_get.png) 

###### 访问方式总结

service的clusterip可以使用，使用方式serviceClusterIP:port（其他方式都不行），所有的节点都可以通过访问serviceClusterIP:port进行访问服务

pod的clusterip进行访问，使用方式podClusterIP:targetport，但是每个节点只能使用自己节点的podClusterIP访问自己的pod

node的ip进行访问，使用方式nodeIP:nodePort（如果nodePort没有指定，创建service会自动生成一个）

## Database Deployment

安装mysql

```
sudo apt-get install mysql-client
```

连接到mysql

通过nodeport暴露32100 端口,，可直接通过 IP 访问mysql。

```
mysql -h 172.31.7.23 -P32100 -uroot -p
```

使用 Deployment 部署 MySQL 的 Pod

![avatar](https://github.com/Dianaaaa/SE419-422/blob/project/report/images/2.PNG) 

后端插入了数据后的情况

![avatar](https://github.com/Dianaaaa/SE419-422/blob/project/report/images/mysql.PNG) 