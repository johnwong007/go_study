package main

import (
"fmt"
"strconv"
)

func main(){
fmt.Println("GO数值和字符串的相互转换!")
var age int = 8;
var myprintln = fmt.Println
myprintln("age is "+strconv.Itoa(age))

var delta string = "2"
//var ageTmp, error = strconv.Atoi(delta)

//ageTmp, error := strconv.Atoi(delta)
//if error != nil{
//println("转换字符串失败")
//}
ageTmp, _ := strconv.Atoi(delta)
var ageNow int = age+ageTmp
myprintln("now age is "+strconv.Itoa(ageNow))
}



