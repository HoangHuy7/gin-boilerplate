# Yêu cầu nghiệp vụ – Phần mềm quản lý vật tư nông nghiệp

## 1. Mục tiêu
- Quản lý bán hàng vật tư nông nghiệp cho cửa hàng nhỏ, thao tác nhanh và đơn giản.
- Ưu tiên các luồng POS (quét barcode → tính tiền → in hóa đơn).
- Theo dõi tồn kho và công nợ rõ ràng.

## 2. Phạm vi
Áp dụng cho một cửa hàng, một kho trong giai đoạn đầu. Hỗ trợ đa tenant theo cấu hình `tenancies`.

## 3. Đối tượng sử dụng
- **Chủ cửa hàng**: xem báo cáo, quản trị dữ liệu.
- **Nhân viên bán hàng**: tạo đơn, cập nhật giao hàng, ghi nợ.

## 4. Chức năng nghiệp vụ chính

### 4.1. Sản phẩm
- Tạo/sửa/xóa sản phẩm.
- Mỗi sản phẩm có barcode duy nhất (có thể nhập tay).
- Quản lý: tên, loại, đơn vị, giá bán, giá nhập, tồn kho, hình ảnh.

### 4.2. Nhập kho và điều chỉnh tồn
- Ghi nhận nhập hàng theo sản phẩm và số lượng.
- Điều chỉnh tồn (tăng/giảm) có ghi chú.
- Tự động ghi log tồn kho theo từng giao dịch.

### 4.3. Bán hàng (POS)
- Quét barcode để thêm sản phẩm vào giỏ.
- Cho phép sửa số lượng, giá bán từng dòng.
- Trả đủ hoặc ghi nợ cho khách.
- Tự động trừ tồn kho sau khi tạo đơn.
- In hóa đơn khổ 80mm (tối giản, dễ đọc).

### 4.4. Khách hàng
- Tạo/sửa/xóa khách.
- Lưu thông tin cơ bản: tên, điện thoại, địa chỉ.
- Xem lịch sử mua hàng và tổng nợ hiện tại.

### 4.5. Công nợ
- Hệ thống tự tính nợ từ đơn hàng.
- Ghi nhận giao dịch trả nợ (trả một phần hoặc toàn bộ).
- Theo dõi tổng nợ theo khách hàng.

### 4.6. Giao hàng trong ngày
- Tạo lịch giao theo đơn.
- Danh sách “Hôm nay cần giao”.
- Cập nhật trạng thái giao (pending/done).

### 4.7. Báo cáo
- Doanh thu theo ngày/tháng/năm.
- Thống kê số lượng bán.
- Theo dõi hàng tồn còn lại.

## 5. Quy tắc nghiệp vụ
- Không cho phép tồn kho âm khi bán hàng.
- Barcode là duy nhất trong hệ thống.
- Tổng tiền đơn = tổng dòng sản phẩm.
- Nợ phát sinh = tổng đơn - số tiền đã trả.

## 6. Yêu cầu phi chức năng
- Thời gian phản hồi API nhanh (POS là luồng chính).
- Ổn định khi dùng liên tục trong giờ cao điểm.
- Dữ liệu rõ ràng, dễ truy vết (log hành động).

## 7. Ngoài phạm vi (giai đoạn đầu)
- Nhiều kho, nhiều chi nhánh.
- Quản lý nhà cung cấp, nhập hàng theo hóa đơn nhà cung cấp.
- Tích hợp kế toán hoặc hóa đơn điện tử.
