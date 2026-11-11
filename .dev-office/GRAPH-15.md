# GRAPH-15 · ฝึก Go (อ่านเฉลย): Gin GET /albums/:id (ข้อ 81)

- สถานะ: done · ทีม: graph · ผู้รับผิดชอบ: Ying
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:56

## งาน

โจทย์ http_apis/081: อ่าน README, main.go และ solution/ แล้วสรุปด้วยคำของตัวเอง: แนวคิดหลัก, โค้ดตัวอย่างสั้นที่ทำแบบนี้ (func Router() *gin.Engine มี GET /albums/:id คืน 200+JSON หรือ 404 · เทสต์ด้วย httptest เคสเจอและไม่เจอ), จุดที่มือใหม่มักพลาด, และเอาไปใช้กับ backend จริงตรงไหน

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Ball
- 11-11 18:12 Ann [meeting] kickoff: Ann: สรุปบทเรียน Gin GET /albums/:id (ข้อ 81) จาก http_apis/081 ให้ครบ 4 หัวข้อ โดยไม่แก้ไฟล์ · Pear: ok · Ball: ok · Ying: ok
- 11-11 18:26 Ball [evidence] อ่าน http_apis/081/README.md, main.go, solution/solution.go แล้ว ข้อนี้คือ Gin path parameter: r.GET('/albums/:id') ใช้ c.Param('id') ค้น album ใน slice เจอตอบ 200+JSON ไม่เจอตอบ 404 โค้ดสำคัญ: for _, a := range albums { if a.ID == id { c.IndentedJSON(200, a); return } }; c.IndentedJSON(404, gin.H{'message':'album not found'}) ควรแยก Router() เพื่อเทสต์ httptest เคส /albums/1 (200) กับ /albums/999
- 11-11 18:26 Ball [handoff] Ball → Ying
- 11-11 18:54 Ying [working] ขออ่านไฟล์เพื่อตรวจ: http_apis/081/README.md, http_apis/081/main.go, http_apis/081/solution/solution.go
- 11-11 18:56 Ying [evidence] ตรวจ http_apis/081 แล้ว สรุปครบ 4 ข้อ: (1) แนวคิดหลักคือ Gin path parameter — register r.GET('/albums/:id') แล้วอ่านค่าด้วย c.Param('id') เพื่อค้น album ใน slice; (2) โค้ดตัวอย่างจาก solution: for _, album := range albums { if id == album.ID { c.IndentedJSON(http.StatusOK, album); return } }; แล้ว c.IndentedJSON(http.StatusNotFound, gin.H{'message':'album not found'}) และใน main ลงทะเบียน router /
- 11-11 18:56 Ann [meeting] review: Ann: สรุปพร้อมให้เจ้าของรีวิว · Ball: อ่าน http_apis/081/README.md, main.go, solution/solution.go แล้ว ข้อนี้คือ Gin path parameter: r.GET('/albums/:id') ใช้ c.Param('id') ค้น album ใน slice เจอตอบ 200+JSON ไม่เจอตอบ 404 โค้ดสำคัญ: for _, a := range albums { if a.ID == id { c.IndentedJSON(200, a); return } }; c.IndentedJSON(404, gin.H{'message':'album not found'}) ควรแยก Router() เพื่อเทสต์ httptest เคส /albums/1 (200) กับ /albums/999 · Ying: ตรวจ http_apis/081 แล้ว สรุปครบ 4 ข้อ: (1) แนวคิดหลักคือ Gin path parameter — register r.GET('/albums/:id') แล้วอ่านค่าด้วย c.Param('id') เพื่อค้น album ใน slice; (2) โค้ดตัวอย่างจาก solution: for _, album := range albums { if id == album.ID { c.IndentedJSON(http.StatusOK, album); return } }; แล้ว c.IndentedJSON(http.StatusNotFound, gin.H{'message':'album not found'}) และใน main ลงทะเบียน router /
