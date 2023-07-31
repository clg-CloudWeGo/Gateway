# Gateway
```
.
├── APILayer
├── CourseManagementService
├── CoursePurchaseSeervice
├── idl
│   ├── course_manage.thrift
│   ├── course_purchase.thrift
│   ├── gateway.thrift
│   └── idlmanager.thrift
├── IDLManagementPlatform
└── README.md
```

## 依赖
- go
- kitex
- hertz
- etcd


## 运行

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
学生服务idl
```shell
curl --location 'http://localhost:6666/idl/add' \
--header 'Content-Type: application/json' \
--data '{
    "name": "StudentService",
    "idl":"namespace go demo\n\nstruct College {\n    1: required string name(go.tag = '\''json:\"name\"'\''),\n    2: string address(go.tag = '\''json:\"address\"'\''),\n}\n\nstruct Student {\n    1: required i32 id(api.body='\''id'\''),\n    2: required string name(api.body='\''name'\''),\n    3: required College college(api.body='\''college'\''),\n    4: optional list<string> email(api.body='\''email'\''),\n}\n\nstruct RegisterResp {\n    1: bool success(api.body='\''success'\''),\n    2: string message(api.body='\''message'\''),\n}\n\nstruct QueryReq {\n    1: required i32 id(api.query='\''id'\'')\n}\n\nservice StudentService {\n    RegisterResp Register(1: Student student)(api.post = '\''/gateway/StudentService/Register'\'')\n    Student Query(1: QueryReq req)(api.get = '\''/gateway/StudentService/Query'\'')\n}"
}'

```
Echo服务idl
```shell
curl --location 'http://localhost:6666/idl/add' \
--header 'Content-Type: application/json' \
--data '{
    "name": "EchoService",
    "idl":"namespace go api\n\nstruct Request {\n    1: string message\n}\n\nstruct Response {\n    1: string message\n}\n\nservice Echo {\n    Response echo(1: Request req)(api.post=\"/gateway/EchoService/Echo\")\n}"
}'

```

### 检测网关功能
调用Echo服务
```shell
curl --location 'http://localhost:8888/gateway/EchoService/Echo' \
--header 'Content-Type: application/json' \
--data '{
  "message": "hello raccoon"
}'
```
调用学生注册服务
```shell
curl -H "Content-Type: application/json" -X POST http://127.0.0.1:8888/gateway/StudentService/Register -d '{"id": 100, "name":"Emma", "college": {"name": "software college", "address": "逸夫"}, "email": ["emma@nju.com"]}'
```
调用学生查询服务（没实现，需要把get改成post）
```shell
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:8888/gateway/StudentService/Query
```

### 清理数据库残余
若不清理，下次注册idl会显示已经注册
```shell
rm IDLManagementPlatform/IDLMessage.db
```