# GRAPH-14 · ฝึก Go (อ่านเฉลย): interface area/perim (ข้อ 64)

- สถานะ: done · ทีม: graph · ผู้รับผิดชอบ: Ball
- Epic: EPIC-3 · ฝึก Go: Backend 2 คนทำแบบฝึก 100-golang-exercices
- เปิด 2026-11-11 13:32 · ปิด 2026-11-11 18:26

## งาน

โจทย์ interfaces/064: อ่าน README, main.go และ solution/ แล้วสรุปด้วยคำของตัวเอง: แนวคิดหลัก, โค้ดตัวอย่างสั้นที่ทำแบบนี้ (type Geometry interface{Area() float64; Perim() float64}; type Rect{W,H float64} · เทสต์ผ่าน interface ไม่เรียก Rect ตรง), จุดที่มือใหม่มักพลาด, และเอาไปใช้กับ backend จริงตรงไหน

## บันทึก

- 11-11 13:32 owner [created] จากงานใหญ่ EPIC-3 · มอบให้ Ball
- 11-11 18:10 Ann [meeting] kickoff: Ann: อ่านเฉลย interfaces/064 แล้วสรุปบทเรียน Go เรื่อง interface (Area/Perim) ให้อ่านแล้วเข้าใจและเอาไปใช้กับ backend ได้ · Pear: ok · Ball: ok
- 11-11 18:26 Ball [evidence] 1) แนวคิด: interface='สัญญา' รวมหลาย method (type geometry มี area()+perim()); type ที่ implement ครบทุก method ก็ใช้ได้แบบ implicit; ถ้าเพิ่ม method ต้องแก้ implementer ทุกตัว 2) โค้ด: geometry interface { area() float64; perim() float64 }, rect struct {width,height float64} มี func (r rect) area() และ func (r rect) perim() float64 { return 2*r.width + 2*r.height }; main เรียก r.area(), r.perim()
- 11-11 18:26 Ann [meeting] review: Ann: สรุปพร้อมให้เจ้าของรีวิว · Ball: 1) แนวคิด: interface='สัญญา' รวมหลาย method (type geometry มี area()+perim()); type ที่ implement ครบทุก method ก็ใช้ได้แบบ implicit; ถ้าเพิ่ม method ต้องแก้ implementer ทุกตัว 2) โค้ด: geometry interface { area() float64; perim() float64 }, rect struct {width,height float64} มี func (r rect) area() และ func (r rect) perim() float64 { return 2*r.width + 2*r.height }; main เรียก r.area(), r.perim()
