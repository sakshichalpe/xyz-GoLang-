package main

import "fmt"

func main() {
	//no memory
	var empdept = make(map[int]string)
	fmt.Println(empdept)

	if empdept == nil {
		fmt.Println("Map is nil")
	} else {
		fmt.Println("Map is not nil")
	}
	empdept[1] : "Sales",
	empdept[2] : "Devp",
	// fmt.Println(empdept)

	// 	//create map usingliteral
	// 	var empdept = map[int]string{1: "sales", 2: "devp"}
	// 	fmt.Println(empdept)
	// 	fmt.Println("dep of 1 is:", empdept[1])

	//     //range for iteration
	//   for k,v := range empdept{
	// 	  fmt.Printf("department of %d is %s", k,v)
	//   }
	//var item = map[string]int{"samosa": 500, "idli": 5001, "Kachori": 499}

	// for k, v := range item {
	// 	fmt.Printf(" Price of item %s is %d", k, v)
	// }

	// if (v < 500){
	// 	for _, v := range item {
	// 		fmt.Printf(" Price of item %s is %d",  v)
	//    }

	// var item = map[string]int{"samosa": 500, "idli": 5001, "Kachori": 499}
	// for k,v := range item{
	// //v, k := item["Kachori"]
	//   if v,k := item["Kachori"] {
	// fmt.Println("Key found value is", k,v)
	// item[k]= v + 500
	//     fmt.Println("increased", k,v)
	}else{
		fmt.Println("No found")
	}
	}

	var string = "Jay jawan jay kisan "

	for i, item := range string {
		fmt.Printf("%d  %s", i, item)
	 }
}
