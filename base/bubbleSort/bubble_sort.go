package main

import "fmt"

func main(){
    bubbleSort()
    selectSort()
}

func bubbleSort(){
    a:= []int {3,5,2,7,4,1}
    for i:=0; i< 6; i++ {
        flag := false
        for j:=0; j<6-i-1; j++ {
            if a[j] > a[j+1] {
                flag = true
                a[j],a[j+1] = a[j+1],a[j]
            }
        }
        if !flag {
            break
        }
        fmt.Println(i,":",a)
    }
}

func selectSort(){
    a:= []int {3,5,2,7,4,1}
    for i:=0; i< 6; i++ {
        flag := false
        for j:=i+1; j<6; j++ {
            if a[i] > a[j] {
                flag = true
                a[i],a[j] = a[j],a[i]
            }
        }
        if !flag {
            break
        }
        fmt.Println(a)
    }
}
