# 🏗️ gin-boilerplate

> **Gin boilerplate "xịn sò" cho anh em, kiến trúc modular, monorepo chuẩn chỉ, không phải cái đống spaghetti code mà người cũ để lại đâu!**

Chào đồng, lại là tôi đây! 👋 Đây là cái bộ **startkit monorepo** tôi làm ra để anh em đỡ phải ngồi setup lại từ đầu mỗi khi "vẽ vời" dự án mới. Nói thật, code cái này là để anh em bớt tạo nghiệp với mấy con monolith to như cái nhà mà sửa một chỗ chết chục chỗ nhé.

## 🌟 Có gì mà khoe? (Highlights)

- **🧩 Modular Architecture**: Chia domain (`iam`, `device`, `notification`) ra đàng hoàng. Mỗi ông một **Server Con** (Child Server) riêng biệt, thằng nào chết thằng ấy tự chịu, không kéo cả lò chết chùm. Kiểu module trong Maven ấy, chắc bro biết rồi (chưa biết thì search Google đi).
- **🏗️ Monorepo Structure**: Cái folder **`internal`** kia là "bảo vật trấn môn" (Core/Shared Library). Logic dùng chung, DTO, router base... nhét hết vào đấy. Nó giống cái Maven parent mà mấy ông Java hay thần thánh hóa ấy.
- **🔐 Authorization (Casbin)**: Đã tích hợp **Casbin** để phân quyền (RBAC). Đừng để user thường vào xoá database là được.
- **⚙️ Config Xịn Xò**: Quản lý config tập trung trong folder `configs`. Đừng hardcode password trong code nữa nhé, quê lắm.
- **📜 Swagger "Tự Động Hóa"**: Tôi gắn sẵn `swaggest` rồi, viết code xong là có document Swagger luôn. Khỏi phải ngồi hì hục viết doc bằng cơm ("chạy bằng cơm") nữa nhé, thời gian đấy để đi chơi với người yêu.
- **🛡️ Production Ready**: Tôi đã gắn sẵn logging, routing xịn (`routerx`), DTO chuẩn cơm mẹ nấu rồi. Anh em chỉ việc clone về, đắp logic nghiệp vụ vào rồi đi nhậu thôi.
- **🔌 Scalable & Extensible**: Chạy bằng **[Gin](https://github.com/gin-gonic/gin)** (nhanh vãi linh hồn), cân được từ cái MVP "làm cho vui" đến hệ thống triệu view (nếu bro đủ trình marketing).

## 📂 Soi "nội thất" (Project Structure)

Nhìn cho kĩ cái cây này, project giờ xịn hơn rồi, nhiều folder hơn nên đừng lác mắt:

```text
.
├── apps                    # 🏢 Khu tập thể cho các Server Con
│   ├── device              # Logic Device
│   ├── iam                 # Logic IAM - Đừng để lộ password là được
│   │   ├── app             # 🔌 Dây điện chằng chịt (Config, DB, Auth)
│   │   │   ├── casbin      # Phân quyền Casbin (Bảo kê)
│   │   │   ├── config      # Load file config
│   │   │   └── database    # Kết nối DB
│   │   └── controller      # Controller nhận request
│   │       └── v1
│   │           └── HelloController.go
│   └── notification        # Logic Notification
├── cmd
│   ├── device              # Cổng vào cho ông Device Server
│   │   └── main.go
│   ├── iam                 # Cổng vào cho ông IAM Server
│   │   └── main.go
│   └── notification        # Cổng vào cho Notification Server
│       └── main.go
├── configs                 # ⚙️ Chỗ để file cấu hình (YAML, Policy)
│   └── iam
│       ├── application.yaml
│       └── casbin
├── internal                # 🧱 Hàng dùng chung (Core) - Cấm táy máy lung tung
│   ├── base
│   ├── dto
│   ├── logger
│   ├── server
│   └── utils               # 🛠️ Đồ nghề lặt vặt (Utils)
├── go.mod
├── go.sum
└── main.go                 # File này để ngắm thôi
```

## 🛠️ Chiến thôi! (Getting Started)

### Cần gì?

- **Go** (bản 1.20+ nha, đừng dùng bản thời tống nữa plzz)

### Cài đặt

Copy paste dòng này vào terminal này (đừng bảo không biết mở terminal nhé):

```bash
git clone https://github.com/HoangHuy7/gin-boilerplate.git
cd gin-boilerplate
go mod download
```

### Lên nhạc (Run Service)

Muốn chạy con nào thì vào `cmd` gọi con đấy dậy. Ví dụ muốn test **IAM** xem login được chưa:

```bash
go run cmd/iam/main.go
```

### 📚 Tài liệu API (Swagger)

Chạy server lên xong thì vào đường link này mà ngắm API, đừng hỏi tôi API có những gì:
- **Link**: `http://localhost:8080/swagger/` (Cổng 8080 hay bao nhiêu tùy bro config nhé)


## 🤝 Góp gạch xây nhà (Contribution)

Anh em thấy tôi code "ngáo" chỗ nào hoặc muốn show trình thì cứ PR mạnh tay vào! Chỉ xin một điều: **đừng làm nát cái folder `internal`** của tôi là được, chỗ đó là vùng cấm bay, sửa phải có não nha bro.

---

Code with ❤️ (and a bit of ☕) by **HoangHuy7**
