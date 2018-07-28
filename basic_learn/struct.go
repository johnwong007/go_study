//创建结构体的六种方式
package main

import "fmt"
import "unsafe"

type peple struct{
    name string;
    age int;
};

type Student struct{
    id int
    name string
    class string
}

func main(){
    println("创建结构体的六种方式")
    //第一种方式创建结构体
    var zhangsan peple;
    zhangsan.name="peple"
    zhangsan.age = 18
    fmt.Printf("%+v\n", zhangsan)
    println("zhangsan占用内存大小：",unsafe.Sizeof(zhangsan))
    //第二种方式创建结构体
    var lisi = peple{"lisi", 20}
    fmt.Printf("%+v\n", lisi)
    //第三种方式创建结构体
    var wangwu = peple{name:"wangwu", age:22}
    fmt.Printf("%#v\n", wangwu)
    //第四种方式创建结构体
    var lily *Student = new(Student)
    lily.id = 101
    lily.name = "lily"
    lily.class = "三年级二班"
    fmt.Printf("%v\n", lily)
    //第五种方式创建结构体
    var hanmeimei = &Student{102, "韩梅梅", "三年级四班"}
    fmt.Printf("%+v\n", hanmeimei)
    //第六种方式创建结构体
    var lilei = &Student{id:103, name:"李雷", class:"三年级一班"}
    fmt.Printf("%#v\n", lilei)
    println("\n结构体作为函数参数进行传递")
    dealfunc(wangwu)
    println("wangwu's name is "+wangwu.name)
    dealfunc1(hanmeimei)
    println("hanmeimei's name is "+hanmeimei.name)
}

func dealfunc(stru peple){
    stru.name = stru.name+"(changed by dealfunc)"
    println(stru.name)
}

func dealfunc1(stru *Student){
    stru.name = stru.name+"(changed by dealfunc1)"
    println(stru.name)
}

