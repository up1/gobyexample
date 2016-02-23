// ในภาษา Go นั้นมีชนิดตัวแปรจำนวนมาก</br>
// ไม่ว่าจะเป็น string, integer, float, boolean</br>
// มาดูตัวอย่างการใช้งานแบบเล็ก ๆ กัน

package main

import "fmt"

func main() {

    // Strings, ทำการเชื่อมต่อข้อมูลด้วย `+`
    fmt.Println("go" + "lang")

    // Integers และ floats
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // Booleans, สามารถใช้งาน operator ต่าง ๆ เหล่านี้ได้เลย
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
