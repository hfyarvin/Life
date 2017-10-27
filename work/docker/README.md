## Docker Hub

### uname -r 内核版本
### wget -qO- https://get.docker.com/ | sh
### apt install docker.io
### docker ps
### docker run -d nginx:v2
### docker stats --all
### docker start **
### docker stop **
### docker rm ** 删容器
### docker rmi ** 删镜像
### docker search redia 在Docker Hub
### docker pull redis:3.2
- [docker pull](https://www.processon.com/view/link/59f2a5ece4b0fd097f99d11a "docker pull inage")
### service docker start
### docker pull ubuntu:14.o4
### docker run -it --rm ubuntu:14.04 bash
- '--rm': 进程退出即删除容器
- '-ti': 分配一个伪终端并进入交互模式

### 虚拟镜像 
- docker images -f dangling=true //-f(filter过滤)
- docker rmi $(docker images -q -f dangling=true)

### 部分镜像
- docker images ubuntu
- docker images --format "{{.ID}}: {{.Repository}}"
- docker images --format "table {{.ID}}\t{{.Repository}}\t{{.Tag}}" //以表格展示

### 理解镜像构成 
- docker run --name webserver -d -p 80:80 nginx //localhost
- docker exec -it webserver bash
- echo '<h1>Hello, Docker!</h1>' > /usr/share/nginx/html/index.html
- exit
- docker diff **
- docker commit [选项] <容器ID或容器名> [<仓库名>[:<标签>]] 黑箱操作
- docker history nginx:v2

### docker-compose up -d
### docker-compose ps

## 新建镜像
- Dockerfile
- docker build -t new_docker_name .     //不能忽略点

## command
- docker attach $ID
- docker stop $ID
- docker attach