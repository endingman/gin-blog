#FROM指定基础镜像（必须有的指令，并且必须是第一条指令）
FROM golang:latest

#格式为 WORKDIR <工作目录路径>
#使用 WORKDIR 指令可以来指定工作目录（或者称为当前目录）
#以后各层的当前目录就被改为指定的目录，如果目录不存在，WORKDIR 会帮你建立目录

#golang:latest 镜像为基础镜像，将工作目录设置为 $GOPATH/src/go-gin-example
WORKDIR $GOPATH/src/go-gin-example


#COPY  格式：

#COPY <源路径>... <目标路径>
#COPY ["<源路径1>",... "<目标路径>"]

#，将当前上下文目录的内容复制到 $GOPATH/src/go-gin-example 中
COPY . $GOPATH/src/go-gin-example


#用于执行命令行命令   格式：RUN <命令>

#go build 编译完毕后，将容器启动程序设置为 ./go-gin-example，也就是我们所编译的可执行文件

#go module模式(不知道为什么设置不了，所以构建是不成功的)
RUN export GO111MODULE=on
#go module 代理(不知道为什么设置不了，所以构建是不成功的)
RUN export GOPROXY=https://goproxy.io

#ENTRYPOINT ["export GO111MODULE=on"]
#ENTRYPOINT ["export GOPROXY=https://goproxy.io"]

#查看验证是设置否正确
#RUN go env
#RUN go version

RUN go build .

#EXPOSE

#格式为 EXPOSE <端口1> [<端口2>...]
#EXPOSE 指令是声明运行时容器提供服务端口，这只是一个声明，在运行时并不会因为这个声明应用就会开启这个端口的服务
EXPOSE 8090


#go-gin-example 在 docker 容器里编译，并没有在宿主机现场编译
ENTRYPOINT ["./go-gin-example"]