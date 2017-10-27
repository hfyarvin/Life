## 三组件
- Docker Client
- Docker Daemon
- Docker Index

## 三基本要素
- DockerFile --> Docker Images --> Docker Containers

## 命令
- docker pull
- docker push

### 例1
- sudo docker pull busybox
- docker run busybox /bin/echo Hello World

### 例2
- sample_job=$(docker run -d busybox /bin/sh -c "while true; do echo Docker; sleep 1; done")
- docker logs $sample_job
- docker stop $sample_job
- docker restart $sample_job
- docker stop $sample_job docker rm $sample_job
- docker commit $sample_job job1

## 创建镜像
- FROM <images>
- MAINTAINER <author name>
- RUN <command>
- ADD <src> <destination>
- CMD command param1 param2  //只允许使用一次
- EXPOSE <port>
- ENTRYPOINT command param1 param2  //默认程序
- WORKDIR /path/to/workdir
- ENV <key><value>
- USER <uid>
- VOLUME ['/data'] 授权访问从容器内到主机的目录
- 

## webhooks --> Docker Registry
- index
- registry
- registry client