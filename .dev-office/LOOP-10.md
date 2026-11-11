# LOOP-10 · ฝึก Go: เลขคี่ (ข้อ 15)

- สถานะ: done · ทีม: loop · ผู้รับผิดชอบ: Bas
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- ไฟล์: practice/bas/015/odd.go, practice/bas/015/odd_test.go
- เทสต์: `go test ./practice/bas/015/`
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:13

## งาน

โจทย์ basic_exercices/015. func IsOdd(n int) bool · เทสต์ 0, 1, 2, -3 (ระวังเลขลบกับ %) · เขียนเองจาก README/main.go ของข้อนั้นก่อน แล้วค่อยเทียบกับ solution/ ตอนรีวิว · ไม่แก้ไฟล์โจทย์เดิม · package practice

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Bas
- 11-11 18:13 Jay [meeting] kickoff: Jay: สร้างฟังก์ชัน IsOdd ใน package practice ให้เทสต์ 0, 1, 2, -3 ผ่านด้วย go test ./practice/bas/015/ · Nut: ok · Bas: ok
- 11-11 18:13 Bas [commit] สร้าง odd.go: func IsOdd(n int) bool { return n%2 != 0 } ทำให้เลขลบ -3 คืน true ถูกต้อง และ odd_test.go เทสต์ 0, 1, 2, -3 ใน package practice (08d5a9c0b0cbbd907d67ae613f478d01279f2d5b)
- 11-11 18:13 Jay [meeting] review: Jay: สรุปพร้อมให้เจ้าของรีวิว · Bas: สร้าง odd.go: func IsOdd(n int) bool { return n%2 != 0 } ทำให้เลขลบ -3 คืน true ถูกต้อง และ odd_test.go เทสต์ 0, 1, 2, -3 ใน package practice
