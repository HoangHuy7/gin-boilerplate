# Dev plan & checklist – Backend quản lý vật tư

## Mục tiêu hiện tại
Hoàn thiện backend theo tài liệu `spec/yeu-cau-nghiep-vu.md`, ưu tiên các luồng nghiệp vụ lõi cho vận hành cửa hàng: POS, tồn kho, công nợ, giao hàng, báo cáo.

## Kế hoạch triển khai
1. **Phase 1 (MVP lõi):** Sản phẩm + bán hàng + tồn kho + công nợ.
2. **Phase 2:** Giao hàng trong ngày + hoàn thiện truy vấn/pagination + ổn định API.
3. **Phase 3:** Báo cáo nâng cao + hardening + tài liệu vận hành.

## Checklist tiến độ

### A. Nghiệp vụ lõi
- [x] Bỏ module Menu chưa hoàn thiện (tránh panic runtime).
- [x] Chuẩn hóa GraphQL sau khi bỏ Menu (regen schema/resolver).
- [x] Tạo tài liệu yêu cầu nghiệp vụ chính thức (`spec/yeu-cau-nghiep-vu.md`).
- [x] Bổ sung `application.example.yaml` để dễ setup môi trường.

### B. Tồn kho
- [x] `createInventoryLog` cập nhật tồn kho sản phẩm theo loại giao dịch.
- [x] Chặn tồn kho âm khi xuất/điều chỉnh.
- [x] Chuẩn hóa loại giao dịch kho (`import`, `sale`, `adjust`).
- [x] Validate dữ liệu đầu vào (type/quantity/product_id).

### C. Đơn hàng & công nợ
- [x] Chặn tạo đơn khi `paid_amount > total_amount`.
- [x] Chặn quantity không hợp lệ trong order item.
- [x] Tạo đơn có phát sinh nợ sẽ cộng `customers.total_debt`.
- [x] Tạo transaction công nợ `borrow` khi đơn phát sinh nợ.
- [x] Xóa đơn sẽ hoàn kho và điều chỉnh lại công nợ khách hàng.

### D. Debt transaction
- [x] `createDebtTransaction` cập nhật `customers.total_debt` theo `borrow/pay`.
- [x] Chặn overpay gây công nợ âm.
- [x] `deleteDebtTransaction` rollback đúng công nợ.
- [x] Validate ID đầu vào tại resolver để tránh panic.

### E. Việc còn lại (next)
- [ ] Thêm pagination/filter cho `inventoryLogs` và `debtTransactions`.
- [ ] Bổ sung test integration cho luồng tạo/xóa đơn + công nợ + tồn kho.
- [ ] Bổ sung endpoint/integration cho in hóa đơn 80mm (nội dung hóa đơn).
- [ ] Tài liệu hóa danh sách mutation/query ưu tiên cho frontend POS.
