# 🏗️ gin-boilerplate

> **Gin boilerplate "xịn sò" cho anh em dev, kiến trúc modular, monorepo chuẩn chỉ, scale thoải mái!**

Chào anh em! 👋 Đây là bộ **startkit monorepo** tâm huyết mình build để anh em đỡ phải setup lại từ đầu mỗi khi làm dự án Go backend. Nói không với code "mì ăn liền", repo này hướng tới style clean, gọn gàng, chia module rõ ràng để team đông người vẫn code vui vẻ không dẫm chân nhau.

## 🌟 Có gì hot? (Highlights)

- **🧩 Modular Architecture**: Chia domain (`iam`, `device`, `notification`) rành mạch. Mỗi ông là một **Server Con** (Child Server) riêng biệt, như kiểu các module trong Maven á. Đỡ phải lo conflict code lung tung.
- **🏗️ Monorepo Structure**: Thư mục **`internal`** là "trái tim" (Core/Shared Library) của cả hệ thống. Nó chứa logic dùng chung, DTO, router base... giống như cái Maven common/parent mà anh em hay dùng bên Java ấy.
- **🛡️ Production Ready**: Đã lắp sẵn đồ chơi: logging, routing xịn (`routerx`), DTO chuẩn bài. Anh em chỉ việc clone về là chiến logic nghiệp vụ luôn.
- **🔌 Scalable & Extensible**: Build trên nền **[Gin](https://github.com/gin-gonic/gin)** (thánh tốc độ), bao cân từ dự án MVP bé tẹo đến hệ thống triệu view.

## 📂 Soi "nội thất" (Project Structure)

```text
.
├── cmd
│   ├── device          # Cửa chính (Entry point) cho ông Device Server
│   │   └── main.go
│   ├── iam             # Cửa chính cho ông IAM Server
│   │   └── main.go
│   └── notification    # Cửa chính cho ông Notification Server
│       └── main.go
├── device              # Logic/Nghiệp vụ của Device (Server con)
├── iam                 # Logic/Nghiệp vụ của IAM (Server con)
│   └── controller
│       ├── Module.go
│       ├── Router.go
│       └── v1
│           └── HelloController.go
├── internal            # 🧱 Hàng dùng chung (Core / Shared Libs) - Đụng vào đây nhớ cẩn thận nha bro
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
├── notification        # Logic/Nghiệp vụ của Notification (Server con)
│   └── controller
│       └── v1
├── go.mod
├── go.sum
└── main.go
```

## 🛠️ Chiến thôi! (Getting Started)

### Cần gì?

- **Go** (bản 1.20 trở lên nha anh em)

### Cài đặt

Kéo hàng về máy:

```bash
git clone https://github.com/HoangHuy7/gin-boilerplate.git
cd gin-boilerplate
go mod download
```

### Lên nhạc (Run Service)

Mỗi module (domain) có file chạy riêng trong folder `cmd`. Ví dụ anh em muốn chạy con **IAM** lên để test:

```bash
go run cmd/iam/main.go
```

## 🤝 Góp gạch xây nhà (Contribution)

Anh em thấy gì hay ho hoặc chỗ nào chuối cứ tự nhiên PR nha! Chỉ cần nhớ quy tắc là giữ cho folder `internal` sạch đẹp, gọn gàng để cả làng dùng chung là được.

---

Code with ❤️ by **HoangHuy7**
