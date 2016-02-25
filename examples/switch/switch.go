// คำสั่ง _Switch_ ใช้สำหรับการจัดการเงื่อนไขที่มีหลายทางเลือก

package main

import "fmt"
import "time"

func main() {

    // การใช้งาน `switch` แบบง่าย ๆ
    i := 2
    fmt.Print("write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    // สามารถใช้ เครื่องหมาย (,) <br/>
    // สำหรับกรณีที่มีมากกว่าหนึ่งเงื่อนไขในกรณีเดียว<br/>
    // และถ้าไม่ตรงกับกรณีใดเลย สามารถกำหนดใำเข้าทำงานในส่วนของ `default` ได้
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("it's the weekend")
    default:
        fmt.Println("it's a weekday")
    }

    // ถ้าในคำสั่ง `switch` ไม่มีชุดคำสั่งใด ๆ เลย<br/>
    // มันคือ การทำงานในอีกรูปแบบหนึ่งของ if/else นั่นเอง<br/>
    // โดยที่ในแต่ละกรณีไม่จำเป็นต้องเป็นค่าคงที่แล้ว
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("it's before noon")
    default:
        fmt.Println("it's after noon")
    }
}

// todo: type switches
