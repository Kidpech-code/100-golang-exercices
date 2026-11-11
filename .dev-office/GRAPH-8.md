# GRAPH-8 · ฝึก Go (อ่านเฉลย): variadic function (ข้อ 30)

- สถานะ: done · ทีม: graph · ผู้รับผิดชอบ: Ann
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:16

## งาน

โจทย์ intermediate_exercices/030: อ่าน README, main.go และ solution/ แล้วสรุปด้วยคำของตัวเอง: แนวคิดหลัก, โค้ดตัวอย่างสั้นที่ทำแบบนี้ (func Sum(n ...int) int · เทสต์ ไม่ส่งอะไรเลย, 1 ตัว, ส่ง slice ด้วย s...), จุดที่มือใหม่มักพลาด, และเอาไปใช้กับ backend จริงตรงไหน

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Ball
- 11-11 18:06 Ann [meeting] kickoff: Ann: อ่านเฉลยข้อ 30 แล้วสรุป variadic function พร้อมตัวอย่างและจุดนำไปใช้ใน backend · Pear: ok · Ball: ok
- 11-11 18:15 Ball [evidence] อ่าน README/main.go/solution ข้อ 30 แล้ว: 1) แนวคิดหลักคือ variadic function ใช้ ... ข้างหน้าพารามิเตอร์ตัวสุดท้าย รับอาร์กิวเมนต์กี่ตัวก็ได้ และข้างในเป็น slice 2) โค้ดตัวอย่าง: func Sum(n ...int) int { total := 0; for _, v := range n { total += v }; return total } ใช้ได้ทั้ง Sum(), Sum(1,2), และ Sum(s...) เมื่อ s เป็น slice 3) มือใหม่พลาดตรงลืม ... ใน signature, ใส่พารามิเตอร์ตัวอื่นต่อท้าย vari
- 11-11 18:15 Ball [handoff] Ball → Ann
- 11-11 18:16 Ann [evidence] ตรวจทาน summary ของ Ball แล้วถูกต้อง: variadic function ใช้ ... กับพารามิเตอร์ตัวสุดท้าย ข้างในเป็น slice; func Sum(n ...int) int ใช้ได้ทั้ง Sum(), Sum(1,2), Sum(s...) เมื่อ s เป็น []int; จุดพลาดคือลืม ... ใน signature และห้ามมีพารามิเตอร์ตัวอื่นต่อท้าย variadic ไม่มีการแก้ไฟล์
- 11-11 18:16 Ann [meeting] review: Ann: สรุปพร้อมให้เจ้าของรีวิว · Ball: อ่าน README/main.go/solution ข้อ 30 แล้ว: 1) แนวคิดหลักคือ variadic function ใช้ ... ข้างหน้าพารามิเตอร์ตัวสุดท้าย รับอาร์กิวเมนต์กี่ตัวก็ได้ และข้างในเป็น slice 2) โค้ดตัวอย่าง: func Sum(n ...int) int { total := 0; for _, v := range n { total += v }; return total } ใช้ได้ทั้ง Sum(), Sum(1,2), และ Sum(s...) เมื่อ s เป็น slice 3) มือใหม่พลาดตรงลืม ... ใน signature, ใส่พารามิเตอร์ตัวอื่นต่อท้าย vari · Ann: ตรวจทาน summary ของ Ball แล้วถูกต้อง: variadic function ใช้ ... กับพารามิเตอร์ตัวสุดท้าย ข้างในเป็น slice; func Sum(n ...int) int ใช้ได้ทั้ง Sum(), Sum(1,2), Sum(s...) เมื่อ s เป็น []int; จุดพลาดคือลืม ... ใน signature และห้ามมีพารามิเตอร์ตัวอื่นต่อท้าย variadic ไม่มีการแก้ไฟล์
