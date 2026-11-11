# GRAPH-6 · ฝึก Go (อ่านเฉลย): เช็กไฟล์มีไหม (ข้อ 16)

- สถานะ: done · ทีม: graph · ผู้รับผิดชอบ: Ball
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:15

## งาน

โจทย์ file_management/016: อ่าน README, main.go และ solution/ แล้วสรุปด้วยคำของตัวเอง: แนวคิดหลัก, โค้ดตัวอย่างสั้นที่ทำแบบนี้ (func Exists(path string) bool ใช้ os.Stat + errors.Is(err, fs.ErrNotExist) · เทสต์ใน t.TempDir()), จุดที่มือใหม่มักพลาด, และเอาไปใช้กับ backend จริงตรงไหน

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Ball
- 11-11 18:04 Ann [meeting] kickoff: Ann: สรุปบทเรียนข้อ 016 เช็กไฟล์มีไหม ให้เข้าใจแนวคิด โค้ดตัวอย่าง จุดพลาด และการใช้ใน backend จริง · Pear: ok · Ball: ok
- 11-11 18:14 Ball [evidence] อ่าน file_management/016 (README, main.go, solution) แล้ว สรุป: ใช้ os.Stat เช็กสถานะไฟล์ ถ้า err == nil แปลว่ามีไฟล์; โค้ดสั้นคือ func Exists(path string) bool { _, err := os.Stat(path); return !errors.Is(err, fs.ErrNotExist) } · เทสต์ควรสร้างไฟล์ใน t.TempDir() จุดพลาดคืออย่าตัดสินด้วย err != nil เพียงอย่างเดียว ต้องเช็ก fs.ErrNotExist (สิทธิ์/error อื่นไม่ใช่ 'ไม่มีไฟล์') · backend ใช้ก่อนอ่าน/寫
- 11-11 18:14 Ball [handoff] Ball → Ball
- 11-11 18:15 Ball [evidence] วิเคราะห์ file_management/016 เสร็จ: จุดพลาดมือใหม่ 3 ข้อ + ตัวอย่างเช็ก config ใน backend ไม่ได้แก้ไฟล์
- 11-11 18:15 Ann [meeting] review: Ann: สรุปพร้อมให้เจ้าของรีวิว · Ball: อ่าน file_management/016 (README, main.go, solution) แล้ว สรุป: ใช้ os.Stat เช็กสถานะไฟล์ ถ้า err == nil แปลว่ามีไฟล์; โค้ดสั้นคือ func Exists(path string) bool { _, err := os.Stat(path); return !errors.Is(err, fs.ErrNotExist) } · เทสต์ควรสร้างไฟล์ใน t.TempDir() จุดพลาดคืออย่าตัดสินด้วย err != nil เพียงอย่างเดียว ต้องเช็ก fs.ErrNotExist (สิทธิ์/error อื่นไม่ใช่ 'ไม่มีไฟล์') · backend ใช้ก่อนอ่าน/寫 · Ball: วิเคราะห์ file_management/016 เสร็จ: จุดพลาดมือใหม่ 3 ข้อ + ตัวอย่างเช็ก config ใน backend ไม่ได้แก้ไฟล์
