# LOOP-6 · แก้เฉลยข้อ 62 ที่ build ไม่ผ่าน

- สถานะ: done · ทีม: loop · ผู้รับผิดชอบ: Prim
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- ไฟล์: interfaces/062/solution/solution.go
- เทสต์: `go vet ./interfaces/062/solution/`
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:29

## งาน

interfaces/062/solution/solution.go: main() เรียก context.Background() กับ doSomething() ที่ไม่มีในไฟล์ (คัดลอกมาจากบท context) โจทย์บอกว่าข้อนี้ไม่ต้องรันอะไร ให้ main ว่าง และเอา import ที่ไม่ใช้ออก

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Bas
- 11-11 18:12 Jay [meeting] kickoff: Jay: แก้ interfaces/062/solution/solution.go ให้ build ผ่าน โดยให้ main ว่างและลบ import ที่ไม่ใช้ · Nut: ok · Bas: ok · Prim: ok
- 11-11 18:17 Bas [commit] ลบ import fmt/math และทำให้ main ว่างใน interfaces/062/solution/solution.go ตามการ์ด (ab27b7ae67c81f7948ed24601457e647e9eec310)
- 11-11 18:17 Bas [handoff] Bas → Prim
- 11-11 18:27 Prim [round] รอบที่ 1 ไม่ผ่าน: no file in your work. work holds a change to a file like this (SEARCH = lines copied exactly from the file): === EDIT: path/to/file.go <<<<<<< SEARCH <old lines> ======= <new lines> >>>>>>> REPLACE ==
- 11-11 18:29 Prim [commit] รัน go vet ./interfaces/062/solution/ ผ่าน: main ว่าง ไม่มี import เกิน และไม่มี context.Background/doSomething ใน diff (0f85a4c93a4e5d72c703bf7bf5f3aa8438a865b3)
- 11-11 18:29 Jay [meeting] review: Jay: สรุปพร้อมให้เจ้าของรีวิว · Bas: ลบ import fmt/math และทำให้ main ว่างใน interfaces/062/solution/solution.go ตามการ์ด · Prim: รัน go vet ./interfaces/062/solution/ ผ่าน: main ว่าง ไม่มี import เกิน และไม่มี context.Background/doSomething ใน diff
