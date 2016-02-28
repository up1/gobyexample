// _Maps_ คือชนิดข้อมูลหนึ่งที่อยู่ในภาษา Go [associative data type](http://en.wikipedia.org/wiki/Associative_array)
// โครงสร้างข้อมูลที่ประกอบไปด้วย key และ value คู่กันเสมอ <br/>
// ในภาษาโปรแกรมอื่น ๆ อาจจะเรียกว่า _hashes_ หรือ _dicts_

package main

import "fmt"

func main() {

    // ทำการสร้าง map ว่าง ๆ ด้วยคำสั่ง `make`:<br/>
    // รูปแบบการใช้งาน `make(map[key-type]val-type)`
    m := make(map[string]int)

    // ทำการกำหนดค่าของ key-value `name[key] = val`
    m["k1"] = 7
    m["k2"] = 13

    // แสดงข้อมูลที่อยู่ใน map
    fmt.Println("map:", m)

    // ทำการดึงข้อมูลจาก map ด้วย key `name[key]`.
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    // ใช้คำสั่ง `len` สำหรับดึงจำนวนข้อมูลใน map
    fmt.Println("len:", len(m))

    // ใช้คำสั่ง `delete` สำหรับลบข้อมูลใน map
    delete(m, "k2")
    fmt.Println("map:", m)

    // ถ้าไม่ต้องการใช้ข้อมูลที่ดึงมาจาก map แล้ว<br/>
    // สามารถทำการ ignore หรือ ไม่สนใจค่านั้น<br/>
    // ด้วยการใช้  _blank identifier_ หรือ `_`
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    // สามารถประกาศ และ กำหนดค่าเริ่มต้นให้กับตัวแปรชนิด map ได้
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}
