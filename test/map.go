package main

import "fmt"

func main() {
    type Map1 map[string]interface{}
    type Map2 map[string]int
    m := Map1{"foo": Map2{"first": 1}, "boo": Map2{"second": 2}}
    //m = map[foo:map[first: 1] boo: map[second: 2]]
    fmt.Println("m:", m)
    for k, v := range m {
        fmt.Println("k:", k, "v:", v)
    }

}
