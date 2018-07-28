package main

import "unsafe"

const(
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
    d = unsafe.Sizeof("a")
    e = unsafe.Sizeof("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
)

func main(){
    println("string内部实现//组成可以理解成此结构体\ntype string struct{\n\nchar* buffer;\n\nsize_tlen;\n\n}")
    println("go语言中指针占8位")
    println(a,b,c,d)
    var arr = [17]int{1,2,3,4,5,6,7,8,9,20}
    println(len(arr), unsafe.Sizeof(arr))
    var arr1 = [17]string{"1"}
    println(len(arr1), unsafe.Sizeof(arr1))
    var tmp = "test"
    var pstr *string = &tmp
    println(unsafe.Sizeof(pstr))
}

