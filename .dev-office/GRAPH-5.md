# GRAPH-5 · ฝึก Go (อ่านเฉลย): for loop บวก 2 (ข้อ 9)

- สถานะ: done · ทีม: graph · ผู้รับผิดชอบ: Ying
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:18

## งาน

โจทย์ basic_exercices/009: อ่าน README, main.go และ solution/ แล้วสรุปด้วยคำของตัวเอง: แนวคิดหลัก, โค้ดตัวอย่างสั้นที่ทำแบบนี้ (func AddTwo(a []int) []int คืน slice ใหม่ ไม่แก้ของเดิม · เทสต์ว่าของเดิมไม่เปลี่ยน), จุดที่มือใหม่มักพลาด, และเอาไปใช้กับ backend จริงตรงไหน

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Ball
- 11-11 18:04 Ann [meeting] kickoff: Ann: ฝึก Go ข้อ 9: อ่าน README, main.go และ solution ของ for loop บวก 2 แล้วสรุปเป็นความรู้ 4 หัวข้อ · Pear: ok · Ball: ok · Ying: ok
- 11-11 18:14 Ball [evidence] อ่าน README, main.go และ solution ของ basic_exercices/009 แล้ว สรุปครบ 4 หัวข้อไว้ใน memo: for loop +2, โค้ด AddTwo ที่คืน slice ใหม่ไม่แก้ของเดิม, จุดพลาด และการประยุกต์ backend
- 11-11 18:14 Ball [handoff] Ball → Ying
- 11-11 18:18 Ying [evidence] ตรวจทานสรุปของ Ball แล้วถูกต้อง: โค้ด func AddTwo(a []int) []int สร้าง slice ใหม่ b แล้ววน for i := range a ใส่ a[i]+2 ลง b จึงไม่แก้ slice เดิม และตรงกับโจทย์ 009 ที่ให้วนลูปบวก 2 ทุกค่า ไม่มีประเด็นค้าง
- 11-11 18:18 Ann [meeting] review: Ann: สรุปพร้อมให้เจ้าของรีวิว · Ball: อ่าน README, main.go และ solution ของ basic_exercices/009 แล้ว สรุปครบ 4 หัวข้อไว้ใน memo: for loop +2, โค้ด AddTwo ที่คืน slice ใหม่ไม่แก้ของเดิม, จุดพลาด และการประยุกต์ backend · Ying: ตรวจทานสรุปของ Ball แล้วถูกต้อง: โค้ด func AddTwo(a []int) []int สร้าง slice ใหม่ b แล้ววน for i := range a ใส่ a[i]+2 ลง b จึงไม่แก้ slice เดิม และตรงกับโจทย์ 009 ที่ให้วนลูปบวก 2 ทุกค่า ไม่มีประเด็นค้าง
