package main

import (
	"fmt"
	"strings"
)

func main() {
	str := GetString("Hello world", "", "")

	fmt.Println(str)

}

func GetString(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == 1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:5], end)
	if e == 1 {
		return
	}
	e += s + e - 1
	return str[s:5]
}

// func hStringOneAr(hString string) string {

// 	h := []byte(hString)
// 	for _, i := range hString {
// 		if 	h

// 	}
// }
