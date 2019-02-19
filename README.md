# README

### 使用

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



### 运行

#### 直接运行

```GO
cd GOPATH/ConfCenter
go run main.go -f ./config/config.toml
```

#### 编译运行

```GO
cd GOPATH
go build ConfCenter

linux/mac环境下
./ConfCenter -f ./src/ConfCenter/config/config.toml
win环境下
ConfCenter.exe -f ./src/ConfCenter/config/config.toml
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

