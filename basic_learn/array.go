package main

import "fmt"
import "unsafe"

func getAverage(arr[10] int) float32{
    var i int
    var sum int = 0
    for i=0; i<len(arr); i++{
        sum += arr[i]
    }
    var average = float32(sum) / float32(len(arr))
    return average
}

func main(){
    println("golang中的数组都是值拷贝而不是引用拷贝。如果你想对数组本身进行操作，那么就操作指向这个数组对象的指针，也就是将值拷贝改为引用拷贝")
    var balance[10] float32 = [10]float32{0.1,0.2}
    var i,j int;
    for i = 0; i<len(balance); i++ {
        fmt.Printf("%f",balance[i])
    }

    var n[10]int
    for j = 0; j < 10 ; j++{
        n[j] = j
    }
    //fmt.Printf("%v", n)
    //fmt.Printf("%+v", n)
    fmt.Printf("%#v", n)

    var average = getAverage(n)
    println(average)
    var pstr = n
    println(unsafe.Sizeof(pstr))
}

