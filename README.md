# README

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

```
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

```
id：
route ： string 
service ：string
servicename ：string 服务的名字，这个是唯一的，注意这里是主键
```

