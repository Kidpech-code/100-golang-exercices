# GRAPH-10 · ฝึก Go (อ่านเฉลย): select กับ timeout (ข้อ 40)

- สถานะ: done · ทีม: graph · ผู้รับผิดชอบ: Ball
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:18

## งาน

โจทย์ concurrency/040: อ่าน README, main.go และ solution/ แล้วสรุปด้วยคำของตัวเอง: แนวคิดหลัก, โค้ดตัวอย่างสั้นที่ทำแบบนี้ (func Wait(work, limit time.Duration) string คืนข้อความจากงาน หรือ "Timeout!!!!" · เทสต์ใช้ ms ไม่เกิน 100ms ต่อเคส ห้ามรอ 10 วินาที), จุดที่มือใหม่มักพลาด, และเอาไปใช้กับ backend จริงตรงไหน

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Ball
- 11-11 18:05 Ann [meeting] kickoff: Ann: สรุปแนวคิด select กับ timeout จากเฉลยข้อ 40 เป็นภาษาไทย · Pear: ok
- 11-11 18:06 Ann [blocked] Ann: อ่านไฟล์ไม่เห็นใน environment: ไม่มี README/main.go/solution ของ concurrency/040 โผล่มาให้อ่าน ไม่ควรเดาข้อมูลจาก repo
- 11-11 18:09 Ann [meeting] stuck: Ann: สรุป select + timeout จาก concurrency/040 ให้ครบ 4 หัวข้อ · Pear: ok · Ball: ok
- 11-11 18:18 Ball [evidence] อ่าน concurrency/040 (README/main/solution) แล้ว สรุป 4 หัวข้อลง memo ตรงกับเฉลย: แนวคิด select + time.After, ตัวอย่าง func Wait(work, limit), จุดที่มือใหม่พลาด และการเอาไปใช้กับ backend จริง
- 11-11 18:18 Ann [meeting] review: Ann: สรุปพร้อมให้เจ้าของรีวิว · Ball: อ่าน concurrency/040 (README/main/solution) แล้ว สรุป 4 หัวข้อลง memo ตรงกับเฉลย: แนวคิด select + time.After, ตัวอย่าง func Wait(work, limit), จุดที่มือใหม่พลาด และการเอาไปใช้กับ backend จริง
