# 🏗️ gin-boilerplate

> **Gin boilerplate "xịn sò" cho anh em, kiến trúc modular, monorepo chuẩn chỉ kèm automation bằng Uber-fx.**

Chào đồng bào! 👋 Đây là cái bộ **startkit monorepo** tôi làm ra để anh em khỏi phải lo chuyện setup. Code chuẩn, tách lớp và tự động hóa tận răng.

## 🌟 Có gì mà khoe? (Highlights)

- **🧩 Modular Architecture**: Chia domain (`iam`, `device`, ...) ra đàng hoàng, độc lập và dễ mở rộng.
- **🏗️ Monorepo Structure**: Folder **`internal`** bảo mật và dùng chung logic cho toàn bộ hệ thống.
- **⚡ Dependency Injection**: Sử dụng **Uber-fx** để tự động kết nối (wiring) các component.
- **🤖 Tự động hóa hoàn toàn**: Cả Controller và OpenAPI đều được đăng ký tự động. Code tới đâu, doc tới đó.
- **🔐 Authorization (Casbin)**: Tích hợp sẵn RBAC/ABAC cực mạnh.
- **📜 OpenAPI/Swagger Tự Động**: Không cần viết comment, chỉ cần định nghĩa DTO là có ngay giao diện Swagger đẹp mắt.

## 📂 Soi "nội thất" (Project Structure)

```text
.
├── apps                    # 🏢 Nghiệp vụ chính (Domain Logic)
│   ├── device
│   ├── iam                 # Quản lý định danh
│   │   ├── app             # Đấu nối module (DB, Auth, Config, Module)
│   │   │   ├── casbin
│   │   │   ├── config
│   │   │   ├── database
│   │   │   └── Module.go
│   │   └── controller      # Xử lý HTTP Request
│   │       ├── v1
│   │       │   └── HelloController.go
│   │       └── Module.go   # Nơi đăng ký controller với Fx
│   └── notification
├── cmd                     # 🚀 Cổng vào thực thi
│   ├── iam/main.go
├── configs                 # ⚙️ Cấu hình hệ thống
├── internal                # 🧱 "Trái tim" hệ thống (Shared Core)
│   ├── base                # Interface chung & base controller
│   ├── dto                 # Định nghĩa dữ liệu truyền tải
│   ├── logger              # Zap Logger xịn sò
│   ├── server              # Core Server, Router & logic OpenAPI
│   └── utils               # Đồ nghề hỗ trợ
├── go.mod
└── main.go
```

## 🤖 Hướng dẫn Tự động hóa (Automation)

Project này sử dụng sức mạnh của **Uber-fx** để giải phóng đôi tay của bạn.

### 1. Đăng ký Controller
Bạn không cần gọi `router.GET` ở khắp nơi. Chỉ cần khai báo trong module của folder `controller`:
```go
fx.Annotate(
    v1.NewHelloController,
    fx.As(new(base.Controller)),
    fx.ResultTags(`group:"controllers"`),
)
```

### 2. Tích hợp OpenAPI Tự động (Không dùng Comment)
Quên việc viết `// @Summary` đi, ở đây chúng ta dùng **Code-First** với `routerx`.

#### Bước 1: Khai báo Endpoint trong Controller
Trong hàm `Register`, hãy mô tả API bằng struct `dto.OpenEndpoint`:
```go
func (this *HelloController) Register(rg *routerx.Routerx) {
    rg.POST(dto.OpenEndpoint{
        Path:        "/create",
        Handler:     this.Create,
        Summary:     "Tạo mới gì đó",
        Request:     &dto.CreatePostRequest{}, // Tự gen schema từ struct luôn!
        Responses:   map[int]any{
            200: gin.H{"status": "success"},
        },
    })
}
```

#### Bước 2: Bật OpenAPI trong Metadata
Đảm bảo biến `EnableOpenAPI` là `true` trong metadata của Controller:
```go
Metadata: dto.Metadata{
    Tag:           "IAM Service",
    EnableOpenAPI: true,
}
```

#### Bước 3: Tận hưởng
Chạy server và truy cập:
`http://localhost:8080/swagger/`

## 🛠️ Chiến thôi! (Getting Started)

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
Code with ❤️ by **HoangHuy7**
