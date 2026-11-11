# LOOP-9 · ฝึก Go: if/else if อุณหภูมิ (ข้อ 13)

- สถานะ: done · ทีม: loop · ผู้รับผิดชอบ: Bas
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- ไฟล์: practice/bas/013/weather.go, practice/bas/013/weather_test.go
- เทสต์: `go test ./practice/bas/013/`
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:59

## งาน

โจทย์ basic_exercices/013. func Weather(t int) string คืน "too cold" (<20), "perfect" (20-30), "so hot" (>30) · เทสต์ขอบ 19,20,30,31 · เขียนเองจาก README/main.go ของข้อนั้นก่อน แล้วค่อยเทียบกับ solution/ ตอนรีวิว · ไม่แก้ไฟล์โจทย์เดิม · package practice

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Bas
- 11-11 18:13 Jay [meeting] kickoff: Jay: สร้าง practice/bas/013/weather.go และ weather_test.go ให้ func Weather คืนค่า 'too cold'/'perfect'/'so hot' ตามเงื่อนไข และ go test ./practice/bas/013/ ผ่าน · Nut: ok · Bas: ok
- 11-11 18:32 Bas [working] ขออ่าน basic_exercices/013/README.md และ basic_exercices/013/main.go ก่อนสร้าง weather.go และ weather_test.go
- 11-11 18:37 Bas [commit] อ่าน basic_exercices/013/README.md และ main.go ของข้อ 013 แล้วสร้าง func Weather คืน too cold/perfect/so hot ตามขอบ 20-30 และเทสต์ครอบคลุม 19,20,30,31 (cebafb2e88fe2356a4a33135a2a8d060ebe96aab)
- 11-11 18:37 Bas [handoff] Bas → Bas
- 11-11 18:46 Bas [rejected] need_help ต้องมี question.text
- 11-11 18:49 Bas [rejected] need_help ต้องมี question.text
- 11-11 18:56 Bas [rejected] need_help ต้องมี question.text
- 11-11 18:56 Jay [blocked] Bas ตอบผิดกติกา 3 ครั้งติด: need_help ต้องมี question.text
- 11-11 18:56 Jay [meeting] stuck: Jay: ฝึก Go if/else: Weather(t int) คืน too cold/perfect/so hot ตามขอบ 20/30 และเทสต์ผ่าน · Bas: ok · Nut: ok
- 11-11 18:59 Bas [commit] ตรวจ weather.go/weather_test.go แล้ว ปรับ Weather เป็น if/else-if ตามโจทย์ (19=too cold, 20-30=perfect, 31=so hot) เทสต์ครบ 4 ขอบแล้ว (aebc1cee2c91294844ee7b248d10ae1846cad7e4)
- 11-11 18:59 Jay [meeting] review: Jay: สรุปพร้อมให้เจ้าของรีวิว · Bas: ตรวจ weather.go/weather_test.go แล้ว ปรับ Weather เป็น if/else-if ตามโจทย์ (19=too cold, 20-30=perfect, 31=so hot) เทสต์ครบ 4 ขอบแล้ว
