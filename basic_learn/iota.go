package main

func main(){
println("iota是个常量，没使用一次，编译器会加1")
const(
a = iota
b
c
)
println(a,b,c)

const(
i = 1<<iota
j = 3<<iota
k
l
)

println(i,j,k,l)

}



