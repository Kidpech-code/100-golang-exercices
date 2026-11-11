# LOOP-19 · ฝึก Go: type switch (ข้อ 68)

- สถานะ: done · ทีม: loop · ผู้รับผิดชอบ: Prim
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- ไฟล์: practice/bas/068/describe.go, practice/bas/068/describe_test.go
- เทสต์: `go test ./practice/bas/068/`
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:33

## งาน

โจทย์ interfaces/068. func Describe(v any) string คืน "int", "string", "bool", หรือ "unknown" · เทสต์ 4 เคส รวม nil · เขียนเองจาก README/main.go ของข้อนั้นก่อน แล้วค่อยเทียบกับ solution/ ตอนรีวิว · ไม่แก้ไฟล์โจทย์เดิม · package practice

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Bas
- 11-11 18:01 Jay [meeting] kickoff: Jay: สร้าง Describe(v any) string ใน practice/bas/068 ให้คืน int/string/bool/unknown และมีเทสต์ 4 เคสรวม nil · Nut: ok · Bas: ok · Prim: ok
- 11-11 18:02 Bas [commit] สร้าง describe.go มี func Describe(v any) string ใช้ type switch คืน int/string/bool/unknown และ describe_test.go ครบ 4 เคสรวม nil (2d7c73c0a38bbabfa7da45714716f112bb3e37fe)
- 11-11 18:02 Bas [handoff] Bas → Prim
- 11-11 18:16 Prim [working] ขออ่าน practice/bas/068/solution/solution.go เพื่อเทียบกับโค้ดที่เขียนไว้
- 11-11 18:17 Prim [working] ไม่พบ practice/bas/068/solution/solution.go ขออ่าน practice/bas/068/README.md และ practice/bas/068/main.go เพื่อเทียบกับโจทย์ก่อนรีวิว
- 11-11 18:25 Prim [round] รอบที่ 1 ไม่ผ่าน: no file in your work. work holds a change to a file like this (SEARCH = lines copied exactly from the file): === EDIT: path/to/file.go <<<<<<< SEARCH <old lines> ======= <new lines> >>>>>>> REPLACE ==
- 11-11 18:33 Prim [commit] เทียบ solution/068 แล้วไม่มีไฟล์ให้เทียบ; describe.go ตรง spec และ diff ไม่แตะไฟล์โจทย์เดิม · เพิ่มคอมเมนต์รีวิว + commit dev/LOOP-19 · ลิงก์เปิด PR: https://github.com/Kidpech-code/100-golang-exercices/pull/new/dev/LOOP-19 `go test ./practice/bas/068/` (de2cea3efa5c49e31160e7c0ce35b64762a479d7)
- 11-11 18:33 Jay [meeting] review: Jay: สรุปพร้อมให้เจ้าของรีวิว · Bas: สร้าง describe.go มี func Describe(v any) string ใช้ type switch คืน int/string/bool/unknown และ describe_test.go ครบ 4 เคสรวม nil · Prim: เทียบ solution/068 แล้วไม่มีไฟล์ให้เทียบ; describe.go ตรง spec และ diff ไม่แตะไฟล์โจทย์เดิม · เพิ่มคอมเมนต์รีวิว + commit dev/LOOP-19 · ลิงก์เปิด PR: https://github.com/Kidpech-code/100-golang-exercices/pull/new/dev/LOOP-19 `go test ./practice/bas/068/`
