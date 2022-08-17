[toc]


### postgresql
```shell
1.拉取镜像
docker pull postgres:11

2.创建数据卷
docker volume create pgdata
# 查看数据卷(mountpoint)
docker volume inspect pgdata

3.启动
docker run --name postgres11 -e POSTGRES_PASSWORD=password -p 5432:5432  -d postgres:11

```


### rabbitmq
```shell
1.拉取镜像
docker pull rabbitmq:management

2.启动服务
docker run -d -p 5672:5672 -p 15672:15672 --name rabbitmq rabbitmq:management

# 默认账号密码 guest/guest (http://ip:15672)
```