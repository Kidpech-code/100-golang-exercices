# LOOP-20 · ฝึก Go: HTTP handler ตอบตาม path (ข้อ 73)

- สถานะ: done · ทีม: loop · ผู้รับผิดชอบ: Bas
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- ไฟล์: practice/bas/073/echo.go, practice/bas/073/echo_test.go
- เทสต์: `go test ./practice/bas/073/`
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:19

## งาน

โจทย์ http_apis/073. func Hello(w http.ResponseWriter, r *http.Request) เขียน "Hello, "+r.URL.Path ห้าม hardcode · เทสต์ด้วย httptest 2 path ไม่เปิด port จริง · เขียนเองจาก README/main.go ของข้อนั้นก่อน แล้วค่อยเทียบกับ solution/ ตอนรีวิว · ไม่แก้ไฟล์โจทย์เดิม · package practice

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Bas
- 11-11 18:03 Jay [meeting] kickoff: Jay: สร้าง HTTP handler ใน practice/bas/073 ตอบ "Hello, " + path ด้วย httptest 2 path โดยไม่เปิด port จริง · Nut: ok · Bas: ok
- 11-11 18:19 Bas [commit] สร้าง echo.go กับ echo_test.go สำหรับ practice/bas/073 (a88a740f5d028660e414ea71894a4f99f69733dd)
- 11-11 18:19 Jay [meeting] review: Jay: สรุปพร้อมให้เจ้าของรีวิว · Bas: สร้าง echo.go กับ echo_test.go สำหรับ practice/bas/073
