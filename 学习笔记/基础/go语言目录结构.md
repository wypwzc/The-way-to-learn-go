# 一个标准的 Go 项目目录

```
my-project/
│
├── cmd/                    # 程序入口（main函数）
│   └── server/
│       └── main.go
│
├── internal/               # 内部业务代码（外部不能引用）
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── model/
│   ├── middleware/
│   ├── router/
│   └── config/
│
├── pkg/                    # 可复用公共代码
│   ├── logger/
│   ├── utils/
│   ├── response/
│   └── validator/
│
├── api/                    # API定义（Swagger、Proto）
│   ├── swagger/
│   └── proto/
│
├── configs/                # 配置文件
│   ├── config.yaml
│   └── config-dev.yaml
│
├── scripts/                # Shell脚本
│   ├── build.sh
│   └── deploy.sh
│
├── deployments/            # Docker/Kubernetes部署文件
│   ├── Dockerfile
│   ├── docker-compose.yml
│   └── k8s/
│
├── docs/                   # 项目文档
│
├── test/                   # 集成测试
│
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

这是目前企业中最常见的结构。

------

# 每个目录的作用

## ① cmd/

程序入口。

例如：

```
cmd/
    server/
        main.go
```

这里一般只有：

```
func main() {
    // 初始化配置
    // 初始化数据库
    // 初始化日志
    // 初始化路由
    // 启动HTTP服务
}
```

**main.go 不应该写业务逻辑。**

------

## ② internal/

这是整个项目最重要的目录。

官方推荐：

```
internal/
```

里面放所有业务代码。

例如：

```
internal/
    handler/
    service/
    repository/
    model/
```

### handler

HTTP接口。

例如：

```
func Login(c *gin.Context)
```

负责：

- 接收请求
- 参数校验
- 返回JSON

**不处理业务逻辑。**

------

### service

业务逻辑。

例如：

```
登录

注册

下单

支付
```

例如：

```
func Login(username,password string)
```

这里才是真正的业务。

------

### repository（DAO）

负责数据库。

例如：

```
SELECT

INSERT

UPDATE
```

例如：

```
db.Create()

db.First()

db.Delete()
```

------

### model

数据结构。

例如：

```
type User struct {
    ID int
    Name string
}
```

------

### router

路由。

例如：

```
r.GET("/login",handler.Login)
```

------

### middleware

中间件。

例如：

```
JWT

日志

跨域

权限
```

------

### config

配置读取。

例如：

```
viper.ReadInConfig()
```

------

# ③ pkg/

公共代码。

例如：

```
utils/

logger/

jwt/

email/

validator/
```

特点：

这些代码以后可以复制到其他项目继续使用。

例如：

```
logger

response

jwt

utils
```

都可以放这里。

------

# ④ api/

一般保存：

```
Swagger

OpenAPI

Proto
```

例如：

```
api/

swagger/

proto/
```

如果项目以后要做gRPC：

```
api/proto/user.proto
```

------

# ⑤ configs/

配置文件。

例如：

```
config.yaml

config-dev.yaml

config-prod.yaml
```

例如：

```
server:
  port: 8080

mysql:
  host: localhost
```

------

# ⑥ deployments/

部署文件。

例如：

```
Dockerfile

docker-compose.yml

k8s/
```

Kubernetes：

```
deployment.yaml

service.yaml

ingress.yaml
```

------

# ⑦ scripts/

脚本。

例如：

```
build.sh

deploy.sh

backup.sh
```

------

# ⑧ docs/

文档。

例如：

```
API.md

README

设计文档
```

------

# ⑨ test/

集成测试。

例如：

```
API测试

数据库测试
```

------

# 根目录文件

## go.mod

项目依赖。

例如：

```
module github.com/demo/shop
```

以及：

```
require (
    github.com/gin-gonic/gin
)
```

------

## go.sum

依赖校验。

不用修改。

------

## Makefile

企业非常常见。

例如：

```
build:
	go build -o app ./cmd/server

run:
	go run ./cmd/server

test:
	go test ./...

fmt:
	go fmt ./...
```

以后只需要：

```
make build

make run

make test
```

------

## README.md

项目说明。

包括：

- 如何启动
- API说明
- Docker部署

------

# 一个HTTP请求是如何流转的？

假设请求：

```
POST /login
```

整个调用链通常是：

```
客户端
      │
      ▼
cmd/server/main.go
      │
      ▼
router/router.go
      │
      ▼
middleware/
      │
      ▼
handler/login.go
      │
      ▼
service/user.go
      │
      ▼
repository/user.go
      │
      ▼
MySQL
```

返回结果时则按相反方向逐层返回。

------

# 对于你来说（Go 后端 + DevOps），建议采用这种目录结构

结合你未来想做 Go 后端开发，并且有 Kubernetes、Docker、CI/CD 的学习方向，建议直接使用下面这种企业级结构：

```
go-demo/
├── cmd/                  # 程序入口
├── internal/
│   ├── handler/          # HTTP接口
│   ├── service/          # 业务逻辑
│   ├── repository/       # 数据访问
│   ├── model/            # 数据模型
│   ├── router/           # 路由注册
│   ├── middleware/       # JWT、日志等中间件
│   └── config/           # 配置加载
├── pkg/                  # 通用工具（日志、响应封装等）
├── configs/              # YAML配置文件
├── deployments/          # Dockerfile、Kubernetes YAML
├── scripts/              # 构建、部署脚本
├── api/                  # Swagger 或 Protobuf
├── docs/
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

这套结构足够支撑从小型项目到中大型后端服务，也与很多开源 Go 项目的组织方式保持一致，后续学习微服务、Docker、Kubernetes 时也能自然衔接。