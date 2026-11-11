# LOOP-15 · ฝึก Go: channel ส่งต่อข้อความ (ข้อ 33)

- สถานะ: done · ทีม: loop · ผู้รับผิดชอบ: Prim
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- ไฟล์: practice/bas/033/relay.go, practice/bas/033/relay_test.go
- เทสต์: `go test -race ./practice/bas/033/`
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:43

## งาน

โจทย์ concurrency/033. func Relay(msg string) string: goroutine f1 ส่ง msg เข้า channel, f2 รับแล้วคืน "I am f2 and "+msg · เทสต์ด้วย -race · เขียนเองจาก README/main.go ของข้อนั้นก่อน แล้วค่อยเทียบกับ solution/ ตอนรีวิว · ไม่แก้ไฟล์โจทย์เดิม · package practice

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Bas
- 11-11 18:19 Jay [meeting] kickoff: Jay: ฝึก Go: สร้าง func Relay(msg string) string ใช้ goroutine f1 ส่ง msg เข้า channel แล้ว f2 รับแล้วคืน "I am f2 and "+msg โดยมีเทสต์ผ่าน -race · Nut: ok · Bas: ok · Prim: ok
- 11-11 18:20 Bas [working] ขออ่านไฟล์โจทย์ concurrency/033/README.md และ concurrency/033/main.go ก่อนเขียน
- 11-11 18:24 Bas [commit] สร้าง practice/bas/033/relay.go และ relay_test.go แล้ว Relay ใช้ f1 ส่ง channel และ f2 รับแล้วคืนข้อความ (a545acba79b3e335eaf1da818b75418ac9ed9498)
- 11-11 18:24 Bas [handoff] Bas → Prim
- 11-11 18:32 Prim [working] ขั้นรีวิว: ขออ่าน concurrency/033/README.md, concurrency/033/main.go และ concurrency/033/solution/solution.go เพื่อเทียบกับ practice/bas/033
- 11-11 18:40 Prim [round] รอบที่ 1 ไม่ผ่าน: no file in your work. work holds a change to a file like this (SEARCH = lines copied exactly from the file): === EDIT: path/to/file.go <<<<<<< SEARCH <old lines> ======= <new lines> >>>>>>> REPLACE ==
- 11-11 18:40 Prim [round] รอบที่ 2 ไม่ผ่าน: no file in your work. work holds a change to a file like this (SEARCH = lines copied exactly from the file): === EDIT: path/to/file.go <<<<<<< SEARCH <old lines> ======= <new lines> >>>>>>> REPLACE ==
- 11-11 18:43 Prim [evidence] รีวิว PR dev/LOOP-15 แล้ว: practice/bas/033/relay.go ตรงกับ concurrency/033/README.md, main.go และ solution/solution.go ใช้ channel ให้ f1 ส่ง → f2 รับแล้วคืนข้อความ; ไม่ได้แก้ concurrency/033/*
- 11-11 18:43 Jay [meeting] review: Jay: สรุปพร้อมให้เจ้าของรีวิว · Bas: สร้าง practice/bas/033/relay.go และ relay_test.go แล้ว Relay ใช้ f1 ส่ง channel และ f2 รับแล้วคืนข้อความ · Prim: รีวิว PR dev/LOOP-15 แล้ว: practice/bas/033/relay.go ตรงกับ concurrency/033/README.md, main.go และ solution/solution.go ใช้ channel ให้ f1 ส่ง → f2 รับแล้วคืนข้อความ; ไม่ได้แก้ concurrency/033/*
