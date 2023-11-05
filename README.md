# Gateway

> 队伍：潜心一志@OK爸
> 
> 成员：
>


> 项目结构见文档结尾
## 环境配置

详见[环境配置文档.md](./环境配置文档.md).

- go
- kitex
- hertz
- etcd
- thriftgo


## 接口文档

### 启动网关，注册中心和服务
启动etcd
```shell
etcd --log-level debug
```

启动网关（默认8888端口）
```shell
cd Gateway
go run .
```

启动IDL管理平台（6666端口）
```shell
cd IDLManagementPlatform
go run .
```

启动Echo服务（8887端口）
```shell
cd EchoService
go run .
```

启动学生服务（8889端口）
```shell
cd StudentService
go run .
```


### 为IDL管理平台添加服务对应的idl
1. 学生服务idl

```shell
curl --location 'http://localhost:6666/idl/add' \
--header 'Content-Type: application/json' \
--data '{
    "name": "StudentService",
    "idl":"namespace go demo\n\nstruct College {\n    1: required string name(go.tag = '\''json:\"name\"'\''),\n    2: string address(go.tag = '\''json:\"address\"'\''),\n}\n\nstruct Student {\n    1: required i32 id(api.body='\''id'\''),\n    2: required string name(api.body='\''name'\''),\n    3: required College college(api.body='\''college'\''),\n    4: optional list<string> email(api.body='\''email'\''),\n}\n\nstruct RegisterResp {\n    1: bool success(api.body='\''success'\''),\n    2: string message(api.body='\''message'\''),\n}\n\nstruct QueryReq {\n    1: required i32 id(api.query='\''id'\'')\n}\n\nservice StudentService {\n    RegisterResp Register(1: Student student)(api.post = '\''/gateway/StudentService/Register'\'')\n    Student Query(1: QueryReq req)(api.get = '\''/gateway/StudentService/Query'\'')\n}"
}'

```


2. Echo服务idl

```shell
curl --location 'http://localhost:6666/idl/add' \
--header 'Content-Type: application/json' \
--data '{
    "name": "EchoService",
    "idl":"namespace go api\n\nstruct Request {\n    1: string message\n}\n\nstruct Response {\n    1: string message\n}\n\nservice Echo {\n    Response echo(1: Request req)(api.post=\"/gateway/EchoService/Echo\")\n}"
}'

```

### 检测网关功能
1. 调用Echo服务

```shell
curl --location 'http://localhost:8888/gateway/EchoService/Echo' \
--header 'Content-Type: application/json' \
--data '{
  "message": "hello raccoon"
}'
```

2. 调用学生注册服务

```shell
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/gateway/StudentService/Register -d '{"id": 100, "name":"Emma", "college": {"name": "software college", "address": "逸夫"}, "email": ["emma@nju.com"]}'
```



### 清理数据库残余
若不清理，下次注册idl会显示已经注册
```shell
rm IDLManagementPlatform/IDLMessage.db
```

## 性能测试
见[测试文档.md](./测试文档.md).

## 项目结构

```
├── EchoService
│   ├── build.sh
│   ├── client
│   │   └── main.go
│   ├── echo.thrift
│   ├── echo_test.go
│   ├── go.mod
│   ├── go.sum
│   ├── handler.go
│   ├── kitex_gen
│   │   └── api
│   │       ├── echo
│   │       │   ├── client.go
│   │       │   ├── echo.go
│   │       │   ├── invoker.go
│   │       │   └── server.go
│   │       ├── echo.go
│   │       ├── k-consts.go
│   │       └── k-echo.go
│   ├── kitex_info.yaml
│   ├── main.go
│   ├── script
│   │   └── bootstrap.sh
│   └── todo.md
├── Gateway
│   ├── biz
│   │   ├── handler
│   │   │   └── hertzSvr
│   │   │       ├── gateway
│   │   │       │   └── gateway.go
│   │   │       ├── idlManager
│   │   │       │   └── idlservice.go
│   │   │       ├── init.go
│   │   │       └── utils
│   │   │           └── utils.go
│   │   ├── model
│   │   │   └── hertzSvr
│   │   │       └── idlManager
│   │   │           └── gateway.go
│   │   └── router
│   │       ├── hertzSvr
│   │       │   └── idlManager
│   │       │       ├── gateway.go
│   │       │       └── middleware.go
│   │       └── register.go
│   ├── build.sh
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── router.go
│   ├── router_gen.go
│   └── script
│       └── bootstrap.sh
├── IDLManagementPlatform
│   ├── IDLMessage.db
│   ├── biz
│   │   ├── handler
│   │   │   └── hertzSvr
│   │   │       └── service
│   │   │           └── idlservice.go
│   │   ├── model
│   │   │   └── hertzSvr
│   │   │       └── service
│   │   │           └── idlmanager.go
│   │   └── router
│   │       ├── hertzSvr
│   │       │   └── service
│   │       │       ├── idlmanager.go
│   │       │       └── middleware.go
│   │       └── register.go
│   ├── build.sh
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── router.go
│   ├── router_gen.go
│   └── script
│       └── bootstrap.sh
├── README.md
├── StudentService
│   ├── build.sh
│   ├── client
│   │   └── main.go
│   ├── foo.db
│   ├── go.mod
│   ├── go.sum
│   ├── handler.go
│   ├── kitex_gen
│   │   └── demo
│   │       ├── k-consts.go
│   │       ├── k-student.go
│   │       ├── student.go
│   │       └── studentservice
│   │           ├── client.go
│   │           ├── invoker.go
│   │           ├── server.go
│   │           └── studentservice.go
│   ├── kitex_info.yaml
│   ├── main.go
│   ├── model
│   │   └── student.go
│   ├── output
│   │   ├── bin
│   │   │   └── student.nju.rpc
│   │   ├── bootstrap.sh
│   │   └── log
│   │       ├── app
│   │       └── rpc
│   ├── script
│   │   └── bootstrap.sh
│   ├── student.thrift
│   └── student_test.go
├── default.etcd
│   └── member
│       ├── snap
│       │   └── db
│       └── wal
│           ├── 0.tmp
│           └── 0000000000000000-0000000000000000.wal
├── idl
│   ├── gateway.thrift
│   └── idlmanager.thrift
└── 测试文档.md

50 directories, 74 files

```
