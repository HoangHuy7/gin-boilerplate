# 🏗️ gin-boilerplate

> **Production-ready Gin boilerplate with modular architecture, monorepo support, and best practices for scalable backend systems.**

This repository serves as a powerful **startkit monorepo** designed to jumpstart your Go backend development. It moves away from monolithic chaos, embracing a clean, domain-driven design that scales with your team and product.

## 🌟 Highlights

- **🧩 Modular Architecture**: Distinct domains (`iam`, `device`, `notification`) functioning as **child servers** (microservices), similar to Maven modules.
- **🏗️ Monorepo Structure**: **`internal`** acts as the **Core/Shared Library** (like a Maven parent/common), holding base logic, DTOs, and router configurations used by all child services.
- **📜 Auto Swagger / OpenAPI**: Automatic API documentation generation using `swaggest/openapi-go`. Just define your DTOs and Controller metadata, and the docs are ready!
- **🛡️ Production Ready**: Pre-configured with logging, robust routing strategies, and standard DTOs.
- **🔌 Scalable & Extensible**: Built on top of [Gin](https://github.com/gin-gonic/gin), ready to grow from a startup MVP to a high-load system.

## 📂 Project Structure

```text
.
├── apps                # 🏢 Container for all Child Servers logic
│   ├── device          # Device Service Logic
│   ├── iam             # IAM Service Logic
│   │   └── controller
│   │       ├── Module.go
│   │       ├── Router.go
│   │       └── v1
│   │           └── HelloController.go
│   └── notification    # Notification Service Logic
│       └── controller
│           └── v1
├── cmd
│   ├── device          # Entry point for Device Child Server
│   │   └── main.go
│   ├── iam             # Entry point for IAM Child Server
│   │   └── main.go
│   └── notification    # Entry point for Notification Child Server
│       └── main.go
├── internal            # 🧱 Core / Shared Libraries (Maven-like Common Module)
│   ├── base
│   │   ├── Base.go
│   │   └── routerx
│   │       └── Routerx.go
│   ├── dto
│   │   └── system.go
│   ├── logger
│   │   └── module.go
│   └── server
│       ├── router.go
│       └── server.go
├── go.mod
├── go.sum
└── main.go
```

## 🛠️ Getting Started

### Prerequisites

- **Go** (1.20 or higher)

### Installation

Clone the repository:

```bash
git clone https://github.com/HoangHuy7/gin-boilerplate.git
cd gin-boilerplate
go mod download
```

### Running a Microservice

Each domain has its own entry point in `cmd/`. For example, to run the **IAM** service:

```bash
go run cmd/iam/main.go
```

### 📚 API Documentation (Swagger)

After running a service, you can access the Swagger UI at:
- **URL**: `http://localhost:8080/swagger/` (Port may vary based on configuration)


## 🤝 Contribution

Contributions are welcome! Focus on keeping the `internal` directory clean and reusable across different domains.

---

Crafted with ❤️ by **HoangHuy7**
