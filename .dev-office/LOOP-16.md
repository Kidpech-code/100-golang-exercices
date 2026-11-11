# LOOP-16 · ฝึก Go: range บน buffered channel (ข้อ 42)

- สถานะ: done · ทีม: loop · ผู้รับผิดชอบ: Bas
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- ไฟล์: practice/bas/042/drain.go, practice/bas/042/drain_test.go
- เทสต์: `go test ./practice/bas/042/`
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 20:27

## งาน

โจทย์ concurrency/042. func Drain(nums []int) []int: ใส่ทุกตัวลง buffered channel, close, range อ่านกลับ · เทสต์ลำดับเท่าเดิม และ slice ว่าง · เขียนเองจาก README/main.go ของข้อนั้นก่อน แล้วค่อยเทียบกับ solution/ ตอนรีวิว · ไม่แก้ไฟล์โจทย์เดิม · package practice

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Bas
- 11-11 18:20 Jay [meeting] kickoff: Jay: ทำ Drain(nums []int) []int โดยใส่ทุกตัวลง buffered channel, close แล้ว range อ่านกลับ ให้เทสต์ผ่านและไม่แก้ไฟล์โจทย์เดิม · Nut: ok · Bas: ok · Prim: ok
- 11-11 18:22 Bas [commit] สร้าง practice/bas/042/drain.go ใช้ buffered channel + close + range คงลำดับเดิม และ drain_test.go เทสต์ slice ว่าง (6b5290a10e9780e1865049afc2ef52a4f483700c)
- 11-11 18:22 Bas [handoff] Bas → Prim
- 11-11 18:34 Prim [round] รอบที่ 1 ไม่ผ่าน: no file in your work. work holds a change to a file like this (SEARCH = lines copied exactly from the file): === EDIT: path/to/file.go <<<<<<< SEARCH <old lines> ======= <new lines> >>>>>>> REPLACE ==
- 11-11 18:42 Prim [round] รอบที่ 2 ไม่ผ่าน: no file in your work. work holds a change to a file like this (SEARCH = lines copied exactly from the file): === EDIT: path/to/file.go <<<<<<< SEARCH <old lines> ======= <new lines> >>>>>>> REPLACE ==
- 11-11 18:50 Prim [round] รอบที่ 3 ไม่ผ่าน: no file in your work. work holds a change to a file like this (SEARCH = lines copied exactly from the file): === EDIT: path/to/file.go <<<<<<< SEARCH <old lines> ======= <new lines> >>>>>>> REPLACE ==
- 11-11 18:52 Prim [commit] ตรวจโค้ด drain.go แล้ว ใช้ buffered channel + close + range ครบ · go test ./practice/bas/042/ ผ่าน · ไม่ได้แก้ไฟล์อื่น (af6eb092c9a5db3763b251e47072dc4290b3e512)
- 11-11 18:52 Prim [handoff] Prim → Bas
- 11-11 18:53 Bas [rejected] need_help ต้องมี question.text
- 11-11 18:55 Bas [rejected] need_help ต้องมี question.text
- 11-11 18:56 Bas [issue] ถาม Jay: ผมมีแค่ read/write/test/commit ไม่มีคำสั่งเปิด PR จะให้เปิด PR จาก dev/LOOP-16 ยังไง หรือระบบจะสร้าง PR ให้หลัง commit/push และลิงก์มาจากไหน?
- 11-11 20:18 owner [answer] repo นี้ไม่เปิด PR (pr=false, repo ฝึก) งานเสร็จบน branch dev/LOOP-16 แล้ว ข้ามขั้นเปิด PR และปิดงานได้เลย ไม่ต้องมีลิงก์ PR
- 11-11 20:20 Bas [round] รอบที่ 1 ไม่ผ่าน: no file in your work. work holds a change to a file like this (SEARCH = lines copied exactly from the file): === EDIT: path/to/file.go <<<<<<< SEARCH <old lines> ======= <new lines> >>>>>>> REPLACE ==
- 11-11 20:23 Bas [round] รอบที่ 2 ไม่ผ่าน: no file in your work. work holds a change to a file like this (SEARCH = lines copied exactly from the file): === EDIT: path/to/file.go <<<<<<< SEARCH <old lines> ======= <new lines> >>>>>>> REPLACE ==
- 11-11 20:27 owner [moved] done

## Issues

- Q1 Bas → owner: ผมมีแค่ read/write/test/commit ไม่มีคำสั่งเปิด PR จะให้เปิด PR จาก dev/LOOP-16 ยังไง หรือระบบจะสร้าง PR ให้หลัง commit/push และลิงก์มาจากไหน?
  - owner: repo นี้ไม่เปิด PR (pr=false, repo ฝึก) งานเสร็จบน branch dev/LOOP-16 แล้ว ข้ามขั้นเปิด PR และปิดงานได้เลย ไม่ต้องมีลิงก์ PR
