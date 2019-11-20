package main                                                                                                                                                               

import "fmt"

type I interface {
    name()
}
type S struct{}

func (*S) name() {
    fmt.Println("a")
}

func main() {
    var value I = &S{}
    //value.name() //可以调用

    var point = &value
    (*point).name() //能调用
    point.name()    //不能调用
}
