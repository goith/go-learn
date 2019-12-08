package main

import "fmt"

func main(){
    bubbleSort()
    selectSort()
    
    insertSort()
}

func bubbleSort(){
    a:= []int {3,5,2,7,4,1}
    for i:=0; i< 6-1; i++ {
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
    for i:=0; i< 6-1; i++ {
        min := a[i]
        minIndex := i
        for j:=i+1; j<6; j++ {
            if min > a[j] {
                min = a[j]
                minIndex = j
            }
        }
        if minIndex != i {
            a[minIndex] = a[i]
            a[i] = min
        }
        fmt.Println(i,":",a)
    }
}

/*
func bubbleSort2(){
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
*/


func insertSort(){
    arr:= []int {3,1,2,7,6,4}
    le :=len(arr)
    for i:=1; i< le; i++ {
        insertVal := arr[i]
        insertIndex := i - 1
        for insertIndex >= 0 && insertVal < arr[insertIndex] {
            arr[insertIndex + 1] = arr[insertIndex]
            insertIndex--
        }
        fmt.Println("........", arr)
        if insertIndex + 1 != i {
            arr[insertIndex + 1] = insertVal
        }
        fmt.Println(i,":",arr)
    }
}
