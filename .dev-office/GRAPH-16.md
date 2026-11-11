# GRAPH-16 · ฝึก Go (อ่านเฉลย): mutex + WaitGroup (ข้อ 89)

- สถานะ: done · ทีม: graph · ผู้รับผิดชอบ: Ball
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:27

## งาน

โจทย์ testing_misc/089: อ่าน README, main.go และ solution/ แล้วสรุปด้วยคำของตัวเอง: แนวคิดหลัก, โค้ดตัวอย่างสั้นที่ทำแบบนี้ (type SafeCounter มี Inc(key) และ Value(key) · เทสต์ 100 goroutine Inc แล้วต้องได้ 100 รันด้วย -race), จุดที่มือใหม่มักพลาด, และเอาไปใช้กับ backend จริงตรงไหน

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Ball
- 11-11 18:11 Ann [meeting] kickoff: Ann: สรุปบทเรียน Go เรื่อง mutex + WaitGroup จาก testing_misc/089 ให้ทีมเข้าใจ SafeCounter และวิธีเอาไปใช้กับ backend จริง · Pear: ok · Ball: ok
- 11-11 18:27 Ball [evidence] 1) แนวคิดหลัก: goroutine แชร์ตัวแปรต้องมี sync กัน race; README ระบุ 'Race conditions occur when... No synchronization mechanism prevents concurrent access'; ต้องใช้ Mutex เพราะ c.n++ เป็น read-modify-write ถ้าไม่ lock จะ lost update (5x100000 ควรได้ 500000 ไม่ใช่ 100000) และต้องใช้ WaitGroup เพราะ wg.Add(5)/Done/Wait ให้ main รอ goroutine ครบก่อนอ่านค่า. 2) โค้ดจริง solution.go: SafeCounter{mu ​​
- 11-11 18:27 Ann [meeting] review: Ann: สรุปพร้อมให้เจ้าของรีวิว · Ball: 1) แนวคิดหลัก: goroutine แชร์ตัวแปรต้องมี sync กัน race; README ระบุ 'Race conditions occur when... No synchronization mechanism prevents concurrent access'; ต้องใช้ Mutex เพราะ c.n++ เป็น read-modify-write ถ้าไม่ lock จะ lost update (5x100000 ควรได้ 500000 ไม่ใช่ 100000) และต้องใช้ WaitGroup เพราะ wg.Add(5)/Done/Wait ให้ main รอ goroutine ครบก่อนอ่านค่า. 2) โค้ดจริง solution.go: SafeCounter{mu ​​
