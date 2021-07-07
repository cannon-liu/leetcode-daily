package main

import "log"

func convertToTitle(columnNumber int) string {
	data := 0
	result := ""
	list := make([]string,0)
	for {
		data = columnNumber % 26
		log.Println("mod is ",data)
		columnNumber = columnNumber / 26
		if columnNumber == 0 {
			break
		}
		ch := 'A'+data
		list = append(list,string(ch))
	}
	log.Println("list is ",list)
	for i:=len(list)-1;i>=0;i-- {
		result = result + list[i]
		log.Println("deal result: ",result)
	}
	return result
}


func main() {
	result := convertToTitle(701)
	log.Println("result is ",result)
}