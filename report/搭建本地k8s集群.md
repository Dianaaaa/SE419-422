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

![avatar]([http://https://github.com/Dianaaaa/SE419-422/tree/project/report/images/1.png))