package main                                                                                                                                                               

import "fmt"

func main() {
    x := [3]int{1, 2, 3}
    func(arr [3]int) {
        arr[0] = 7
        fmt.Println("传数组函数内, x:", arr) //prints [7 2 3]
    }(x)
    fmt.Println("传数组函数处理完成后, x:", x) //prints [1 2 3] (not ok if you need [7 2 3])
    fmt.Println("")

    y := [3]int{1, 2, 3}
    func(arr *[3]int) {
        (*arr)[0] = 7
        fmt.Println("传数组指针函数内, y:", arr) //prints &[7 2 3]
    }(&y)
    fmt.Println("传数组指针函数处理完成后, y:", y) //prints [7 2 3]
    fmt.Println("")

    z := []int{1, 2, 3}
    func(arr []int) {
        arr[0] = 7
        fmt.Println("传切片函数内, z:", arr) //prints [7 2 3]
    }(z)
    fmt.Println("传片函数处理完成后, z:", z) //prints [7 2 3]
}
