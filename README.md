<p align="center">
  <img src="https://raw.githubusercontent.com/HoangHuy7/gin-boilerplate/main/.github/assets/banner.png" alt="Gin Monorepo Banner" width="100%">
</p>

<h1 align="center">Gin Monorepo Boilerplate</h1>

<p align="center">
  <em>Production-ready Gin boilerplate with modular architecture, monorepo support, and Uber-fx powered automation.</em>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat-square&logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/Gin-Framework-blue?style=flat-square" alt="Gin Framework">
  <img src="https://img.shields.io/badge/Uber--fx-DI-red?style=flat-square" alt="Uber-fx">
  <img src="https://img.shields.io/badge/Casbin-Auth-orange?style=flat-square" alt="Casbin">
</p>

---

## 🟦 Highlights

- **Modular Architecture**: Domains like `iam`, `device`, and `notification` function as independent modules.
- **Monorepo Structure**: The `internal` directory holds shared logic (Core), DTOs, and server configurations.
- **Dependency Injection**: Powered by Uber-fx for clean lifestyle management and automatic component wiring.
- **Automated Registration**: Controllers and OpenAPI documentation are registered automatically.
- **Authorization**: Built-in Casbin support for RBAC/ABAC.
- **Auto Swagger / OpenAPI**: Reflection-based Swagger generation. No manual documentation required.

## 🟦 Project Structure

```text
.
├── apps                    # Micro-apps / Domain Logic
│   ├── device
│   ├── iam                 # Identity & Access Management
│   │   ├── app             # App Wiring (DB, Auth, Config, Module)
│   │   │   ├── casbin
│   │   │   ├── config
│   │   │   ├── database
│   │   │   └── Module.go
│   │   └── controller      # HTTP Handlers
│   │       ├── v1
│   │       │   └── HelloController.go
│   │       └── Module.go   # Fx registration logic
│   └── notification
├── cmd                     # Execution Entry Points
│   ├── iam/main.go
├── configs                 # App Configurations (YAML, Casbin)
├── internal                # Shared Core Library
│   ├── base                # Base interfaces (Controller, etc.)
│   ├── dto                 # Shared DTOs & Search/Metadata schemas
│   ├── logger              # Zap-based logging
│   ├── server              # Core HTTP server, Router & OpenAPI logic
│   └── utils               # Shared utilities
├── go.mod
└── main.go
```

## 🟦 Automation Logic

The boilerplate uses [Uber-fx](https://github.com/uber-go/fx) to handle dependency injection and lifecycle.

### 1. Controller Registration
In `apps/iam/controller/Module.go`, controllers are annotated to a group:
```go
fx.Annotate(
    v1.NewHelloController,
    fx.As(new(base.Controller)),
    fx.ResultTags(`group:"controllers"`),
)
```
The `internal/server` consumes this group to mount all routes automatically.

### 2. Automatic OpenAPI Integration
This project uses a code-first reflection approach via `routerx`.

#### Step 1: Define the endpoint in Controller
In your `Register` method, use `dto.OpenEndpoint` to describe your API:
```go
func (this *HelloController) Register(rg *routerx.Routerx) {
    rg.POST(dto.OpenEndpoint{
        Path:        "/json",
        Handler:     this.JSON,
        Summary:     "Create something",
        Request:     &dto.CreatePostRequest{},
        Responses:   map[int]any{
            200: gin.H{"status": "ok"},
        },
    })
}
```

#### Step 2: Enable in Metadata
Ensure `EnableOpenAPI: true` is set in your controller's metadata:
```go
Metadata: dto.Metadata{
    Tag:           "User Management",
    EnableOpenAPI: true,
}
```

#### Step 3: Access it
Run your service and navigate to: `http://localhost:8080/swagger/`

## 🟦 Getting Started

### Installation
```bash
git clone https://github.com/HoangHuy7/gin-boilerplate.git
cd gin-boilerplate
go mod download
```

### Running a Service
```bash
go run cmd/iam/main.go
```

---
<p align="center">Crafted with ❤️ by <b>HoangHuy7</b></p>
