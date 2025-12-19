# 🏗️ gin-boilerplate

> **Gin boilerplate "xịn sò" cho anh em, kiến trúc modular, monorepo chuẩn chỉ kèm automation bằng Uber-fx.**

Chào đồng bào! 👋 Đây là cái bộ **startkit monorepo** tôi làm ra để anh em đỡ phải ngồi setup lại từ đầu mỗi khi "vẽ vời" dự án mới. Code cái này là để anh em bớt tạo nghiệp với mấy con monolith to như cái nhà mà sửa một chỗ chết chục chỗ nhé.

## 🌟 Có gì mà khoe? (Highlights)

- **🧩 Modular Architecture**: Chia domain (`iam`, `device`, `notification`) ra đàng hoàng. Mỗi ông một module riêng biệt, tách bạch logic.
- **🏗️ Monorepo Structure**: Folder **`internal`** là "bảo vật trấn môn" (Core/Shared Library). Logic dùng chung, DTO, router base... nhét hết vào đấy.
- **⚡ Dependency Injection**: Sử dụng **Uber-fx** để quản lý lifecycle và tự động hóa việc kết nối các component.
- **🤖 Controller Tự Động Hóa**: Không cần phải khai báo router thủ công cho từng controller. Chỉ cần ném vào module là nó tự chạy. 
- **� Authorization (Casbin)**: Đã tích hợp **Casbin** để phân quyền (RBAC/ABAC) chuẩn chỉ.
- **� Swagger "Tự Động Hóa"**: Sử dụng `swaggest/openapi-go` để gen Swagger từ code. Viết xong là có doc luôn, không phải "chạy bằng cơm".

## 📂 Soi "nội thất" (Project Structure)

```text
.
├── apps                    # 🏢 Module nghiệp vụ / Domain Logic
│   ├── device              # Logic Device
│   ├── iam                 # Logic IAM (Identity & Access)
│   │   ├── app             # App wiring (Config, DB, Auth)
│   │   └── controller      # Controller nhận request
│   └── notification        # Logic Notification
├── cmd                     # 🚀 File thực thi (Entry Points)
│   ├── device/main.go
│   ├── iam/main.go
│   └── notification/main.go
├── configs                 # ⚙️ Cấu hình (YAML, Policy)
├── internal                # 🧱 Hàng dùng chung (Core) - Cấm táy máy lung tung
│   ├── base                # Interface gốc (Controller, ...)
│   ├── logger              # Zap Logger xịn sò
│   ├── server              # Core HTTP Server & Router logic
│   └── utils               # Đồ nghề lặt vặt
├── go.mod
└── main.go                 # File này để ngắm thôi
```

## 🤖 Cách Automation hoạt động

Cái project này tận dụng [Uber-fx](https://github.com/uber-go/fx) để tự động hóa việc đăng ký Controller mà không cần code tay từng dòng router.

### 1. Tại Core Router (`internal/server`)
Hàm `NewRouter` trong `internal/server/router.go` được thiết kế để nhận một list các controller qua DI:
```go
func NewRouter(controllers []base.Controller, ...) *Router
```

### 2. Tự động đăng ký (`apps/iam/controller`)
Trong file `Module.go` của từng module (ví dụ `apps/iam/controller/Module.go`), chúng ta sử dụng **Group Tags**:
```go
fx.Annotate(
    v1.NewHelloController,
    fx.As(new(base.Controller)),
    fx.ResultTags(`group:"controllers"`), // Gom vào tập đoàn "controllers"
)
```
Sau đó inject cả tập đoàn này vào `NewRouter`:
```go
fx.Annotate(
    server.NewRouter,
    fx.ParamTags(`group:"controllers"`), // Gọi cả hội controller ra
)
```

### 3. Kích hoạt (`cmd/iam`)
Trong file `main.go`, chỉ cần gọi cái module controller đó ra là xong:
```go
fx.New(
    app.Module,
    controller.Module, // Phép thuật nằm ở đây
    // ...
    fx.Invoke(server.RunServer),
).Run()
```

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

### 📚 Tài liệu API (Swagger)
Chạy server lên xong thì vào: `http://localhost:8080/swagger/` (Cổng tùy theo config nhé).

## 🤝 Góp gạch xây nhà (Contribution)
Anh em nhớ giữ cái folder `internal` sạch sẽ. Thêm cái gì mới thì nhớ check xem có dùng chung được cho các module khác không nhé.

---
Code with ❤️ by **HoangHuy7**
