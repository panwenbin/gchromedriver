package main

import (
	"fmt"
)

func t(a interface{}) {
	if a == nil {
		fmt.Println("null")
	}
	//if e, ok := a.(chrome.EmptyResponse); ok {
	//	fmt.Println(e)
	//}
}


func main() {
	//a := struct {
	//	Value struct {
	//		Error string `json:"error"`
	//	} `json:"value"`
	//}{}
	//a := chrome.EmptyResponse{}
	a := ""
	t(a)
}