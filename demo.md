# Demo Project - Gin Monorepo Boilerplate

## 📋 Tổng quan

Đây là một boilerplate project xây dựng bằng **Go (Golang)** với kiến trúc **Monorepo** và **Microservices**, sử dụng framework **Gin** kết hợp với **Uber-fx** cho dependency injection.

## 🏗️ Kiến trúc hệ thống

### Cấu trúc thư mục chính

```
.
├── apps/                   # Business logic của từng domain
│   ├── gas/               # Module xử lý khí gas
│   ├── device/            # Module quản lý thiết bị
│   └── notification/      # Module gửi thông báo
│
├── cmd/                    # Entry points cho các service
│   ├── gas/main.go
│   ├── device/main.go
│   └── notification/main.go
│
├── configs/                # File cấu hình YAML
│   ├── gas/
│   ├── device/
│   └── notification/
│
├── internal/               # Shared code dùng chung
│   ├── base/              # Base interfaces & utilities
│   ├── dto/               # Data Transfer Objects
│   ├── logger/            # Logging module (Zap)
│   ├── server/            # HTTP Server & Router
│   └── utils/             # Utility functions
│
├── shares/                 # Entities dùng chung
│   └── entities/          # Database models
│
└── go.mod                  # Go module definition
```

## 🔧 Công nghệ sử dụng

| Component | Technology |
|-----------|------------|
| Framework | Gin Web Framework |
| DI        | Uber-fx |
| ORM       | GORM |
| Logging   | Zap Logger |
| Config    | Viper |
| Auth      | Casbin (RBAC/ABAC) |
| API Docs  | Swagger/OpenAPI (auto-generated) |

## 🚀 Các service chính

### 1. Gas Service (`apps/gas/`)
- Quản lý thông tin khí gas
- Xử lý dữ liệu từ cảm biến
- Theo dõi mức tiêu thụ

### 2. Device Service (`apps/device/`)
- Quản lý thiết bị IoT
- Đăng ký và theo dõi trạng thái thiết bị
- Xử lý kết nối device-to-server

### 3. Notification Service (`apps/notification/`)
- Gửi thông báo push/email/SMS
- Quản lý template thông báo
- History logs

## ⚙️ Tính năng nổi bật

### 1. Dependency Injection tự động
Sử dụng Uber-fx để tự động wire các dependencies:
```go
fx.Provide(
    NewDatabase,
    NewRepository,
    NewService,
    fx.Annotate(
        NewController,
        fx.As(new(base.Controller)),
        fx.ResultTags(`group:"controllers"`),
    ),
)
```

### 2. Auto Register Controllers
Tất cả controllers được đăng ký tự động vào router mà không cần config thủ công.

### 3. Swagger tự động sinh
Chỉ cần define endpoint với `dto.OpenEndpoint`, Swagger docs sẽ được sinh tự động:
```go
rg.POST(dto.OpenEndpoint{
    Path:      "/api/v1/customers",
    Handler:   this.CreateCustomer,
    Summary:   "Create new customer",
    Request:   &dto.CreateCustomerRequest{},
    Responses: map[int]any{200: CustomerResponse{}},
})
```

### 4. Multi-environment Config
Hỗ trợ environment variables trong file YAML:
```yaml
database:
  host: ${DB_HOST}
  port: ${DB_PORT}
  password: ${DB_PASSWORD}
```

## 📖 Hướng dẫn chạy project

### Cài đặt
```bash
git clone <repository-url>
cd <project-folder>
go mod download
```

### Chạy service
```bash
# Chạy gas service
go run cmd/gas/main.go

# Chạy device service
go run cmd/device/main.go

# Chạy notification service
go run cmd/notification/main.go
```

### Truy cập Swagger UI
Sau khi chạy service, truy cập:
```
http://localhost:8080/swagger/
```

## 🔐 Bảo mật

- **Casbin**: Hỗ trợ RBAC (Role-Based Access Control) và ABAC (Attribute-Based Access Control)
- **JWT Authentication**: Xác thực người dùng qua JWT tokens
- **Environment Variables**: Lưu trữ secrets trong biến môi trường

## 📝 Quy tắc phát triển

Xem file [CONTRIBUTING.md](./CONTRIBUTING.md) để biết chi tiết về coding conventions và quy trình đóng góp.

## 📄 License

Xem file [LICENSE](./LICENSE) để biết thông tin về giấy phép sử dụng.

---

<p align="center"><b>Developed with ❤️ using Go & Gin Framework</b></p>
