# LOOP-11 · ฝึก Go: map (ข้อ 21)

- สถานะ: done · ทีม: loop · ผู้รับผิดชอบ: Bas
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- ไฟล์: practice/bas/021/letters.go, practice/bas/021/letters_test.go
- เทสต์: `go test ./practice/bas/021/`
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:32

## งาน

โจทย์ intermediate_exercices/021. func Letters() map[int]string คืน {1:A,2:B,3:C} · เทสต์ทุก key และ key ที่ไม่มีคืน "" · เขียนเองจาก README/main.go ของข้อนั้นก่อน แล้วค่อยเทียบกับ solution/ ตอนรีวิว · ไม่แก้ไฟล์โจทย์เดิม · package practice

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Bas
- 11-11 18:14 Jay [meeting] kickoff: Jay: สร้าง letters.go ที่คืน map {1:A,2:B,3:C} พร้อมเทสต์ใน practice/bas/021 ให้ go test ผ่าน · Nut: ok · Bas: ok
- 11-11 18:32 Bas [commit] สร้าง practice/bas/021/letters.go คืน map {1:A,2:B,3:C} และ letters_test.go เทสต์ทุก key กับ key ที่ไม่มีค่า (0a72f1615f631c4791329531e89d6e6a8d86ff5b)
- 11-11 18:32 Jay [meeting] review: Jay: สรุปพร้อมให้เจ้าของรีวิว · Bas: สร้าง practice/bas/021/letters.go คืน map {1:A,2:B,3:C} และ letters_test.go เทสต์ทุก key กับ key ที่ไม่มีค่า
