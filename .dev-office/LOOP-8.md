# LOOP-8 · ฝึก Go: ต่อสตริง (ข้อ 4)

- สถานะ: done · ทีม: loop · ผู้รับผิดชอบ: Bas
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- ไฟล์: practice/bas/004/concat.go, practice/bas/004/concat_test.go
- เทสต์: `go test ./practice/bas/004/`
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:27

## งาน

โจทย์ basic_exercices/004. func Concat(a, b string) string คืน a+" "+b · เทสต์ table-driven 3 เคส รวมสตริงว่าง · เขียนเองจาก README/main.go ของข้อนั้นก่อน แล้วค่อยเทียบกับ solution/ ตอนรีวิว · ไม่แก้ไฟล์โจทย์เดิม · package practice

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Bas
- 11-11 18:13 Jay [meeting] kickoff: Jay: สร้าง practice/bas/004 ให้มีฟังก์ชัน Concat(a,b string) string คืน a+" "+b พร้อมเทสต์ table-driven 3 เคสรวมสตริงว่าง และ go test ./practice/bas/004/ ผ่าน · Nut: ok · Bas: ok
- 11-11 18:18 Bas [working] ขอดู basic_exercices/004/README.md กับ basic_exercices/004/main.go ก่อนตามการ์ด
- 11-11 18:27 Bas [commit] อ่าน README กับ main.go ของข้อ 004 แล้วสร้าง practice/bas/004/concat.go และ concat_test.go (table-driven 3 เคสรวมสตริงว่าง) (b5eb1dc987ac32acee63536ddc27c4bef3a41495)
- 11-11 18:27 Jay [meeting] review: Jay: สรุปพร้อมให้เจ้าของรีวิว · Bas: อ่าน README กับ main.go ของข้อ 004 แล้วสร้าง practice/bas/004/concat.go และ concat_test.go (table-driven 3 เคสรวมสตริงว่าง)
