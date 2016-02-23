// ในภาษา Go นั้นสนับสนุนเรื่องของค่าคงที่<br/>
// ซึ่งค่าคงที่ประกอบไปด้วย character, string, boolean และตัวเลข

package main

import "fmt"
import "math"

// สามารถประกาศค่าคงที่ด้วย `const`
const s string = "constant"

func main() {
    fmt.Println(s)

    // โดยจะมีรูปแบบการใช้งานเหมือนกับ `var`
    const n = 500000000

    // สามารถใช้งาน operation ต่าง ๆ ทางคณิตศาสตร์ได้
    const d = 3e20 / n
    fmt.Println(d)

    // โดยที่ค่าคงที่จะไม่มีชนิดของข้อมูลไปจนกว่า จะถูก cast type
    fmt.Println(int64(d))

    // ค่าคงที่ที่มีข้อมูลเป็นตัวเลข
    // จะทำการกำหนดชนิดข้อมูลที่เหมาะสมให้</br>
    // ตัวอย่างเช่นการใช้งาน `math.Sin` จะกำหนดชนิดข้อมูลเป็น `float64` ให้
    fmt.Println(math.Sin(n))
}
