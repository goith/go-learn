package main

import "fmt"

//单链表部分操作

type A interface {
	Add()
	Update()
	ShowList()
}
type Item struct {
	no       int
	name     string
	nickname string
	next     *Item
}

func main() {
	it := &Item{} //初始化单链表Head

	n1 := &Item{no: 1, name: "宋江", nickname: "及时雨"}
	n2 := &Item{no: 2, name: "卢俊义", nickname: "玉麒麟"}
	n3 := &Item{no: 3, name: "吴用", nickname: "智多星"}
	n4 := &Item{no: 4, name: "公孙胜", nickname: "入云龙"}

	it.Add(n1)
	it.Add(n2)
	it.Add(n3)
	it.Add(n4)
	it.ShowList()

	fmt.Println("修改后的单链表：")
	newN2 := &Item{no: 2, name: "小卢", nickname: "小玉"}
	it.Update(newN2)
	it.ShowList()
	//删除后
	fmt.Println("删除后的单链表：")
	it.Del(5)
	it.ShowList()

	fmt.Println("反转后的单链表：")
	it.Reverse()
	it.ShowList()
}

func (it *Item) Add(row *Item) {
	temp := it
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = row
	return
}

func (it *Item) Update(row *Item) {
	if it.next == nil {
		fmt.Println("单链表为空")
		return
	}
	for {

		if it.next.no == row.no {
			it.next.name = row.name
			it.next.nickname = row.nickname
			break
		}
		it = it.next
	}
}

func (it *Item) Del(no int) {
	if it.next == nil {
		fmt.Println("单链表为空")
		return
	}
	temp := it.next
	flag := false
	for {
		if temp.next == nil {
			break
		}
		if temp.next.no == no {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("未找到要删除的项：", no)
	}
}

func (it *Item) ShowList() {
	for {
		fmt.Printf("%p, no:%d, name:%s, nickname:%s \n", it, it.no, it.name, it.nickname)
		if it.next == nil {
			break
		}
		it = it.next
	}
}

func (it *Item) Reverse() {
	if it.next == nil {
		fmt.Println("单链表为空")
		return
	}
	cur := it.next
	next := new(Item)
	rL := new(Item)
	for cur != nil {
		next = cur.next
		cur.next = rL.next
		rL.next = cur
		cur = next
	}
	it.next = rL.next
}
