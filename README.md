# README

```go
ConfCenter 是一个基于go语言开发微服务配置中心。
特点：
轻量级
基于go语言强大的并发能力毋庸质疑
跨语言，跨平台
轻松上云，docker镜像小于三十m
所有语言都可以轻松使用
占用内存小
轻松对远程调用
...
```



### 运行

```go
cd $GOPATH
go get github.com/EpochCloud/ConfCenter
cd src/github.com/EpochCloud/ConfCenter
go install
cd $GOPATH
mv ConfCenter $GOPATH   //这里如果是windows是ConfCenter.exe
win环境
ConfCenter.exe -f ./src/ConfCenter/config/config.toml
linux/mac环境
./ConfCenter -f ./src/ConfCenter/config/config.toml
```

### 数据库表结构

#### configuration表

```go
id 
ip：string ：需要启动的ip
port：string ：需要启动的端口
timeout ： int ：api控制的超时时间
loglevel ： string ：日志的级别
logpath ： string ：日志的打印路径
modification ： uint64 ：是否被覆盖，如果被覆盖是1，没有覆盖是0
bufpool ： int ：池子容量
```

#### service表

```go
id：
route ： string 
service ：string
servicename ：string 服务的名字，这个是唯一的，注意这里是主键
```

#### allservice表

```go
id:
route:string 
ip   :string
port :string
srvname:string //服务名字
srv  : string //服务配置
```

### 使用

#### 接口/service_operation

###### GET

```go
get请求主要是查看添加的服务的所有的配置的详细信息
```

###### POST

```go
post方法用于添加服务的所有的配置
```

###### patch

```go
patch方法用于更新服务的配置
```



