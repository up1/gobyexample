# ขั้นตอนการทำงาน<br/>
# 1. ทำการ copy code ไปไว้ในไฟล์ `hello-world.go` <br/>
# 2. ทำการ run โปรแกรมด้วยคำสั่ง `go run`.
$ go run hello-world.go
hello world

# ในบางครั้ง เราอาจต้องการ build โปรแกรม
# ให้อยู่ในรูปแบบ binaries
# ก็สามารถทำได้ด้วยคำสั่ง `go build`
$ go build hello-world.go
$ ls
hello-world	hello-world.go

# สามารถเรียกใช้งาน binary ได้ดังนี้
$ ./hello-world
hello world

# โดยในตอนนี้ เราสามารถทำการ run และ build
# โปรแกรมที่พัฒนาด้วยภาษา Go ได้แล้ว <br/><br/>
# ดังนั้น มาเริ่มเรียนรู้ความสามารถอื่น ๆ กัน
