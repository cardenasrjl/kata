package main

import (
	"fmt"
	"strconv"
	"strings"
)

/***

Run-length encoding is a fast and simple method of encoding strings. The basic idea is to represent repeated successive characters as a single count and character. For example, the string "AAAABBBCCDAA" would be encoded as "4A3B2C1D2A".

Implement run-length encoding and decoding. You can assume the string to be encoded have no digits and consists solely of alphabetic characters. You can assume the string to be decoded is valid.


*/


func main () {
	
	string := "munddasdfasdfasdfasdfasdfooooaaaaadddddddddddddddddad"
	
	print(decode(encode(string)))

	
}

func encode(s string) string {
	arrData := strings.Split(s,"")
	current := ""
	count := 0
	encoded := ""
	for _,v := range arrData {
		
		if current == v {
			count++
			continue
		}

		if current != v   {
			if count > 0 {
				encoded += fmt.Sprintf("%d%s", count, current)
			}
			count=1
			current = v
		}
		
	}

	encoded += fmt.Sprintf("%d%s",count,current)
	
	return encoded
}


func decode(s string) string {
	arrData := strings.Split(s,"")
	decoded := ""
	for i,v := range arrData {
		if value, err := strconv.Atoi(v); err == nil {
			decoded += repeatString(arrData[i+1], value) 
		}
	}
	
	return decoded
}

func repeatString(s string ,n int) string {
	result := ""
	for i := 0;i <n;i++ {
		result+=s
	}
	return result
}
