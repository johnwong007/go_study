package main

import "encoding/json"
import "fmt"
//import "os"

type Response1 struct {
    Page int
    Fruits []string
}

type Response2 struct {
    Page int   //'json:"page"'
    Fruits []string //'json:"fruits"'
}


func main(){
    var res1 = Response1{}
    var res2 = Response2{}
    fmt.Printf("res1:%+v\n", res1)
    fmt.Printf("res2:%+v\n", res2)

    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple":5,"lettuce":7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    res2B, _ := json.Marshal(res2)
    fmt.Println(string(res2B))

    byt := []byte("{\"num\":6.13, \"strs\":[\"a\",\"b\"]}")
    var dat map[string] interface{}

    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)




}
