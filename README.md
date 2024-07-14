# Struct

- Cấu trúc struct được lưu trong bộ nhớ RAM, mỗi khi update dữ liệu thì nó sẽ sao chép dữ liệu của địa chỉ trước và nhận các sự thay đổi gán địa chỉ mới

```
namPointer := &nam // Lấy địa chỉ của biến nam
namPointer := nam // sao chep va địa chỉ mới

False
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

True
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
```

- Cần dùng Point Types
  - int
  - float
  - string
  - bool
  - structs
- Còn lại không cần !!
