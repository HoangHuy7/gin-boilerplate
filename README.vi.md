<p align="center">
  <img src="https://raw.githubusercontent.com/HoangHuy7/gin-boilerplate/main/.github/assets/banner.jpg" alt="Gin Monorepo Banner">
</p>

<h1 align="center">Gin Monorepo Boilerplate</h1>

<p align="center">
  <em>Bộ startkit monorepo xịn sò cho Go backend, kiến trúc modular, monorepo chuẩn chỉ kèm automation bằng Uber-fx.</em>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat-square&logo=go" alt="Go Version">
  <img src="https://img.shields.io/badge/Gin-Framework-blue?style=flat-square" alt="Gin Framework">
  <img src="https://img.shields.io/badge/Uber--fx-DI-red?style=flat-square" alt="Uber-fx">
  <img src="https://img.shields.io/badge/Casbin-Auth-orange?style=flat-square" alt="Casbin">
  <img src="https://img.shields.io/badge/GORM-ORM-blue?style=flat-square" alt="GORM">
  <img src="https://img.shields.io/badge/Swagger-OpenAPI-85EA2D?style=flat-square&logo=swagger" alt="Swagger">
</p>

---

## 🟦 Highlights

- **Modular Architecture**: Chia domain (`iam`, `device`, ...) rõ ràng, độc lập và dễ quản lý.
- **Monorepo Structure**: Thư mục `internal` chứa toàn bộ logic dùng chung, đảm bảo tính nhất quán.
- **Dependency Injection**: Sử dụng Uber-fx để tự động hóa việc kết nối các component.
- **Automatic Automation**: Cả Controller và OpenAPI đều được đăng ký tự động qua hệ thống module.
- **Authorization**: Tích hợp sẵn Casbin cho việc phân quyền RBAC/ABAC.
- **OpenAPI/Swagger**: Tự động sinh documentation từ code, không cần viết comment thủ công.

## 🟦 Cấu trúc dự án

```text
.
├── apps                    # Nghiệp vụ chính (Domain Logic)
│   ├── device
│   ├── iam                 # Quản lý định danh
│   │   ├── app             # Cấu hình module (DB, Auth, Config, Module)
│   │   │   ├── casbin
│   │   │   ├── config
│   │   │   ├── database
│   │   │   └── Module.go
│   │   └── controller      # Xử lý HTTP Request
│   │       ├── v1
│   │       │   └── HelloController.go
│   │       └── Module.go   # Đăng ký controller với Fx
│   └── notification
├── cmd                     # Entry points thực thi
│   ├── iam/main.go
├── configs                 # Cấu hình hệ thống (YAML, Policy)
├── internal                # Shared Core (Trái tim hệ thống)
│   ├── base                # Base interfaces
│   ├── dto                 # Data Transfer Objects
│   ├── logger              # Zap Logger
│   ├── server              # Core Server & OpenAPI logic
│   └── utils               # Đồ nghề hỗ trợ
├── go.mod
```

## 🟦 Hướng dẫn Tự động hóa

Project tận dụng sức mạnh của Uber-fx để giải phóng việc khai báo router thủ công.

### 1. Đăng ký Controller
Chỉ cần khai báo controller trong module tương ứng, hệ thống sẽ tự động nhận diện:
```go
fx.Annotate(
    v1.NewHelloController,
    fx.As(new(base.Controller)),
    fx.ResultTags(`group:"controllers"`),
)
```

### 2. Tích hợp OpenAPI Tự động
Sử dụng phương pháp Code-First với `routerx` để sinh Swagger UI mà không cần viết annotation phức tạp.

#### Bước 1: Khai báo Endpoint
Mô tả API trực tiếp trong hàm `Register` của Controller:
```go
func (this *HelloController) Register(rg *routerx.Routerx) {
    rg.POST(dto.OpenEndpoint{
        Path:        "/create",
        Handler:     this.Create,
        Summary:     "Tạo mới dữ liệu",
        Request:     &dto.CreatePostRequest{},
        Responses:   map[int]any{
            200: gin.H{"status": "success"},
        },
    })
}
```

#### Bước 2: Kích hoạt trong Metadata
Đảm bảo `EnableOpenAPI: true` trong biến metadata của Controller.

#### Bước 3: Kiểm tra
Chạy server và truy cập: `http://localhost:8080/swagger/`

## 🟦 Bắt đầu

### Cài đặt
```bash
git clone https://github.com/HoangHuy7/gin-boilerplate.git
cd gin-boilerplate
go mod download
```

### Chạy Service
```bash
go run cmd/iam/main.go
```

---
<p align="center">Code with ❤️ by <b>HoangHuy7</b></p>
