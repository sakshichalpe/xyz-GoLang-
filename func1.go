package main
import f "fmt"

func add(){
    
	var num1 int
    var num2 int
f.Println("enter 1st the number ")
f.Scan(&num1)
f.Println("enter 2nd the number ")
f.Scan(&num2)
add := num1 + num2
f.Println(add)
}
func main(){
add()
}