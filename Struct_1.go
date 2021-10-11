package main

import "fmt"

type item struct {
	itemid    int
	itemName  string
	itemPrice float64
	itemQ     int
}

func main() {
	var i item = item{}
	fmt.Println("itemid = ", i.itemid)
	fmt.Println("item name = ", i.itemName)
	fmt.Println("itemPrice = ", i.itemPrice)
	fmt.Println("Quntity = ", i.itemQ)

	i.itemid = 1000
	i.itemName = "ABC"
	i.itemPrice = 600
	i.itemQ = 20

	fmt.Println("Item ", i)

	var i2 item = item{itemName: "XYZ", itemPrice: 20, itemid: 200, itemQ: 200}
	fmt.Println(i2)
	var i3 item = item{itemName: "MNO", itemPrice: 660, itemid: 780, itemQ: 240}
	fmt.Println(i3)
	var i4 item = item{itemName: "XYZ", itemPrice: 20, itemid: 200, itemQ: 200}
	fmt.Println(i4)

	p1 := &i2
	fmt.Println("Items refered by pointr :", *(p1))
	fmt.Println("Item id refered by pointer", *&(p1).itemid)
	fmt.Println("Item id refered by pointer", *&(p1).itemid, ",", (*p1).itemName, ",", (*p1).itemPrice, ",", (*p1).itemQ)

	p2 := getItem(300, "LMN", 20, 50)
	fmt.Println(p2)

	p3 := getItem(400, "HIJ", 2600, 30)
	fmt.Println(p3)

	var arr [4]item
	arr[0] = i2
	arr[1] = i3
	arr[2] = i4

	for i := 0; i < len(arr); i++ {
		fmt.Println("Array index : ", i, ",", arr[i])
	}
	fmt.Println(arr)

	i3.displaygrade()
	ptr := &i3
	ptr.modifyname("XYZ")
	fmt.Println(i3)
}

func getItem(itemid int, itemName string, itemPrice float64, itemQ int) *item {
	item1 := item{itemid, itemName, itemPrice, itemQ}
	return &item1
}

func (ireciever item) displaygrade() {
	if ireciever.itemPrice < 500 {
		fmt.Println("Item is A graded")
	} else {
		fmt.Println("Item is B graded")
	}
}

func (p *item) modifyname(newname string) {
	p.itemName = newname
	fmt.Println(" chnaged")
}
