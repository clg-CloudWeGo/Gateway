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


```shell
cd Gateway
go run .
```

```shell
curl --location 'http://localhost:6666/idl/add' \
--header 'Content-Type: application/json' \
--data '{
    "name": "StudentService",
    "idl":"namespace go demo\n\nstruct College {\n    1: required string name(go.tag = '\''json:\"name\"'\''),\n    2: string address(go.tag = '\''json:\"address\"'\''),\n}\n\nstruct Student {\n    1: required i32 id(api.body='\''id'\''),\n    2: required string name(api.body='\''name'\''),\n    3: required College college(api.body='\''college'\''),\n    4: optional list<string> email(api.body='\''email'\''),\n}\n\nstruct RegisterResp {\n    1: bool success(api.body='\''success'\''),\n    2: string message(api.body='\''message'\''),\n}\n\nstruct QueryReq {\n    1: required i32 id(api.query='\''id'\'')\n}\n\nservice StudentService {\n    RegisterResp Register(1: Student student)(api.post = '\''/add-student-info'\'')\n    Student Query(1: QueryReq req)(api.get = '\''/query'\'')\n}"
}'

```

```shell
curl --location 'http://localhost:6666/idl/add' \
--header 'Content-Type: application/json' \
--data '{
    "name": "EchoService",
    "idl":"namespace go api\n\nstruct Request {\n  1: string message\n}\n\nstruct Response {\n  1: string message\n}\n\nservice Echo {\n    Response echo(1: Request req)\n}"
}'

```

```shell
curl --location 'http://localhost:8888/gateway/EchoService/Echo' \
--header 'Content-Type: application/json' \
--data '{
  "message": "hello raccoon",
}'
```