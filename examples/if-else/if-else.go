// การตรวจสอบเงื่อนไขด้วย `if` และ `else` เป็นสิ่งที่ง่ายมาก

package main

import "fmt"

func main() {

    // ตัวอย่างการใช้งานขั้นพื้นฐาน
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    // สามารถมีคำสั่ง `if` โดยไม่ต้องมี `else` ได้
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    // ในแต่ละ statement สามารถทำการประกาศและกำหนดค่าให้ตัวแปร
    // รวมทั้งทำการตรวจสอบเงื่อนไขได้ด้วย
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}

// ในภาษา Go นั้น <br/>
// จะเห็นได้ว่าในเงื่อนไขไม่มีการครอบด้วยวงเล็บ<br/>
// แต่ต้องมีปีกกาเปิดและปิดเสมอนะ
