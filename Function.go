package main

import (
	"fmt"
	f "fmt"
	s "strings"
)

func main() {

	var x string = "CompuCom Company"
	var x1 string = "compucom Company"
	fmt.Println(x)
	fmt.Println(x1)

	fmt.Println("Lengthof String is:", len(x))
	fmt.Println("Ascii value of A is ", "A"[0])
	fmt.Println("Converting string Upercase:", s.ToUpper(x))
	fmt.Println("Converting string lower case :", s.ToLower(x))

	fmt.Println("Compare two string ", s.Compare("x", "x1"))
	fmt.Println(s.Compare("x1", "x"))
	fmt.Println(s.Compare("x", "x"))

	f.Println("Contains to check does it contains", s.Contains("Compucom", "com"))
	f.Println(s.Contains("Compucom", "Pu"))
	f.Println(s.ContainsAny("Compare", "om-pu"))

	concati := f.Sprintf("%s %s", x, x1)
	f.Println(concati)

	f.Println("Repeat Function")
	f.Println(s.Repeat(x, 4))

	f.Println("Replace Function")
	f.Println(s.Replace(x, "o", "Z", 3))

	f.Println("Count Words :")
	f.Println(s.Count(x, "c"))

	f.Println("Count No of Words in string x")
	var arr []string = s.Split(x, "")
	f.Println(len(arr))
	for i, v := range arr {
		fmt.Println("Index : ", i, "value : ", v)
	}
	f.Println("Substring")
	substring := x[5:]
	fmt.Println(substring)

	var strData = "TQ"
	f.Println()
	for i := 0; i < len(strData); i++ {
		fmt.Printf("Character at %d Index Position = %c\n", i, strData[i])

		f.Println()
		f.Println("Concatinate two string:", x+x1)

	}
}
