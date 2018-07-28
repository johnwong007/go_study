package main

func plusPlus(a, b, c int) (int,int) {
    return a+b+c,a*b*c
}

func main(){
    const(
    a = iota
    b
    c
    )
    var sum,ji = plusPlus(a,b,c)
    println(sum,ji)
}
