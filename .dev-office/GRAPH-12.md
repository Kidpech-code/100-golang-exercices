# GRAPH-12 · ฝึก Go (อ่านเฉลย): JSON ซ้อนกัน (ข้อ 51)

- สถานะ: done · ทีม: graph · ผู้รับผิดชอบ: Ying
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:37

## งาน

โจทย์ json/051: อ่าน README, main.go และ solution/ แล้วสรุปด้วยคำของตัวเอง: แนวคิดหลัก, โค้ดตัวอย่างสั้นที่ทำแบบนี้ (type Dimensions{Height, Weight string}; Human มี Dimensions; func ParseDims(b []byte) (Dimensions, error) · เทสต์ JSON ซ้อน และ JSON ไม่มี dimensions), จุดที่มือใหม่มักพลาด, และเอาไปใช้กับ backend จริงตรงไหน

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Ball
- 11-11 18:10 Ann [meeting] kickoff: Ann: อ่านเฉลย json/051 แล้วสรุปเรื่อง JSON ซ้อนกัน (nested JSON) ด้วยคำของตัวเองให้เห็นแนวคิดและวิธีเอาไปใช้กับ backend · Pear: ok · Ball: ok · Ying: ok
- 11-11 18:25 Ball [evidence] อ่านเฉลย json/051 ครบ 3 ไฟล์แล้ว สรุปครบ 4 หัวข้อใน memo: nested JSON -> nested struct, เฉลยจริงใช้ Dimensions/Human + Unmarshal, จุดพลาดเรื่อง exported fields/error handling, ใช้กับ backend ตรง parse body ของ API ไม่แก้ไฟล์เพราะเป็นการวิเคราะห์ล้วน
- 11-11 18:25 Ball [handoff] Ball → Ying
- 11-11 18:37 Ying [evidence] ตรวจ json/051/solution/solution.go แล้ว สรุปของ Ball ตรงเฉลย: มี type Dimensions{Height, Weight string} และ type Human{Name, Description, Dimensions Dimensions} จริง, unmarshal แล้ว print ผ่าน human.Dimensions.*; ข้อแนะนำ ParseDims เป็นของเพิ่มตอนใช้ backend ไม่ได้อ้างว่าเป็นฟังก์ชันในเฉลย จึงไม่พบ mismatch
- 11-11 18:37 Ann [meeting] review: Ann: สรุปพร้อมให้เจ้าของรีวิว · Ball: อ่านเฉลย json/051 ครบ 3 ไฟล์แล้ว สรุปครบ 4 หัวข้อใน memo: nested JSON -> nested struct, เฉลยจริงใช้ Dimensions/Human + Unmarshal, จุดพลาดเรื่อง exported fields/error handling, ใช้กับ backend ตรง parse body ของ API ไม่แก้ไฟล์เพราะเป็นการวิเคราะห์ล้วน · Ying: ตรวจ json/051/solution/solution.go แล้ว สรุปของ Ball ตรงเฉลย: มี type Dimensions{Height, Weight string} และ type Human{Name, Description, Dimensions Dimensions} จริง, unmarshal แล้ว print ผ่าน human.Dimensions.*; ข้อแนะนำ ParseDims เป็นของเพิ่มตอนใช้ backend ไม่ได้อ้างว่าเป็นฟังก์ชันในเฉลย จึงไม่พบ mismatch
