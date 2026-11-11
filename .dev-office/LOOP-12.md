# LOOP-12 · ฝึก Go: slice กลาง 3 ตัว (ข้อ 27)

- สถานะ: done · ทีม: loop · ผู้รับผิดชอบ: Prim
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- ไฟล์: practice/bas/027/middle.go, practice/bas/027/middle_test.go
- เทสต์: `go test ./practice/bas/027/`
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:40

## งาน

โจทย์ intermediate_exercices/027. func Middle3(s []int) []int คืน 3 ตัวกลาง; len<3 คืน s เดิม · เทสต์ len 5, 6, 2 · เขียนเองจาก README/main.go ของข้อนั้นก่อน แล้วค่อยเทียบกับ solution/ ตอนรีวิว · ไม่แก้ไฟล์โจทย์เดิม · package practice

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Bas
- 11-11 18:29 Jay [meeting] kickoff: Jay: สร้าง Middle3 ใน practice/bas/027 ให้คืน 3 ตัวกลาง (len<3 คืน slice เดิม) พร้อมเทสต์ และรัน go test ./practice/bas/027/ ให้ผ่าน · Nut: ok · Bas: ok · Prim: ok
- 11-11 18:36 Bas [commit] สร้าง practice/bas/027/middle.go และ middle_test.go ครอบกรณี len 5, 6, 2 แล้ว ส่งให้ระบบรัน go test ./practice/bas/027/ (085b1b36bd88cf7453810c432cbf7c7d46cf4083)
- 11-11 18:36 Bas [handoff] Bas → Prim
- 11-11 18:37 Prim [round] รอบที่ 1 ไม่ผ่าน: no file in your work. work holds a change to a file like this (SEARCH = lines copied exactly from the file): === EDIT: path/to/file.go <<<<<<< SEARCH <old lines> ======= <new lines> >>>>>>> REPLACE ==
- 11-11 18:40 Prim [commit] รีวิว middle.go กับ middle_test.go แล้ว ครอบ len 5/6/2 ตรงโจทย์ · Middle3 ไม่เขียนทับ slice เดิม และเพิ่มเทสต์ยืนยัน · go test ./practice/bas/027/ ผ่าน (4d05d14090132807466b07f57cececa1c1c17aec)
- 11-11 18:40 Jay [meeting] review: Jay: สรุปพร้อมให้เจ้าของรีวิว · Bas: สร้าง practice/bas/027/middle.go และ middle_test.go ครอบกรณี len 5, 6, 2 แล้ว ส่งให้ระบบรัน go test ./practice/bas/027/ · Prim: รีวิว middle.go กับ middle_test.go แล้ว ครอบ len 5/6/2 ตรงโจทย์ · Middle3 ไม่เขียนทับ slice เดิม และเพิ่มเทสต์ยืนยัน · go test ./practice/bas/027/ ผ่าน
