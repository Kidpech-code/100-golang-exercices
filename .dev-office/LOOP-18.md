# LOOP-18 · ฝึก Go: JSON ไม่รู้โครงสร้าง (ข้อ 53)

- สถานะ: done · ทีม: loop · ผู้รับผิดชอบ: Bas
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- ไฟล์: practice/bas/053/birds.go, practice/bas/053/birds_test.go
- เทสต์: `go test ./practice/bas/053/`
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:18

## งาน

โจทย์ json/053. func Birds(b []byte) (map[string]any, error) อ่าน key "birds" ด้วย type assertion · เทสต์ ไม่มี key/ชนิดผิด ต้องคืน error ไม่ panic · เขียนเองจาก README/main.go ของข้อนั้นก่อน แล้วค่อยเทียบกับ solution/ ตอนรีวิว · ไม่แก้ไฟล์โจทย์เดิม · package practice

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Bas
- 11-11 18:02 Jay [meeting] kickoff: Jay: สร้าง func Birds(b []byte) (map[string]any, error) อ่าน key "birds" ด้วย type assertion และเทสต์ครอบกรณีไม่มี key/ชนิดผิดต้องคืน error ไม่ panic · Nut: ok · Bas: ok
- 11-11 18:18 Bas [commit] สร้าง birds.go อ่าน JSON แล้วใช้ type assertion ดึง key birds เป็น map[string]any; คืน error เมื่อ key หาย/ชนิดผิด/JSON ไม่ valid พร้อม birds_test.go ครอบเคส valid, no key, wrong type, invalid JSON (51fe2767376bbe3189492d5aa231ba56515d0396)
- 11-11 18:18 Jay [meeting] review: Jay: สรุปพร้อมให้เจ้าของรีวิว · Bas: สร้าง birds.go อ่าน JSON แล้วใช้ type assertion ดึง key birds เป็น map[string]any; คืน error เมื่อ key หาย/ชนิดผิด/JSON ไม่ valid พร้อม birds_test.go ครอบเคส valid, no key, wrong type, invalid JSON
