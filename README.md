# 🏗️ gin-boilerplate

> **Production-ready Gin boilerplate with modular architecture, monorepo support, and Uber-fx powered automation.**

This repository is a **startkit monorepo** designed for scalable Go backend development. It features a domain-driven design that separates core logic from application-specific modules.

## 🌟 Highlights

- **🧩 Modular Architecture**: Domains like `iam`, `device`, and `notification` function as independent modules.
- **🏗️ Monorepo Structure**: The `internal` directory holds shared logic (Core), DTOs, and server configurations.
- **⚡ Dependency Injection**: Powered by **Uber-fx** for clean lifecycle management and automatic component wiring.
- **🤖 Automated Controller Registration**: No more manual routing for every controller. Just provide it to the module, and it's live.
- **🔐 Authorization**: Built-in **Casbin** support for RBAC/ABAC.
- **📜 Auto Swagger / OpenAPI**: Reflection-based Swagger generation using `swaggest/openapi-go`.

## 📂 Project Structure

```text
.
├── apps                    # 🏢 Micro-apps / Domain Logic
│   ├── device              # Device Domain
│   ├── iam                 # Identity & Access Management
│   │   ├── app             # App-specific wiring (DB, Auth, Config)
│   │   └── controller      # HTTP Handlers
│   └── notification        # Notification Domain
├── cmd                     # 🚀 Execution Entry Points
│   ├── device/main.go
│   ├── iam/main.go
│   └── notification/main.go
├── configs                 # ⚙️ App Configurations (YAML, Casbin)
├── internal                # 🧱 Shared Core Library
│   ├── base                # Base interfaces (Controller, etc.)
│   ├── logger              # Zap-based logging
│   ├── server              # Core HTTP server & Router logic
│   └── utils               # Shared utilities
├── go.mod
└── main.go                 # Root entry (optional/bridge)
```

## 🤖 How Automation Works

The boilerplate uses [Uber-fx](https://github.com/uber-go/fx) to automate the wiring of dependencies, specifically for controllers.

### 1. The Core Router (`internal/server`)
The `NewRouter` function in `internal/server/router.go` is designed to receive a list of controllers via dependency injection:
```go
func NewRouter(controllers []base.Controller, ...) *Router
```

### 2. Automatic Registration (`apps/iam/controller`)
In your module's controller package (e.g., `apps/iam/controller/Module.go`), you register controllers using **Group Tags**:
```go
fx.Annotate(
    v1.NewHelloController,
    fx.As(new(base.Controller)),
    fx.ResultTags(`group:"controllers"`), // Adds to the "controllers" group
)
```
And then inject that group into the `NewRouter`:
```go
fx.Annotate(
    server.NewRouter,
    fx.ParamTags(`group:"controllers"`), // Injects all controllers from the group
)
```

### 3. Wiring it up (`cmd/iam`)
In the `main.go` of your service, simply include the controller module:
```go
fx.New(
    app.Module,
    controller.Module, // Automation happens here
    // ...
    fx.Invoke(server.RunServer),
).Run()
```

## 🛠️ Getting Started

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

### 📚 API Documentation
Access Swagger UI at: `http://localhost:8080/swagger/` (Port depends on your config).

## 🤝 Contribution
Keep the `internal` directory clean and reusable. If you add a new shared utility, ensure it follows the base interfaces.

---
Crafted with ❤️ by **HoangHuy7**
