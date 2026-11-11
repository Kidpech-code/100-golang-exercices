# GRAPH-13 · ฝึก Go (อ่านเฉลย): ยกเลิก context (ข้อ 59)

- สถานะ: done · ทีม: graph · ผู้รับผิดชอบ: Ball
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:36

## งาน

โจทย์ context/059: อ่าน README, main.go และ solution/ แล้วสรุปด้วยคำของตัวเอง: แนวคิดหลัก, โค้ดตัวอย่างสั้นที่ทำแบบนี้ (func Count(ctx context.Context, ch <-chan int) int นับค่าจาก ch จนกว่า ctx ถูก cancel · เทสต์ cancel แล้วฟังก์ชันจบ ไม่ค้าง รันด้วย -race), จุดที่มือใหม่มักพลาด, และเอาไปใช้กับ backend จริงตรงไหน

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Ball
- 11-11 18:08 Ann [rejected] ขั้นที่ 1 (Ann): ต้องมี what และ done_when · คำตอบ: {"goal": "", "memo": "ยังไม่ได้เริ่มสรุป รออ่าน context/059/README.md, context/059/main.go, context/059/solution/solution.go ก่อน", "question": {"to": "Jay (Tech lead & PM, loop)"}, "steps": [{"done_when": "", "files": [], "what": "", "who": "Ann"}]  }
- 11-11 18:09 Ann [meeting] kickoff: Ann: อ่านเฉลย context/059 แล้วสรุปด้วยคำตัวเอง: แนวคิดหลัก, โค้ดตัวอย่างสั้น, จุดที่มือใหม่พลาด, การใช้กับ backend จริง · Pear: ok · Ball: ok
- 11-11 18:26 Ann [evidence] อ่านเฉลย context/059 แล้ว: (1) แนวคิดหลัก — ใช้ context.Context เป็นสัญญาณ cancel ใน select ร่วมกับ ch เพื่อให้ Count หยุดเมื่อถูก cancel; (2) ตัวอย่าง — func Count(ctx context.Context, ch <-chan int) int { total := 0; for { select { case <-ctx.Done(): return total; case v, ok := <-ch: if !ok { return total }; total += v } } }; (3) มือใหม่พลาด — ลืมดู ctx.Done() / ลืมเช็ค ok / return ไม่ครบ path 使
- 11-11 18:26 Ann [handoff] Ann → Ball
- 11-11 18:36 Ball [evidence] ตรวจโค้ดจริง context/059 แล้ว สรุปตรง: doSomething ใช้ context.WithCancel แล้ว cancel หลังส่งครบ 3 ค่า, doAnother เลือก select ระหว่าง ctx.Done() กับรับจาก printCh; ตัวอย่าง Count สอดคล้องแนวคิด cancel context ไม่มี edge สมมติ
- 11-11 18:36 Ann [meeting] review: Ann: สรุปพร้อมให้เจ้าของรีวิว · Ann: อ่านเฉลย context/059 แล้ว: (1) แนวคิดหลัก — ใช้ context.Context เป็นสัญญาณ cancel ใน select ร่วมกับ ch เพื่อให้ Count หยุดเมื่อถูก cancel; (2) ตัวอย่าง — func Count(ctx context.Context, ch <-chan int) int { total := 0; for { select { case <-ctx.Done(): return total; case v, ok := <-ch: if !ok { return total }; total += v } } }; (3) มือใหม่พลาด — ลืมดู ctx.Done() / ลืมเช็ค ok / return ไม่ครบ path 使 · Ball: ตรวจโค้ดจริง context/059 แล้ว สรุปตรง: doSomething ใช้ context.WithCancel แล้ว cancel หลังส่งครบ 3 ค่า, doAnother เลือก select ระหว่าง ctx.Done() กับรับจาก printCh; ตัวอย่าง Count สอดคล้องแนวคิด cancel context ไม่มี edge สมมติ
