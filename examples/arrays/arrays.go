// ในภาษาโก _อาเรย์_ คือตัวเลขลำดับของสมาชิกที่มีความยาวคงที่

package main

import "fmt"

func main() {

    // สร้างอาเรย์ `a` ชนิด int ขนาด 5 element
    // โดยที่ค่าเริ่มต้นในแต่ละ element คือ `0`
    var a [5]int
    fmt.Println("emp:", a)

    // เราสามารถ set และ get ค่าจากอาเรย์ได้โดยใช้ index
    // set โดยใช้รูปแบบ `array[index] = value`,
    // และ get ค่าโดยใช้รูปแบบ `array[index]`.
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // ใช้ฟังชั่น `len` เมื่อเราอยากรู้ขนาดของอาเรย์.
    fmt.Println("len:", len(a))

    // สำหรับการประกาศและกำหนดค่าเริ่มต้นให้กับอาเรย์ในหนึ่งบรรทัดจะใช้รูปแบบ.
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // ตัวแปรชนิดอาเรย์จะเป็นแบบหนึ่งมิติ แต่เราก็สามารถประกอบให้เป็นหลายๆ มิติได้เช่นกัน
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
