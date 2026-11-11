# GRAPH-11 · ฝึก Go (อ่านเฉลย): เทียบเวลา (ข้อ 48)

- สถานะ: done · ทีม: graph · ผู้รับผิดชอบ: Ball
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 19:03

## งาน

โจทย์ time/048: อ่าน README, main.go และ solution/ แล้วสรุปด้วยคำของตัวเอง: แนวคิดหลัก, โค้ดตัวอย่างสั้นที่ทำแบบนี้ (func Compare(a, b time.Time) (diff time.Duration, equal, after bool) ใช้ Sub/Equal/After · เทสต์เวลาคงที่ ไม่ใช้ time.Now), จุดที่มือใหม่มักพลาด, และเอาไปใช้กับ backend จริงตรงไหน

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Ball
- 11-11 18:07 Ann [meeting] kickoff: Ann: สรุปโจทย์ time/048 การเทียบเวลาใน Go ให้อ่านแล้วเข้าใจครบ 4 หัวข้อ · Pear: ok · Ball: ok · Ying: ok
- 11-11 18:25 Ball [rejected] ส่งงานให้ Ann ซึ่งไม่อยู่ในสายงานของ Ball
- 11-11 18:34 Ball [evidence] GRAPH-11 draft สรุป time/048 (4 หัวข้อ): 1) แนวคิด: เปรียบเทียบเวลาด้วย time.Time methods — Sub() หา duration, Equal() เทียบว่าเท่ากัน, After()/Before() ว่าอันไหนหลัง/ก่อน. 2) โค้ดตัวอย่าง: func Compare(a, b time.Time) (diff time.Duration, equal, after bool) { return b.Sub(a), a.Equal(b), a.After(b) } และเทสต์ใช้เวลาคงที่อย่าง time.Date(...) ไม่ใช้ time.Now() เพื่อให้ deterministic. 3) จุดพลาด: ใช
- 11-11 18:34 Ball [handoff] Ball → Ying
- 11-11 18:35 Ying [working] ขออ่านไฟล์ time/048/README.md, time/048/main.go, time/048/solution/solution.go เพื่อตรวจ draft สรุป 4 หัวข้อ
- 11-11 18:48 Ying [rejected] need_help ต้องมี question.text
- 11-11 18:54 Ying [rejected] คำตอบยาวจนถูกตัดก่อนถึง work: summary และ memo ให้สั้น ใช้ === EDIT แทนการเขียนทั้งไฟล์
- 11-11 19:00 Ying [rejected] คำตอบยาวจนถูกตัดก่อนถึง work: summary และ memo ให้สั้น ใช้ === EDIT แทนการเขียนทั้งไฟล์
- 11-11 19:00 Ann [blocked] Ying ตอบผิดกติกา 3 ครั้งติด: คำตอบยาวจนถูกตัดก่อนถึง work: summary และ memo ให้สั้น ใช้ === EDIT แทนการเขียนทั้งไฟล์
- 11-11 19:01 Ann [meeting] stuck: Ann: สรุปแบบฝึก time/048 ให้ครบ 4 หัวข้อ สั้นพอใส่ใน summary+memo โดยไม่แก้ไฟล์ · Ying: ok · Pear: ok
- 11-11 19:02 Ball [rejected] ส่งงานให้ Ann ซึ่งไม่อยู่ในสายงานของ Ball
- 11-11 19:02 Ball [rejected] ส่งงานให้ Ann ซึ่งไม่อยู่ในสายงานของ Ball
- 11-11 19:03 Ball [evidence] 1) แนวคิดหลัก: เปรียบเทียบเวลาใน Go ด้วย Sub (ผลต่าง Duration), Equal (เท่ากัน), After/Before (ก่อน-หลัง) จาก time/048 README+solution<br>2) func Compare(a, b time.Time) (diff time.Duration, equal, after bool) { return a.Sub(b), a.Equal(b), a.After(b) } — ควรเทสต์ด้วยเวลาคงที่ ไม่ใช้ time.Now<br>3) จุดพลาด: ใช้ == แทน Equal (location ต่างแต่ instant เดียวกัน) และสลับข้าง After; ใช้ใน backend ตรวจ 
- 11-11 19:03 Ann [meeting] review: Ann: สรุปพร้อมให้เจ้าของรีวิว · Ball: 1) แนวคิดหลัก: เปรียบเทียบเวลาใน Go ด้วย Sub (ผลต่าง Duration), Equal (เท่ากัน), After/Before (ก่อน-หลัง) จาก time/048 README+solution<br>2) func Compare(a, b time.Time) (diff time.Duration, equal, after bool) { return a.Sub(b), a.Equal(b), a.After(b) } — ควรเทสต์ด้วยเวลาคงที่ ไม่ใช้ time.Now<br>3) จุดพลาด: ใช้ == แทน Equal (location ต่างแต่ instant เดียวกัน) และสลับข้าง After; ใช้ใน backend ตรวจ 
