# GRAPH-9 · ฝึก Go (อ่านเฉลย): รอ goroutine แรกเสร็จก่อน (ข้อ 36)

- สถานะ: done · ทีม: graph · ผู้รับผิดชอบ: Ball
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:18

## งาน

โจทย์ concurrency/036: อ่าน README, main.go และ solution/ แล้วสรุปด้วยคำของตัวเอง: แนวคิดหลัก, โค้ดตัวอย่างสั้นที่ทำแบบนี้ (func RunInOrder() []string: task ทำก่อน task2 เสมอ ใช้ channel ส่งสัญญาณ ไม่ใช้ time.Sleep · เทสต์ได้ ["task","task2"] และรันด้วย -race), จุดที่มือใหม่มักพลาด, และเอาไปใช้กับ backend จริงตรงไหน

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Ball
- 11-11 18:06 Ann [meeting] kickoff: Ann: สรุปบทเรียน Go ข้อ concurrency/036: ใช้ channel ส่งสัญญาณเพื่อให้ task รอ task2 เสมอ ไม่ใช้ time.Sleep · Pear: ok · Ball: ok
- 11-11 18:18 Ball [evidence] แนวคิดหลัก: goroutine ทำงานพร้อมกัน แต่ถ้าต้องการลำดับต้องใช้ channel ส่งสัญญาณว่า task แรกเสร็จแล้ว แล้วค่อยเริ่ม task2; main รอรับจาก done ด้วย <-done จึงเริ่ม goroutine ถัดไป การใช้ time.Sleep ไม่รับประกันลำดับและช้า unnecessary. โค้ดตัวอย่างสั้น func RunInOrder() []string { done := make(chan bool, 1); go func(){ task(); done <- true }(); <-done; task2(); return []string{...} } — ใช้ channel สร
- 11-11 18:18 Ann [meeting] review: Ann: สรุปพร้อมให้เจ้าของรีวิว · Ball: แนวคิดหลัก: goroutine ทำงานพร้อมกัน แต่ถ้าต้องการลำดับต้องใช้ channel ส่งสัญญาณว่า task แรกเสร็จแล้ว แล้วค่อยเริ่ม task2; main รอรับจาก done ด้วย <-done จึงเริ่ม goroutine ถัดไป การใช้ time.Sleep ไม่รับประกันลำดับและช้า unnecessary. โค้ดตัวอย่างสั้น func RunInOrder() []string { done := make(chan bool, 1); go func(){ task(); done <- true }(); <-done; task2(); return []string{...} } — ใช้ channel สร
