# แม้ว่า slices โครงสร้างภายในจะต่างกับ arrays
# แต่เวลาใช้ `fmt.Println` จะแสดงผลคล้ายกัน
$ go run slices.go
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]

# ลองเข้าไปอ่านบล็อก
# [โพสต์ที่ยอดเยี่ยม]
# (http://blog.golang.org/2011/01/go-slices-usage-and-internals.html)
# จากทีม Go สำหรับรายละเอียดที่มากขึ้น
# เกี่ยวกับการออกแบบและสร้าง slices ใน Go

# ตอนนี้เราได้เรียนรู้เกี่ยวกับ arrays และ slices
# กันไปแล้ว
# ต่อไปเราดูอีกโครงสร้างข้อมูลที่สำคัญที่ Go
# มีมาให้แล้วอีกอันคือ maps
