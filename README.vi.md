# 🏗️ gin-boilerplate

> **Gin boilerplate chuẩn Production với kiến trúc module, hỗ trợ monorepo và các best practices cho hệ thống backend quy mô lớn.**

Repo này là một **startkit monorepo** mạnh mẽ, được thiết kế để khởi động nhanh quá trình phát triển Go backend của bạn. Dự án giúp tổ chức source code rành mạch, tránh sự hỗn độn của kiến trúc monolithic truyền thống, hướng tới thiết kế domain-driven rõ ràng, dễ dàng mở rộng theo sự phát triển của team và sản phẩm.

## 🌟 Điểm Nổi Bật

- **🧩 Kiến Trúc Modular**: Các domain (`iam`, `device`, `notification`) hoạt động như các **Server Con** (Child Servers/Microservices), tương tự như các module trong Maven.
- **🏗️ Cấu Trúc Monorepo**: Thư mục **`internal`** đóng vai trò là **Core / Shared Library** (giống như Maven common/parent), chứa các logic nền tảng, DTO và cấu hình dùng chung cho tất cả các server con.
- **🛡️ Sẵn Sàng Cho Production**: Cấu hình sẵn logging, chiến lược routing (`routerx`), và các mô hình dữ liệu (DTO) chuẩn.
- **🔌 Khả Năng Mở Rộng**: Xây dựng trên nền tảng [Gin](https://github.com/gin-gonic/gin), sẵn sàng phát triển từ MVP startup đến hệ thống chịu tải cao.

## 📂 Cấu Trúc Dự Án

```text
.
├── cmd
│   ├── device          # Entry point cho Server Con Device
│   │   └── main.go
│   ├── iam             # Entry point cho Server Con IAM
│   │   └── main.go
│   └── notification    # Entry point cho Server Con Notification
│       └── main.go
├── device              # Device Service (Server Con)
├── iam                 # IAM Service (Server Con)
│   └── controller
│       ├── Module.go
│       ├── Router.go
│       └── v1
│           └── HelloController.go
├── internal            # 🧱 Core / Thư viện dùng chung (Tương tự Maven Common)
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
├── notification        # Notification Service (Server Con)
│   └── controller
│       └── v1
├── go.mod
├── go.sum
└── main.go
```

## 🛠️ Hướng Dẫn Bắt Đầu

### Yêu cầu

- **Go** (1.20 trở lên)

### Cài đặt

Clone repository:

```bash
git clone https://github.com/HoangHuy7/gin-boilerplate.git
cd gin-boilerplate
go mod download
```

### Chạy Service

Mỗi domain (module) có entry point riêng nằm trong thư mục `cmd`. Ví dụ để chạy **IAM Service**:

```bash
go run cmd/iam/main.go
```

## 🤝 Đóng Góp

Mọi sự đóng góp đều được hoan nghênh! Hãy đảm bảo code trong thư mục `internal` luôn gọn gàng và có tính tái sử dụng cao.

---

Phát triển bởi **HoangHuy7**
