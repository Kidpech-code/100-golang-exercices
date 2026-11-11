# LOOP-7 · แก้เฉลยข้อ 25 ที่ใช้ Println กับ %d

- สถานะ: done · ทีม: loop · ผู้รับผิดชอบ: Bas
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- ไฟล์: intermediate_exercices/025/solution/solution.go
- เทสต์: `go vet ./intermediate_exercices/025/solution/`
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:18

## งาน

intermediate_exercices/025/solution/solution.go: fmt.Println("... %d", x) ไม่แทนค่า ให้ใช้ fmt.Printf กับ %d สำหรับค่า และ %p สำหรับ address พร้อม \n

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Bas
- 11-11 18:12 Jay [meeting] kickoff: Jay: แก้เฉลยข้อ 25 ให้ใช้ fmt.Printf พร้อม %d, %p และ \n แทน fmt.Println ที่ไม่แทนค่า · Nut: ok · Bas: ok
- 11-11 18:18 Bas [commit] เปลี่ยน fmt.Println เป็น fmt.Printf ในเฉลยข้อ 25 โดยใช้ %d สำหรับค่า และ %p สำหรับ address/pointer พร้อม เติมตัวขึ้นบรรทัดใหม่ (c4ad21a50687a938aaac5b08cd5bad2b9a0bd24b)
- 11-11 18:18 Jay [meeting] review: Jay: สรุปพร้อมให้เจ้าของรีวิว · Bas: เปลี่ยน fmt.Println เป็น fmt.Printf ในเฉลยข้อ 25 โดยใช้ %d สำหรับค่า และ %p สำหรับ address/pointer พร้อม เติมตัวขึ้นบรรทัดใหม่
