package main

import f "fmt"

type mytype int

type Item struct {
	itemName  string
	itemPrice int
}

func (ireceiver Item) updateprice(itemPrice int) {
	if ireceiver.itemPrice < 10 {
		//	ireceiver.itemPrice += itemPrice
		f.Println("Great ")
		//f.Println("New Price ",ireceiver.itemName , ireceiver.itemPrice)
	} else {
		//f.Println("No change ",ireceiver.itemName , ireceiver.itemPrice)
		f.Println("Not Great ")
	}
}

func (p *Item) modifyName(newName string) {
	p.itemName = newName
	f.Println("changed")
}
func main() {
	var i1 Item = Item{itemName: "A", itemPrice: 10}
	f.Println(i1)
	var i2 Item = Item{itemName: "Y", itemPrice: 20}
	f.Println(i2)
	//i1.modifyName()

}
